package trips

import (
	"fmt"
	"main/config"
	"main/lib"
	shapes_coordinates "main/services/geo/shapes"
	stops_coordinates "main/services/geo/stops"
	"main/types"
	registry "main/validations"
	validations "main/validations/trips/validations"
	"slices"
)

func init() {
	registry.Register("trips", RunValidations)
}

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Trips Validations...")

	// Pre-compute stop_times data per trip_id for performance
	// This avoids N+1 queries in stop_sequence and shape_id validations
	lib.AppLogger.Debug("Pre-computing stop_times data per trip...")
	tripStopTimesCache := make(map[string][]types.StopTimeRaw)

	err := gtfs.IterateStopTimes(func(i int, rawStopTime types.StopTimeRaw) error {
		if rawStopTime.TripId == "" {
			return nil
		}
		tripStopTimesCache[rawStopTime.TripId] = append(tripStopTimesCache[rawStopTime.TripId], rawStopTime)
		return nil
	})
	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error pre-computing trip stop times: %v", err))
	}
	lib.AppLogger.Debug(fmt.Sprintf("Pre-computed stop_times for %d trips", len(tripStopTimesCache)))

	// Pre-compute stops cache (stop_id -> coordinates) for stop coordinates validation
	lib.AppLogger.Debug("Pre-computing stops cache...")
	stopsCache := make(map[string]types.StopCoordinatesValidation)
	err = gtfs.IterateStops(func(_ int, stop types.StopRaw) error {
		if stop.StopId == "" || stop.StopLat == "" || stop.StopLon == "" {
			return nil
		}
		if _, exists := stopsCache[stop.StopId]; !exists {
			stopsCache[stop.StopId] = types.StopCoordinatesValidation{
				StopId:  stop.StopId,
				StopLat: stop.StopLat,
				StopLon: stop.StopLon,
			}
		}
		return nil
	})
	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error pre-computing stops cache: %v", err))
	}
	lib.AppLogger.Debug(fmt.Sprintf("Pre-computed stops cache for %d stops", len(stopsCache)))

	// Pre-compute shape IDs used by trips, then load chunked shape data
	lib.AppLogger.Debug("Pre-computing shape IDs from trips...")
	shapeIdsByTrips := make(map[string]struct{})
	err = gtfs.IterateTrips(func(_ int, rawTrip types.TripRaw) error {
		if rawTrip.ShapeId != "" {
			shapeIdsByTrips[rawTrip.ShapeId] = struct{}{}
		}
		return nil
	})
	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error pre-computing shape IDs: %v", err))
	}
	lib.AppLogger.Debug("Pre-loading shape chunked cache...")
	shapeChunkedCache := make(map[string]*shapes_coordinates.ShapeChunkedData)
	for shapeID := range shapeIdsByTrips {
		shapes, err := gtfs.GetShapesByShapeId(shapeID)
		if err != nil || len(shapes) == 0 {
			continue
		}
		shapeCoords := make([]types.ShapeCoordinatesValidation, 0, len(shapes))
		for _, shape := range shapes {
			if shape.ShapePtLat == "" || shape.ShapePtLon == "" {
				continue
			}
			shapeCoords = append(shapeCoords, types.ShapeCoordinatesValidation{
				ShapeId:    shape.ShapeId,
				ShapePtLat: shape.ShapePtLat,
				ShapePtLon: shape.ShapePtLon,
				ShapePtSeq: shape.ShapePtSequence,
			})
		}
		if data := shapes_coordinates.BuildShapeChunkedData(shapeCoords); data != nil {
			shapeChunkedCache[shapeID] = data
		}
	}
	lib.AppLogger.Debug(fmt.Sprintf("Pre-loaded %d shapes into chunked cache", len(shapeChunkedCache)))

	// Pre-compute stop-to-shape distance cache for all unique (stop_id, shape_id) pairs
	// This avoids O(trips × stops × points) Haversine calls by computing once per unique pair
	lib.AppLogger.Debug("Pre-computing stop-to-shape distance cache...")
	tripToShapeID := make(map[string]string)
	_ = gtfs.IterateTrips(func(_ int, rawTrip types.TripRaw) error {
		if rawTrip.ShapeId != "" && rawTrip.TripId != "" {
			tripToShapeID[rawTrip.TripId] = rawTrip.ShapeId
		}
		return nil
	})
	stopToShapeIDs := make(map[string]map[string]struct{})
	for tripID, stopTimes := range tripStopTimesCache {
		shapeID := tripToShapeID[tripID]
		if shapeID == "" {
			continue
		}
		for _, st := range stopTimes {
			if st.StopId == "" {
				continue
			}
			if stopToShapeIDs[st.StopId] == nil {
				stopToShapeIDs[st.StopId] = make(map[string]struct{})
			}
			stopToShapeIDs[st.StopId][shapeID] = struct{}{}
		}
	}
	stopToShapeIDsSlice := make(map[string][]string)
	var stopCoordinatesForCache []types.StopCoordinatesValidation
	for stopID, shapeSet := range stopToShapeIDs {
		stop, ok := stopsCache[stopID]
		if !ok {
			continue
		}
		shapeIDs := make([]string, 0, len(shapeSet))
		for sid := range shapeSet {
			shapeIDs = append(shapeIDs, sid)
		}
		stopToShapeIDsSlice[stopID] = shapeIDs
		stopCoordinatesForCache = append(stopCoordinatesForCache, stop)
	}
	stopClosestShapePointsCache := stops_coordinates.BuildStopClosestShapePointsDistanceMapPerStopShape(stopCoordinatesForCache, stopToShapeIDsSlice, shapeChunkedCache)
	lib.AppLogger.Debug(fmt.Sprintf("Pre-computed stop-to-shape distance cache for %d violations", len(stopClosestShapePointsCache)))

	// Caches for GetRowsById to avoid repeated DB queries (many trips share route_id and service_id)
	routeRowsCache := make(map[string][]int)
	calendarRowsCache := make(map[string][]int)
	calendarDatesRowsCache := make(map[string][]int)

	// Create progress tracker
	tracker := lib.CreateProgressTracker(gtfs, "trips", config.ProgressThresholdLarge)
	var tripsGroupedByPattern types.TripGroupedByPattern = make(types.TripGroupedByPattern)
	var tripsGroupedByShapeId types.TripGroupedByShapeId = make(types.TripGroupedByShapeId)

	err = gtfs.IterateTrips(func(i int, rawTrips types.TripRaw) error {
		tracker.Track()
		trip := validations.ParseTrips(rawTrips, i)

		if trip == (types.Trip{}) {
			return nil
		}

		var tripRules *types.TripsRules
		if rules != nil {
			tripRules = &rules.Trips
		}

		// Validate trip_id
		validations.TripIdValidation(&trip, i, &gtfs)

		// Validate shape_id (pass cached stop_times data and route cache)
		validations.ShapeIdValidation(&trip, i, &gtfs, tripRules, tripStopTimesCache, routeRowsCache)

		// Validate route_id (pass route cache)
		validations.RouteIdValidation(&trip, i, &gtfs, routeRowsCache)

		// Validate service_id (pass calendar caches)
		validations.ServiceIdValidation(&trip, i, &gtfs, calendarRowsCache, calendarDatesRowsCache)

		// Validate trip_headsign
		validations.TripHeadsignValidation(&trip, i, &gtfs, tripRules)

		// Validate trip_short_name
		validations.TripShortNameValidation(&trip, i, &gtfs, tripRules)

		// Validate direction_id
		validations.DirectionIdValidation(&trip, i, &gtfs, tripRules)

		// Validate block_id
		validations.BlockIdValidation(&trip, i, &gtfs, tripRules)

		// Validate wheelchair_accessible
		validations.WheelchairAccessibleValidation(&trip, i, &gtfs, tripRules)

		// Validate bikes_allowed
		validations.BikesAllowedValidation(&trip, i, &gtfs, tripRules)

		// Validate stop_times.stop_sequence (pass cached stop_times data)
		groupHash := validations.StopSequenceValidation(&trip, i, &gtfs, tripRules, tripStopTimesCache)

		// Validate stop_coordinates (pass precomputed stop-to-shape distance cache)
		validations.StopCoordinatesByTripIdValidation(&trip, i, &gtfs, tripStopTimesCache, stopsCache, stopClosestShapePointsCache, tripRules)

		// CMET SPECIFIC VALIDATIONS
		hasPatternId := validations.PatternIdValidation(&trip, i, &gtfs, tripRules)
		if hasPatternId {
			group := tripsGroupedByPattern[*trip.PatternId]
			group.Trips = append(group.Trips, trip)

			if !slices.Contains(group.Hash, groupHash) {
				group.Hash = append(group.Hash, groupHash)
			}

			tripsGroupedByPattern[*trip.PatternId] = group
		}

		// Validate direction_id matches pattern_id
		validations.DirectionPatternIdMatchValidation(&trip, i, &gtfs, tripRules)

		// Validate shape_id is unique per pattern_id
		if trip.ShapeId != nil {
			group := tripsGroupedByShapeId[*trip.ShapeId]
			group.Trips = append(group.Trips, trip)

			if !slices.Contains(group.Hash, groupHash) {
				group.Hash = append(group.Hash, groupHash)
			}

			tripsGroupedByShapeId[*trip.ShapeId] = group
		}
		// Validate trip_id_limit_characters
		validations.TripIdLimitCharactersValidation(&trip, i, &gtfs, tripRules)

		// Validate pattern_id_format
		validations.PatternIdFormatValidation(&trip, i, &gtfs, tripRules)

		return nil
	})

	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error iterating trips: %v", err))
	} else {
		lib.AppLogger.Info(fmt.Sprintf("Completed trips.txt validation: %d rows processed", tracker.GetProcessedCount()))
	}

	validations.ValidatePatternGroups(tripsGroupedByPattern, tripsGroupedByShapeId, &gtfs)
}

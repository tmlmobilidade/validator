package trips

import (
	"strconv"

	"main/lib"
	"main/services"
	stops_coordinates "main/services/geo/stops"
	"main/types"
)

/*
# Attributes
  - File: [trips.txt]
  - Field: stop_coordinates_by_trip_id
  - Presence: optional
  - Type: coordinates

# Description
Validate if the stop_lat and stop_lon are valid.
*/

func StopCoordinatesByTripIdValidation(trip *types.Trip, row int, gtfs *types.Gtfs, tripStopTimesCache map[string][]types.StopTimeRaw, stopsCache map[string]types.StopCoordinatesValidation, stopClosestShapePointsCache map[string]types.StopClosestShapePointsInfo, rules *types.TripsRules) []types.StopCoordinatesValidation {
	ctx := lib.NewValidationContext("coordinates", "trips.txt", "coordinates_validation", row, services.AppMessageService)
	if rules != nil && rules.StopCoordinatesByTripId.Severity != "" {
		ctx.WithSeverity(rules.StopCoordinatesByTripId.Severity)
	}

	stopTimesRaw, exists := tripStopTimesCache[*trip.TripId]
	if !exists {
		stopTimes, err := gtfs.GetRowsById("stop_times", *trip.TripId)
		if err != nil {
			return []types.StopCoordinatesValidation{}
		}
		for _, row := range stopTimes {
			stopTimeRaw, err := gtfs.GetStopTime(row)
			if err != nil {
				continue
			}
			stopTimesRaw = append(stopTimesRaw, stopTimeRaw)
		}
	}

	if ctx.ShouldSkip() {
		return []types.StopCoordinatesValidation{}
	}

	if trip.ShapeId == nil || *trip.ShapeId == "" {
		return []types.StopCoordinatesValidation{}
	}

	stopCoordinates := make([]types.StopCoordinatesValidation, 0)
	seenStops := make(map[string]struct{})

	for _, stopTimeRaw := range stopTimesRaw {
		stopId := stopTimeRaw.StopId
		if _, seen := seenStops[stopId]; seen {
			continue
		}
		seenStops[stopId] = struct{}{}
		stop, ok := stopsCache[stopId]
		if !ok {
			continue
		}
		stopCoordinates = append(stopCoordinates, stop)
	}

	// Look up precomputed violations from cache (keyed by "stop_id|shape_id")
	for _, stop := range stopCoordinates {
		key := stops_coordinates.StopShapeCacheKey(stop.StopId, *trip.ShapeId)
		if info, ok := stopClosestShapePointsCache[key]; ok {
			stopLat, _ := strconv.ParseFloat(stop.StopLat, 64)
			stopLon, _ := strconv.ParseFloat(stop.StopLon, 64)
			ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("coordinates_validation.invalid_distance_to_shape", stopLat, stopLon, info.ShapeID, info.ClosestShapePtSeq, info.ClosestShapePtLat, info.ClosestShapePtLon, info.DistanceMeters))
		}
	}

	return stopCoordinates
}

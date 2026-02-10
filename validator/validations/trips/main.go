package trips

import (
	"fmt"
	"main/config"
	"main/lib"
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

	// Create progress tracker
	tracker := lib.CreateProgressTracker(gtfs, "trips", config.ProgressThresholdLarge)
	var tripsGroupedByPattern types.TripGroupedByPattern = make(types.TripGroupedByPattern)

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

		// Validate shape_id (pass cached stop_times data)
		validations.ShapeIdValidation(&trip, i, &gtfs, tripRules, tripStopTimesCache)

		// Validate route_id
		validations.RouteIdValidation(&trip, i, &gtfs)

		// Validate service_id
		validations.ServiceIdValidation(&trip, i, &gtfs)

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

		// Validate trip_id_limit_characters
		validations.TripIdLimitCharactersValidation(&trip, i, tripRules)

		// Validate stop_times.stop_sequence (pass cached stop_times data)
		groupHash := validations.StopSequenceValidation(&trip, i, &gtfs, tripRules, tripStopTimesCache)

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

		return nil
	})

	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error iterating trips: %v", err))
	} else {
		lib.AppLogger.Info(fmt.Sprintf("Completed trips.txt validation: %d rows processed", tracker.GetProcessedCount()))
	}

	//Validate pattern_id_group
	validations.PatternIdGroupValidation(tripsGroupedByPattern, &gtfs)
}

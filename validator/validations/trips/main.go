package trips

import (
	"main/lib"
	"main/types"
	validations "main/validations/trips/validations"
	"slices"
)

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Trips Validations...")

	var tripsGroupedByPattern types.TripGroupedByPattern = make(types.TripGroupedByPattern)

	for i, rawTrips := range gtfs.Trip {
		trip := validations.ParseTrips(rawTrips, i)

		if trip == (types.Trip{}) {
			continue
		}

		// Validate trip_id
		validations.TripIdValidation(&trip, i, &gtfs)

		// Validate shape_id
		validations.ShapeIdValidation(&trip, i, &gtfs, &rules.Trips)

		// Validate route_id
		validations.RouteIdValidation(&trip, i, &gtfs)

		// Validate service_id
		validations.ServiceIdValidation(&trip, i, &gtfs)

		// Validate trip_headsign
		validations.TripHeadsignValidation(&trip, i, &gtfs, &rules.Trips)

		// Validate trip_short_name
		validations.TripShortNameValidation(&trip, i, &gtfs, &rules.Trips)

		// Validate direction_id
		validations.DirectionIdValidation(&trip, i, &gtfs, &rules.Trips)

		// Validate block_id
		validations.BlockIdValidation(&trip, i, &gtfs, &rules.Trips)

		// Validate wheelchair_accessible
		validations.WheelchairAccessibleValidation(&trip, i, &gtfs, &rules.Trips)

		// Validate bikes_allowed
		validations.BikesAllowedValidation(&trip, i, &gtfs, &rules.Trips)

		// Validate stop_times.stop_sequence
		groupHash := validations.StopSequenceValidation(&trip, i, &gtfs, &rules.Trips)

		// CMET SPECIFIC VALIDATIONS
		hasPatternId := validations.PatternIdValidation(&trip, i, &gtfs, &rules.Trips)
		if hasPatternId {
			group := tripsGroupedByPattern[*trip.PatternId]
			group.Trips = append(group.Trips, trip)

			if !slices.Contains(group.Hash, groupHash) {
				group.Hash = append(group.Hash, groupHash)
			}

			tripsGroupedByPattern[*trip.PatternId] = group
		}
	}

	//Validate pattern_id_group
	validations.PatternIdGroupValidation(tripsGroupedByPattern, &gtfs)
}

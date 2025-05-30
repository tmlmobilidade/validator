package trips

import (
	"main/lib"
	"main/types"
	validations "main/validations/trips/validations"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Trips Validations...")

	var tripsGroupedByPattern types.TripGroupedByPattern = make(types.TripGroupedByPattern)

	for i, rawTrips := range gtfs.Files["trips"] {
		trip := validations.ParseTrips(rawTrips, i)

		if trip == (types.Trip{}) {
			continue
		}

		// Validate trip_id
		validations.TripIdValidation(&trip, i, &gtfs)

		// Validate shape_id
		validations.ShapeIdValidation(nil, &trip, i, &gtfs)

		// Validate route_id
		validations.RouteIdValidation(&trip, i, &gtfs)

		// Validate service_id
		validations.ServiceIdValidation(&trip, i, &gtfs)

		// Validate trip_headsign
		validations.TripHeadsignValidation(nil, &trip, i, &gtfs)

		// Validate trip_short_name
		validations.TripShortNameValidation(nil, &trip, i, &gtfs)

		// Validate direction_id
		validations.DirectionIdValidation(nil, &trip, i, &gtfs)

		// Validate block_id
		validations.BlockIdValidation(nil, &trip, i, &gtfs)

		// Validate wheelchair_accessible
		validations.WheelchairAccessibleValidation(nil, &trip, i, &gtfs)

		// Validate bikes_allowed
		validations.BikesAllowedValidation(nil, &trip, i, &gtfs)

		// Validate stop_times.stop_sequence
		validations.StopSequenceValidation(&trip, i, &gtfs)

		
		// CMET SPECIFIC VALIDATIONS
		hasPatternId := validations.PatternIdValidation(lib.Ptr(types.SEVERITY_ERROR), &trip, i, &gtfs)
		if hasPatternId {
			tripsGroupedByPattern[*trip.PatternId] = append(tripsGroupedByPattern[*trip.PatternId], trip)
		}
	}

	//Validate pattern_id_group
	validations.PatternIdGroupValidation(tripsGroupedByPattern, &gtfs)
}

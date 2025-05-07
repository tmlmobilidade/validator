package trips

import (
	"main/lib"
	"main/types"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Trips Validations...")

	for i, rawTrips := range gtfs.Files["trips"] {
		trip := ParseTrips(rawTrips, i, &gtfs)

		// Validate trip_id
		TripIdValidation(&trip, i, &gtfs)

		// Validate shape_id
		ShapeIdValidation(nil, &trip, i, &gtfs)

		// Validate route_id
		RouteIdValidation(&trip, i, &gtfs)

		// Validate service_id
		ServiceIdValidation(&trip, i, &gtfs)

		// Validate trip_headsign
		TripHeadsignValidation(nil, &trip, i, &gtfs)

		// Validate trip_short_name
		TripShortNameValidation(nil, &trip, i, &gtfs)
		
		// Validate direction_id
		DirectionIdValidation(nil, &trip, i, &gtfs)
		
		// Validate block_id
		BlockIdValidation(nil, &trip, i, &gtfs)
		
		// Validate wheelchair_accessible
		WheelchairAccessibleValidation(nil, &trip, i, &gtfs)

		// Validate bikes_allowed
		BikesAllowedValidation(nil, &trip, i, &gtfs)
	}
}

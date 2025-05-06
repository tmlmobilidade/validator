package trips

import (
	"main/lib"
	"main/types"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Trips Validations...")

	for i, rawTrips := range gtfs.Files["trips"] {
		trip := ParseTrips(rawTrips, i, &gtfs)

		// Validate Trip ID
		TripIdValidation(&trip, i, &gtfs)

		// Validate Route Id
		RouteIdValidation(&trip, i, &gtfs)

		// Validate Service ID
		ServiceIdValidation(&trip, i, &gtfs)

		// Trip Headsign Validation
		TripHeadsignValidation(nil, &trip, i, &gtfs)

		// Trip Short Name Validation
		TripShortNameValidation(nil, &trip, i, &gtfs)

		// Bikes Allowed Validation
		BikesAllowedValidation(nil, &trip, i, &gtfs)

		// Wheelchair Accessible Validation
		WheelchairAccessibleValidation(nil, &trip, i, &gtfs)

		// Direction ID Validation
		DirectionIdValidation(nil, &trip, i, &gtfs)

		// Block ID Validation
		BlockIdValidation(nil, &trip, i, &gtfs)
		
	}
}

package trips

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes
  - File: [trips.txt]
  - Field: trip_id
  - Presence: Required (TML Rule)
  - Type: Foreign Key referencing trips.trip_id

# Description

Ensures the trip_id is less than or equal to 32 characters.

[trips.txt]: https://gtfs.org/schedule/reference/#trips
*/

func TripIdLimitCharactersValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules) {
	ctx := lib.NewValidationContext("trip_id_limit_characters", "trips.txt", "trip_id_limit_characters_validation", "trip_id_limit_characters_rule", row, services.AppMessageService)
	if rules != nil && rules.TripIdLimitCharacters.Severity != "" {
		ctx.WithSeverity(rules.TripIdLimitCharacters.Severity)
	}

	if trip.TripId == nil {
		return
	}

	if len(*trip.TripId) > 36 {
		ctx.AddError(ctx.GetTranslatedMessage("trip_id_limit_characters_validation.too_long"))
		return
	}

}

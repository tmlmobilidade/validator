package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	"slices"
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

func TripIdLimitCharactersValidation(trip *types.Trip, row int, rules *types.TripsRules) {
	ctx := lib.NewValidationContext("trip_id", "trips.txt", "trip_id_limit_characters_validation", row, services.AppMessageService)
	if rules != nil && rules.TripIdLimitCharacters.Severity != "" {
		ctx.WithSeverity(rules.TripIdLimitCharacters.Severity)
	}

	if trip.TripId == nil {
		return
	}

	if len(*trip.TripId) > 31 {
		ctx.AddError(ctx.GetTranslatedMessage("trip_id_limit_characters_validation.too_long"))
		return
	}

	// Validate rules
	if rules != nil && rules.TripIdLimitCharacters.Options != nil {
		if slices.Contains(*rules.TripIdLimitCharacters.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.TripIdLimitCharacters.Options, *trip.TripId) {
			ctx.AddError(ctx.GetTranslatedMessage("trip_id_limit_characters_validation.not_allowed", *trip.TripId))
			return
		}
	}
}

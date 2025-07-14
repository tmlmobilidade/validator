package trips

import (
	"fmt"
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

- File: [trips.txt]
- Field: bikes_allowed
- Presence: Optional
- Type: Enum

# Description

Indicates whether bikes are allowed.

Valid options are:

  - 0 or empty - No bike information for the trip.
  - 1 - Vehicle being used on this particular trip can accommodate at least one bicycle.
  - 2 - No bicycles are allowed on this trip.

[trips.txt]: https://gtfs.org/schedule/reference/#tripstxt
*/
func BikesAllowedValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.BikesAllowed.Severity != "" {
		s = rules.BikesAllowed.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "bikes_allowed",
			FileName:     "trips.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "bikes_allowed_validation",
		})
	}

	// 1. Validate bikes_allowed is required
	if trip.BikesAllowed == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"bikes_allowed_validation.required",
				"bikes_allowed_validation.recommended",
			),
		)

		addMessage(message, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("bikes_allowed_validation.forbidden"), s)
		return
	}

	// 2. Validate bikes_allowed is 0 or 1 if it exists
	if trip.BikesAllowed != nil {
		validBikesAllowed := map[int]bool{0: true, 1: true, 2: true}
		if !validBikesAllowed[*trip.BikesAllowed] {
			addMessage(i18n.AppTranslator.Get("bikes_allowed_validation.invalid"), s)
			return
		}
	}

	// 3. Validate Rule Options
	if rules != nil && rules.BikesAllowed.Options != nil {
		if slices.Contains(*rules.BikesAllowed.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.BikesAllowed.Options, fmt.Sprintf("%d", *trip.BikesAllowed)) {
			addMessage(i18n.AppTranslator.Get("bikes_allowed_validation.not_allowed", map[string]interface{}{"value": *trip.BikesAllowed}), s)
			return
		}
	}
}

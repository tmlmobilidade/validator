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
- Field: wheelchair_accessible
- Presence: Optional
- Type: Enum

# Description

Indicates wheelchair accessibility. Valid options are:

  - 0 or empty - No accessibility information for the trip.
  - 1 - Vehicle being used on this particular trip can accommodate at least one rider in a wheelchair.
  - 2 - No riders in wheelchairs can be accommodated on this trip.

[trips.txt]: https://gtfs.org/schedule/reference/#tripstxt
*/
func WheelchairAccessibleValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.WheelchairAccessible.Severity != "" {
		s = rules.WheelchairAccessible.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "wheelchair_accessible",
			FileName:     "trips.txt",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
			ValidationID: "wheelchair_accessible_validation",
		})
	}

	// 1. Validate wheelchair_accessible is required
	if trip.WheelchairAccessible == nil {
		if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"wheelchair_accessible_validation.required",
				"wheelchair_accessible_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	// 2. Validate wheelchair_accessible is forbidden
	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("wheelchair_accessible_validation.forbidden"), s)
		return
	}

	// 3. Validate wheelchair_accessible is 0 or 1 if it exists
	if trip.WheelchairAccessible != nil {
		validWheelchairAccessible := map[int]bool{0: true, 1: true, 2: true}
		if !validWheelchairAccessible[*trip.WheelchairAccessible] {
			addMessage(i18n.AppTranslator.Get("wheelchair_accessible_validation.invalid"), s)
			return
		}
	}

	// Validate Rule Options
	if rules != nil && rules.WheelchairAccessible.Options != nil {
		if slices.Contains(*rules.WheelchairAccessible.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.WheelchairAccessible.Options, fmt.Sprintf("%d", *trip.WheelchairAccessible)) {
			addMessage(i18n.AppTranslator.Get("wheelchair_accessible_validation.not_allowed", map[string]interface{}{"value": *trip.WheelchairAccessible}), s)
			return
		}
	}
}

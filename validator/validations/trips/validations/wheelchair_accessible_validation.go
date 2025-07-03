package trips

import (
	"fmt"
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

	// 1. Validate wheelchair_accessible is 0 or 1 if it exists
	if trip.WheelchairAccessible != nil {
		validWheelchairAccessible := map[int]bool{0: true, 1: true, 2: true}
		if !validWheelchairAccessible[*trip.WheelchairAccessible] {
			addMessage("Invalid wheelchair_accessible value. Valid values are 0, 1, and 2.", s)
			return
		}
	}

	// 2. Validate wheelchair_accessible is required
	if s == types.SEVERITY_IGNORE {
		return
	}

	if trip.WheelchairAccessible == nil {
		addMessage(lib.IfThenElse(s == types.SEVERITY_ERROR, "wheelchair_accessible is required", "wheelchair_accessible is recommended"), s)
		return
	}

	// Validate Rule Options
	if rules != nil && rules.WheelchairAccessible.Options != nil {
		if slices.Contains(*rules.WheelchairAccessible.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.WheelchairAccessible.Options, fmt.Sprintf("%d", *trip.WheelchairAccessible)) {
			addMessage(fmt.Sprintf("wheelchair_accessible is not allowed: %d", *trip.WheelchairAccessible), s)
			return
		}
	}
}

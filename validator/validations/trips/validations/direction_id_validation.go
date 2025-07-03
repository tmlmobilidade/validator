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
  - Field: direction_id
  - Presence: Optional
  - Type: Enum

# Description

Indicates the direction of travel for a trip.
This field should not be used in routing; it provides a way to separate trips by direction when publishing time tables.

Valid options are:

  - 0 - Travel in one direction (e.g. outbound travel).
  - 1 - Travel in the opposite direction (e.g. inbound travel).

# Example

The `trip_headsign` and `direction_id` fields may be used together to assign a name to travel in each direction for a set of trips. A `trips.txt` file could contain these records for use in time tables:

	trip_id,...,trip_headsign,direction_id
	1234,...,Airport,0
	1505,...,Downtown,1

[trips.txt]: https://gtfs.org/schedule/reference/#tripstxt
*/
func DirectionIdValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.DirectionId.Severity != "" {
		s = rules.DirectionId.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "direction_id",
			FileName:     "trips.txt",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
			ValidationID: "direction_id_validation",
		})
	}

	// 1. Validate direction_id is 0 or 1 if it exists
	if trip.DirectionId != nil {
		validDirectionIds := map[int]bool{0: true, 1: true}
		if !validDirectionIds[*trip.DirectionId] {
			addMessage("Invalid direction_id value. Valid values are 0 and 1.", s)
			return
		}
	}

	// 2. Validate direction_id is required
	if s == types.SEVERITY_IGNORE {
		return
	}

	if trip.DirectionId == nil {
		addMessage(lib.IfThenElse(s == types.SEVERITY_ERROR, "direction_id is required", "direction_id is recommended"), s)
		return
	}

	// Validate Rule Options
	if rules != nil && rules.DirectionId.Options != nil {
		if slices.Contains(*rules.DirectionId.Options, types.ALL_OPTIONS) {
			return
		}

		if slices.Contains(*rules.DirectionId.Options, fmt.Sprintf("%d", *trip.DirectionId)) {
			return
		}

		addMessage(fmt.Sprintf("direction_id is not allowed: %d", *trip.DirectionId), s)
	}
}

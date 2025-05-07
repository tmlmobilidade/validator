package trips

import (
	"main/lib"
	"main/services"
	"main/types"
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
func DirectionIdValidation(severity *types.Severity, trip *types.Trip, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	// 1. Validate direction_id is 0 or 1 if it exists
	if trip.DirectionId != nil {
		validDirectionIds := map[int]bool{0: true, 1: true}
		if !validDirectionIds[*trip.DirectionId] {
			message := types.Message{
				Field: "direction_id",
				FileName: "trips.txt",
				Message: "Invalid direction_id value. Valid values are 0 and 1.",
				Rows: []int{row},
				Severity: s,
				ValidationID: "direction_id_validation",
			}
			services.AppMessageService.AddMessage(message)
			return
		}
	}
	
	// 2. Validate direction_id is required
	if s == types.SEVERITY_IGNORE {
		return;
	}
	
	if trip.DirectionId == nil {
		message := types.Message{
			Field: "direction_id",
			FileName: "trips.txt",
			Message: lib.IfThenElse(s == types.SEVERITY_ERROR, "Direction ID is required", "Direction ID is recommended"),
			Rows: []int{row},
			Severity: s,
			ValidationID: "direction_id_validation",
		}
		services.AppMessageService.AddMessage(message)
		return;
	}
}
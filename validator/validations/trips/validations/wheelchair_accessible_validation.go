package trips

import (
	"main/lib"
	"main/services"
	"main/types"
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
func WheelchairAccessibleValidation(severity *types.Severity, trip *types.Trip, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	// 1. Validate wheelchair_accessible is 0 or 1 if it exists
	if trip.WheelchairAccessible != nil {
		validWheelchairAccessible := map[int]bool{0: true, 1: true, 2: true}
		if !validWheelchairAccessible[*trip.WheelchairAccessible] {
			message := types.Message{
				Field: "wheelchair_accessible",
				FileName: "trips.txt",
				Message: "Invalid wheelchair_accessible value. Valid values are 0, 1, and 2.",
				Rows: []int{row},
				Severity: s,
				ValidationID: "wheelchair_accessible_validation",
			}
			services.AppMessageService.AddMessage(message)
			return
		}
	}
	
	// 2. Validate wheelchair_accessible is required
	if s == types.SEVERITY_IGNORE {
		return;
	}
	
	if trip.WheelchairAccessible == nil {
		message := types.Message{
			Field: "wheelchair_accessible",
			FileName: "trips.txt",
			Message: lib.IfThenElse(s == types.SEVERITY_ERROR, "wheelchair_accessible is required", "wheelchair_accessible is recommended"),
			Rows: []int{row},
			Severity: s,
			ValidationID: "wheelchair_accessible_validation",
		}
		services.AppMessageService.AddMessage(message)
		return;
	}
}
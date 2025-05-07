package trips

import (
	"main/lib"
	"main/services"
	"main/types"
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
func BikesAllowedValidation(severity *types.Severity, trip *types.Trip, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	// 1. Validate bikes_allowed is 0 or 1 if it exists
	if trip.BikesAllowed != nil {
		validBikesAllowed := map[int]bool{0: true, 1: true, 2: true}
		if !validBikesAllowed[*trip.BikesAllowed] {
			message := types.Message{
				Field: "bikes_allowed",
				FileName: "trips.txt",
				Message: "Invalid bikes_allowed value. Valid values are 0, 1, and 2.",
				Rows: []int{row},
				Severity: s,
				ValidationID: "bikes_allowed_validation",
			}
			services.AppMessageService.AddMessage(message)
			return
		}
	}
	
	// 2. Validate bikes_allowed is required
	if s == types.SEVERITY_IGNORE {
		return;
	}
	
	if trip.BikesAllowed == nil {
		message := types.Message{
			Field: "bikes_allowed",
			FileName: "trips.txt",
			Message: lib.IfThenElse(s == types.SEVERITY_ERROR, "bikes_allowed is required", "bikes_allowed is recommended"),
			Rows: []int{row},
			Severity: s,
			ValidationID: "bikes_allowed_validation",
		}
		services.AppMessageService.AddMessage(message)
		return;
	}
}
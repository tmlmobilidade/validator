package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

 - File: [stop_times.txt]
 - Field: drop_off_type
 - Presence: Conditionally Forbidden
 - Type: Enum

# Description

Indicates drop off method.

Valid options are:

  - 0 or empty - Regularly scheduled drop off.
  - 1 - No drop off available.
  - 2 - Must phone agency to arrange drop off.
  - 3 - Must coordinate with driver to arrange drop off.

Conditionally Forbidden:

 - drop_off_type=0 forbidden if start_pickup_drop_off_window or end_pickup_drop_off_window are defined.
 - Optional otherwise.

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func DropOffTypeValidation(severity *types.Severity, stopTime *types.StopTime, row int) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "drop_off_type",
			FileName:     "stop_times.txt",
			ValidationID: "drop_off_type_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	if stopTime.DropOffType == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}
		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "drop_off_type is recommended", "drop_off_type is required")
		addMessage(warn, s)
		return
	}

	// Validate values
	dt := *stopTime.DropOffType
	if dt < 0 || dt > 3 {
		addMessage("drop_off_type must be 0, 1, 2, or 3.", types.SEVERITY_ERROR)
		return
	}

	// drop_off_type=0 forbidden if start_pickup_drop_off_window or end_pickup_drop_off_window are defined
	if dt == 0 && ((stopTime.StartPickupDropOffWindow != nil && *stopTime.StartPickupDropOffWindow != "") || (stopTime.EndPickupDropOffWindow != nil && *stopTime.EndPickupDropOffWindow != "")) {
		addMessage("drop_off_type=0 is forbidden if start_pickup_drop_off_window or end_pickup_drop_off_window are defined.", types.SEVERITY_ERROR)
		return
	}
} 
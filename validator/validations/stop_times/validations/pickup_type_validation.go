package stop_times

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

  - File: [stop_times.txt]
  - Field: pickup_type
  - Presence: Conditionally Required
  - Type: Enum

# Description

Indicates pickup method.

Valid options are:

  - 0 or empty - Regularly scheduled pickup.
  - 1 - No pickup available.
  - 2 - Must phone agency to arrange pickup.
  - 3 - Must coordinate with driver to arrange pickup.

Conditionally Forbidden:

  - pickup_type=0 is forbidden if start_pickup_drop_off_window or end_pickup_drop_off_window are defined.
  - pickup_type=3 is forbidden if start_pickup_drop_off_window or end_pickup_drop_off_window are defined.
  - pickup_type is optional otherwise.

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func PickupTypeValidation(stopTime *types.StopTime, row int, rules *types.StopTimesRules) {

	s := types.SEVERITY_IGNORE
	if rules != nil && rules.PickupType.Severity != "" {
		s = rules.PickupType.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "pickup_type",
			FileName:     "stop_times.txt",
			ValidationID: "pickup_type_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	// Validate presence
	if stopTime.PickupType == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "pickup_type is recommended", "pickup_type is required")
		addMessage(warn, s)
		return
	}

	// Validate values
	pt := *stopTime.PickupType
	if pt < 0 || pt > 3 {
		addMessage("pickup_type must be 0, 1, 2, or 3.", types.SEVERITY_ERROR)
		return
	}

	// pickup_type=0 or 3 forbidden if start_pickup_drop_off_window or end_pickup_drop_off_window are defined
	if (pt == 0 || pt == 3) && ((stopTime.StartPickupDropOffWindow != nil && *stopTime.StartPickupDropOffWindow != "") || (stopTime.EndPickupDropOffWindow != nil && *stopTime.EndPickupDropOffWindow != "")) {
		addMessage("pickup_type 0 or 3 is forbidden if start_pickup_drop_off_window or end_pickup_drop_off_window are defined.", types.SEVERITY_ERROR)
		return
	}

	// Validate Rule Options
	if rules != nil && rules.PickupType.Options != nil {
		if slices.Contains(*rules.PickupType.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.PickupType.Options, fmt.Sprintf("%d", pt)) {
			addMessage(fmt.Sprintf("pickup_type is not allowed: %d", pt), s)
			return
		}
	}
}

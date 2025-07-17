package stop_times

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
func DropOffTypeValidation(stopTime *types.StopTime, row int, rules *types.StopTimesRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.DropOffType.Severity != "" {
		s = rules.DropOffType.Severity
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
		if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
			return
		}
		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, i18n.AppTranslator.Get("drop_off_type_validation.recommended"), i18n.AppTranslator.Get("drop_off_type_validation.required"))
		addMessage(warn, s)
		return
	}

	// Validate values
	dt := *stopTime.DropOffType
	if dt < 0 || dt > 3 {
		addMessage(i18n.AppTranslator.Get("drop_off_type_validation.invalid"), types.SEVERITY_ERROR)
		return
	}

	// drop_off_type=0 forbidden if start_pickup_drop_off_window or end_pickup_drop_off_window are defined
	if dt == 0 && ((stopTime.StartPickupDropOffWindow != nil && *stopTime.StartPickupDropOffWindow != "") || (stopTime.EndPickupDropOffWindow != nil && *stopTime.EndPickupDropOffWindow != "")) {
		addMessage(i18n.AppTranslator.Get("drop_off_type_validation.forbidden_zero_with_window"), types.SEVERITY_ERROR)
		return
	}

	// Validate Rule Options
	if rules != nil && rules.DropOffType.Options != nil {
		if slices.Contains(*rules.DropOffType.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.DropOffType.Options, fmt.Sprintf("%d", dt)) {
			addMessage(i18n.AppTranslator.Get("drop_off_type_validation.not_allowed", fmt.Sprintf("%d", dt)), s)
			return
		}
	}
}

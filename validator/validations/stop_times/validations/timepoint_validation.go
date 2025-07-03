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
  - Field: timepoint
  - Presence: Optional
  - Type: Enum

# Description

Indicates if arrival and departure times for a stop are strictly adhered to by the vehicle or if they are instead approximate and/or interpolated times. This field allows a GTFS producer to provide interpolated stop-times, while indicating that the times are approximate.

Valid options are:

  - 0 - Times are considered approximate.
  - 1 - Times are considered exact.

All records of [stop_times.txt] with defined arrival or departure times should have timepoint values populated. If no timepoint values are provided, all times are considered exact.

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func TimepointValidation(stopTime *types.StopTime, row int, rules *types.StopTimesRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.Timepoint.Severity != "" {
		s = rules.Timepoint.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "timepoint",
			FileName:     "stop_times.txt",
			ValidationID: "timepoint_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	if stopTime.Timepoint == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}
		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "timepoint is recommended", "timepoint is required")
		addMessage(warn, s)
		return
	}

	tp := *stopTime.Timepoint
	if tp != 0 && tp != 1 {
		addMessage("timepoint must be 0 or 1.", types.SEVERITY_ERROR)
		return
	}

	// Validate Rule Options
	if rules != nil && rules.Timepoint.Options != nil {
		if slices.Contains(*rules.Timepoint.Options, types.ALL_OPTIONS) {
			return
		}

		if slices.Contains(*rules.Timepoint.Options, fmt.Sprintf("%d", tp)) {
			return
		}

		addMessage(fmt.Sprintf("timepoint is not allowed: %d", tp), s)
		return
	}
}

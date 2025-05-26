package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

 - File: [stop_times.txt]
 - Field: stop_headsign
 - Presence: Optional
 - Type: String

# Description

Text that appears on signage identifying the trip's destination to riders.

This field overrides the default `trips.trip_headsign` when the headsign changes between stops.

If the headsign is displayed for an entire trip, `trips.trip_headsign` should be used instead.

A `stop_headsign` value specified for one `stop_time` does not apply to subsequent `stop_times` in the same trip.

If you want to override the `trip_headsign` for multiple `stop_times` in the same trip, the `stop_headsign` value must be repeated in each `stop_time` row.

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func StopHeadsignValidation(severity *types.Severity, stopTime *types.StopTime, row int) {

	s := types.SEVERITY_IGNORE
	if *severity != types.SEVERITY_IGNORE {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_headsign",
			FileName:     "stop_times.txt",
			ValidationID: "stop_headsign_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	if stopTime.StopHeadsign == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "stop_headsign is recommended", "stop_headsign is required")
		addMessage(warn, s)
		return
	}
}
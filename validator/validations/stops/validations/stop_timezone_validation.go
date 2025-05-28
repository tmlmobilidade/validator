/*
# Attributes

 - File: [stops.txt]
 - Field: stop_timezone
 - Presence: Optional
 - Type: Timezone

# Description

Timezone of the location.

If the location has a parent station, it inherits the parent station's timezone instead of applying its own.

Stations and parentless stops with empty stop_timezone inherit the timezone specified by `agency.agency_timezone`.

The times provided in stop_times.txt are in the timezone specified by `agency.agency_timezone`, not stop_timezone. This ensures that the time values in a trip always increase over the course of a trip, regardless of which timezones the trip crosses.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/

package stops

import (
	"main/lib"
	"main/services"
	"main/types"
)

// StopTimezoneValidation validates the stop_timezone field in stops.txt
func StopTimezoneValidation(severity *types.Severity, stop *types.Stop, row int) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_timezone",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "stop_timezone_validation",
		})
	}

	if stop.StopTimezone == nil || *stop.StopTimezone == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}
		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "stop_timezone is required", "stop_timezone is recommended")
		addMessage(warn, s)
		return
	}

	err := lib.ValidateTimezone(*stop.StopTimezone)
	if err != "" {
		addMessage(err, types.SEVERITY_ERROR)
		return
	}
} 
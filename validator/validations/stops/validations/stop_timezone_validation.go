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
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

// StopTimezoneValidation validates the stop_timezone field in stops.txt
func StopTimezoneValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.StopTimezone.Severity != "" {
		s = rules.StopTimezone.Severity
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
		if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
			return
		}
		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"stop_timezone_validation.required",
				"stop_timezone_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("stop_timezone_validation.forbidden"), s)
		return
	}

	if !lib.ValidateTimezone(*stop.StopTimezone) {
		addMessage(i18n.AppTranslator.Get("stop_timezone_validation.invalid", *stop.StopTimezone), types.SEVERITY_ERROR)
		return
	}

	// Validate rules
	if rules != nil && rules.StopTimezone.Options != nil {
		if slices.Contains(*rules.StopTimezone.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.StopTimezone.Options, *stop.StopTimezone) {
			addMessage(i18n.AppTranslator.Get("stop_timezone_validation.not_allowed", *stop.StopTimezone), s)
			return
		}
	}
}

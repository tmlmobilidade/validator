package stop_times

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [stop_times.txt]
  - Field: end_pickup_drop_off_window
  - Presence: Conditionally Required
  - Type: Time

# Description

Time that on-demand service ends in a GeoJSON location, location group, or stop.

Conditionally Required:
  - Required if stop_times.location_group_id or stop_times.location_id is defined.
  - Required if start_pickup_drop_off_window is defined.
  - Forbidden if arrival_time or departure_time is defined.
  - Optional otherwise.

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func EndPickupDropOffWindowValidation(stopTime *types.StopTime, row int, rules *types.StopTimesRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.EndPickupDropOffWindow.Severity != "" {
		s = rules.EndPickupDropOffWindow.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "end_pickup_drop_off_window",
			FileName:     "stop_times.txt",
			ValidationID: "end_pickup_drop_off_window_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	// Forbidden if arrival_time or departure_time are defined
	if (stopTime.ArrivalTime != nil && *stopTime.ArrivalTime != "") || (stopTime.DepartureTime != nil && *stopTime.DepartureTime != "") {
		if stopTime.EndPickupDropOffWindow != nil && *stopTime.EndPickupDropOffWindow != "" {
			addMessage(i18n.AppTranslator.Get("end_pickup_drop_off_window_validation.forbidden_with_time"), types.SEVERITY_ERROR)
		}
		return
	}

	required := false
	// Required if location_group_id or location_id is defined
	if (stopTime.LocationGroupId != nil && *stopTime.LocationGroupId != "") || (stopTime.LocationId != nil && *stopTime.LocationId != "") {
		required = true
	}
	// Required if start_pickup_drop_off_window is defined
	if stopTime.StartPickupDropOffWindow != nil && *stopTime.StartPickupDropOffWindow != "" {
		required = true
	}

	if required {
		if stopTime.EndPickupDropOffWindow == nil || *stopTime.EndPickupDropOffWindow == "" {
			addMessage(i18n.AppTranslator.Get("end_pickup_drop_off_window_validation.required_conditional"), types.SEVERITY_ERROR)
			return
		}
	}

	// Validate time format if present
	if stopTime.EndPickupDropOffWindow != nil && *stopTime.EndPickupDropOffWindow != "" {
		if !lib.ValidateTime(*stopTime.EndPickupDropOffWindow) {
			addMessage(i18n.AppTranslator.Get("end_pickup_drop_off_window_validation.invalid_time"), types.SEVERITY_ERROR)
			return
		}
	}

	// Optional
	if stopTime.EndPickupDropOffWindow == nil && s != types.SEVERITY_IGNORE {
		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, i18n.AppTranslator.Get("end_pickup_drop_off_window_validation.required"), i18n.AppTranslator.Get("end_pickup_drop_off_window_validation.recommended"))
		addMessage(warn, s)
		return
	}
}

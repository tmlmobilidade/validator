package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [stop_times.txt]
  - Field: departure_time
  - Presence: Conditionally Required
  - Type: Time

# Description

Departure time from the stop (defined by `stop_times.stop_id`) for a specific trip (defined by `stop_times.trip_id`) in the time zone specified by `agency.agency_timezone`, not `stops.stop_timezone`.

If there are not separate times for arrival and departure at a stop, `arrival_time` and `departure_time` should be the same.

For times occurring after midnight on the service day, enter the time as a value greater than 24:00:00 in HH:MM:SS.

If exact arrival and departure times (timepoint=1) are not available, estimated or interpolated arrival and departure times (timepoint=0) should be provided.

Conditionally Required:

  - Required for the first and last stop in a trip (defined by `stop_times.stop_sequence`).
  - Required for `timepoint=1`.
  - Forbidden when `start_pickup_drop_off_window` or `end_pickup_drop_off_window` are defined.
  - Optional otherwise.

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func DepartureTimeValidation(stopTime *types.StopTime, row int, gtfs *types.Gtfs, rules *types.StopTimesRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.DepartureTime.Severity != "" {
		s = rules.DepartureTime.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "departure_time",
			FileName:     "stop_times.txt",
			ValidationID: "departure_time_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	// Required for timepoint=1
	if stopTime.Timepoint != nil && *stopTime.Timepoint == 1 && stopTime.DepartureTime == nil {
		addMessage("departure_time is required for timepoint=1.", types.SEVERITY_ERROR)
		return
	}

	// Forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined
	if (stopTime.StartPickupDropOffWindow != nil || stopTime.EndPickupDropOffWindow != nil) && stopTime.DepartureTime != nil {
		addMessage("departure_time is forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined.", types.SEVERITY_ERROR)
		return
	}

	// Validate time
	if stopTime.DepartureTime != nil {
		if err := lib.ValidateTime(*stopTime.DepartureTime); err != "" {
			addMessage(err, types.SEVERITY_ERROR)
			return
		}
	}

	// Optional
	if stopTime.DepartureTime == nil && s != types.SEVERITY_IGNORE {
		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "departure_time is required.", "departure_time is recommended.")
		addMessage(warn, s)
		return
	}
}

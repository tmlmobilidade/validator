package stop_times

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"math"
	"strconv"
)

/*
# Attributes

  - File: [stop_times.txt]
  - Field: arrival_time
  - Presence: Conditionally Required
  - Type: Time

# Description

Arrival time at the stop (defined by `stop_times.stop_id`) for a specific trip (defined by `stop_times.trip_id`) in the time zone specified by `agency.agency_timezone`, not `stops.stop_timezone`.

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
func ArrivalTimeValidation(stopTime *types.StopTime, row int, gtfs *types.Gtfs, rules *types.StopTimesRules, tripStopSequences map[string]types.TripStopSequence) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.ArrivalTime.Severity != "" {
		s = rules.ArrivalTime.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "arrival_time",
			FileName:     "stop_times.txt",
			ValidationID: "arrival_time_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	// Required for timepoint=1
	if stopTime.Timepoint != nil && *stopTime.Timepoint == 1 && stopTime.ArrivalTime == nil {
		addMessage(i18n.AppTranslator.Get("arrival_time_validation.required_timepoint"), types.SEVERITY_ERROR)
		return
	}

	// Forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined
	if (stopTime.StartPickupDropOffWindow != nil || stopTime.EndPickupDropOffWindow != nil) && stopTime.ArrivalTime != nil {
		addMessage(i18n.AppTranslator.Get("arrival_time_validation.forbidden_with_window"), types.SEVERITY_ERROR)
		return
	}

	// Required for the first and last stop in a trip
	if stopTime.StopSequence != nil && stopTime.TripId != nil {
		// Use cached trip stop sequences instead of querying database
		seq, exists := tripStopSequences[*stopTime.TripId]
		if !exists {
			// Fallback to database query if not in cache (shouldn't happen)
			stopTimes, err := gtfs.GetRowsById("stop_times", *stopTime.TripId)
			if err != nil || len(stopTimes) == 0 {
				addMessage(i18n.AppTranslator.Get("arrival_time_validation.invalid_trip_id"), types.SEVERITY_ERROR)
				return
			}

			minStopSequence := math.MaxInt
			maxStopSequence := math.MinInt

			for _, row := range stopTimes {
				stopTimeRaw, err := gtfs.GetStopTime(row)
				if err != nil {
					continue
				}
				stopSequence, err := strconv.Atoi(stopTimeRaw.StopSequence)
				if err != nil {
					addMessage(i18n.AppTranslator.Get("arrival_time_validation.invalid_stop_sequence"), types.SEVERITY_ERROR)
					return
				}

				if stopSequence < minStopSequence {
					minStopSequence = stopSequence
				}
				if stopSequence > maxStopSequence {
					maxStopSequence = stopSequence
				}
			}

			if *stopTime.StopSequence == minStopSequence && stopTime.ArrivalTime == nil {
				addMessage(i18n.AppTranslator.Get("arrival_time_validation.required_first_stop"), types.SEVERITY_ERROR)
				return
			}

			if *stopTime.StopSequence == maxStopSequence && stopTime.ArrivalTime == nil {
				addMessage(i18n.AppTranslator.Get("arrival_time_validation.required_last_stop"), types.SEVERITY_ERROR)
				return
			}
		} else {
			// Use cached values - much faster!
			stopSequence := *stopTime.StopSequence

			if stopSequence == seq.Min && stopTime.ArrivalTime == nil {
				addMessage(i18n.AppTranslator.Get("arrival_time_validation.required_first_stop"), types.SEVERITY_ERROR)
				return
			}

			if stopSequence == seq.Max && stopTime.ArrivalTime == nil {
				addMessage(i18n.AppTranslator.Get("arrival_time_validation.required_last_stop"), types.SEVERITY_ERROR)
				return
			}
		}
	}

	// Validate time
	if stopTime.ArrivalTime != nil {
		if !lib.ValidateTime(*stopTime.ArrivalTime) {
			addMessage(i18n.AppTranslator.Get("arrival_time_validation.invalid_time"), types.SEVERITY_ERROR)
			return
		}
	}

	// Optional
	if stopTime.ArrivalTime == nil && s != types.SEVERITY_IGNORE {
		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, i18n.AppTranslator.Get("arrival_time_validation.required"), i18n.AppTranslator.Get("arrival_time_validation.recommended"))
		addMessage(warn, s)
		return
	}
}

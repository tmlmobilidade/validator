package stop_times

import (
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
func ArrivalTimeValidation(severity *types.Severity, stopTime *types.StopTime, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
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
		addMessage("arrival_time is required for timepoint=1.", types.SEVERITY_ERROR)
		return
	}

	// Forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined
	if (stopTime.StartPickupDropOffWindow != nil || stopTime.EndPickupDropOffWindow != nil) && stopTime.ArrivalTime != nil {
		addMessage("arrival_time is forbidden when start_pickup_drop_off_window or end_pickup_drop_off_window are defined.", types.SEVERITY_ERROR)
		return
	}

	// Required for the first and last stop in a trip
	if stopTime.StopSequence != nil {
		stopTimes, exists := gtfs.IdMap["stop_times"][*stopTime.TripId]
		if !exists {
			addMessage("trip_id must reference a valid trip_id from trips.txt.", types.SEVERITY_ERROR)
			return
		}

		minStopSequence := math.MaxInt
		maxStopSequence := math.MinInt

		for _, row := range stopTimes {
			stopSequence, err := strconv.Atoi(gtfs.StopTime[row].StopSequence)
			if err != nil {
				addMessage("stop_sequence must be an integer.", types.SEVERITY_ERROR)
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
			addMessage("arrival_time is required for the first stop in a trip.", types.SEVERITY_ERROR)
			return
		}

		if *stopTime.StopSequence == maxStopSequence && stopTime.ArrivalTime == nil {
			addMessage("arrival_time is required for the last stop in a trip.", types.SEVERITY_ERROR)
			return
		}
	}

	// Validate time
	if stopTime.ArrivalTime != nil {
		if err := lib.ValidateTime(*stopTime.ArrivalTime); err != "" {
			addMessage(err, types.SEVERITY_ERROR)
			return
		}
	}
	
	// Optional
	if stopTime.ArrivalTime == nil && s != types.SEVERITY_IGNORE {
		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "arrival_time is required.", "arrival_time is recommended.")
		addMessage(warn, s)
		return
	}
}
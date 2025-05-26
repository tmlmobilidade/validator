package stop_times

import (
	"main/services"
	"main/types"
	"strconv"
)

/*
# Attributes

 - File: [stop_times.txt]
 - Field: stop_sequence
 - Presence: Required
 - Type: Non-negative Integer

# Description

Order of stops, location groups, or GeoJSON locations for a particular trip. The values must increase along the trip but do not need to be consecutive.

# Example

The first location on the trip could have a `stop_sequence=1`, the second location on the trip could have a `stop_sequence=23`, the third location could have a `stop_sequence=40`, and so on.

Travel within the same location group or GeoJSON location requires two records in `stop_times.txt` with the same `location_group_id` or `location_id`.

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func StopSequenceValidation(stopTime *types.StopTime, row int, gtfs *types.Gtfs) {
	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_sequence",
			FileName:     "stop_times.txt",
			ValidationID: "stop_sequence_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	// Required
	if stopTime.StopSequence == nil {
		addMessage("stop_sequence is required.", types.SEVERITY_ERROR)
		return
	}

	// Non-negative integer
	if *stopTime.StopSequence < 0 {
		addMessage("stop_sequence must be a non-negative integer.", types.SEVERITY_ERROR)
		return
	}

	// Check each trip's stop times for pickup/dropoff windows
	sequence := make([]int, 0)
	if rows, exists := gtfs.IdMap["stop_times"][*stopTime.TripId]; exists {
		for _, row := range rows {
			seqStr := gtfs.Files["stop_times"][row]["stop_sequence"]
			seq, err := strconv.Atoi(seqStr)
			if err != nil {
				addMessage("stop_sequence must be a non-negative integer.", types.SEVERITY_ERROR)
				return
			}
			sequence = append(sequence, seq)
		}
	}

	// Check if the sequence is increasing
	for i := 1; i < len(sequence); i++ {
		if sequence[i] <= sequence[i-1] {
			addMessage("stop_sequence values must increase along the trip.", types.SEVERITY_ERROR)
			return
		}
	}
} 
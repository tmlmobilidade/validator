package trips

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
func StopSequenceValidation(trip *types.Trip, row int, gtfs *types.Gtfs) {
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

	// Check each trip's stop times for pickup/dropoff windows
	stopSequences := make([]int, 0)
	shapeDistTraveledSequences := make([]float64, 0)

	stopTimes := gtfs.IdMap["stop_times"][trip.TripId]
	for _, row := range stopTimes {
		
		stopSequence, err := strconv.Atoi(gtfs.Files["stop_times"][row]["stop_sequence"])
		if err != nil {
			addMessage("stop_sequence must be a non-negative integer.", types.SEVERITY_ERROR)
			return
		}

		shapeDistTraveled, err := strconv.ParseFloat(gtfs.Files["stop_times"][row]["shape_dist_traveled"], 64)
		if err != nil {
			addMessage("shape_dist_traveled must be a float.", types.SEVERITY_ERROR)
			return
		}

		stopSequences = append(stopSequences, stopSequence)
		shapeDistTraveledSequences = append(shapeDistTraveledSequences, shapeDistTraveled)
	}

	// Bubble Sort based on stopSequences, but also reorders shapeDistTraveledSequences to match
	for i := range stopSequences {
		for j := range stopSequences[:len(stopSequences)-i-1] {
			if stopSequences[j] > stopSequences[j+1] {
				stopSequences[j], stopSequences[j+1] = stopSequences[j+1], stopSequences[j]
				shapeDistTraveledSequences[j], shapeDistTraveledSequences[j+1] = shapeDistTraveledSequences[j+1], shapeDistTraveledSequences[j]
			}
		}
	}

	// Check if the stopSequences and shapeDistTraveledSequences are increasing
	for i := 1; i < len(stopSequences); i++ {
		if stopSequences[i] <= stopSequences[i-1] {
			addMessage("stop_sequence values must increase along the trip", types.SEVERITY_ERROR)
			return
		}

		if shapeDistTraveledSequences[i] <= shapeDistTraveledSequences[i-1] {
			addMessage("shape_dist_traveled values must increase along the trip", types.SEVERITY_ERROR)
			return
		}
	}
	
} 
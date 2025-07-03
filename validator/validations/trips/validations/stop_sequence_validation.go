package trips

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	"sort"
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
func StopSequenceValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules) (stopSequenceHash string) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.StopSequence.Severity != "" {
		s = rules.StopSequence.Severity
	}

	if s == types.SEVERITY_IGNORE {
		return
	}

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

	if trip.TripId == nil {
		return
	}

	// Check each trip's stop times for pickup/dropoff windows

	stopSequences := make([]types.StopTime, 0)
	hash := ""

	stopTimes := gtfs.IdMap["stop_times"][*trip.TripId]
	for _, row := range stopTimes {

		stopSequence, err := strconv.Atoi(gtfs.StopTime[row].StopSequence)
		if err != nil {
			addMessage("stop_sequence must be a non-negative integer.", types.SEVERITY_ERROR)
			return
		}

		shapeDistTraveled := -1.0
		if gtfs.StopTime[row].ShapeDistTraveled != "" {
			shapeDistTraveled, err = strconv.ParseFloat(gtfs.StopTime[row].ShapeDistTraveled, 64)
			if err != nil {
				addMessage("shape_dist_traveled must be a float.", types.SEVERITY_ERROR)
				return
			}
		}

		stopId := gtfs.StopTime[row].StopId

		stopSequences = append(stopSequences, types.StopTime{
			StopSequence:      &stopSequence,
			ShapeDistTraveled: &shapeDistTraveled,
			StopId:            &stopId,
		})

		hash += fmt.Sprintf("%s-%v-%v", stopId, shapeDistTraveled, stopSequence)
	}

	stopSequences = lib.RemoveDuplicates(stopSequences)
	sort.Slice(stopSequences, func(i, j int) bool {
		return *stopSequences[i].StopSequence < *stopSequences[j].StopSequence
	})

	for i, stopSequence := range stopSequences {
		if i > 0 {
			if *stopSequence.StopSequence <= *stopSequences[i-1].StopSequence {
				addMessage("stop_sequence values must increase along the trip ('"+*trip.TripId+"')", types.SEVERITY_ERROR)
				return
			}

			if *stopSequence.ShapeDistTraveled >= 0 && *stopSequences[i-1].ShapeDistTraveled >= 0 {
				if *stopSequence.ShapeDistTraveled < *stopSequences[i-1].ShapeDistTraveled {
					addMessage("shape_dist_traveled values must increase along the trip ('"+*trip.TripId+"')", types.SEVERITY_ERROR)
					return
				}
			}
		}
	}

	stopSequenceHash = lib.Hash(hash)
	return
}

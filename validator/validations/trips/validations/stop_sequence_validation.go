package trips

import (
	"fmt"
	"main/i18n"
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
func StopSequenceValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules, tripStopTimesCache map[string][]types.StopTimeRaw) (stopSequenceHash string) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.StopSequence.Severity != "" {
		s = rules.StopSequence.Severity
	}

	if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
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
	// Use cached stop_times data instead of querying database

	stopSequences := make([]types.StopTime, 0)
	hash := ""

	// Use cached stop_times data if available
	stopTimesRaw, exists := tripStopTimesCache[*trip.TripId]
	if !exists {
		// Fallback to database query if not in cache (shouldn't happen)
		stopTimes, err := gtfs.GetRowsById("stop_times", *trip.TripId)
		if err != nil {
			return
		}
		for _, row := range stopTimes {
			stopTimeRaw, err := gtfs.GetStopTime(row)
			if err != nil {
				continue
			}
			stopTimesRaw = append(stopTimesRaw, stopTimeRaw)
		}
	}

	// Process cached stop_times data
	for _, stopTimeRaw := range stopTimesRaw {
		stopSequence, err := strconv.Atoi(stopTimeRaw.StopSequence)
		if err != nil {
			addMessage(i18n.AppTranslator.Get("stop_sequence_validation.invalid_sequence"), types.SEVERITY_ERROR)
			return
		}

		shapeDistTraveled := -1.0
		if stopTimeRaw.ShapeDistTraveled != "" {
			shapeDistTraveled, err = strconv.ParseFloat(stopTimeRaw.ShapeDistTraveled, 64)
			if err != nil {
				addMessage(i18n.AppTranslator.Get("stop_sequence_validation.invalid_shape_dist"), types.SEVERITY_ERROR)
				return
			}
		}

		stopId := stopTimeRaw.StopId

		stopSequences = append(stopSequences, types.StopTime{
			StopSequence:      &stopSequence,
			ShapeDistTraveled: &shapeDistTraveled,
			StopId:            &stopId,
		})
	}

	stopSequences = lib.RemoveDuplicates(stopSequences)
	sort.Slice(stopSequences, func(i, j int) bool {
		return *stopSequences[i].StopSequence < *stopSequences[j].StopSequence
	})

	for _, stopSequence := range stopSequences {
		hash += fmt.Sprintf("%s-%v-%v", *stopSequence.StopId, *stopSequence.ShapeDistTraveled, *stopSequence.StopSequence)
	}

	for i, stopSequence := range stopSequences {
		if i > 0 {
			if *stopSequence.StopSequence <= *stopSequences[i-1].StopSequence {
				addMessage(i18n.AppTranslator.Get("stop_sequence_validation.non_increasing_sequence", *trip.TripId), types.SEVERITY_ERROR)
				return
			}

			if *stopSequence.ShapeDistTraveled >= 0 && *stopSequences[i-1].ShapeDistTraveled >= 0 {
				if *stopSequence.ShapeDistTraveled < *stopSequences[i-1].ShapeDistTraveled {
					addMessage(i18n.AppTranslator.Get("stop_sequence_validation.non_increasing_shape_dist", *trip.TripId), types.SEVERITY_ERROR)
					return
				}
			}
		}
	}

	stopSequenceHash = lib.Hash(hash)
	return
}

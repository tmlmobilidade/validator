package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	"sort"
	"strconv"
	"strings"
)

type TimeSequenceStop struct {
	Row           int
	StopSequence  int
	ArrivalTime   *string
	DepartureTime *string
}

/*
# Attributes

  - File: [stop_times.txt]
  - Fields: arrival_time, departure_time, stop_sequence
  - Presence: Conditionally Required
  - Type: Time

# Description

Checks that arrival_time and departure_time do not go backwards between consecutive
stops of the same trip when ordered by stop_sequence.

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func ArrivalDepartureTimeSequenceValidation(stopTimesByTrip map[string][]TimeSequenceStop, rules *types.StopTimesRules) {
	for tripId, stopTimes := range stopTimesByTrip {
		sort.Slice(stopTimes, func(i, j int) bool {
			if stopTimes[i].StopSequence == stopTimes[j].StopSequence {
				return stopTimes[i].Row < stopTimes[j].Row
			}
			return stopTimes[i].StopSequence < stopTimes[j].StopSequence
		})

		for i := 1; i < len(stopTimes); i++ {
			validateStopTimePair(tripId, stopTimes[i-1], stopTimes[i], rules)
		}
	}
}

func validateStopTimePair(tripId string, previous TimeSequenceStop, current TimeSequenceStop, rules *types.StopTimesRules) {
	ctx := lib.NewValidationContext("arrival_time", "stop_times.txt", "arrival_departure_time_non_decreasing_by_stop_sequence", current.Row, services.AppMessageService)
	if rules != nil && rules.ArrivalDepartureSequence.Severity != "" {
		ctx.WithSeverity(rules.ArrivalDepartureSequence.Severity)
	}

	if ctx.ShouldSkip() {
		return
	}

	previousTime, previousTimeLabel, ok := lastStopTime(previous)
	if !ok {
		return
	}

	currentTime, currentTimeLabel, ok := firstStopTime(current)
	if !ok {
		return
	}

	if currentTime < previousTime {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage(
			"arrival_departure_time_sequence_validation.decreasing",
			tripId,
			previous.StopSequence,
			previousTimeLabel,
			current.StopSequence,
			currentTimeLabel,
		))
	}
}

func firstStopTime(stopTime TimeSequenceStop) (int, string, bool) {
	if stopTime.ArrivalTime != nil {
		seconds, ok := parseStopTimeSeconds(*stopTime.ArrivalTime)
		if ok {
			return seconds, *stopTime.ArrivalTime, true
		}
	}

	if stopTime.DepartureTime != nil {
		seconds, ok := parseStopTimeSeconds(*stopTime.DepartureTime)
		if ok {
			return seconds, *stopTime.DepartureTime, true
		}
	}

	return 0, "", false
}

func lastStopTime(stopTime TimeSequenceStop) (int, string, bool) {
	if stopTime.DepartureTime != nil {
		seconds, ok := parseStopTimeSeconds(*stopTime.DepartureTime)
		if ok {
			return seconds, *stopTime.DepartureTime, true
		}
	}

	if stopTime.ArrivalTime != nil {
		seconds, ok := parseStopTimeSeconds(*stopTime.ArrivalTime)
		if ok {
			return seconds, *stopTime.ArrivalTime, true
		}
	}

	return 0, "", false
}

func parseStopTimeSeconds(value string) (int, bool) {
	if !lib.ValidateTime(value) {
		return 0, false
	}

	parts := strings.Split(value, ":")
	if len(parts) != 3 {
		return 0, false
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, false
	}
	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, false
	}
	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, false
	}

	return hours*3600 + minutes*60 + seconds, true
}

package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestStopSequenceValidation_Required(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{StopSequence: nil}
	gtfs := &types.Gtfs{}
	validations.StopSequenceValidation(stopTime, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing stop_sequence should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_NegativeValue(t *testing.T) {
	services.AppMessageService.Clear()
	neg := -1
	stopTime := &types.StopTime{StopSequence: &neg}
	gtfs := &types.Gtfs{}
	validations.StopSequenceValidation(stopTime, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Negative stop_sequence should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	seq := 2
	tripId := "T1"
	stopTime := &types.StopTime{StopSequence: &seq, TripId: &tripId}
	gtfs := &types.Gtfs{
		Files: map[string][]map[string]string{
			"stop_times": {
				{"stop_sequence": "1"},
				{"stop_sequence": "2"},
				{"stop_sequence": "5"},
			},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"T1": {0, 1, 2},
			},
		},
	}
	validations.StopSequenceValidation(stopTime, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid stop_sequence should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_NotIncreasing(t *testing.T) {
	services.AppMessageService.Clear()
	seq := 2
	tripId := "T1"
	stopTime := &types.StopTime{StopSequence: &seq, TripId: &tripId}
	gtfs := &types.Gtfs{
		Files: map[string][]map[string]string{
			"stop_times": {
				{"stop_sequence": "1"},
				{"stop_sequence": "2"},
				{"stop_sequence": "2"},
			},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"T1": {0, 1, 2},
			},
		},
	}
	validations.StopSequenceValidation(stopTime, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Non-increasing stop_sequence should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestStopSequenceValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{TripId: lib.Ptr("T1")}
	gtfs := &types.Gtfs{
		StopTime: []types.StopTimeRaw{
			{StopSequence: "1", ShapeDistTraveled: "0"},
			{StopSequence: "3", ShapeDistTraveled: "20"},
			{StopSequence: "2", ShapeDistTraveled: "10"},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"T1": {0, 1, 2},
			},
		},
	}
	validations.StopSequenceValidation(trip, 0, gtfs, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, make(map[string][]types.StopTimeRaw))
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid stop_sequence and shape_dist_traveled should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_NonIncreasingStopSequence(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{TripId: lib.Ptr("T2")}
	gtfs := &types.Gtfs{
		StopTime: []types.StopTimeRaw{
			{StopSequence: "1", ShapeDistTraveled: "0"},
			{StopSequence: "2", ShapeDistTraveled: "10"},
			{StopSequence: "2", ShapeDistTraveled: "20"},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"T2": {0, 1, 2},
			},
		},
	}
	validations.StopSequenceValidation(trip, 0, gtfs, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, make(map[string][]types.StopTimeRaw))
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Non-increasing stop_sequence should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_NonIncreasingShapeDistTraveled(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{TripId: lib.Ptr("T3")}
	gtfs := &types.Gtfs{
		StopTime: []types.StopTimeRaw{
			{StopSequence: "1", ShapeDistTraveled: "0"},
			{StopSequence: "2", ShapeDistTraveled: "10"},
			{StopSequence: "3", ShapeDistTraveled: "5"},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"T3": {0, 1, 2},
			},
		},
	}
	validations.StopSequenceValidation(trip, 0, gtfs, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, make(map[string][]types.StopTimeRaw))
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Non-increasing shape_dist_traveled should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_InvalidStopSequence(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{TripId: lib.Ptr("T4")}
	gtfs := &types.Gtfs{
		StopTime: []types.StopTimeRaw{
			{StopSequence: "A", ShapeDistTraveled: "0"},
			{StopSequence: "2", ShapeDistTraveled: "10"},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"T4": {0, 1},
			},
		},
	}
	validations.StopSequenceValidation(trip, 0, gtfs, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, make(map[string][]types.StopTimeRaw))
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid stop_sequence should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_InvalidShapeDistTraveled(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{TripId: lib.Ptr("T5")}
	gtfs := &types.Gtfs{

		StopTime: []types.StopTimeRaw{
			{StopSequence: "1", ShapeDistTraveled: "foo"},
			{StopSequence: "2", ShapeDistTraveled: "10"},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"T5": {0, 1},
			},
		},
	}
	validations.StopSequenceValidation(trip, 0, gtfs, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, make(map[string][]types.StopTimeRaw))
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid shape_dist_traveled should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_ValidInput_NoShapeDistTraveled(t *testing.T) {
	services.AppMessageService.Clear()

	trip := &types.Trip{TripId: lib.Ptr("T6")}
	gtfs := &types.Gtfs{
		StopTime: []types.StopTimeRaw{
			{StopSequence: "1"},
			{StopSequence: "2"},
			{StopSequence: "3"},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"T6": {0, 1, 2},
			},
		},
	}
	validations.StopSequenceValidation(trip, 0, gtfs, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, make(map[string][]types.StopTimeRaw))

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid stop_sequence and shape_dist_traveled should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

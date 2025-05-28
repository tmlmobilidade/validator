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
	trip := &types.Trip{TripId: "T1"}
	gtfs := &types.Gtfs{
		Files: map[string][]map[string]string{
			"stop_times": {
				{"stop_sequence": "1", "shape_dist_traveled": "0"},
				{"stop_sequence": "3", "shape_dist_traveled": "20"},
				{"stop_sequence": "2", "shape_dist_traveled": "10"},
			},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"T1": {0, 1, 2},
			},
		},
	}
	validations.StopSequenceValidation(trip, 0, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid stop_sequence and shape_dist_traveled should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_NonIncreasingStopSequence(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{TripId: "T2"}
	gtfs := &types.Gtfs{
		Files: map[string][]map[string]string{
			"stop_times": {
				{"stop_sequence": "1", "shape_dist_traveled": "0"},
				{"stop_sequence": "2", "shape_dist_traveled": "10"},
				{"stop_sequence": "2", "shape_dist_traveled": "20"},
			},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"T2": {0, 1, 2},
			},
		},
	}
	validations.StopSequenceValidation(trip, 0, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Non-increasing stop_sequence should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_NonIncreasingShapeDistTraveled(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{TripId: "T3"}
	gtfs := &types.Gtfs{
		Files: map[string][]map[string]string{
			"stop_times": {
				{"stop_sequence": "1", "shape_dist_traveled": "0"},
				{"stop_sequence": "2", "shape_dist_traveled": "10"},
				{"stop_sequence": "3", "shape_dist_traveled": "5"},
			},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"T3": {0, 1, 2},
			},
		},
	}
	validations.StopSequenceValidation(trip, 0, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Non-increasing shape_dist_traveled should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_InvalidStopSequence(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{TripId: "T4"}
	gtfs := &types.Gtfs{
		Files: map[string][]map[string]string{
			"stop_times": {
				{"stop_sequence": "A", "shape_dist_traveled": "0"},
				{"stop_sequence": "2", "shape_dist_traveled": "10"},
			},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"T4": {0, 1},
			},
		},
	}
	validations.StopSequenceValidation(trip, 0, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid stop_sequence should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_InvalidShapeDistTraveled(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{TripId: "T5"}
	gtfs := &types.Gtfs{
		Files: map[string][]map[string]string{
			"stop_times": {
				{"stop_sequence": "1", "shape_dist_traveled": "foo"},
				{"stop_sequence": "2", "shape_dist_traveled": "10"},
			},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"T5": {0, 1},
			},
		},
	}
	validations.StopSequenceValidation(trip, 0, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid shape_dist_traveled should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_ValidInput_NoShapeDistTraveled(t *testing.T) {
	services.AppMessageService.Clear()
	
	trip := &types.Trip{TripId: "T6"}
	gtfs := &types.Gtfs{
		Files: map[string][]map[string]string{
			"stop_times": {
				{"stop_sequence": "1"},
				{"stop_sequence": "2"},
				{"stop_sequence": "3"},
			},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"T6": {0, 1, 2},
			},
		},
	}
	validations.StopSequenceValidation(trip, 0, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid stop_sequence and shape_dist_traveled should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
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
	tripStopTimesCache := map[string][]types.StopTimeRaw{
		"T1": {
			{TripId: "T1", StopSequence: "1", ShapeDistTraveled: "0"},
			{TripId: "T1", StopSequence: "3", ShapeDistTraveled: "20"},
			{TripId: "T1", StopSequence: "2", ShapeDistTraveled: "10"},
		},
	}
	validations.StopSequenceValidation(trip, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
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
	tripStopTimesCache := map[string][]types.StopTimeRaw{
		"T2": {
			{TripId: "T2", StopSequence: "1", ShapeDistTraveled: "0"},
			{TripId: "T2", StopSequence: "2", ShapeDistTraveled: "10"},
			{TripId: "T2", StopSequence: "2", ShapeDistTraveled: "20"},
		},
	}
	validations.StopSequenceValidation(trip, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
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
	tripStopTimesCache := map[string][]types.StopTimeRaw{
		"T3": {
			{TripId: "T3", StopSequence: "1", ShapeDistTraveled: "0"},
			{TripId: "T3", StopSequence: "2", ShapeDistTraveled: "10"},
			{TripId: "T3", StopSequence: "3", ShapeDistTraveled: "5"},
		},
	}
	validations.StopSequenceValidation(trip, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
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
	tripStopTimesCache := map[string][]types.StopTimeRaw{
		"T4": {
			{TripId: "T4", StopSequence: "A", ShapeDistTraveled: "0"},
			{TripId: "T4", StopSequence: "2", ShapeDistTraveled: "10"},
		},
	}
	validations.StopSequenceValidation(trip, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
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
	tripStopTimesCache := map[string][]types.StopTimeRaw{
		"T5": {
			{TripId: "T5", StopSequence: "1", ShapeDistTraveled: "foo"},
			{TripId: "T5", StopSequence: "2", ShapeDistTraveled: "10"},
		},
	}
	validations.StopSequenceValidation(trip, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
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
	tripStopTimesCache := map[string][]types.StopTimeRaw{
		"T6": {
			{TripId: "T6", StopSequence: "1"},
			{TripId: "T6", StopSequence: "2"},
			{TripId: "T6", StopSequence: "3"},
		},
	}
	validations.StopSequenceValidation(trip, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid stop_sequence and shape_dist_traveled should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_ConsecutiveStopIds_ValidDifferentStopIds(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{TripId: lib.Ptr("T7")}
	tripStopTimesCache := map[string][]types.StopTimeRaw{
		"T7": {
			{TripId: "T7", StopSequence: "1", StopId: "S1"},
			{TripId: "T7", StopSequence: "2", StopId: "S2"},
			{TripId: "T7", StopSequence: "3", StopId: "S3"},
		},
	}
	validations.StopSequenceValidation(trip, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Different stop_ids consecutively (when ordered by stop_sequence) should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_ConsecutiveStopIds_InvalidConsecutiveStopIds(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{TripId: lib.Ptr("T8")}
	tripStopTimesCache := map[string][]types.StopTimeRaw{
		"T8": {
			{TripId: "T8", StopSequence: "1", StopId: "S1"},
			{TripId: "T8", StopSequence: "2", StopId: "S1"}, // Same stop_id consecutively
			{TripId: "T8", StopSequence: "3", StopId: "S2"},
		},
	}
	validations.StopSequenceValidation(trip, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Same stop_id consecutively (when ordered by stop_sequence) should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_ConsecutiveStopIds_SameStopIdDifferentTrips(t *testing.T) {
	services.AppMessageService.Clear()
	trip1 := &types.Trip{TripId: lib.Ptr("T9")}
	trip2 := &types.Trip{TripId: lib.Ptr("T10")}
	tripStopTimesCache := map[string][]types.StopTimeRaw{
		"T9": {
			{TripId: "T9", StopSequence: "1", StopId: "S1"},
			{TripId: "T9", StopSequence: "2", StopId: "S2"},
		},
		"T10": {
			{TripId: "T10", StopSequence: "1", StopId: "S1"}, // Same stop_id but different trip
			{TripId: "T10", StopSequence: "2", StopId: "S2"},
		},
	}
	validations.StopSequenceValidation(trip1, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
	validations.StopSequenceValidation(trip2, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Same stop_id in different trips should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_ConsecutiveStopIds_FirstStopInTrip(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{TripId: lib.Ptr("T11")}
	tripStopTimesCache := map[string][]types.StopTimeRaw{
		"T11": {
			{TripId: "T11", StopSequence: "1", StopId: "S1"},
		},
	}
	validations.StopSequenceValidation(trip, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "First stop in a trip should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_ConsecutiveStopIds_MissingStopId(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{TripId: lib.Ptr("T12")}
	tripStopTimesCache := map[string][]types.StopTimeRaw{
		"T12": {
			{TripId: "T12", StopSequence: "1", StopId: ""}, // Empty stop_id
			{TripId: "T12", StopSequence: "2", StopId: "S1"},
		},
	}
	validations.StopSequenceValidation(trip, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing/empty stop_id should skip consecutive stop_id validation",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_ConsecutiveStopIds_ValidAfterNonConsecutive(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{TripId: lib.Ptr("T13")}
	tripStopTimesCache := map[string][]types.StopTimeRaw{
		"T13": {
			{TripId: "T13", StopSequence: "1", StopId: "S1"},
			{TripId: "T13", StopSequence: "2", StopId: "S2"},
			{TripId: "T13", StopSequence: "3", StopId: "S1"}, // Same as first, but not consecutive
		},
	}
	validations.StopSequenceValidation(trip, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Same stop_id appearing non-consecutively (when ordered by stop_sequence) should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_ConsecutiveStopIds_OrderedByStopSequence(t *testing.T) {
	services.AppMessageService.Clear()
	// Test case from user's example: stop_sequences out of order in file, but correct when sorted
	trip := &types.Trip{TripId: lib.Ptr("T14")}
	tripStopTimesCache := map[string][]types.StopTimeRaw{
		"T14": {
			{TripId: "T14", StopSequence: "2", StopId: "050107"}, // Appears first in file
			{TripId: "T14", StopSequence: "1", StopId: "050113"}, // Appears second in file
			{TripId: "T14", StopSequence: "3", StopId: "050109"}, // Appears third in file
		},
	}
	validations.StopSequenceValidation(trip, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Stop_ids should be checked consecutively when ordered by stop_sequence, not file order",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopSequenceValidation_ConsecutiveStopIds_OrderedByStopSequence_Invalid(t *testing.T) {
	services.AppMessageService.Clear()
	// Test case where stop_ids are consecutive when ordered by stop_sequence
	trip := &types.Trip{TripId: lib.Ptr("T15")}
	tripStopTimesCache := map[string][]types.StopTimeRaw{
		"T15": {
			{TripId: "T15", StopSequence: "2", StopId: "S1"}, // Appears first in file
			{TripId: "T15", StopSequence: "1", StopId: "S1"}, // Same stop_id, appears second in file
			{TripId: "T15", StopSequence: "3", StopId: "S2"},
		},
	}
	validations.StopSequenceValidation(trip, 0, &types.Gtfs{}, &types.TripsRules{StopSequence: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, tripStopTimesCache)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Consecutive stop_ids when ordered by stop_sequence should error even if out of order in file",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

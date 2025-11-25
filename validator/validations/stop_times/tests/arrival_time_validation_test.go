package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestArrivalTimeValidation_MissingForTimepoint1(t *testing.T) {
	services.AppMessageService.Clear()

	timepoint := 1
	stopTime := &types.StopTime{
		Timepoint:   &timepoint,
		ArrivalTime: nil,
	}

	validations.ArrivalTimeValidation(stopTime, 1, &types.Gtfs{}, nil, make(map[string]types.TripStopSequence))

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing arrival_time for timepoint=1 should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestArrivalTimeValidation_ForbiddenWithPickupDropOffWindow(t *testing.T) {
	services.AppMessageService.Clear()

	startWindow := "08:00:00"
	arrival := "09:00:00"

	stopTime := &types.StopTime{
		StartPickupDropOffWindow: &startWindow,
		ArrivalTime:              &arrival,
	}

	gtfs := &types.Gtfs{}

	validations.ArrivalTimeValidation(stopTime, 2, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "arrival_time is forbidden when start_pickup_drop_off_window is defined",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestArrivalTimeValidation_MissingForFirstStop(t *testing.T) {
	services.AppMessageService.Clear()

	stopSeq := 1
	tripId := "trip1"
	stopTime := &types.StopTime{
		StopSequence: &stopSeq,
		TripId:       &tripId,
		ArrivalTime:  nil,
	}

	gtfs := &types.Gtfs{
		StopTime: []types.StopTimeRaw{
			{StopSequence: "1", TripId: "trip1"},
			{StopSequence: "5", TripId: "trip1"},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"trip1": {0, 1},
			},
		},
	}

	validations.ArrivalTimeValidation(stopTime, 3, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing arrival_time for first stop in trip should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestArrivalTimeValidation_MissingForLastStop(t *testing.T) {
	services.AppMessageService.Clear()

	stopSeq := 5
	tripId := "trip1"
	stopTime := &types.StopTime{
		StopSequence: &stopSeq,
		TripId:       &tripId,
		ArrivalTime:  nil,
	}

	gtfs := &types.Gtfs{
		StopTime: []types.StopTimeRaw{
			{StopSequence: "1"},
			{StopSequence: "5"},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"trip1": {0, 1},
			},
		},
	}

	validations.ArrivalTimeValidation(stopTime, 4, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing arrival_time for last stop in trip should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestArrivalTimeValidation_InvalidTripId(t *testing.T) {
	services.AppMessageService.Clear()

	stopSeq := 1
	tripId := "invalid_trip"
	stopTime := &types.StopTime{
		StopSequence: &stopSeq,
		TripId:       &tripId,
		ArrivalTime:  nil,
	}

	gtfs := &types.Gtfs{
		StopTime: []types.StopTimeRaw{},
		IdMap: map[string]map[string][]int{
			"stop_times": {},
		},
	}

	validations.ArrivalTimeValidation(stopTime, 5, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid trip_id should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestArrivalTimeValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()

	stopSeq := 2
	tripId := "trip1"
	arrival := "10:00:00"
	stopTime := &types.StopTime{
		StopSequence: &stopSeq,
		TripId:       &tripId,
		ArrivalTime:  &arrival,
	}

	gtfs := &types.Gtfs{
		StopTime: []types.StopTimeRaw{
			{StopSequence: "1"},
			{StopSequence: "2"},
			{StopSequence: "5"},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"trip1": {0, 1, 2},
			},
		},
	}

	validations.ArrivalTimeValidation(stopTime, 6, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid input should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestArrivalTimeValidation_InvalidStopSequence(t *testing.T) {
	services.AppMessageService.Clear()

	stopSeq := 1
	tripId := "trip1"
	stopTime := &types.StopTime{
		StopSequence: &stopSeq,
		TripId:       &tripId,
		ArrivalTime:  nil,
	}

	gtfs := &types.Gtfs{
		StopTime: []types.StopTimeRaw{
			{StopSequence: "INVALID"},
			{StopSequence: "2"},
			{StopSequence: "5"},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"trip1": {0, 1, 2},
			},
		},
	}

	validations.ArrivalTimeValidation(stopTime, 7, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid stop_sequence should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestArrivalTimeValidation_InvalidTime(t *testing.T) {
	services.AppMessageService.Clear()

	stopSeq := 1
	tripId := "trip1"
	arrival := "INVALID"
	stopTime := &types.StopTime{
		StopSequence: &stopSeq,
		TripId:       &tripId,
		ArrivalTime:  &arrival,
	}

	gtfs := &types.Gtfs{
		StopTime: []types.StopTimeRaw{
			{StopSequence: "1"},
			{StopSequence: "2"},
			{StopSequence: "5"},
		},
		IdMap: map[string]map[string][]int{
			"stop_times": {
				"trip1": {0, 1, 2},
			},
		},
	}

	validations.ArrivalTimeValidation(stopTime, 8, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid time should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestArrivalTimeValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}

	severity := types.SEVERITY_ERROR
	validations.ArrivalTimeValidation(stopTime, 1, &types.Gtfs{}, &types.StopTimesRules{ArrivalTime: types.RuleConfig{Severity: severity}}, make(map[string]validations.TripStopSequence))

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Severity error should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestArrivalTimeValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}

	severity := types.SEVERITY_WARNING
	validations.ArrivalTimeValidation(stopTime, 1, &types.Gtfs{}, &types.StopTimesRules{ArrivalTime: types.RuleConfig{Severity: severity}}, make(map[string]validations.TripStopSequence))

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Severity warning should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

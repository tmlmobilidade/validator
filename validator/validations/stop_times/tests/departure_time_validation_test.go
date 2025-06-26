package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestDepartureTimeValidation_MissingForTimepoint1(t *testing.T) {
	services.AppMessageService.Clear()

	timepoint := 1
	stopTime := &types.StopTime{
		Timepoint:   &timepoint,
		DepartureTime: nil,
	}
	gtfs := &types.Gtfs{}

	validations.DepartureTimeValidation(nil, stopTime, 1, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing departure_time for timepoint=1 should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDepartureTimeValidation_ForbiddenWithPickupDropOffWindow(t *testing.T) {
	services.AppMessageService.Clear()

	startWindow := "08:00:00"
	departure := "09:00:00"
	
	stopTime := &types.StopTime{
		StartPickupDropOffWindow: &startWindow,
		DepartureTime: &departure,
	}
	
	gtfs := &types.Gtfs{}
	
	validations.DepartureTimeValidation(nil, stopTime, 2, gtfs)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "departure_time is forbidden when start_pickup_drop_off_window is defined",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDepartureTimeValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()

	stopSeq := 2
	tripId := "trip1"
	departure := "10:00:00"
	stopTime := &types.StopTime{
		StopSequence: &stopSeq,
		TripId:      &tripId,
		DepartureTime: &departure,
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

	validations.DepartureTimeValidation(nil, stopTime, 6, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid input should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDepartureTimeValidation_InvalidTime(t *testing.T) {
	services.AppMessageService.Clear()

	stopSeq := 1
	tripId := "trip1"
	departure := "INVALID"
	stopTime := &types.StopTime{
		StopSequence: &stopSeq,
		TripId:      &tripId,
		DepartureTime: &departure,
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

	validations.DepartureTimeValidation(nil, stopTime, 8, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid time should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDepartureTimeValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	
	severity := types.SEVERITY_ERROR
	validations.DepartureTimeValidation(&severity, stopTime, 1, &types.Gtfs{})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Severity error should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDepartureTimeValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	
	severity := types.SEVERITY_WARNING
	validations.DepartureTimeValidation(&severity, stopTime, 1, &types.Gtfs{})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Severity warning should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestStartPickupDropOffWindowValidation_RequiredWithLocationGroupId(t *testing.T) {
	services.AppMessageService.Clear()
	locationGroupId := "LG1"
	stopTime := &types.StopTime{
		LocationGroupId: &locationGroupId,
	}

	validations.StartPickupDropOffWindowValidation(stopTime, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing start_pickup_drop_off_window with location_group_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStartPickupDropOffWindowValidation_RequiredWithLocationId(t *testing.T) {
	services.AppMessageService.Clear()
	locationId := "L1"
	stopTime := &types.StopTime{
		LocationId: &locationId,
	}

	validations.StartPickupDropOffWindowValidation(stopTime, 2, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing start_pickup_drop_off_window with location_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStartPickupDropOffWindowValidation_RequiredWithEndPickupDropOffWindow(t *testing.T) {
	services.AppMessageService.Clear()
	endWindow := "10:00:00"
	stopTime := &types.StopTime{
		EndPickupDropOffWindow: &endWindow,
	}

	validations.StartPickupDropOffWindowValidation(stopTime, 3, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing start_pickup_drop_off_window with end_pickup_drop_off_window should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStartPickupDropOffWindowValidation_ForbiddenWithArrivalTime(t *testing.T) {
	services.AppMessageService.Clear()
	arrival := "08:00:00"
	startWindow := "07:00:00"
	stopTime := &types.StopTime{
		ArrivalTime:              &arrival,
		StartPickupDropOffWindow: &startWindow,
	}

	validations.StartPickupDropOffWindowValidation(stopTime, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "start_pickup_drop_off_window should be forbidden if arrival_time is defined",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStartPickupDropOffWindowValidation_ForbiddenWithDepartureTime(t *testing.T) {
	services.AppMessageService.Clear()
	departure := "09:00:00"
	startWindow := "07:00:00"
	stopTime := &types.StopTime{
		DepartureTime:            &departure,
		StartPickupDropOffWindow: &startWindow,
	}

	validations.StartPickupDropOffWindowValidation(stopTime, 5, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "start_pickup_drop_off_window should be forbidden if departure_time is defined",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStartPickupDropOffWindowValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	locationGroupId := "LG1"
	startWindow := "07:00:00"
	stopTime := &types.StopTime{
		LocationGroupId:          &locationGroupId,
		StartPickupDropOffWindow: &startWindow,
	}

	validations.StartPickupDropOffWindowValidation(stopTime, 6, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid start_pickup_drop_off_window should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStartPickupDropOffWindowValidation_InvalidTime(t *testing.T) {
	services.AppMessageService.Clear()
	locationGroupId := "LG1"
	startWindow := "INVALID"
	stopTime := &types.StopTime{
		LocationGroupId:          &locationGroupId,
		StartPickupDropOffWindow: &startWindow,
	}

	validations.StartPickupDropOffWindowValidation(stopTime, 7, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid time for start_pickup_drop_off_window should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStartPickupDropOffWindowValidation_Optional(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}

	validations.StartPickupDropOffWindowValidation(stopTime, 8, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Optional start_pickup_drop_off_window should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

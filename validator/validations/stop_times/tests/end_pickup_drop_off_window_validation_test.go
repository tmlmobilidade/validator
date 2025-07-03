package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestEndPickupDropOffWindowValidation_RequiredWithLocationGroupId(t *testing.T) {
	services.AppMessageService.Clear()
	locationGroupId := "LG1"
	stopTime := &types.StopTime{
		LocationGroupId: &locationGroupId,
	}

	validations.EndPickupDropOffWindowValidation(stopTime, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing end_pickup_drop_off_window with location_group_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_RequiredWithLocationId(t *testing.T) {
	services.AppMessageService.Clear()
	locationId := "L1"
	stopTime := &types.StopTime{
		LocationId: &locationId,
	}

	validations.EndPickupDropOffWindowValidation(stopTime, 2, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing end_pickup_drop_off_window with location_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_RequiredWithStartPickupDropOffWindow(t *testing.T) {
	services.AppMessageService.Clear()
	startWindow := "07:00:00"
	stopTime := &types.StopTime{
		StartPickupDropOffWindow: &startWindow,
	}

	validations.EndPickupDropOffWindowValidation(stopTime, 3, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing end_pickup_drop_off_window with start_pickup_drop_off_window should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_ForbiddenWithArrivalTime(t *testing.T) {
	services.AppMessageService.Clear()
	arrival := "08:00:00"
	endWindow := "10:00:00"
	stopTime := &types.StopTime{
		ArrivalTime:            &arrival,
		EndPickupDropOffWindow: &endWindow,
	}

	validations.EndPickupDropOffWindowValidation(stopTime, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "end_pickup_drop_off_window should be forbidden if arrival_time is defined",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_ForbiddenWithDepartureTime(t *testing.T) {
	services.AppMessageService.Clear()
	departure := "09:00:00"
	endWindow := "10:00:00"
	stopTime := &types.StopTime{
		DepartureTime:          &departure,
		EndPickupDropOffWindow: &endWindow,
	}

	validations.EndPickupDropOffWindowValidation(stopTime, 5, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "end_pickup_drop_off_window should be forbidden if departure_time is defined",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	locationGroupId := "LG1"
	endWindow := "10:00:00"
	stopTime := &types.StopTime{
		LocationGroupId:        &locationGroupId,
		EndPickupDropOffWindow: &endWindow,
	}

	validations.EndPickupDropOffWindowValidation(stopTime, 6, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid end_pickup_drop_off_window should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_InvalidTime(t *testing.T) {
	services.AppMessageService.Clear()
	locationGroupId := "LG1"
	endWindow := "INVALID"
	stopTime := &types.StopTime{
		LocationGroupId:        &locationGroupId,
		EndPickupDropOffWindow: &endWindow,
	}

	validations.EndPickupDropOffWindowValidation(stopTime, 7, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid time for end_pickup_drop_off_window should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_Optional(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}

	validations.EndPickupDropOffWindowValidation(stopTime, 8, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Optional end_pickup_drop_off_window should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	severity := types.SEVERITY_ERROR
	validations.EndPickupDropOffWindowValidation(stopTime, 9, &types.StopTimesRules{EndPickupDropOffWindow: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "end_pickup_drop_off_window missing with severity error should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	severity := types.SEVERITY_WARNING
	validations.EndPickupDropOffWindowValidation(stopTime, 10, &types.StopTimesRules{EndPickupDropOffWindow: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "end_pickup_drop_off_window missing with severity warning should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

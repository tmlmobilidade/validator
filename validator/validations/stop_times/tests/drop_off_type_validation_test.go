package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestDropOffTypeValidation_ValidValues(t *testing.T) {
	services.AppMessageService.Clear()
	for _, val := range []int{0, 1, 2, 3} {
		stopTime := &types.StopTime{DropOffType: &val}
		validations.DropOffTypeValidation(stopTime, 1, nil)
	}
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid enum values should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDropOffTypeValidation_InvalidEnum(t *testing.T) {
	services.AppMessageService.Clear()
	val := 5
	stopTime := &types.StopTime{DropOffType: &val}
	validations.DropOffTypeValidation(stopTime, 2, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid enum value should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDropOffTypeValidation_ForbiddenZeroWithStartWindow(t *testing.T) {
	services.AppMessageService.Clear()
	val := 0
	startWindow := "07:00:00"
	stopTime := &types.StopTime{DropOffType: &val, StartPickupDropOffWindow: &startWindow}
	validations.DropOffTypeValidation(stopTime, 3, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "drop_off_type=0 forbidden with start_pickup_drop_off_window",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDropOffTypeValidation_ForbiddenZeroWithEndWindow(t *testing.T) {
	services.AppMessageService.Clear()
	val := 0
	endWindow := "10:00:00"
	stopTime := &types.StopTime{DropOffType: &val, EndPickupDropOffWindow: &endWindow}
	validations.DropOffTypeValidation(stopTime, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "drop_off_type=0 forbidden with end_pickup_drop_off_window",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDropOffTypeValidation_AllowedWithStartWindowIfOne(t *testing.T) {
	services.AppMessageService.Clear()
	val := 1
	startWindow := "07:00:00"
	stopTime := &types.StopTime{DropOffType: &val, StartPickupDropOffWindow: &startWindow}
	validations.DropOffTypeValidation(stopTime, 5, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "drop_off_type=1 allowed with start_pickup_drop_off_window",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDropOffTypeValidation_OptionalNotPresent(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	validations.DropOffTypeValidation(stopTime, 6, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "drop_off_type not present should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDropOffTypeValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	severity := types.SEVERITY_ERROR
	validations.DropOffTypeValidation(stopTime, 7, &types.StopTimesRules{DropOffType: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "drop_off_type missing with severity error should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDropOffTypeValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	severity := types.SEVERITY_WARNING
	validations.DropOffTypeValidation(stopTime, 8, &types.StopTimesRules{DropOffType: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "drop_off_type missing with severity warning should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

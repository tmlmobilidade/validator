package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestContinuousDropOffValidation_ValidValues(t *testing.T) {
	services.AppMessageService.Clear()
	for _, val := range []int{0, 1, 2, 3} {
		stopTime := &types.StopTime{ContinuousDropOff: &val}
		validations.ContinuousDropOffValidation(stopTime, 1, nil)
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

func TestContinuousDropOffValidation_InvalidEnum(t *testing.T) {
	services.AppMessageService.Clear()
	val := 5
	stopTime := &types.StopTime{ContinuousDropOff: &val}
	validations.ContinuousDropOffValidation(stopTime, 2, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid enum value should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousDropOffValidation_ForbiddenWithStartWindow(t *testing.T) {
	services.AppMessageService.Clear()
	val := 0
	startWindow := "07:00:00"
	stopTime := &types.StopTime{ContinuousDropOff: &val, StartPickupDropOffWindow: &startWindow}
	validations.ContinuousDropOffValidation(stopTime, 3, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "continuous_drop_off=0 forbidden with start_pickup_drop_off_window",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousDropOffValidation_ForbiddenWithEndWindow(t *testing.T) {
	services.AppMessageService.Clear()
	val := 2
	endWindow := "10:00:00"
	stopTime := &types.StopTime{ContinuousDropOff: &val, EndPickupDropOffWindow: &endWindow}
	validations.ContinuousDropOffValidation(stopTime, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "continuous_drop_off=2 forbidden with end_pickup_drop_off_window",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousDropOffValidation_AllowedWithStartWindowIfOne(t *testing.T) {
	services.AppMessageService.Clear()
	val := 1
	startWindow := "07:00:00"
	stopTime := &types.StopTime{ContinuousDropOff: &val, StartPickupDropOffWindow: &startWindow}
	validations.ContinuousDropOffValidation(stopTime, 5, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "continuous_drop_off=1 allowed with start_pickup_drop_off_window",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousDropOffValidation_OptionalNotPresent(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	validations.ContinuousDropOffValidation(stopTime, 6, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "continuous_drop_off not present should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousDropOffValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	severity := types.SEVERITY_ERROR
	validations.ContinuousDropOffValidation(stopTime, 7, &types.StopTimesRules{ContinuousDropOff: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "continuous_drop_off missing with severity error should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousDropOffValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	severity := types.SEVERITY_WARNING
	validations.ContinuousDropOffValidation(stopTime, 8, &types.StopTimesRules{ContinuousDropOff: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "continuous_drop_off missing with severity warning should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

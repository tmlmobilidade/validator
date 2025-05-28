package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestContinuousPickupValidation_ValidValues(t *testing.T) {
	services.AppMessageService.Clear()
	for _, val := range []int{0, 1, 2, 3} {
		stopTime := &types.StopTime{ContinuousPickup: &val}
		validations.ContinuousPickupValidation(nil, stopTime, 1)
	}
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid enum values should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousPickupValidation_InvalidEnum(t *testing.T) {
	services.AppMessageService.Clear()
	val := 5
	stopTime := &types.StopTime{ContinuousPickup: &val}
	validations.ContinuousPickupValidation(nil, stopTime, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid enum value should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousPickupValidation_ForbiddenWithStartWindow(t *testing.T) {
	services.AppMessageService.Clear()
	val := 0
	startWindow := "07:00:00"
	stopTime := &types.StopTime{ContinuousPickup: &val, StartPickupDropOffWindow: &startWindow}
	validations.ContinuousPickupValidation(nil, stopTime, 3)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "continuous_pickup=0 forbidden with start_pickup_drop_off_window",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousPickupValidation_ForbiddenWithEndWindow(t *testing.T) {
	services.AppMessageService.Clear()
	val := 2
	endWindow := "10:00:00"
	stopTime := &types.StopTime{ContinuousPickup: &val, EndPickupDropOffWindow: &endWindow}
	validations.ContinuousPickupValidation(nil, stopTime, 4)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "continuous_pickup=2 forbidden with end_pickup_drop_off_window",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousPickupValidation_AllowedWithStartWindowIfOne(t *testing.T) {
	services.AppMessageService.Clear()
	val := 1
	startWindow := "07:00:00"
	stopTime := &types.StopTime{ContinuousPickup: &val, StartPickupDropOffWindow: &startWindow}
	validations.ContinuousPickupValidation(nil, stopTime, 5)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "continuous_pickup=1 allowed with start_pickup_drop_off_window",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousPickupValidation_OptionalNotPresent(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	validations.ContinuousPickupValidation(nil, stopTime, 6)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "continuous_pickup not present should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousPickupValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	severity := types.SEVERITY_ERROR
	validations.ContinuousPickupValidation(&severity, stopTime, 7)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "continuous_pickup missing with severity error should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousPickupValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	severity := types.SEVERITY_WARNING
	validations.ContinuousPickupValidation(&severity, stopTime, 8)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "continuous_pickup missing with severity warning should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
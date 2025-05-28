package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestPickupTypeValidation_ForbiddenZeroWithStartWindow(t *testing.T) {
	services.AppMessageService.Clear()
	
	pt := 0
	startWindow := "07:00:00"
	stopTime := &types.StopTime{
		PickupType: &pt,
		StartPickupDropOffWindow: &startWindow,
	}
	
	validations.PickupTypeValidation(nil, stopTime, 1)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "pickup_type=0 forbidden with start_pickup_drop_off_window",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPickupTypeValidation_ForbiddenThreeWithEndWindow(t *testing.T) {
	services.AppMessageService.Clear()

	pt := 3
	endWindow := "10:00:00"
	stopTime := &types.StopTime{
		PickupType: &pt,
		EndPickupDropOffWindow: &endWindow,
	}
	
	validations.PickupTypeValidation(nil, stopTime, 2)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "pickup_type=3 forbidden with end_pickup_drop_off_window",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPickupTypeValidation_ValidValues(t *testing.T) {
	services.AppMessageService.Clear()

	pt := 2
	stopTime := &types.StopTime{
		PickupType: &pt,
	}
	
	validations.PickupTypeValidation(nil, stopTime, 3)
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "pickup_type=2 should be valid",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPickupTypeValidation_InvalidEnum(t *testing.T) {
	services.AppMessageService.Clear()

	pt := 5
	stopTime := &types.StopTime{
		PickupType: &pt,
	}
	
	validations.PickupTypeValidation(nil, stopTime, 4)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "pickup_type=5 should error as invalid enum",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPickupTypeValidation_OptionalNotPresent(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	
	validations.PickupTypeValidation(nil, stopTime, 5)
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "pickup_type not present should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 

func TestPickupTypeValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	
	stopTime := &types.StopTime{}

	severity := types.SEVERITY_ERROR
	validations.PickupTypeValidation(&severity, stopTime, 6)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "pickup_type=0 should error with severity error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPickupTypeValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	
	stopTime := &types.StopTime{}

	severity := types.SEVERITY_WARNING
	validations.PickupTypeValidation(&severity, stopTime, 7)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "pickup_type=0 should warn with severity warning",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}	
}
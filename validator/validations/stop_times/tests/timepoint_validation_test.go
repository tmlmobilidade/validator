package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestTimepointValidation_ValidValues(t *testing.T) {
	services.AppMessageService.Clear()
	for _, val := range []int{0, 1} {
		stopTime := &types.StopTime{Timepoint: &val}
		validations.TimepointValidation(nil, stopTime, 1)
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

func TestTimepointValidation_InvalidEnum(t *testing.T) {
	services.AppMessageService.Clear()
	val := 2
	stopTime := &types.StopTime{Timepoint: &val}
	validations.TimepointValidation(nil, stopTime, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid enum value should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTimepointValidation_OptionalNotPresent(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	validations.TimepointValidation(nil, stopTime, 3)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Optional timepoint not present should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTimepointValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	severity := types.SEVERITY_ERROR
	validations.TimepointValidation(&severity, stopTime, 4)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "timepoint missing with severity error should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTimepointValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	severity := types.SEVERITY_WARNING
	validations.TimepointValidation(&severity, stopTime, 5)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "timepoint missing with severity warning should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
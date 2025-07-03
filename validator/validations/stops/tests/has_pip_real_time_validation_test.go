package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestHasPipRealTimeValidation_MissingHasPipRealTime_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasPipRealTime: nil}
	validations.HasPipRealTimeValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_PipRealTime with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasPipRealTimeValidation_MissingHasPipRealTime_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasPipRealTime: nil}
	severity := types.SEVERITY_ERROR
	validations.HasPipRealTimeValidation(stop, 2, &types.StopsRules{HasPipRealTime: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_PipRealTime with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasPipRealTimeValidation_MissingHasPipRealTime_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasPipRealTime: nil}
	severity := types.SEVERITY_WARNING
	validations.HasPipRealTimeValidation(stop, 3, &types.StopsRules{HasPipRealTime: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing has_PipRealTime with severity WARNING should warn")
	}
}

func TestHasPipRealTimeValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	val := 1
	stop := &types.Stop{HasPipRealTime: &val}
	validations.HasPipRealTimeValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_PipRealTime should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasPipRealTimeValidation_InvalidInput_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	val := 9999
	stop := &types.Stop{HasPipRealTime: &val}
	severity := types.SEVERITY_ERROR
	validations.HasPipRealTimeValidation(stop, 5, &types.StopsRules{HasPipRealTime: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid has_PipRealTime should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasPipRealTimeValidation_ValidInput_WithOptions(t *testing.T) {
	services.AppMessageService.Clear()
	val := 1
	stop := &types.Stop{HasPipRealTime: &val}
	severity := types.SEVERITY_ERROR
	validations.HasPipRealTimeValidation(stop, 7, &types.StopsRules{HasPipRealTime: types.RuleConfig{Severity: severity, Options: &[]string{"1", "2"}}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_PipRealTime with severity ERROR should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasPipRealTimeValidation_InValidInput_WithOptions(t *testing.T) {
	services.AppMessageService.Clear()
	val := 5
	stop := &types.Stop{HasPipRealTime: &val}
	severity := types.SEVERITY_ERROR
	validations.HasPipRealTimeValidation(stop, 7, &types.StopsRules{HasPipRealTime: types.RuleConfig{Severity: severity, Options: &[]string{"1", "2"}}})
	if services.AppMessageService.GetSummary().TotalErrors != 1 {
		t.Error("Valid has_PipRealTime with severity ERROR should error")
	}
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_PipRealTime with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

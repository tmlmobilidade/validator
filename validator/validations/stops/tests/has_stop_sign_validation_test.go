package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestHasStopSignValidation_MissingHasStopSign_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasStopSign: nil}
	validations.HasStopSignValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_StopSign with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasStopSignValidation_MissingHasStopSign_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasStopSign: nil}
	severity := types.SEVERITY_ERROR
	validations.HasStopSignValidation(stop, 2, &types.StopsRules{HasStopSign: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_StopSign with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasStopSignValidation_MissingHasStopSign_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasStopSign: nil}
	severity := types.SEVERITY_WARNING
	validations.HasStopSignValidation(stop, 3, &types.StopsRules{HasStopSign: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing has_StopSign with severity WARNING should warn")
	}
}

func TestHasStopSignValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	val := 1
	stop := &types.Stop{HasStopSign: &val}
	validations.HasStopSignValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_StopSign should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasStopSignValidation_InvalidInput_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	val := 9999
	stop := &types.Stop{HasStopSign: &val}
	severity := types.SEVERITY_ERROR
	validations.HasStopSignValidation(stop, 5, &types.StopsRules{HasStopSign: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid has_StopSign should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasStopSignValidation_ValidInput_WithOptions(t *testing.T) {
	services.AppMessageService.Clear()
	val := 1
	stop := &types.Stop{HasStopSign: &val}
	severity := types.SEVERITY_ERROR
	validations.HasStopSignValidation(stop, 7, &types.StopsRules{HasStopSign: types.RuleConfig{Severity: severity, Options: &[]string{"1", "2"}}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_StopSign with severity ERROR should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasStopSignValidation_InValidInput_WithOptions(t *testing.T) {
	services.AppMessageService.Clear()
	val := 5
	stop := &types.Stop{HasStopSign: &val}
	severity := types.SEVERITY_ERROR
	validations.HasStopSignValidation(stop, 7, &types.StopsRules{HasStopSign: types.RuleConfig{Severity: severity, Options: &[]string{"1", "2"}}})
	if services.AppMessageService.GetSummary().TotalErrors != 1 {
		t.Error("Valid has_StopSign with severity ERROR should error")
	}
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_StopSign with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

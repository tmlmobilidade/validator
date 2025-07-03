package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestHasNetworkMapValidation_MissingHasNetworkMap_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasNetworkMap: nil}
	validations.HasNetworkMapValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_NetworkMap with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasNetworkMapValidation_MissingHasNetworkMap_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasNetworkMap: nil}
	severity := types.SEVERITY_ERROR
	validations.HasNetworkMapValidation(stop, 2, &types.StopsRules{HasNetworkMap: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_NetworkMap with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasNetworkMapValidation_MissingHasNetworkMap_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasNetworkMap: nil}
	severity := types.SEVERITY_WARNING
	validations.HasNetworkMapValidation(stop, 3, &types.StopsRules{HasNetworkMap: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing has_NetworkMap with severity WARNING should warn")
	}
}

func TestHasNetworkMapValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	val := 1
	stop := &types.Stop{HasNetworkMap: &val}
	validations.HasNetworkMapValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_NetworkMap should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasNetworkMapValidation_InvalidInput_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	val := 9999
	stop := &types.Stop{HasNetworkMap: &val}
	severity := types.SEVERITY_ERROR
	validations.HasNetworkMapValidation(stop, 5, &types.StopsRules{HasNetworkMap: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid has_NetworkMap should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasNetworkMapValidation_ValidInput_WithOptions(t *testing.T) {
	services.AppMessageService.Clear()
	val := 1
	stop := &types.Stop{HasNetworkMap: &val}
	severity := types.SEVERITY_ERROR
	validations.HasNetworkMapValidation(stop, 7, &types.StopsRules{HasNetworkMap: types.RuleConfig{Severity: severity, Options: &[]string{"1", "2"}}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_NetworkMap with severity ERROR should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasNetworkMapValidation_InValidInput_WithOptions(t *testing.T) {
	services.AppMessageService.Clear()
	val := 5
	stop := &types.Stop{HasNetworkMap: &val}
	severity := types.SEVERITY_ERROR
	validations.HasNetworkMapValidation(stop, 7, &types.StopsRules{HasNetworkMap: types.RuleConfig{Severity: severity, Options: &[]string{"1", "2"}}})
	if services.AppMessageService.GetSummary().TotalErrors != 1 {
		t.Error("Valid has_NetworkMap with severity ERROR should error")
	}
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_NetworkMap with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

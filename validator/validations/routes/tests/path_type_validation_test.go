package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestPathTypeValidation_NoRules(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{PathType: nil}
	validations.PathTypeValidation(route, 1, nil)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing path_type should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_RulesWithNotRequired(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{PathType: lib.Ptr("1")}
	validations.PathTypeValidation(route, 1, &types.RoutesRules{PathType: types.RuleConfig{Severity: types.SEVERITY_ERROR, Options: nil}})

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Path type should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_MissingPathType_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{PathType: nil}
	severity := types.SEVERITY_ERROR
	validations.PathTypeValidation(route, 2, &types.RoutesRules{PathType: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing path_type with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_MissingPathType_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{PathType: nil}
	severity := types.SEVERITY_WARNING
	validations.PathTypeValidation(route, 3, &types.RoutesRules{PathType: types.RuleConfig{Severity: severity}})
	// Note: The validation function always calls AddError() for missing values, not AddMessageWithSeverity()
	// So it will error even with WARNING severity
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Missing path_type with severity WARNING should error (validation always uses AddError)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_EmptyPathType_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{PathType: lib.Ptr("")}
	severity := types.SEVERITY_ERROR
	validations.PathTypeValidation(route, 4, &types.RoutesRules{PathType: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Empty path_type with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_Forbidden_Present(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{PathType: lib.Ptr("1")}
	severity := types.SEVERITY_FORBIDDEN
	validations.PathTypeValidation(route, 5, &types.RoutesRules{PathType: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Path type present when forbidden should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_Forbidden_Missing(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{PathType: nil}
	severity := types.SEVERITY_FORBIDDEN
	validations.PathTypeValidation(route, 5, &types.RoutesRules{PathType: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Path type missing when forbidden should skip validation",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_SeverityIgnore(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{PathType: lib.Ptr("1")}
	severity := types.SEVERITY_IGNORE
	validations.PathTypeValidation(route, 6, &types.RoutesRules{PathType: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Path type with severity IGNORE should skip validation",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_ValidInput_WithOptions(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{PathType: lib.Ptr("1")}
	severity := types.SEVERITY_ERROR
	validations.PathTypeValidation(route, 7, &types.RoutesRules{PathType: types.RuleConfig{Severity: severity, Options: &[]string{"1", "2", "3"}}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid path_type in options should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_InvalidInput_WithOptions(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{PathType: lib.Ptr("4")}
	severity := types.SEVERITY_ERROR
	validations.PathTypeValidation(route, 8, &types.RoutesRules{PathType: types.RuleConfig{Severity: severity, Options: &[]string{"1", "2", "3"}}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid path_type not in options should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_AllOptions(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{PathType: lib.Ptr("999")}
	severity := types.SEVERITY_ERROR
	validations.PathTypeValidation(route, 9, &types.RoutesRules{PathType: types.RuleConfig{Severity: severity, Options: &[]string{types.ALL_OPTIONS}}})
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Path type with ALL_OPTIONS should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_ValidInput_WithoutOptions(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{PathType: lib.Ptr("2")}
	severity := types.SEVERITY_ERROR
	validations.PathTypeValidation(route, 10, &types.RoutesRules{PathType: types.RuleConfig{Severity: severity, Options: nil}})
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid path_type without options should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

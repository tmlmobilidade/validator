package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestPathTypeValidation_MissingPathType(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{PathType: nil}
	validations.PathTypeValidation(route, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing path_type should not error (optional field)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_EmptyPathType(t *testing.T) {
	services.AppMessageService.Clear()
	emptyPathType := ""
	route := &types.Route{PathType: &emptyPathType}
	validations.PathTypeValidation(route, 2, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Empty path_type should not error (optional field)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_InvalidPathType(t *testing.T) {
	services.AppMessageService.Clear()
	invalidPathType := "4"
	route := &types.Route{PathType: &invalidPathType}
	validations.PathTypeValidation(route, 3, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid path_type should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_ValidPathType(t *testing.T) {
	services.AppMessageService.Clear()
	validTypes := []string{"1", "2", "3"}
	for i, v := range validTypes {
		route := &types.Route{PathType: &v}
		validations.PathTypeValidation(route, i+4, nil)
		assertion := lib.AssertionMessage{
			Expected: 0,
			Actual:   services.AppMessageService.GetSummary().TotalErrors,
			Message:  "Valid path_type should not error",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Errorf("path_type %s: %s", v, assert)
		}
		services.AppMessageService.Clear()
	}
}

func TestPathTypeValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	invalidPathType := "99"
	route := &types.Route{PathType: &invalidPathType}
	severity := types.SEVERITY_ERROR
	validations.PathTypeValidation(route, 7, &types.RoutesRules{PathType: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid path_type should error with severity error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_RuleOptions(t *testing.T) {
	services.AppMessageService.Clear()
	pathType := "3"
	route := &types.Route{PathType: &pathType}
	validations.PathTypeValidation(route, 8, &types.RoutesRules{PathType: types.RuleConfig{Options: &[]string{"1", "2"}}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "path_type not in rule options should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_RuleOptionsValid(t *testing.T) {
	services.AppMessageService.Clear()
	pathType := "2"
	route := &types.Route{PathType: &pathType}
	validations.PathTypeValidation(route, 9, &types.RoutesRules{PathType: types.RuleConfig{Options: &[]string{"1", "2", "3"}}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "path_type in rule options should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_RuleOptionsAllOptions(t *testing.T) {
	services.AppMessageService.Clear()
	pathType := "3"
	route := &types.Route{PathType: &pathType}
	validations.PathTypeValidation(route, 10, &types.RoutesRules{PathType: types.RuleConfig{Options: &[]string{types.ALL_OPTIONS}}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "path_type with ALL_OPTIONS rule should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPathTypeValidation_Forbidden(t *testing.T) {
	services.AppMessageService.Clear()
	pathType := "1"
	route := &types.Route{PathType: &pathType}
	validations.PathTypeValidation(route, 11, &types.RoutesRules{PathType: types.RuleConfig{Severity: types.SEVERITY_FORBIDDEN}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "path_type with forbidden severity should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

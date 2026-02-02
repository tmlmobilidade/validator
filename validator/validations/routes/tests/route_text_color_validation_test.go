package routes

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestRouteTextColorValidation_MissingColor(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteTextColor: nil}
	validations.RouteTextColorValidation(route, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing route_text_color should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteTextColorValidation_EmptyColor(t *testing.T) {
	services.AppMessageService.Clear()
	empty := ""
	route := &types.Route{RouteTextColor: &empty}
	validations.RouteTextColorValidation(route, 2, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Empty route_text_color should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteTextColorValidation_InvalidColor(t *testing.T) {
	services.AppMessageService.Clear()
	color := "ZZZZZZ"
	route := &types.Route{RouteTextColor: &color}
	validations.RouteTextColorValidation(route, 3, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid route_text_color should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteTextColorValidation_ValidColor(t *testing.T) {
	services.AppMessageService.Clear()
	color := "123ABC"
	route := &types.Route{RouteTextColor: &color}
	validations.RouteTextColorValidation(route, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid route_text_color should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteTextColorValidation_NoColor_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteTextColor: nil}
	severity := types.SEVERITY_WARNING
	validations.RouteTextColorValidation(route, 5, &types.RoutesRules{RouteTextColor: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "No route_text_color should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteTextColorValidation_NoColor_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteTextColor: nil}
	severity := types.SEVERITY_ERROR
	validations.RouteTextColorValidation(route, 6, &types.RoutesRules{RouteTextColor: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "No route_text_color should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAllRouteTextColorValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericColorTestCases("route_text_color") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var severity types.Severity
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}

			validations.RouteTextColorValidation(&types.Route{RouteTextColor: tc.Color}, tc.Row, &types.RoutesRules{RouteTextColor: types.RuleConfig{Severity: severity}})
			if tc.ExpectedErrors > 0 {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name)
			}
		})
	}
}

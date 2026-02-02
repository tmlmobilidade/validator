package routes

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestAllRouteShortNameValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("route_short_name") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var severity types.Severity
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}
			validations.RouteShortNameValidation(&types.Route{RouteShortName: tc.Value}, tc.Row, &types.RoutesRules{RouteShortName: types.RuleConfig{Severity: severity}})
			if tc.ExpectedErrors > 0 {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name)
			}
		})
	}
}

func TestRouteShortNameValidation_TooLong(t *testing.T) {
	services.AppMessageService.Clear()
	longName := "This is a very long route short name that exceeds 12 characters"
	route := &types.Route{RouteShortName: &longName}
	validations.RouteShortNameValidation(route, 1, nil)
	summary := services.AppMessageService.GetSummary()
	if summary.TotalWarnings != 1 {
		t.Errorf("Expected 1 warning for route_short_name > 12 characters, got %d warnings", summary.TotalWarnings)
	}
}

func TestRouteShortNameValidation_Exactly12Characters(t *testing.T) {
	services.AppMessageService.Clear()
	name := "123456789012" // Exactly 12 characters
	route := &types.Route{RouteShortName: &name}
	validations.RouteShortNameValidation(route, 1, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Route short name with exactly 12 characters should not warn")
}

func TestRouteShortNameValidation_13Characters(t *testing.T) {
	services.AppMessageService.Clear()
	name := "1234567890123" // 13 characters
	route := &types.Route{RouteShortName: &name}
	validations.RouteShortNameValidation(route, 1, nil)
	summary := services.AppMessageService.GetSummary()
	if summary.TotalWarnings != 1 {
		t.Errorf("Expected 1 warning for route_short_name with 13 characters, got %d warnings", summary.TotalWarnings)
	}
}

func TestRouteShortNameValidation_BothShortAndLongNameMissing(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteShortName: nil, RouteLongName: nil}
	validations.RouteShortNameValidation(route, 1, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Both route_short_name and route_long_name missing should error")
}

func TestRouteShortNameValidation_ShortNameMissing_LongNamePresent(t *testing.T) {
	services.AppMessageService.Clear()
	longName := "Long Route Name"
	route := &types.Route{RouteShortName: nil, RouteLongName: &longName}
	validations.RouteShortNameValidation(route, 1, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "route_short_name missing but route_long_name present should not error")
}

func TestRouteShortNameValidation_ShortNameEmpty_LongNamePresent(t *testing.T) {
	services.AppMessageService.Clear()
	empty := ""
	longName := "Long Route Name"
	route := &types.Route{RouteShortName: &empty, RouteLongName: &longName}
	validations.RouteShortNameValidation(route, 1, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "route_short_name empty but route_long_name present should not error")
}

func TestRouteShortNameValidation_WithOptions_NotAllowed(t *testing.T) {
	services.AppMessageService.Clear()
	allowedOptions := []string{"1", "2", "3", "A", "B"}
	shortName := "X"
	route := &types.Route{RouteShortName: &shortName}
	rules := &types.RoutesRules{
		RouteShortName: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
			Options:  &allowedOptions,
		},
	}
	validations.RouteShortNameValidation(route, 1, rules)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Not allowed route_short_name should error")
}

func TestRouteShortNameValidation_WithOptions_AllOptions(t *testing.T) {
	services.AppMessageService.Clear()
	allOptions := []string{types.ALL_OPTIONS}
	shortName := "Any Name"
	route := &types.Route{RouteShortName: &shortName}
	rules := &types.RoutesRules{
		RouteShortName: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
			Options:  &allOptions,
		},
	}
	validations.RouteShortNameValidation(route, 1, rules)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "ALL_OPTIONS should allow any route_short_name")
}

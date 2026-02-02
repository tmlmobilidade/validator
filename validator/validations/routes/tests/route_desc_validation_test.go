package routes

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestAllRouteDescValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("route_desc") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var routeDesc *string
			if tc.Value != nil {
				routeDesc = tc.Value
			}

			var severity types.Severity
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}

			validations.RouteDescValidation(&types.Route{RouteDesc: routeDesc}, tc.Row, &types.RoutesRules{RouteDesc: types.RuleConfig{Severity: severity}})
			assertion := lib.AssertionMessage{
				Expected: tc.ExpectedErrors,
				Actual:   services.AppMessageService.GetSummary().TotalErrors,
				Message:  tc.Name,
			}
			assertionWarnings := lib.AssertionMessage{
				Expected: tc.ExpectedWarnings,
				Actual:   services.AppMessageService.GetSummary().TotalWarnings,
				Message:  tc.Name,
			}
			if assert := lib.Assert(assertion); assert != "" {
				t.Errorf("%s: %s", tc.Name, assert)
			}
			if assert := lib.Assert(assertionWarnings); assert != "" {
				t.Errorf("%s: %s", tc.Name, assert)
			}
		})
	}

	for _, tc := range test_helpers.GetGenericSeverityTestCases("route_desc") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var routeDesc *string
			if tc.Value != nil {
				if ptr, ok := tc.Value.(*string); ok && ptr != nil {
					routeDesc = ptr
				}
			}
			validations.RouteDescValidation(&types.Route{RouteDesc: routeDesc}, tc.Row, &types.RoutesRules{RouteDesc: types.RuleConfig{Severity: tc.Severity}})
			assertion := lib.AssertionMessage{
				Expected: tc.ExpectedErrors,
				Actual:   services.AppMessageService.GetSummary().TotalErrors,
				Message:  tc.Name,
			}
			assertionWarnings := lib.AssertionMessage{
				Expected: tc.ExpectedWarnings,
				Actual:   services.AppMessageService.GetSummary().TotalWarnings,
				Message:  tc.Name,
			}
			if assert := lib.Assert(assertion); assert != "" {
				t.Errorf("%s: %s", tc.Name, assert)
			}
			if assert := lib.Assert(assertionWarnings); assert != "" {
				t.Errorf("%s: %s", tc.Name, assert)
			}
		})
	}
}
func TestRouteDescValidation_DuplicateShortName(t *testing.T) {
	services.AppMessageService.Clear()
	desc := "Route A"
	shortName := "Route A"
	route := &types.Route{
		RouteDesc:      &desc,
		RouteShortName: &shortName,
	}
	validations.RouteDescValidation(route, 1, nil)
	summary := services.AppMessageService.GetSummary()
	if summary.TotalWarnings != 1 {
		t.Errorf("Expected 1 warning for route_desc matching route_short_name, got %d warnings", summary.TotalWarnings)
	}
}

func TestRouteDescValidation_DuplicateBothShortAndLongName(t *testing.T) {
	services.AppMessageService.Clear()
	desc := "Route A"
	shortName := "Route A"
	longName := "Route A"
	route := &types.Route{
		RouteDesc:      &desc,
		RouteShortName: &shortName,
		RouteLongName:  &longName,
	}
	validations.RouteDescValidation(route, 1, nil)
	summary := services.AppMessageService.GetSummary()
	if summary.TotalWarnings != 2 {
		t.Errorf("Expected 2 warnings for route_desc matching both route_short_name and route_long_name, got %d warnings", summary.TotalWarnings)
	}
}

func TestRouteDescValidation_NoDuplicates(t *testing.T) {
	services.AppMessageService.Clear()
	desc := "A route description"
	shortName := "Route A"
	longName := "Main Street Express"
	route := &types.Route{
		RouteDesc:      &desc,
		RouteShortName: &shortName,
		RouteLongName:  &longName,
	}
	validations.RouteDescValidation(route, 1, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Unique route_desc should not warn")
}

func TestRouteDescValidation_RequiredWhenShortNameEmpty(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{
		RouteDesc:      nil,
		RouteShortName: nil,
	}
	rules := &types.RoutesRules{
		RouteDesc: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
		},
	}
	validations.RouteDescValidation(route, 1, rules)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "route_desc should be required when route_short_name is empty")
}

func TestRouteDescValidation_OptionalWhenShortNamePresent(t *testing.T) {
	services.AppMessageService.Clear()
	shortName := "Route A"
	route := &types.Route{
		RouteDesc:      nil,
		RouteShortName: &shortName,
	}
	validations.RouteDescValidation(route, 1, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "route_desc should be optional when route_short_name is present")
}

func TestRouteDescValidation_ShortNameEmptyString_Required(t *testing.T) {
	services.AppMessageService.Clear()
	empty := ""
	route := &types.Route{
		RouteDesc:      nil,
		RouteShortName: &empty,
	}
	rules := &types.RoutesRules{
		RouteDesc: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
		},
	}
	validations.RouteDescValidation(route, 1, rules)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "route_desc should be required when route_short_name is empty string")
}

func TestRouteDescValidation_Forbidden(t *testing.T) {
	services.AppMessageService.Clear()
	desc := "Some description"
	route := &types.Route{
		RouteDesc:      &desc,
		RouteShortName: lib.Ptr("Route A"),
	}
	rules := &types.RoutesRules{
		RouteDesc: types.RuleConfig{
			Severity: types.SEVERITY_FORBIDDEN,
		},
	}
	validations.RouteDescValidation(route, 1, rules)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "FORBIDDEN severity should error when route_desc is present")
}

func TestRouteDescValidation_WithOptions_NotAllowing(t *testing.T) {
	services.AppMessageService.Clear()
	allowedOptions := []string{"Description 1", "Description 2"}
	desc := "Invalid Description"
	route := &types.Route{
		RouteDesc:      &desc,
		RouteShortName: lib.Ptr("Route A"),
	}
	rules := &types.RoutesRules{
		RouteDesc: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
			Options:  &allowedOptions,
		},
	}
	validations.RouteDescValidation(route, 1, rules)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Not allowing route_desc should error")
}

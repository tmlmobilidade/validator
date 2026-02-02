package routes

import (
	"fmt"
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestAllContinuousDropOffValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetContinuousPickupDropOffValidOptions()
	validOptionsStrings := make([]string, len(validOptions))
	for i, opt := range validOptions {
		validOptionsStrings[i] = fmt.Sprintf("%d", opt)
	}
	for _, tc := range test_helpers.GetGenericEnumIntTestCases("continuous_drop_off", validOptions) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var severity types.Severity
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}

			var continuousDropOff *string
			if tc.Value != nil {
				if ptr, ok := tc.Value.(*int); ok && ptr != nil {
					valStr := fmt.Sprintf("%d", *ptr)
					continuousDropOff = &valStr
				}
			}

			validations.ContinuousDropOffValidation(&types.Route{ContinuousDropOff: continuousDropOff}, tc.Row, &types.Gtfs{}, &types.RoutesRules{ContinuousDropOff: types.RuleConfig{Severity: severity, Options: &validOptionsStrings}}, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}

	for _, tc := range test_helpers.GetGenericSeverityTestCases("continuous_drop_off") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var continuousDropOff *string
			if tc.Value != nil {
				if ptr, ok := tc.Value.(*string); ok && ptr != nil {
					continuousDropOff = ptr
				}
			}

			validations.ContinuousDropOffValidation(&types.Route{ContinuousDropOff: continuousDropOff}, tc.Row, &types.Gtfs{}, &types.RoutesRules{ContinuousDropOff: types.RuleConfig{Severity: tc.Severity, Options: &validOptionsStrings}}, nil)
			if tc.ExpectedErrors > 0 {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name)
			}
		})
	}
}

// Test that continuous_drop_off=1 is valid if the route has no trips with pickup/dropoff windows
func TestContinuousDropOffValidation_ValidWithoutWindows(t *testing.T) {
	services.AppMessageService.Clear()
	routesWithWindows := map[string]bool{}
	validations.ContinuousDropOffValidation(&types.Route{RouteId: lib.Ptr("MY_ROUTE_ID"), ContinuousDropOff: lib.Ptr("1")}, 1, &types.Gtfs{}, nil, routesWithWindows)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "ContinuousDropOffValidation should not error without windows")
}

func TestContinuousDropOffValidation_ForbiddenValue0WithWindows(t *testing.T) {
	services.AppMessageService.Clear()
	continuousDropOff := "0"
	routeId := "ROUTE1"
	route := &types.Route{
		RouteId:           &routeId,
		ContinuousDropOff: &continuousDropOff,
	}
	routesWithWindows := map[string]bool{routeId: true}
	rules := &types.RoutesRules{
		ContinuousDropOff: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
		},
	}
	validations.ContinuousDropOffValidation(route, 1, &types.Gtfs{}, rules, routesWithWindows)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "continuous_drop_off=0 should be forbidden when route has windows")
}

func TestContinuousDropOffValidation_ForbiddenValue2WithWindows(t *testing.T) {
	services.AppMessageService.Clear()
	continuousDropOff := "2"
	routeId := "ROUTE1"
	route := &types.Route{
		RouteId:           &routeId,
		ContinuousDropOff: &continuousDropOff,
	}
	routesWithWindows := map[string]bool{routeId: true}
	rules := &types.RoutesRules{
		ContinuousDropOff: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
		},
	}
	validations.ContinuousDropOffValidation(route, 1, &types.Gtfs{}, rules, routesWithWindows)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "continuous_drop_off=2 should be forbidden when route has windows")
}

func TestContinuousDropOffValidation_ForbiddenValue3WithWindows(t *testing.T) {
	services.AppMessageService.Clear()
	continuousDropOff := "3"
	routeId := "ROUTE1"
	route := &types.Route{
		RouteId:           &routeId,
		ContinuousDropOff: &continuousDropOff,
	}
	routesWithWindows := map[string]bool{routeId: true}
	rules := &types.RoutesRules{
		ContinuousDropOff: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
		},
	}
	validations.ContinuousDropOffValidation(route, 1, &types.Gtfs{}, rules, routesWithWindows)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "continuous_drop_off=3 should be forbidden when route has windows")
}

func TestContinuousDropOffValidation_ValidValue1WithWindows(t *testing.T) {
	services.AppMessageService.Clear()
	continuousDropOff := "1"
	routeId := "ROUTE1"
	route := &types.Route{
		RouteId:           &routeId,
		ContinuousDropOff: &continuousDropOff,
	}
	routesWithWindows := map[string]bool{routeId: true}
	rules := &types.RoutesRules{
		ContinuousDropOff: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
		},
	}
	validations.ContinuousDropOffValidation(route, 1, &types.Gtfs{}, rules, routesWithWindows)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "continuous_drop_off=1 should be valid even with windows")
}

func TestContinuousDropOffValidation_NoRouteId(t *testing.T) {
	services.AppMessageService.Clear()
	continuousDropOff := "0"
	route := &types.Route{
		RouteId:           nil,
		ContinuousDropOff: &continuousDropOff,
	}
	routesWithWindows := map[string]bool{"ROUTE1": true}
	rules := &types.RoutesRules{
		ContinuousDropOff: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
		},
	}
	validations.ContinuousDropOffValidation(route, 1, &types.Gtfs{}, rules, routesWithWindows)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "continuous_drop_off validation should skip window check when route_id is nil")
}

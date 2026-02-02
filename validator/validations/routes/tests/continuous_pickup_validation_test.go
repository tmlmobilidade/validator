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

func TestContinuousPickupValidation_MissingContinuousPickup(t *testing.T) {
	services.AppMessageService.Clear()
	routeId := "MY_ROUTE_ID"
	route := &types.Route{RouteId: &routeId, ContinuousPickup: nil, ContinuousDropOff: nil}
	gtfs := &types.Gtfs{}
	routesWithWindows := make(map[string]bool)
	validations.ContinuousPickupValidation(route, 1, gtfs, nil, routesWithWindows)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing continuous_pickup should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousPickupValidation_MissingRequiredContinuousPickup(t *testing.T) {
	services.AppMessageService.Clear()
	routeId := "MY_ROUTE_ID"
	route := &types.Route{RouteId: &routeId, ContinuousPickup: nil}
	gtfs := &types.Gtfs{}

	severity := types.SEVERITY_ERROR
	routesWithWindows := make(map[string]bool)
	validations.ContinuousPickupValidation(route, 1, gtfs, &types.RoutesRules{ContinuousPickup: types.RuleConfig{Severity: severity}}, routesWithWindows)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing required continuous_pickup should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousPickupValidation_ForbiddenValueWithPickupWindow(t *testing.T) {
	services.AppMessageService.Clear()
	continuousPickup := "2"
	continuousDropOff := "1"
	routeId := "MY_ROUTE_ID"
	route := &types.Route{RouteId: &routeId, ContinuousPickup: &continuousPickup, ContinuousDropOff: &continuousDropOff}
	// Simulate GTFS with a trip and stop_times with pickup window
	gtfs := &types.Gtfs{
		Trip: []types.TripRaw{
			{TripId: "MY_TRIP_ID"},
		},
		StopTime: []types.StopTimeRaw{
			{StartPickupDropOffWindow: "08:00:00", EndPickupDropOffWindow: "09:00:00"},
		},
		IdMap: map[string]map[string][]int{
			"trips": {
				"MY_ROUTE_ID": {0},
			},
			"stop_times": {
				"MY_TRIP_ID": {0},
			},
		},
	}
	severity := types.SEVERITY_ERROR
	routesWithWindows := make(map[string]bool)
	routesWithWindows[routeId] = true // Route has trips with pickup/dropoff windows
	validations.ContinuousPickupValidation(route, 2, gtfs, &types.RoutesRules{ContinuousPickup: types.RuleConfig{Severity: severity}}, routesWithWindows)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Forbidden continuous_pickup value with pickup window should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousPickupValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	continuousPickup := "1"
	routeId := "MY_ROUTE_ID"
	route := &types.Route{RouteId: &routeId, ContinuousPickup: &continuousPickup}
	// Simulate GTFS with a trip and stop_times without pickup window
	gtfs := &types.Gtfs{
		Trip: []types.TripRaw{
			{TripId: "MY_TRIP_ID"},
		},
		StopTime: []types.StopTimeRaw{
			{StartPickupDropOffWindow: "", EndPickupDropOffWindow: ""},
		},
		IdMap: map[string]map[string][]int{
			"trips": {
				"MY_ROUTE_ID": {0},
			},
			"stop_times": {
				"MY_TRIP_ID": {0},
			},
		},
	}
	severity := types.SEVERITY_ERROR
	routesWithWindows := make(map[string]bool)
	validations.ContinuousPickupValidation(route, 3, gtfs, &types.RoutesRules{ContinuousPickup: types.RuleConfig{Severity: severity}}, routesWithWindows)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid continuous_pickup with no pickup window should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousPickupValidation_ValidInputWithPickupWindow(t *testing.T) {
	services.AppMessageService.Clear()
	continuousPickup := "1"
	routeId := "MY_ROUTE_ID"
	route := &types.Route{RouteId: &routeId, ContinuousPickup: &continuousPickup}
	// Simulate GTFS with a trip and stop_times with pickup window
	gtfs := &types.Gtfs{
		Trip: []types.TripRaw{
			{TripId: "MY_TRIP_ID"},
		},
		StopTime: []types.StopTimeRaw{
			{StartPickupDropOffWindow: "08:00:00", EndPickupDropOffWindow: "09:00:00"},
		},
		IdMap: map[string]map[string][]int{
			"trips": {
				"MY_ROUTE_ID": {0},
			},
			"stop_times": {
				"MY_TRIP_ID": {0},
			},
		},
	}
	severity := types.SEVERITY_ERROR
	routesWithWindows := make(map[string]bool)
	routesWithWindows[routeId] = true // Route has trips with pickup/dropoff windows, but continuous_pickup is "1" so it returns early
	validations.ContinuousPickupValidation(route, 4, gtfs, &types.RoutesRules{ContinuousPickup: types.RuleConfig{Severity: severity}}, routesWithWindows)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid continuous_pickup with pickup window should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

// Test that continuous_pickup=0 is forbidden if the route has trips with pickup/dropoff windows
func TestAllContinuousPickupValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetContinuousPickupDropOffValidOptions()
	validOptionsStrings := make([]string, len(validOptions))
	for i, opt := range validOptions {
		validOptionsStrings[i] = fmt.Sprintf("%d", opt)
	}

	for _, tc := range test_helpers.GetGenericEnumIntTestCases("continuous_pickup", validOptions) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var severity types.Severity
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}

			// Convert tc.Value (*int or nil) to *string for ContinuousPickup field
			var continuousPickup *string
			if tc.Value != nil {
				if ptr, ok := tc.Value.(*int); ok && ptr != nil {
					valStr := fmt.Sprintf("%d", *ptr)
					continuousPickup = &valStr
				}
			}

			validations.ContinuousPickupValidation(
				&types.Route{ContinuousPickup: continuousPickup},
				tc.Row,
				&types.Gtfs{},
				&types.RoutesRules{
					ContinuousPickup: types.RuleConfig{
						Severity: severity,
						Options:  &validOptionsStrings,
					},
				},
				nil,
			)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
}

func TestContinuousPickupValidation_ForbiddenValueWithWindows(t *testing.T) {
	services.AppMessageService.Clear()
	continuousPickup := "0"
	routeId := "ROUTE1"
	route := &types.Route{
		RouteId:          &routeId,
		ContinuousPickup: &continuousPickup,
	}
	routesWithWindows := map[string]bool{routeId: true}
	rules := &types.RoutesRules{
		ContinuousPickup: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
		},
	}
	validations.ContinuousPickupValidation(route, 1, &types.Gtfs{}, rules, routesWithWindows)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "continuous_pickup=0 should be forbidden when route has windows")
}

func TestContinuousPickupValidation_ForbiddenValue2WithWindows(t *testing.T) {
	services.AppMessageService.Clear()
	continuousPickup := "2"
	routeId := "ROUTE1"
	route := &types.Route{
		RouteId:          &routeId,
		ContinuousPickup: &continuousPickup,
	}
	routesWithWindows := map[string]bool{routeId: true}
	rules := &types.RoutesRules{
		ContinuousPickup: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
		},
	}
	validations.ContinuousPickupValidation(route, 1, &types.Gtfs{}, rules, routesWithWindows)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "continuous_pickup=2 should be forbidden when route has windows")
}

func TestContinuousPickupValidation_ForbiddenValue3WithWindows(t *testing.T) {
	services.AppMessageService.Clear()
	continuousPickup := "3"
	routeId := "ROUTE1"
	route := &types.Route{
		RouteId:          &routeId,
		ContinuousPickup: &continuousPickup,
	}
	routesWithWindows := map[string]bool{routeId: true}
	rules := &types.RoutesRules{
		ContinuousPickup: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
		},
	}
	validations.ContinuousPickupValidation(route, 1, &types.Gtfs{}, rules, routesWithWindows)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "continuous_pickup=3 should be forbidden when route has windows")
}

func TestContinuousPickupValidation_ValidValue1WithWindows(t *testing.T) {
	services.AppMessageService.Clear()
	continuousPickup := "1"
	routeId := "ROUTE1"
	route := &types.Route{
		RouteId:          &routeId,
		ContinuousPickup: &continuousPickup,
	}
	routesWithWindows := map[string]bool{routeId: true}
	rules := &types.RoutesRules{
		ContinuousPickup: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
		},
	}
	validations.ContinuousPickupValidation(route, 1, &types.Gtfs{}, rules, routesWithWindows)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "continuous_pickup=1 should be valid even with windows")
}

func TestContinuousPickupValidation_NoRouteId(t *testing.T) {
	services.AppMessageService.Clear()
	continuousPickup := "0"
	route := &types.Route{
		RouteId:          nil,
		ContinuousPickup: &continuousPickup,
	}
	routesWithWindows := map[string]bool{"ROUTE1": true}
	rules := &types.RoutesRules{
		ContinuousPickup: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
		},
	}
	validations.ContinuousPickupValidation(route, 1, &types.Gtfs{}, rules, routesWithWindows)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "continuous_pickup validation should skip window check when route_id is nil")
}

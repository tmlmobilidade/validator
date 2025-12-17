package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestContinuousDropOffValidation_MissingContinuousDropOff(t *testing.T) {
	services.AppMessageService.Clear()
	routeId := "MY_ROUTE_ID"
	route := &types.Route{RouteId: &routeId, ContinuousDropOff: nil}
	gtfs := &types.Gtfs{}
	routesWithWindows := make(map[string]bool)
	validations.ContinuousDropOffValidation(route, 1, gtfs, nil, routesWithWindows)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing continuous_pickup should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousDropOffValidation_MissingRequiredContinuousDropOff(t *testing.T) {
	services.AppMessageService.Clear()
	routeId := "MY_ROUTE_ID"
	route := &types.Route{RouteId: &routeId, ContinuousDropOff: nil}
	gtfs := &types.Gtfs{}

	severity := types.SEVERITY_ERROR
	routesWithWindows := make(map[string]bool)
	validations.ContinuousDropOffValidation(route, 1, gtfs, &types.RoutesRules{ContinuousDropOff: types.RuleConfig{Severity: severity}}, routesWithWindows)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing required continuous_pickup should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousDropOffValidation_ForbiddenValueWithDropOffWindow(t *testing.T) {
	services.AppMessageService.Clear()
	continuousDropOff := "2"
	routeId := "MY_ROUTE_ID"
	route := &types.Route{RouteId: &routeId, ContinuousDropOff: &continuousDropOff}
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
	validations.ContinuousDropOffValidation(route, 2, gtfs, &types.RoutesRules{ContinuousDropOff: types.RuleConfig{Severity: severity}}, routesWithWindows)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Forbidden continuous_pickup value with pickup window should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousDropOffValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	continuousDropOff := "1"
	routeId := "MY_ROUTE_ID"
	route := &types.Route{RouteId: &routeId, ContinuousDropOff: &continuousDropOff}
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
	validations.ContinuousDropOffValidation(route, 3, gtfs, &types.RoutesRules{ContinuousDropOff: types.RuleConfig{Severity: severity}}, routesWithWindows)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid continuous_pickup with no pickup window should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContinuousDropOffValidation_ValidInputWithDropOffWindow(t *testing.T) {
	services.AppMessageService.Clear()
	continuousDropOff := "1"
	routeId := "MY_ROUTE_ID"
	route := &types.Route{RouteId: &routeId, ContinuousDropOff: &continuousDropOff}
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
	routesWithWindows[routeId] = true // Route has trips with pickup/dropoff windows, but continuous_drop_off is "1" so it returns early
	validations.ContinuousDropOffValidation(route, 4, gtfs, &types.RoutesRules{ContinuousDropOff: types.RuleConfig{Severity: severity}}, routesWithWindows)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid continuous_pickup with pickup window should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

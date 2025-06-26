package routes

import (
	"main/lib"
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
	validations.ContinuousPickupValidation(nil, route, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing continuous_pickup should not error",
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
	validations.ContinuousPickupValidation(&severity, route, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing required continuous_pickup should error",
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
	validations.ContinuousPickupValidation(&severity, route, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Forbidden continuous_pickup value with pickup window should error",
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
	validations.ContinuousPickupValidation(&severity, route, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid continuous_pickup with no pickup window should not error",
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
	validations.ContinuousPickupValidation(&severity, route, 4, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid continuous_pickup with pickup window should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
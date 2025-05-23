package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestRouteIdValidation_MissingRouteId(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteId: nil}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"routes": {}}}
	validations.RouteIdValidation(route, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing route_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteIdValidation_DuplicateRouteId(t *testing.T) {
	services.AppMessageService.Clear()
	routeId := "R1"
	route := &types.Route{RouteId: &routeId}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"routes": {"R1": {1, 2}}}}
	validations.RouteIdValidation(route, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Duplicate route_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteIdValidation_ValidRouteId(t *testing.T) {
	services.AppMessageService.Clear()
	routeId := "R2"
	route := &types.Route{RouteId: &routeId}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"routes": {"R2": {3}}}}
	validations.RouteIdValidation(route, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid route_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
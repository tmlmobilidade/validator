package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestRouteIdValidation_Required(t *testing.T) {
	trip := &types.Trip{RouteId: nil}
	gtfs := &types.Gtfs{}
	validations.RouteIdValidation(trip, 1, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Route ID is required",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestRouteIdValidation_ValidForeignKey(t *testing.T) {
	trip := &types.Trip{RouteId: lib.Ptr("route1")}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"routes": {"route1": {1}}}}
	validations.RouteIdValidation(trip, 2, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Route ID references a valid route_id, should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestRouteIdValidation_InvalidForeignKey(t *testing.T) {
	trip := &types.Trip{RouteId: lib.Ptr("invalid")}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"routes": {}}}
	validations.RouteIdValidation(trip, 3, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Route ID must reference a valid route_id from routes.txt.",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
} 
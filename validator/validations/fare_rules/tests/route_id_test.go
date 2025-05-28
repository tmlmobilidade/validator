package fare_rules

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_rules/validations"
	"testing"
)

func TestRouteIdValidation_MissingRouteId(t *testing.T) {
	services.AppMessageService.Clear()
	fareRule := &types.FareRule{RouteId: nil}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"routes": {},
		},
	}
	validations.RouteIdValidation(fareRule, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing route_id (optional) should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteIdValidation_InvalidRouteId(t *testing.T) {
	services.AppMessageService.Clear()
	invalidRouteId := "INVALID"
	fareRule := &types.FareRule{RouteId: &invalidRouteId}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"routes": {},
		},
	}
	validations.RouteIdValidation(fareRule, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid route_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteIdValidation_ValidRouteId(t *testing.T) {
	services.AppMessageService.Clear()
	validRouteId := "ROUTE1"
	fareRule := &types.FareRule{RouteId: &validRouteId}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"routes": {"ROUTE1": {1}},
		},
	}
	validations.RouteIdValidation(fareRule, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid route_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
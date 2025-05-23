package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestRouteTypeValidation_MissingType(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteType: nil}
	validations.RouteTypeValidation(route, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing route_type should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteTypeValidation_InvalidType(t *testing.T) {
	services.AppMessageService.Clear()
	typeVal := 99
	route := &types.Route{RouteType: &typeVal}
	validations.RouteTypeValidation(route, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid route_type should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteTypeValidation_ValidType(t *testing.T) {
	services.AppMessageService.Clear()
	validTypes := []int{0, 1, 2, 3, 4, 5, 6, 7, 11, 12}
	for i, v := range validTypes {
		route := &types.Route{RouteType: &v}
		validations.RouteTypeValidation(route, i+3)
		assertion := lib.AssertionMessage{
			Expected: 0,
			Actual: services.AppMessageService.GetSummary().TotalErrors,
			Message: "Valid route_type should not error",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Errorf("route_type %d: %s", v, assert)
		}
		services.AppMessageService.Clear()
	}
} 
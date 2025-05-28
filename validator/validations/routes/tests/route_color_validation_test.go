package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestRouteColorValidation_MissingColor(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteColor: nil}
	validations.RouteColorValidation(nil, route, 1)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing route_color should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteColorValidation_EmptyColor(t *testing.T) {
	services.AppMessageService.Clear()
	empty := ""
	route := &types.Route{RouteColor: &empty}
	validations.RouteColorValidation(nil, route, 2)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Empty route_color should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteColorValidation_InvalidColor(t *testing.T) {
	services.AppMessageService.Clear()
	color := "ZZZZZZ"
	route := &types.Route{RouteColor: &color}
	validations.RouteColorValidation(nil, route, 3)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid route_color should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteColorValidation_ValidColor(t *testing.T) {
	services.AppMessageService.Clear()
	color := "123ABC"
	route := &types.Route{RouteColor: &color}
	validations.RouteColorValidation(nil, route, 4)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid route_color should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 

func TestRouteColorValidation_NoColor_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteColor: nil}
	severity := types.SEVERITY_WARNING
	validations.RouteColorValidation(&severity, route, 5)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "No route_color should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteColorValidation_NoColor_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteColor: nil}
	severity := types.SEVERITY_ERROR
	validations.RouteColorValidation(&severity, route, 6)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "No route_color should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
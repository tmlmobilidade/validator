package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestRouteTextColorValidation_MissingColor(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteTextColor: nil}
	validations.RouteTextColorValidation(nil, route, 1)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing route_text_color should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteTextColorValidation_EmptyColor(t *testing.T) {
	services.AppMessageService.Clear()
	empty := ""
	route := &types.Route{RouteTextColor: &empty}
	validations.RouteTextColorValidation(nil, route, 2)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Empty route_text_color should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteTextColorValidation_InvalidColor(t *testing.T) {
	services.AppMessageService.Clear()
	color := "ZZZZZZ"
	route := &types.Route{RouteTextColor: &color}
	validations.RouteTextColorValidation(nil, route, 3)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid route_text_color should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteTextColorValidation_ValidColor(t *testing.T) {
	services.AppMessageService.Clear()
	color := "123ABC"
	route := &types.Route{RouteTextColor: &color}
	validations.RouteTextColorValidation(nil, route, 4)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid route_text_color should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 

func TestRouteTextColorValidation_NoColor_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteTextColor: nil}
	severity := types.SEVERITY_WARNING
	validations.RouteTextColorValidation(&severity, route, 5)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "No route_text_color should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteTextColorValidation_NoColor_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteTextColor: nil}
	severity := types.SEVERITY_ERROR
	validations.RouteTextColorValidation(&severity, route, 6)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "No route_text_color should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestRouteDescValidation_MissingDesc(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteDesc: nil}
	validations.RouteDescValidation(route, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Missing route_desc should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteDescValidation_DuplicateShortName(t *testing.T) {
	services.AppMessageService.Clear()
	desc := "32"
	shortName := "32"
	route := &types.Route{RouteDesc: &desc, RouteShortName: &shortName}
	validations.RouteDescValidation(route, 2, nil)
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("route_desc duplicating route_short_name should warn")
	}
}

func TestRouteDescValidation_DuplicateLongName(t *testing.T) {
	services.AppMessageService.Clear()
	desc := "Main Street Express"
	longName := "Main Street Express"
	route := &types.Route{RouteDesc: &desc, RouteLongName: &longName}
	validations.RouteDescValidation(route, 3, nil)
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("route_desc duplicating route_long_name should warn")
	}
}

func TestRouteDescValidation_UniqueDesc(t *testing.T) {
	services.AppMessageService.Clear()
	desc := "A trains operate between Inwood-207 St and Far Rockaway."
	shortName := "A"
	longName := "Inwood-Far Rockaway"
	route := &types.Route{RouteDesc: &desc, RouteShortName: &shortName, RouteLongName: &longName}
	validations.RouteDescValidation(route, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Unique route_desc should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteDescValidation_MissingDesc_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteDesc: nil}
	severity := types.SEVERITY_WARNING
	validations.RouteDescValidation(route, 5, &types.RoutesRules{RouteDesc: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing route_desc should warn")
	}
}

func TestRouteDescValidation_MissingDesc_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteDesc: nil}
	severity := types.SEVERITY_ERROR
	validations.RouteDescValidation(route, 6, &types.RoutesRules{RouteDesc: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalErrors != 1 {
		t.Error("Missing route_desc should error")
	}
}

package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestRouteUrlValidation_MissingUrl(t *testing.T) {
	services.AppMessageService.Clear()

	route := &types.Route{RouteUrl: nil}
	validations.RouteUrlValidation(route, 1, nil, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Missing route_url should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteUrlValidation_InvalidUrl(t *testing.T) {
	services.AppMessageService.Clear()

	url := "THIS IS NOT A URL"
	route := &types.Route{RouteUrl: &url}

	validations.RouteUrlValidation(route, 2, nil, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid route_url should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteUrlValidation_SameAsAgencyUrl(t *testing.T) {
	services.AppMessageService.Clear()

	agencyUrl := "https://agency.com"
	agencyId := "my_agency_id"

	route := &types.Route{RouteUrl: &agencyUrl, AgencyId: &agencyId}
	gtfs := &types.Gtfs{
		Agency: []types.AgencyRaw{
			{AgencyId: agencyId, AgencyUrl: agencyUrl},
		},
		IdMap: types.GtfsIdMap{
			"agency": {
				agencyId: {0},
			},
		},
	}

	validations.RouteUrlValidation(route, 3, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "route_url same as agency_url should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteUrlValidation_ValidDifferentUrl(t *testing.T) {
	services.AppMessageService.Clear()
	url := "https://route.com"
	route := &types.Route{RouteUrl: &url}
	validations.RouteUrlValidation(route, 4, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Valid route_url different from agency_url should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteUrlValidation_MissingUrl_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteUrl: nil}
	severity := types.SEVERITY_WARNING
	validations.RouteUrlValidation(route, 5, nil, &types.RoutesRules{RouteUrl: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing route_url should warn")
	}
}

func TestRouteUrlValidation_MissingUrl_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteUrl: nil}
	severity := types.SEVERITY_ERROR
	validations.RouteUrlValidation(route, 6, nil, &types.RoutesRules{RouteUrl: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalErrors != 1 {
		t.Error("Missing route_url should error")
	}
}

func TestRouteUrlValidation_RuleOptions_Valid(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteUrl: lib.Ptr("https://example.com")}
	validations.RouteUrlValidation(route, 7, nil, &types.RoutesRules{RouteUrl: types.RuleConfig{Options: &[]string{"https://example.com"}}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid route_url should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteUrlValidation_RuleOptions_Invalid_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteUrl: lib.Ptr("https://not-in-list.com")}
	validations.RouteUrlValidation(route, 8, nil, &types.RoutesRules{RouteUrl: types.RuleConfig{Options: &[]string{"https://route.com", "https://route2.com"}, Severity: types.SEVERITY_ERROR}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid route_url should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteUrlValidation_RuleOptions_Invalid_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteUrl: lib.Ptr("https://not-in-list.com")}
	validations.RouteUrlValidation(route, 8, nil, &types.RoutesRules{RouteUrl: types.RuleConfig{Options: &[]string{"https://route.com", "https://route2.com"}, Severity: types.SEVERITY_WARNING}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Invalid route_url should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

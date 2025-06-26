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
	validations.RouteUrlValidation(nil, route, 1, nil)
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Missing route_url should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteUrlValidation_InvalidUrl(t *testing.T) {
	services.AppMessageService.Clear()
	
	url := "THIS IS NOT A URL"
	route := &types.Route{RouteUrl: &url}
	
	validations.RouteUrlValidation(nil, route, 2, nil)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid route_url should error",
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
	
	validations.RouteUrlValidation(nil, route, 3, gtfs)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "route_url same as agency_url should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteUrlValidation_ValidDifferentUrl(t *testing.T) {
	services.AppMessageService.Clear()
	url := "https://route.com"
	route := &types.Route{RouteUrl: &url}
	validations.RouteUrlValidation(nil, route, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Valid route_url different from agency_url should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 

func TestRouteUrlValidation_MissingUrl_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteUrl: nil}
	severity := types.SEVERITY_WARNING
	validations.RouteUrlValidation(&severity, route, 5, nil)
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing route_url should warn")
	}
}

func TestRouteUrlValidation_MissingUrl_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteUrl: nil}
	severity := types.SEVERITY_ERROR
	validations.RouteUrlValidation(&severity, route, 6, nil)
	if services.AppMessageService.GetSummary().TotalErrors != 1 {
		t.Error("Missing route_url should error")
	}
}
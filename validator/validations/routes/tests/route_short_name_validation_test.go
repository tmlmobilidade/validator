package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestRouteShortNameValidation_BothNamesMissing(t *testing.T) {
	services.AppMessageService.Clear()
	
	route := &types.Route{RouteShortName: nil, RouteLongName: nil}
	
	validations.RouteShortNameValidation(nil, route, 1)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing both names should produce a warning",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteShortNameValidation_MissingShortName_LongNamePresent(t *testing.T) {
	services.AppMessageService.Clear()
	
	longName := "Long Route Name"
	route := &types.Route{RouteShortName: nil, RouteLongName: &longName}
	
	validations.RouteShortNameValidation(nil, route, 2)
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing short name with long name present should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	
}

func TestRouteShortNameValidation_MissingShortName_LongNameMissing(t *testing.T) {
	services.AppMessageService.Clear()
	
	route := &types.Route{RouteShortName: nil}
	
	validations.RouteShortNameValidation(nil, route, 3)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing short name with empty long name should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteShortNameValidation_ShortNamePresent_LongNameMissing(t *testing.T) {
	services.AppMessageService.Clear()
	shortName := "32"
	longName := ""
	route := &types.Route{RouteShortName: &shortName, RouteLongName: &longName}
	validations.RouteShortNameValidation(nil, route, 4)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Short name present, long name missing should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteShortNameValidation_ShortNameTooLong(t *testing.T) {
	services.AppMessageService.Clear()
	shortName := "1234567890123" // 13 chars
	longName := ""
	route := &types.Route{RouteShortName: &shortName, RouteLongName: &longName}
	validations.RouteShortNameValidation(nil, route, 5)
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Short name > 12 chars should warn")
	}
}

func TestRouteShortNameValidation_ShortNameValidLength(t *testing.T) {
	services.AppMessageService.Clear()
	shortName := "Green"
	longName := ""
	route := &types.Route{RouteShortName: &shortName, RouteLongName: &longName}
	validations.RouteShortNameValidation(nil, route, 6)
	if services.AppMessageService.GetSummary().TotalWarnings != 0 {
		t.Error("Short name <= 12 chars should not warn")
	}
}

func TestRouteShortNameValidation_MissingShortName_LongNamePresent_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	longName := "Long Route Name"
	route := &types.Route{RouteShortName: nil, RouteLongName: &longName}
	severity := types.SEVERITY_ERROR
	validations.RouteShortNameValidation(&severity, route, 7)
	if services.AppMessageService.GetSummary().TotalErrors != 1 {
		t.Error("Missing short name with long name present should error")
	}
}

func TestRouteShortNameValidation_MissingShortName_LongNamePresent_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	longName := "Long Route Name"
	route := &types.Route{RouteShortName: nil, RouteLongName: &longName}
	severity := types.SEVERITY_WARNING
	validations.RouteShortNameValidation(&severity, route, 8)
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing short name with long name present should warn")
	}
}
package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestRouteLongNameValidation_BothNamesMissing(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteShortName: nil, RouteLongName: nil}
	validations.RouteLongNameValidation(route, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing both names should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteLongNameValidation_MissingLongName_ShortNamePresent(t *testing.T) {
	services.AppMessageService.Clear()
	shortName := "32"
	route := &types.Route{RouteShortName: &shortName, RouteLongName: nil}
	validations.RouteLongNameValidation(route, 2, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing long name with short name present should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteLongNameValidation_MissingLongName_ShortNameMissing(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteShortName: nil, RouteLongName: nil}
	validations.RouteLongNameValidation(route, 3, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing long name and short name should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteLongNameValidation_LongNamePresent_ShortNameMissing(t *testing.T) {
	services.AppMessageService.Clear()
	longName := "Main Street Express"
	route := &types.Route{RouteShortName: nil, RouteLongName: &longName}
	validations.RouteLongNameValidation(route, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Long name present, short name missing should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteLongNameValidation_BothNamesPresent(t *testing.T) {
	services.AppMessageService.Clear()
	shortName := "32"
	longName := "Main Street Express"
	route := &types.Route{RouteShortName: &shortName, RouteLongName: &longName}
	validations.RouteLongNameValidation(route, 5, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Both names present should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

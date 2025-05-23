package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestRouteSortOrderValidation_Missing(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteSortOrder: nil}
	validations.RouteSortOrderValidation(nil, route, 1)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing route_sort_order should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteSortOrderValidation_Negative(t *testing.T) {
	services.AppMessageService.Clear()
	val := -1
	route := &types.Route{RouteSortOrder: &val}
	validations.RouteSortOrderValidation(nil, route, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Negative route_sort_order should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteSortOrderValidation_Zero(t *testing.T) {
	services.AppMessageService.Clear()
	val := 0
	route := &types.Route{RouteSortOrder: &val}
	validations.RouteSortOrderValidation(nil, route, 3)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Zero route_sort_order should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteSortOrderValidation_Positive(t *testing.T) {
	services.AppMessageService.Clear()
	val := 5
	route := &types.Route{RouteSortOrder: &val}
	validations.RouteSortOrderValidation(nil, route, 4)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Positive route_sort_order should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 

func TestRouteSortOrderValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteSortOrder: nil}
	severity := types.SEVERITY_WARNING
	validations.RouteSortOrderValidation(&severity, route, 5)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "No route_sort_order should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteSortOrderValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteSortOrder: nil}
	severity := types.SEVERITY_ERROR
	validations.RouteSortOrderValidation(&severity, route, 6)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "No route_sort_order should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
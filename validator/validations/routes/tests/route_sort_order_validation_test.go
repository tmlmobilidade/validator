package routes

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestRouteSortOrderValidation_Negative(t *testing.T) {
	services.AppMessageService.Clear()
	val := -1
	route := &types.Route{RouteSortOrder: &val}
	validations.RouteSortOrderValidation(route, 2, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Negative route_sort_order should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRouteSortOrderValidation_Zero(t *testing.T) {
	services.AppMessageService.Clear()
	val := 0
	route := &types.Route{RouteSortOrder: &val}
	validations.RouteSortOrderValidation(route, 3, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Zero route_sort_order should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAllRouteSortOrderValidationTestCases(t *testing.T) {
	negativeValue := -1
	zeroValue := 0
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("route_sort_order") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var severity types.Severity
			var rules *types.RoutesRules
			if tc.Name == "Recommended_Missing" {
				rules = nil
			} else {
				if tc.ExpectedWarnings > 0 {
					severity = types.SEVERITY_WARNING
				} else {
					severity = types.SEVERITY_ERROR
				}
				rules = &types.RoutesRules{RouteSortOrder: types.RuleConfig{Severity: severity}}
			}

			var routeSortOrder *int
			if tc.Value != nil {
				if tc.Name == "Valid_Present" {
					val := 1
					routeSortOrder = &val
				} else {
					routeSortOrder = &zeroValue
				}
			}

			if tc.Name == "Invalid_Value" {
				routeSortOrder = &negativeValue
			}

			validations.RouteSortOrderValidation(&types.Route{RouteSortOrder: routeSortOrder}, tc.Row, rules)
			if tc.Name == "Recommended_Missing" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name, types.SEVERITY_WARNING)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
			}
		})
	}

	for _, tc := range test_helpers.GetGenericSeverityTestCases("route_sort_order") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var routeSortOrder *int
			if tc.Value != nil {
				// For forbidden present case, use a valid integer value
				val := 1
				routeSortOrder = &val
			}

			validations.RouteSortOrderValidation(&types.Route{RouteSortOrder: routeSortOrder}, tc.Row, &types.RoutesRules{RouteSortOrder: types.RuleConfig{Severity: tc.Severity}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name, types.SEVERITY_WARNING)
		})
	}
}

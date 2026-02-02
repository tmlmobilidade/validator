package routes

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"strconv"
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
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("route_sort_order") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var severity types.Severity
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}

			var routeSortOrder *int
			if tc.Value != nil {
				// Convert string pointer to int pointer if value is set
				// For "Valid_Present" case, use a valid integer value
				if tc.Name == "Valid_Present" {
					val := 1
					routeSortOrder = &val
				} else {
					val, err := strconv.Atoi(*tc.Value)
					if err == nil {
						routeSortOrder = &val
					}
				}
			}

			validations.RouteSortOrderValidation(&types.Route{RouteSortOrder: routeSortOrder}, tc.Row, &types.RoutesRules{RouteSortOrder: types.RuleConfig{Severity: severity}})
			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
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
			if tc.ExpectedErrors > 0 {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name)
			}
		})
	}
}

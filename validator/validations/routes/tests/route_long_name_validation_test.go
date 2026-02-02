package routes

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestAllRouteLongNameValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("route_long_name") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var severity types.Severity
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}

			var routeLongName *string
			if tc.Value != nil {
				routeLongName = tc.Value
			}

			validations.RouteLongNameValidation(&types.Route{RouteLongName: routeLongName}, tc.Row, &types.RoutesRules{RouteLongName: types.RuleConfig{Severity: severity}})
			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
		})
	}
}

func TestRouteLongNameValidation_MissingLongNameAndShortName(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{RouteShortName: nil, RouteLongName: nil}
	validations.RouteLongNameValidation(route, 1, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Missing long name and short name should error")
}

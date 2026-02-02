package routes

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestAllRouteColorValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericColorTestCases("route_color") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var rules *types.RoutesRules
			// For Nil_Color_Optional, don't pass rules (field is truly optional)
			// For other cases, set severity based on expected errors
			if tc.Name != "Nil_Color_Optional" {
				var severity types.Severity
				if tc.ExpectedErrors > 0 {
					severity = types.SEVERITY_ERROR
				} else {
					severity = types.SEVERITY_WARNING
				}
				rules = &types.RoutesRules{RouteColor: types.RuleConfig{Severity: severity}}
			}

			validations.RouteColorValidation(&types.Route{RouteColor: tc.Color}, tc.Row, rules)

			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
		})
	}
}

package routes

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestAllRouteTextColorValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericColorTestCases("route_text_color") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var severity types.Severity
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}

			validations.RouteTextColorValidation(&types.Route{RouteTextColor: tc.Color}, tc.Row, &types.RoutesRules{RouteTextColor: types.RuleConfig{Severity: severity}})
			if tc.ExpectedErrors > 0 {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name)
			}
		})
	}
}

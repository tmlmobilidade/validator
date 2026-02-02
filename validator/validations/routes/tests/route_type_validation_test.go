package routes

import (
	"fmt"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestAllRouteTypeValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetRouteTypeValidOptions()
	validOptionsStrings := make([]string, len(validOptions))
	for i, opt := range validOptions {
		validOptionsStrings[i] = fmt.Sprintf("%d", opt)
	}
	for _, tc := range test_helpers.GetGenericEnumIntTestCases("route_type", validOptions) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var severity types.Severity
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}

			var routeType *int
			if tc.Value != nil {
				if ptr, ok := tc.Value.(*int); ok && ptr != nil {
					routeType = ptr
				}
			}

			validations.RouteTypeValidation(&types.Route{RouteType: routeType}, tc.Row, &types.RoutesRules{RouteType: types.RuleConfig{Severity: severity, Options: &validOptionsStrings}})
			if tc.ExpectedErrors > 0 {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name)
			}
		})
	}
}

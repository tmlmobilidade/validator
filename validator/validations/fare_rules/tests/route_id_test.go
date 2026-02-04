package fare_rules

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/fare_rules/validations"
	"testing"
)

func TestAllRouteIdValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericForeignKeyTestCases("route_id") {
		if tc.Name == "ForeignKey_Invalid" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var routeId *string
			if tc.Id != nil {
				routeId = tc.Id
			}
			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"routes": map[string][]int{*routeId: {1}}}}.ToGtfs()
			validations.RouteIdValidation(&types.FareRule{RouteId: routeId}, tc.Row, &gtfs, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
		})
	}
	for _, tc := range test_helpers.GetGenericSeverityTestCases("route_id") {
		if tc.Name != "Severity_Ignore_Missing" && tc.Name != "Severity_Forbidden_Missing" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"routes": {"MY_ROUTE_ID": {1}}}}.ToGtfs()
			validations.RouteIdValidation(&types.FareRule{RouteId: tc.Value.(*string)}, tc.Row, &gtfs, &types.FareRulesRules{RouteId: types.RuleConfig{Severity: tc.Severity}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name, types.SEVERITY_WARNING)
		})
	}
}

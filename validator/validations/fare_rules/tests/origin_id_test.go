package fare_rules

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/fare_rules/validations"
	"testing"
)

func TestAllOriginIdValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericForeignKeyTestCases("origin_id") {
		if tc.Name == "ForeignKey_Invalid" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var originId *string
			if tc.Id != nil {
				originId = tc.Id
			}
			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": map[string][]int{*originId: {1}}}}.ToGtfs()
			validations.OriginIdValidation(&types.FareRule{OriginId: originId}, tc.Row, &gtfs, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
	for _, tc := range test_helpers.GetGenericSeverityTestCases("origin_id") {
		if tc.Name != "Severity_Ignore_Missing" && tc.Name != "Severity_Forbidden_Missing" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": {"MY_STOP_ID": {1}}}}.ToGtfs()
			validations.OriginIdValidation(&types.FareRule{OriginId: tc.Value.(*string)}, tc.Row, &gtfs, &types.FareRulesRules{OriginId: types.RuleConfig{Severity: tc.Severity}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
}

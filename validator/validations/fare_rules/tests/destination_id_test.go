package fare_rules

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/fare_rules/validations"
	"testing"
)

func TestAllDestinationIdValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericForeignKeyTestCases("destination_id") {
		if tc.Name == "ForeignKey_Invalid" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var destinationId *string
			if tc.Id != nil {
				destinationId = tc.Id
			}

			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": map[string][]int{*destinationId: {1}}}}.ToGtfs()
			validations.DestinationIdValidation(&types.FareRule{DestinationId: destinationId}, tc.Row, &gtfs, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
	for _, tc := range test_helpers.GetGenericSeverityTestCases("destination_id") {
		if tc.Name != "Severity_Ignore_Missing" && tc.Name != "Severity_Forbidden_Missing" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var destinationId *string
			if tc.Value != nil {
				destinationId = tc.Value.(*string)
			} else {
				destinationId = nil
			}
			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": {"DEST1": {1}}}}.ToGtfs()
			validations.DestinationIdValidation(&types.FareRule{DestinationId: destinationId}, tc.Row, &gtfs, &types.FareRulesRules{DestinationId: types.RuleConfig{Severity: tc.Severity}})
			expectedTotalMessages := tc.ExpectedErrors
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
		})
	}
}

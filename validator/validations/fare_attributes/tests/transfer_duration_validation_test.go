package fare_attributes

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/fare_attributes/validations"
	"testing"
)

func TestAllTransferDurationValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetTransferDurationValidOptions()
	negativeTransferDuration := -1
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("transfer_duration") {
		if tc.Name == "Recommended_Missing" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var transferDuration *int
			if tc.Value != nil {
				transferDuration = &validOptions[tc.Row-1]
			}

			if tc.Name == "Invalid_Value" {
				transferDuration = &negativeTransferDuration
			}

			if tc.Name == "Required" {
				transferDuration = nil
			}

			var rules *types.FareAttributesRules
			if tc.Name == "Required" {
				rules = &types.FareAttributesRules{TransferDuration: types.RuleConfig{Severity: types.SEVERITY_ERROR}}
			} else {
				rules = nil
			}

			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"fare_attributes": map[string][]int{"transfer_duration": []int{1}}}}.ToGtfs()
			validations.TransferDurationValidation(&types.FareAttribute{TransferDuration: transferDuration}, tc.Row, &gtfs, rules)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}

	for _, tc := range test_helpers.GetGenericSeverityTestCases("transfer_duration") {
		if tc.Name != "Severity_Ignore_Missing" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"fare_attributes": map[string][]int{"transfer_duration": []int{1}}}}.ToGtfs()
			validations.TransferDurationValidation(&types.FareAttribute{TransferDuration: &validOptions[tc.Row-1]}, tc.Row, &gtfs, nil)
			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
		})
	}
	t.Run("NegativeTransferDuration", func(t *testing.T) {
		services.AppMessageService.Clear()
		gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"fare_attributes": map[string][]int{"transfer_duration": []int{1}}}}.ToGtfs()
		validations.TransferDurationValidation(&types.FareAttribute{TransferDuration: &negativeTransferDuration}, 1, &gtfs, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Negative transfer duration should error")
	})
}

package fare_rules

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/fare_rules/validations"
	"testing"
)

func TestAllFareIdValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("fare_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var fareId *string
			if tc.Value != nil {
				fareId = tc.Value
			}
			fareRule := &types.FareRule{FareId: fareId}
			gtfs := &types.Gtfs{
				IdMap: types.GtfsIdMap{
					"fare_attributes": map[string][]int{},
				},
			}
			if fareId != nil && tc.Name == "Valid_Present" {
				gtfs.IdMap["fare_attributes"][*fareId] = []int{1}
			}
			validations.FareIdValidation(fareRule, tc.Row, gtfs, nil)
			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
		})
	}
}

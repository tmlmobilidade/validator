package fare_attributes

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/fare_attributes/validations"
	"testing"
)

func TestAllFareIdValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericIdTestCases("fare_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"fare_attributes": tc.ExistingIds}}.ToGtfs()
			var fareId *string
			if tc.Id != nil {
				fareId = tc.Id
			}
			fareAttribute := &types.FareAttribute{FareId: fareId}
			validations.FareIdValidation(fareAttribute, tc.Row, &gtfs)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
		})
	}
}

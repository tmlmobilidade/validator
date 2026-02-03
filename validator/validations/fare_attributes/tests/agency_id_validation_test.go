package fare_attributes

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/fare_attributes/validations"
	"testing"
)

func TestAllAgencyIdValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericIdTestCases("agency_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			agencyIdMap := make(map[string][]int)
			if tc.ExistingIds != nil {
				agencyIdMap = tc.ExistingIds
			}
			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"agency": agencyIdMap}}.ToGtfs()
			validations.AgencyIdValidation(&types.FareAttribute{AgencyId: tc.Id}, tc.Row, &gtfs, &types.FareAttributesRules{AgencyId: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("agency_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var severity types.Severity
			if tc.Name == "Recommended_Missing" {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}

			var agencyId *string
			if tc.Value != nil {
				agencyId = tc.Value
			}

			if tc.Name == "Invalid_Value" {
				agencyId = nil
			}

			agencyIdMap := make(map[string][]int)
			if tc.Value != nil && *tc.Value != "" {
				agencyIdMap[*tc.Value] = []int{1}
			}
			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"agency": agencyIdMap}}.ToGtfs()
			validations.AgencyIdValidation(&types.FareAttribute{AgencyId: agencyId}, tc.Row, &gtfs, &types.FareAttributesRules{AgencyId: types.RuleConfig{Severity: severity}})
			expectedTotal := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotal, tc.Name)
		})
	}
}

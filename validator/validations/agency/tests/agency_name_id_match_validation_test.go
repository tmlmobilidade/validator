package agency

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAllAgencyNameIdMatchValidationTestCases(t *testing.T) {
	fieldName := "agency_name_id_match"

	compare := []types.Compare{
		{Key: "AGENCY1", Value: "Metro Transit"},
		{Key: "AGENCY2", Value: "City Bus"},
		{Key: "AGENCY3", Value: "Regional Rail"},
	}

	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases(fieldName) {
		if tc.Name == "Invalid_Id" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var severity types.Severity
			if tc.Name == "Recommended_Missing" {
				severity = types.SEVERITY_WARNING
			} else {
				severity = types.SEVERITY_ERROR
			}

			var agencyId *string
			var agencyName *string
			if tc.Name == "Invalid_Value" {
				agencyId = nil
				agencyName = nil
			} else {
				agencyId = tc.Value
				agencyName = tc.Value
			}

			validations.AgencyNameIdMatchValidation(&types.Agency{AgencyId: agencyId, AgencyName: agencyName}, tc.Row, &types.AgencyRules{AgencyNameIdMatch: types.RuleConfig{Severity: severity}})
			if tc.Name == "Recommended_Missing" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, 1, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}
	for _, tc := range test_helpers.GetGenericSeverityTestCases(fieldName) {
		if tc.Name != "Severity_Ignore_Missing" && tc.Name != "Severity_Forbidden_Missing" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			validations.AgencyNameIdMatchValidation(&types.Agency{AgencyId: lib.Ptr("AGENCY1"), AgencyName: lib.Ptr("Metro Transit")}, tc.Row, &types.AgencyRules{AgencyNameIdMatch: types.RuleConfig{Severity: tc.Severity}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}

	t.Run("WithCompare_Match", func(t *testing.T) {
		services.AppMessageService.Clear()
		validations.AgencyNameIdMatchValidation(&types.Agency{AgencyId: lib.Ptr("AGENCY1"), AgencyName: lib.Ptr("Metro Transit")}, 1, &types.AgencyRules{AgencyNameIdMatch: types.RuleConfig{Severity: types.SEVERITY_ERROR, Compare: &compare}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Matching agency_id and agency_name should not error")
	})
	t.Run("WithCompare_NoMatch", func(t *testing.T) {
		services.AppMessageService.Clear()
		validations.AgencyNameIdMatchValidation(&types.Agency{AgencyId: lib.Ptr("AGENCY1"), AgencyName: lib.Ptr("Wrong Name")}, 1, &types.AgencyRules{AgencyNameIdMatch: types.RuleConfig{Severity: types.SEVERITY_ERROR, Compare: &compare}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Non-matching agency_id and agency_name should error")
	})
	t.Run("WithCompare_NoMatch_Warning", func(t *testing.T) {
		services.AppMessageService.Clear()
		validations.AgencyNameIdMatchValidation(&types.Agency{AgencyId: lib.Ptr("AGENCY1"), AgencyName: lib.Ptr("Wrong Name")}, 1, &types.AgencyRules{AgencyNameIdMatch: types.RuleConfig{Severity: types.SEVERITY_WARNING, Compare: &compare}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Non-matching agency_id and agency_name should warn")
	})
	t.Run("WithCompare_MultipleMatches", func(t *testing.T) {
		services.AppMessageService.Clear()
		validations.AgencyNameIdMatchValidation(&types.Agency{AgencyId: lib.Ptr("AGENCY1"), AgencyName: lib.Ptr("Metro Transit")}, 1, &types.AgencyRules{AgencyNameIdMatch: types.RuleConfig{Severity: types.SEVERITY_ERROR, Compare: &compare}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Matching agency_id and agency_name should not error even with multiple entries")
	})
}

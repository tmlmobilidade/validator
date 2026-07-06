package agency

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAllAgencyNameValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("agency_name") {
		if tc.Name == "Invalid_Value" || tc.Name == "Recommended_Missing" {
			continue
		}

		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			agency := &types.Agency{AgencyName: tc.Value}
			severity := types.SEVERITY_WARNING
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			}

			validations.AgencyNameValidation(agency, tc.Row, &types.AgencyRules{AgencyName: types.RuleConfig{Severity: severity}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name, types.SEVERITY_WARNING)
		})
	}

	t.Run("MissingNameAlwaysErrors", func(t *testing.T) {
		services.AppMessageService.Clear()
		agency := &types.Agency{AgencyName: nil}

		validations.AgencyNameValidation(agency, 2, &types.AgencyRules{AgencyName: types.RuleConfig{Severity: types.SEVERITY_WARNING}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "MissingNameAlwaysErrors", types.SEVERITY_ERROR)
	})

	t.Run("NotAllowedOption", func(t *testing.T) {
		services.AppMessageService.Clear()
		name := "Other Agency"
		options := []string{"Allowed Agency"}
		agency := &types.Agency{AgencyName: &name}

		validations.AgencyNameValidation(agency, 1, &types.AgencyRules{AgencyName: types.RuleConfig{Severity: types.SEVERITY_ERROR, Options: &options}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "NotAllowedOption", types.SEVERITY_ERROR)
	})

	t.Run("AllowedOption", func(t *testing.T) {
		services.AppMessageService.Clear()
		name := "Allowed Agency"
		options := []string{name}
		agency := &types.Agency{AgencyName: &name}

		validations.AgencyNameValidation(agency, 1, &types.AgencyRules{AgencyName: types.RuleConfig{Severity: types.SEVERITY_ERROR, Options: &options}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "AllowedOption", types.SEVERITY_ERROR)
	})
}

package agency

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAllAgencyPhoneValidationTestCases(t *testing.T) {
	fieldName := "agency_phone"

	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases(fieldName) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var severity types.Severity
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}

			var agencyPhone *string
			if tc.Name == "Valid_Present" {
				value := test_helpers.GetValidPhoneNumbers()[0]
				agencyPhone = &value
			} else {
				agencyPhone = tc.Value
			}

			validations.AgencyPhoneValidation(&types.Agency{AgencyPhone: agencyPhone}, tc.Row, &types.AgencyRules{AgencyPhone: types.RuleConfig{Severity: severity}})
			if tc.Name == "Recommended_Missing" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}
}

package agency

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAllAgencyTimezoneValidationTestCases(t *testing.T) {
	fieldName := "agency_timezone"

	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases(fieldName) {
		// Skip Recommended_Missing - agency_timezone is always required per GTFS spec
		if tc.Name == "Recommended_Missing" {
			continue
		}

		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var severity types.Severity
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}

			var agencyTimezone *string
			if tc.Name == "Valid_Present" {
				value := test_helpers.GetValidTimezones()[0]
				agencyTimezone = &value
			} else {
				agencyTimezone = tc.Value
			}

			validations.AgencyTimezoneValidation(&types.Agency{AgencyTimezone: agencyTimezone}, tc.Row, &types.AgencyRules{AgencyTimezone: types.RuleConfig{Severity: severity}})
			if tc.Name == "Recommended_Missing" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, 1, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}
}

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

	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases(fieldName) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var severity types.Severity
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}

			validations.AgencyNameIdMatchValidation(&types.Agency{AgencyId: tc.Value, AgencyName: tc.Value}, tc.Row, &types.AgencyRules{AgencyNameIdMatch: types.RuleConfig{Severity: severity}})
			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
		})
	}
}

func TestAgencyNameIdMatchValidation_WithCompare_Match(t *testing.T) {
	services.AppMessageService.Clear()
	compare := []types.Compare{
		{Key: "AGENCY1", Value: "Metro Transit"},
		{Key: "AGENCY2", Value: "City Bus"},
		{Key: "AGENCY3", Value: "Regional Rail"},
	}
	agency := &types.Agency{
		AgencyId:   lib.Ptr("AGENCY1"),
		AgencyName: lib.Ptr("Metro Transit"),
	}
	rules := &types.AgencyRules{
		AgencyNameIdMatch: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
			Compare:  &compare,
		},
	}
	validations.AgencyNameIdMatchValidation(agency, 1, rules)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Matching agency_id and agency_name should not error")
}

func TestAgencyNameIdMatchValidation_WithCompare_NoMatch(t *testing.T) {
	services.AppMessageService.Clear()
	compare := []types.Compare{
		{Key: "AGENCY1", Value: "Metro Transit"},
		{Key: "AGENCY2", Value: "City Bus"},
	}
	agency := &types.Agency{
		AgencyId:   lib.Ptr("AGENCY1"),
		AgencyName: lib.Ptr("Wrong Name"),
	}
	rules := &types.AgencyRules{
		AgencyNameIdMatch: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
			Compare:  &compare,
		},
	}
	validations.AgencyNameIdMatchValidation(agency, 1, rules)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Non-matching agency_id and agency_name should error")
}

func TestAgencyNameIdMatchValidation_WithCompare_NoMatch_Warning(t *testing.T) {
	services.AppMessageService.Clear()
	compare := []types.Compare{
		{Key: "AGENCY1", Value: "Metro Transit"},
	}
	agency := &types.Agency{
		AgencyId:   lib.Ptr("AGENCY1"),
		AgencyName: lib.Ptr("Wrong Name"),
	}
	rules := &types.AgencyRules{
		AgencyNameIdMatch: types.RuleConfig{
			Severity: types.SEVERITY_WARNING,
			Compare:  &compare,
		},
	}
	validations.AgencyNameIdMatchValidation(agency, 1, rules)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Non-matching agency_id and agency_name should warn")
}

func TestAgencyNameIdMatchValidation_WithCompare_MultipleMatches(t *testing.T) {
	services.AppMessageService.Clear()
	compare := []types.Compare{
		{Key: "AGENCY1", Value: "Metro Transit"},
		{Key: "AGENCY2", Value: "City Bus"},
		{Key: "AGENCY3", Value: "Metro Transit"}, // Same name, different ID
	}
	agency := &types.Agency{
		AgencyId:   lib.Ptr("AGENCY1"),
		AgencyName: lib.Ptr("Metro Transit"),
	}
	rules := &types.AgencyRules{
		AgencyNameIdMatch: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
			Compare:  &compare,
		},
	}
	validations.AgencyNameIdMatchValidation(agency, 1, rules)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Matching agency_id and agency_name should not error even with multiple entries")
}

func TestAgencyNameIdMatchValidation_Forbidden(t *testing.T) {
	services.AppMessageService.Clear()
	agency := &types.Agency{
		AgencyId:   lib.Ptr("AGENCY1"),
		AgencyName: lib.Ptr("Metro Transit"),
	}
	rules := &types.AgencyRules{
		AgencyNameIdMatch: types.RuleConfig{
			Severity: types.SEVERITY_FORBIDDEN,
		},
	}
	validations.AgencyNameIdMatchValidation(agency, 1, rules)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Forbidden severity should error when both fields are present")
}

func TestAgencyNameIdMatchValidation_Ignore(t *testing.T) {
	services.AppMessageService.Clear()
	agency := &types.Agency{
		AgencyId:   lib.Ptr("AGENCY1"),
		AgencyName: lib.Ptr("Metro Transit"),
	}
	rules := &types.AgencyRules{
		AgencyNameIdMatch: types.RuleConfig{
			Severity: types.SEVERITY_IGNORE,
		},
	}
	validations.AgencyNameIdMatchValidation(agency, 1, rules)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "IGNORE severity should skip validation")
}

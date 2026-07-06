package stops

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestAllShelterMaintainerValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("shelter_maintainer") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			stop := &types.Stop{ShelterMaintainer: tc.Value}
			severity := types.SEVERITY_WARNING
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			}

			validations.ShelterMaintainerValidation(stop, tc.Row, &types.StopsRules{ShelterMaintainer: types.RuleConfig{Severity: severity}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name, types.SEVERITY_WARNING)
		})
	}

	t.Run("NotAllowedOption", func(t *testing.T) {
		services.AppMessageService.Clear()
		maintainer := "Other Maintainer"
		options := []string{"Allowed Maintainer"}
		stop := &types.Stop{ShelterMaintainer: &maintainer}

		validations.ShelterMaintainerValidation(stop, 1, &types.StopsRules{ShelterMaintainer: types.RuleConfig{Severity: types.SEVERITY_ERROR, Options: &options}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "NotAllowedOption", types.SEVERITY_ERROR)
	})

	t.Run("AllowedOption", func(t *testing.T) {
		services.AppMessageService.Clear()
		maintainer := "Allowed Maintainer"
		options := []string{maintainer}
		stop := &types.Stop{ShelterMaintainer: &maintainer}

		validations.ShelterMaintainerValidation(stop, 1, &types.StopsRules{ShelterMaintainer: types.RuleConfig{Severity: types.SEVERITY_ERROR, Options: &options}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "AllowedOption", types.SEVERITY_ERROR)
	})
}

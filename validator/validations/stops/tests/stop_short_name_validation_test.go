package stops

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestAllStopShortNameValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("stop_short_name") {
		if tc.Name == "Invalid_Value" {
			continue
		}

		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			stop := &types.Stop{StopShortName: tc.Value}
			severity := types.SEVERITY_WARNING
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			}

			validations.StopShortNameValidation(stop, tc.Row, &types.StopsRules{StopShortName: types.RuleConfig{Severity: severity}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name, types.SEVERITY_WARNING)
		})
	}

	t.Run("NotAllowedOption", func(t *testing.T) {
		services.AppMessageService.Clear()
		shortName := "OTHER"
		options := []string{"ALLOWED"}
		stop := &types.Stop{StopShortName: &shortName}

		validations.StopShortNameValidation(stop, 1, &types.StopsRules{StopShortName: types.RuleConfig{Severity: types.SEVERITY_ERROR, Options: &options}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "NotAllowedOption", types.SEVERITY_ERROR)
	})
}

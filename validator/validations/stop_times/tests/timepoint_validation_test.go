package stop_times

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestAllTimepointValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetTimepointValidOptions()
	for _, tc := range test_helpers.GetGenericEnumIntTestCases("timepoint", validOptions) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var timepoint *int
			if tc.Value != nil {
				if ptr, ok := tc.Value.(*int); ok {
					timepoint = ptr
				}
			}

			var rules *types.StopTimesRules
			if tc.ExpectedErrors > 0 {
				rules = &types.StopTimesRules{Timepoint: types.RuleConfig{Severity: types.SEVERITY_ERROR}}
			} else {
				rules = &types.StopTimesRules{Timepoint: types.RuleConfig{Severity: types.SEVERITY_WARNING}}
			}
			stopTime := &types.StopTime{Timepoint: timepoint}
			validations.TimepointValidation(stopTime, tc.Row, rules)
			if tc.ExpectedWarnings > 0 {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}
}

package stop_times

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestAllContinuousDropOffValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetContinuousPickupDropOffValidOptions()
	for _, tc := range test_helpers.GetGenericEnumIntTestCases("continuous_drop_off", validOptions) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var continuousDropOff *int
			if tc.Value != nil {
				if ptr, ok := tc.Value.(*int); ok {
					continuousDropOff = ptr
				}
			}

			var rules *types.StopTimesRules
			if tc.Name == "Missing_Value_Required" {
				rules = &types.StopTimesRules{ContinuousDropOff: types.RuleConfig{Severity: types.SEVERITY_ERROR}}
			} else if tc.Row == 2 {
				rules = &types.StopTimesRules{ContinuousDropOff: types.RuleConfig{Severity: types.SEVERITY_WARNING}}
			}

			stopTime := &types.StopTime{ContinuousDropOff: continuousDropOff}
			validations.ContinuousDropOffValidation(stopTime, tc.Row, rules)
			if tc.Name == "Missing_Value_Required" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, 1, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}
	t.Run("Forbidden_WithStartWindow", func(t *testing.T) {
		services.AppMessageService.Clear()
		validations.ContinuousDropOffValidation(&types.StopTime{ContinuousDropOff: lib.Ptr(0), StartPickupDropOffWindow: lib.Ptr("07:00:00")}, 3, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Forbidden_WithStartWindow")
	})
	t.Run("Forbidden_WithEndWindow", func(t *testing.T) {
		services.AppMessageService.Clear()
		validations.ContinuousDropOffValidation(&types.StopTime{ContinuousDropOff: lib.Ptr(2), EndPickupDropOffWindow: lib.Ptr("10:00:00")}, 4, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Forbidden_WithEndWindow")
	})
	t.Run("Allowed_WithStartWindowIfOne", func(t *testing.T) {
		services.AppMessageService.Clear()
		validations.ContinuousDropOffValidation(&types.StopTime{ContinuousDropOff: lib.Ptr(1), StartPickupDropOffWindow: lib.Ptr("07:00:00")}, 5, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Allowed_WithStartWindowIfOne")
	})
}

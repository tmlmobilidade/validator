package stop_times

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestAllContinuousPickupValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetContinuousPickupDropOffValidOptions()
	for _, tc := range test_helpers.GetGenericEnumIntTestCases("continuous_pickup", validOptions) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var continuousPickup *int
			if tc.Value != nil {
				if ptr, ok := tc.Value.(*int); ok {
					continuousPickup = ptr
				}
			}

			var rules *types.StopTimesRules
			if tc.Name == "Missing_Value_Required" {
				rules = &types.StopTimesRules{ContinuousPickup: types.RuleConfig{Severity: types.SEVERITY_ERROR}}
			} else if tc.Row == 2 {
				rules = &types.StopTimesRules{ContinuousPickup: types.RuleConfig{Severity: types.SEVERITY_WARNING}}
			}

			stopTime := &types.StopTime{ContinuousPickup: continuousPickup}
			validations.ContinuousPickupValidation(stopTime, tc.Row, rules)
			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
		})
	}
	t.Run("Forbidden_WithStartWindow", func(t *testing.T) {
		services.AppMessageService.Clear()
		validations.ContinuousPickupValidation(&types.StopTime{ContinuousPickup: lib.Ptr(0), StartPickupDropOffWindow: lib.Ptr("07:00:00")}, 3, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Forbidden_WithStartWindow")
	})
	t.Run("Forbidden_WithEndWindow", func(t *testing.T) {
		services.AppMessageService.Clear()
		validations.ContinuousPickupValidation(&types.StopTime{ContinuousPickup: lib.Ptr(2), EndPickupDropOffWindow: lib.Ptr("10:00:00")}, 4, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Forbidden_WithEndWindow")
	})
	t.Run("Allowed_WithStartWindowIfOne", func(t *testing.T) {
		services.AppMessageService.Clear()
		validations.ContinuousPickupValidation(&types.StopTime{ContinuousPickup: lib.Ptr(1), StartPickupDropOffWindow: lib.Ptr("07:00:00")}, 5, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Allowed_WithStartWindowIfOne")
	})
}

package stop_times

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestAllPickupBookingRuleIdValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericForeignKeyTestCases("pickup_booking_rule_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var pickupBookingRuleId *string
			if tc.Id != nil {
				pickupBookingRuleId = tc.Id
			} else {
				pickupBookingRuleId = nil
			}

			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"booking_rules": {*pickupBookingRuleId: {1}}}}.ToGtfs()
			if tc.Name == "ForeignKey_Invalid" {
				gtfs = test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"booking_rules": {}}}.ToGtfs()
			}
			validations.PickupBookingRuleIdValidation(&types.StopTime{PickupBookingRuleId: pickupBookingRuleId}, tc.Row, &gtfs, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
	for _, tc := range test_helpers.GetGenericSeverityTestCases("pickup_booking_rule_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			validations.PickupBookingRuleIdValidation(&types.StopTime{}, tc.Row, &types.Gtfs{}, &types.StopTimesRules{PickupBookingRuleId: types.RuleConfig{Severity: tc.Severity}})
			if tc.Name == "Severity_Warning_Missing" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}

	t.Run("Optional_NotPresent", func(t *testing.T) {
		services.AppMessageService.Clear()
		validations.PickupBookingRuleIdValidation(&types.StopTime{}, 4, &types.Gtfs{}, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Optional_NotPresent")
	})

	t.Run("Missing_BookingRulesIndex", func(t *testing.T) {
		services.AppMessageService.Clear()
		stopTime := &types.StopTime{PickupBookingRuleId: lib.Ptr("BR1")}
		gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"booking_rules": {"BR1": {1}}}}.ToGtfs()
		validations.PickupBookingRuleIdValidation(stopTime, 1, &gtfs, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Missing_BookingRulesIndex")
	})
}

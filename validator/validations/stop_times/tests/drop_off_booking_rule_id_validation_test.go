package stop_times

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestAllDropOffBookingRuleIdValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericForeignKeyTestCases("drop_off_booking_rule_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var severity types.Severity
			if tc.Name == "Severity_Error_Missing" {
				severity = types.SEVERITY_ERROR
			} else if tc.Name == "Severity_Warning_Missing" {
				severity = types.SEVERITY_WARNING
			}
			dropOffBookingRuleId := tc.Id

			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"booking_rules": {}}}.ToGtfs()
			if tc.Name == "ForeignKey_Present" {
				gtfs = test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"booking_rules": {*dropOffBookingRuleId: {1}}}}.ToGtfs()
			}
			validations.DropOffBookingRuleIdValidation(&types.StopTime{DropOffBookingRuleId: dropOffBookingRuleId}, tc.Row, &gtfs, &types.StopTimesRules{DropOffBookingRuleId: types.RuleConfig{Severity: severity}})
			if tc.Name == "Severity_Error_Missing" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, 1, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}
	for _, tc := range test_helpers.GetGenericSeverityTestCases("drop_off_booking_rule_id") {
		if tc.Name != "Severity_Error_Missing" && tc.Name != "Severity_Warning_Missing" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			validations.DropOffBookingRuleIdValidation(&types.StopTime{}, tc.Row, &types.Gtfs{}, &types.StopTimesRules{DropOffBookingRuleId: types.RuleConfig{Severity: tc.Severity}})
			if tc.Name == "Severity_Warning_Missing" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}

	t.Run("Optional_NotPresent", func(t *testing.T) {
		services.AppMessageService.Clear()
		validations.DropOffBookingRuleIdValidation(&types.StopTime{}, 4, &types.Gtfs{}, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Optional_NotPresent")
	})

	t.Run("Missing_BookingRulesIndex", func(t *testing.T) {
		services.AppMessageService.Clear()
		validations.DropOffBookingRuleIdValidation(&types.StopTime{DropOffBookingRuleId: lib.Ptr("BR1")}, 3, &types.Gtfs{}, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "MissingBookingRulesIndex")
	})
}

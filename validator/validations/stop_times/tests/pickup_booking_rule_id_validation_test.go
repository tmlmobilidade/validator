package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestPickupBookingRuleIdValidation_ValidForeignKey(t *testing.T) {
	services.AppMessageService.Clear()
	id := "BR1"
	stopTime := &types.StopTime{PickupBookingRuleId: &id}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"booking_rules": {
				"BR1": {0},
			},
		},
	}
	validations.PickupBookingRuleIdValidation(stopTime, 1, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid pickup_booking_rule_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPickupBookingRuleIdValidation_InvalidForeignKey(t *testing.T) {
	services.AppMessageService.Clear()
	id := "INVALID"
	stopTime := &types.StopTime{PickupBookingRuleId: &id}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"booking_rules": {},
		},
	}
	validations.PickupBookingRuleIdValidation(stopTime, 2, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid pickup_booking_rule_id foreign key should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPickupBookingRuleIdValidation_MissingBookingRulesIndex(t *testing.T) {
	services.AppMessageService.Clear()
	id := "BR1"
	stopTime := &types.StopTime{PickupBookingRuleId: &id}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{},
	}
	validations.PickupBookingRuleIdValidation(stopTime, 3, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing booking_rules index should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPickupBookingRuleIdValidation_OptionalNotPresent(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	gtfs := &types.Gtfs{}
	validations.PickupBookingRuleIdValidation(stopTime, 4, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Optional pickup_booking_rule_id not present should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPickupBookingRuleIdValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	gtfs := &types.Gtfs{}
	severity := types.SEVERITY_ERROR
	validations.PickupBookingRuleIdValidation(stopTime, 5, gtfs, &types.StopTimesRules{PickupBookingRuleId: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "pickup_booking_rule_id missing with severity error should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPickupBookingRuleIdValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	gtfs := &types.Gtfs{}
	severity := types.SEVERITY_WARNING
	validations.PickupBookingRuleIdValidation(stopTime, 6, gtfs, &types.StopTimesRules{PickupBookingRuleId: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "pickup_booking_rule_id missing with severity warning should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

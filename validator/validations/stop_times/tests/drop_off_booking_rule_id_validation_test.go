package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestDropOffBookingRuleIdValidation_ValidForeignKey(t *testing.T) {
	services.AppMessageService.Clear()
	id := "BR1"
	stopTime := &types.StopTime{DropOffBookingRuleId: &id}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"booking_rules": {
				"BR1": {0},
			},
		},
	}
	validations.DropOffBookingRuleIdValidation(nil, stopTime, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid drop_off_booking_rule_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDropOffBookingRuleIdValidation_InvalidForeignKey(t *testing.T) {
	services.AppMessageService.Clear()
	id := "INVALID"
	stopTime := &types.StopTime{DropOffBookingRuleId: &id}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"booking_rules": {},
		},
	}
	validations.DropOffBookingRuleIdValidation(nil, stopTime, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid drop_off_booking_rule_id foreign key should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDropOffBookingRuleIdValidation_MissingBookingRulesIndex(t *testing.T) {
	services.AppMessageService.Clear()
	id := "BR1"
	stopTime := &types.StopTime{DropOffBookingRuleId: &id}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{},
	}
	validations.DropOffBookingRuleIdValidation(nil, stopTime, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing booking_rules index should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDropOffBookingRuleIdValidation_OptionalNotPresent(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	gtfs := &types.Gtfs{}
	validations.DropOffBookingRuleIdValidation(nil, stopTime, 4, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Optional drop_off_booking_rule_id not present should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDropOffBookingRuleIdValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	gtfs := &types.Gtfs{}
	severity := types.SEVERITY_ERROR
	validations.DropOffBookingRuleIdValidation(&severity, stopTime, 5, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "drop_off_booking_rule_id missing with severity error should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDropOffBookingRuleIdValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	gtfs := &types.Gtfs{}
	severity := types.SEVERITY_WARNING
	validations.DropOffBookingRuleIdValidation(&severity, stopTime, 6, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "drop_off_booking_rule_id missing with severity warning should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
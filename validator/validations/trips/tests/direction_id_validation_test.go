package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestDirectionIdValidation_ValidValues(t *testing.T) {
	severity := types.SEVERITY_ERROR
	for _, val := range []int{0, 1} {
		trip := &types.Trip{DirectionId: &val}
		gtfs := &types.Gtfs{}
		validations.DirectionIdValidation(trip, 1, gtfs, &types.TripsRules{DirectionId: types.RuleConfig{Severity: severity}})
		assertion := lib.AssertionMessage{
			Expected: 0,
			Actual:   services.AppMessageService.GetSummary().TotalErrors,
			Message:  "Valid direction_id value should not error",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Error(assert)
		}
		services.AppMessageService.Clear()
	}
}

func TestDirectionIdValidation_InvalidValue(t *testing.T) {
	severity := types.SEVERITY_ERROR
	invalid := 2
	trip := &types.Trip{DirectionId: &invalid}
	gtfs := &types.Gtfs{}
	validations.DirectionIdValidation(trip, 2, gtfs, &types.TripsRules{DirectionId: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid direction_id value should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDirectionIdValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	trip := &types.Trip{DirectionId: nil}
	gtfs := &types.Gtfs{}
	validations.DirectionIdValidation(trip, 3, gtfs, &types.TripsRules{DirectionId: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Direction ID is required",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDirectionIdValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	trip := &types.Trip{DirectionId: nil}
	gtfs := &types.Gtfs{}
	validations.DirectionIdValidation(trip, 4, gtfs, &types.TripsRules{DirectionId: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Direction ID is recommended",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDirectionIdValidation_Ignore(t *testing.T) {
	severity := types.SEVERITY_IGNORE
	trip := &types.Trip{DirectionId: nil}
	gtfs := &types.Gtfs{}
	validations.DirectionIdValidation(trip, 5, gtfs, &types.TripsRules{DirectionId: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Direction ID is ignored, no error or warning should be reported",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

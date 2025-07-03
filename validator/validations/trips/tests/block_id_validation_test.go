package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestBlockIdValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	trip := &types.Trip{BlockId: nil}
	gtfs := &types.Gtfs{}
	validations.BlockIdValidation(trip, 1, gtfs, &types.TripsRules{BlockId: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "block_id is required",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestBlockIdValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	trip := &types.Trip{BlockId: nil}
	gtfs := &types.Gtfs{}
	validations.BlockIdValidation(trip, 2, gtfs, &types.TripsRules{BlockId: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "block_id is recommended",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestBlockIdValidation_Ignore(t *testing.T) {
	severity := types.SEVERITY_IGNORE
	trip := &types.Trip{BlockId: nil}
	gtfs := &types.Gtfs{}
	validations.BlockIdValidation(trip, 3, gtfs, &types.TripsRules{BlockId: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "block_id is ignored, no error or warning should be reported",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

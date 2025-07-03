package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestTripShortNameValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	trip := &types.Trip{TripShortName: nil}
	gtfs := &types.Gtfs{}
	validations.TripShortNameValidation(trip, 1, gtfs, &types.TripsRules{TripShortName: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Trip short name is required",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestTripShortNameValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	trip := &types.Trip{TripShortName: nil}
	gtfs := &types.Gtfs{}
	validations.TripShortNameValidation(trip, 2, gtfs, &types.TripsRules{TripShortName: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Trip short name is recommended",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestTripShortNameValidation_Ignore(t *testing.T) {
	severity := types.SEVERITY_IGNORE
	trip := &types.Trip{TripShortName: nil}
	gtfs := &types.Gtfs{}
	validations.TripShortNameValidation(trip, 3, gtfs, &types.TripsRules{TripShortName: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Trip short name is ignored, no error or warning should be reported",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestTripShortNameValidation_Present(t *testing.T) {
	severity := types.SEVERITY_ERROR
	short := "T123"
	trip := &types.Trip{TripShortName: &short}
	gtfs := &types.Gtfs{}
	validations.TripShortNameValidation(trip, 4, gtfs, &types.TripsRules{TripShortName: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Trip short name present, no error or warning should be reported",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

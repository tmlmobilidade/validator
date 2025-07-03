package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestTripHeadsignValidation_Required(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	trip := &types.Trip{TripHeadsign: nil}
	gtfs := &types.Gtfs{}
	validations.TripHeadsignValidation(trip, 1, gtfs, &types.TripsRules{TripHeadsign: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Trip headsign is required",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestTripHeadsignValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	trip := &types.Trip{TripHeadsign: nil}
	gtfs := &types.Gtfs{}
	validations.TripHeadsignValidation(trip, 2, gtfs, &types.TripsRules{TripHeadsign: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Trip headsign is recommended",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestTripHeadsignValidation_Ignore(t *testing.T) {
	severity := types.SEVERITY_IGNORE
	trip := &types.Trip{TripHeadsign: nil}
	gtfs := &types.Gtfs{}
	validations.TripHeadsignValidation(trip, 3, gtfs, &types.TripsRules{TripHeadsign: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Trip headsign is ignored, no error or warning should be reported",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestTripHeadsignValidation_Present(t *testing.T) {
	severity := types.SEVERITY_ERROR
	head := "Downtown"
	trip := &types.Trip{TripHeadsign: &head}
	gtfs := &types.Gtfs{}
	validations.TripHeadsignValidation(trip, 4, gtfs, &types.TripsRules{TripHeadsign: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Trip headsign present, no error or warning should be reported",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

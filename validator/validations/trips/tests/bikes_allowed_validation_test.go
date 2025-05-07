package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestBikesAllowedValidation_ValidValues(t *testing.T) {
	severity := types.SEVERITY_ERROR
	for _, val := range []int{0, 1, 2} {
		trip := &types.Trip{BikesAllowed: &val}
		gtfs := &types.Gtfs{}
		validations.BikesAllowedValidation(&severity, trip, 1, gtfs)
		assertion := lib.AssertionMessage{
			Expected: 0,
			Actual: services.AppMessageService.GetSummary().TotalErrors,
			Message: "Valid bikes_allowed value should not error",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Error(assert)
		}
		services.AppMessageService.Clear()
	}
}

func TestBikesAllowedValidation_InvalidValue(t *testing.T) {
	severity := types.SEVERITY_ERROR
	invalid := 3
	trip := &types.Trip{BikesAllowed: &invalid}
	gtfs := &types.Gtfs{}
	validations.BikesAllowedValidation(&severity, trip, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid bikes_allowed value should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestBikesAllowedValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	trip := &types.Trip{BikesAllowed: nil}
	gtfs := &types.Gtfs{}
	validations.BikesAllowedValidation(&severity, trip, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "bikes_allowed is required",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestBikesAllowedValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	trip := &types.Trip{BikesAllowed: nil}
	gtfs := &types.Gtfs{}
	validations.BikesAllowedValidation(&severity, trip, 4, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "bikes_allowed is recommended",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestBikesAllowedValidation_Ignore(t *testing.T) {
	severity := types.SEVERITY_IGNORE
	trip := &types.Trip{BikesAllowed: nil}
	gtfs := &types.Gtfs{}
	validations.BikesAllowedValidation(&severity, trip, 5, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message: "bikes_allowed is ignored, no error or warning should be reported",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
} 
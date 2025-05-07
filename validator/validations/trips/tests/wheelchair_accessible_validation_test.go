package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestWheelchairAccessibleValidation_ValidValues(t *testing.T) {
	severity := types.SEVERITY_ERROR
	for _, val := range []int{0, 1, 2} {
		trip := &types.Trip{WheelchairAccessible: &val}
		gtfs := &types.Gtfs{}
		validations.WheelchairAccessibleValidation(&severity, trip, 1, gtfs)
		assertion := lib.AssertionMessage{
			Expected: 0,
			Actual: services.AppMessageService.GetSummary().TotalErrors,
			Message: "Valid wheelchair_accessible value should not error",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Error(assert)
		}
		services.AppMessageService.Clear()
	}
}

func TestWheelchairAccessibleValidation_InvalidValue(t *testing.T) {
	severity := types.SEVERITY_ERROR
	invalid := 3
	trip := &types.Trip{WheelchairAccessible: &invalid}
	gtfs := &types.Gtfs{}
	validations.WheelchairAccessibleValidation(&severity, trip, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid wheelchair_accessible value should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestWheelchairAccessibleValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	trip := &types.Trip{WheelchairAccessible: nil}
	gtfs := &types.Gtfs{}
	validations.WheelchairAccessibleValidation(&severity, trip, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "wheelchair_accessible is required",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestWheelchairAccessibleValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	trip := &types.Trip{WheelchairAccessible: nil}
	gtfs := &types.Gtfs{}
	validations.WheelchairAccessibleValidation(&severity, trip, 4, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "wheelchair_accessible is recommended",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestWheelchairAccessibleValidation_Ignore(t *testing.T) {
	severity := types.SEVERITY_IGNORE
	trip := &types.Trip{WheelchairAccessible: nil}
	gtfs := &types.Gtfs{}
	validations.WheelchairAccessibleValidation(&severity, trip, 5, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message: "wheelchair_accessible is ignored, no error or warning should be reported",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
} 
package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestShapeDistTraveledValidation_ValidNonNegative(t *testing.T) {
	services.AppMessageService.Clear()
	val := 5.25
	stopTime := &types.StopTime{ShapeDistTraveled: &val}
	gtfs := &types.Gtfs{}
	validations.ShapeDistTraveledValidation(nil, stopTime, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid non-negative shape_dist_traveled should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeDistTraveledValidation_NegativeValue(t *testing.T) {
	services.AppMessageService.Clear()
	val := -1.0
	stopTime := &types.StopTime{ShapeDistTraveled: &val}
	gtfs := &types.Gtfs{}
	validations.ShapeDistTraveledValidation(nil, stopTime, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Negative shape_dist_traveled should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeDistTraveledValidation_OptionalNotPresent(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	gtfs := &types.Gtfs{}
	validations.ShapeDistTraveledValidation(nil, stopTime, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Optional shape_dist_traveled not present should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeDistTraveledValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	gtfs := &types.Gtfs{}
	severity := types.SEVERITY_ERROR
	validations.ShapeDistTraveledValidation(&severity, stopTime, 4, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "shape_dist_traveled missing with severity error should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeDistTraveledValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	gtfs := &types.Gtfs{}
	severity := types.SEVERITY_WARNING
	validations.ShapeDistTraveledValidation(&severity, stopTime, 5, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "shape_dist_traveled missing with severity warning should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
package shapes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/shapes/validations"
	"testing"
)

func TestShapePtLatValidation_MissingShapePtLat(t *testing.T) {
	services.AppMessageService.Clear()
	shape := &types.Shape{ShapePtLat: nil}
	validations.ShapePtLatValidation(shape, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing shape_pt_lat should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapePtLatValidation_InvalidLatitude(t *testing.T) {
	services.AppMessageService.Clear()
	invalid := float32(100.0)
	shape := &types.Shape{ShapePtLat: &invalid}
	validations.ShapePtLatValidation(shape, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid latitude should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapePtLatValidation_ValidLatitude(t *testing.T) {
	services.AppMessageService.Clear()
	valid := float32(37.61956)
	shape := &types.Shape{ShapePtLat: &valid}
	validations.ShapePtLatValidation(shape, 3)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid latitude should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
package shapes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/shapes/validations"
	"testing"
)

func TestShapePtLonValidation_MissingShapePtLon(t *testing.T) {
	services.AppMessageService.Clear()
	shape := &types.Shape{ShapePtLon: nil}
	validations.ShapePtLonValidation(shape, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing shape_pt_lon should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapePtLonValidation_InvalidLongitude(t *testing.T) {
	services.AppMessageService.Clear()
	invalid := float32(200.0)
	shape := &types.Shape{ShapePtLon: &invalid}
	validations.ShapePtLonValidation(shape, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid longitude should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapePtLonValidation_ValidLongitude(t *testing.T) {
	services.AppMessageService.Clear()
	valid := float32(-122.48161)
	shape := &types.Shape{ShapePtLon: &valid}
	validations.ShapePtLonValidation(shape, 3)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid longitude should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
package shapes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/shapes/validations"
	"testing"
)

func TestShapePtSequenceValidation_MissingShapePtSequence(t *testing.T) {
	services.AppMessageService.Clear()
	shape := &types.Shape{ShapePtSequence: nil}
	validations.ShapePtSequenceValidation(shape, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing shape_pt_sequence should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapePtSequenceValidation_NegativeValue(t *testing.T) {
	services.AppMessageService.Clear()
	neg := -1
	shape := &types.Shape{ShapePtSequence: &neg}
	validations.ShapePtSequenceValidation(shape, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Negative shape_pt_sequence should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapePtSequenceValidation_ValidValue(t *testing.T) {
	services.AppMessageService.Clear()
	val := 6
	shape := &types.Shape{ShapePtSequence: &val}
	validations.ShapePtSequenceValidation(shape, 3)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid shape_pt_sequence should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
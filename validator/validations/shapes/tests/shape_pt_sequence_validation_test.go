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

func TestShapePtSequenceGroupValidation_NonIncreasingSequence(t *testing.T) {
	services.AppMessageService.Clear()
	shapeId := "A_shp"
	shapes := []types.Shape{
		{ShapeId: &shapeId, ShapePtSequence: intPtr(0)},
		{ShapeId: &shapeId, ShapePtSequence: intPtr(6)},
		{ShapeId: &shapeId, ShapePtSequence: intPtr(5)}, // Not increasing
	}
	validations.ShapePtSequenceGroupValidation(shapes)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Non-increasing shape_pt_sequence should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapePtSequenceGroupValidation_IncreasingSequence(t *testing.T) {
	services.AppMessageService.Clear()
	shapeId := "A_shp"
	shapes := []types.Shape{
		{ShapeId: &shapeId, ShapePtSequence: intPtr(0)},
		{ShapeId: &shapeId, ShapePtSequence: intPtr(6)},
		{ShapeId: &shapeId, ShapePtSequence: intPtr(11)},
	}
	validations.ShapePtSequenceGroupValidation(shapes)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Increasing shape_pt_sequence should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func intPtr(i int) *int { return &i } 
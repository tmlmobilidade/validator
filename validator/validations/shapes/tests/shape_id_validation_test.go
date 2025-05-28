package shapes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/shapes/validations"
	"testing"
)

func TestShapeIdValidation_MissingShapeId(t *testing.T) {
	services.AppMessageService.Clear()
	shape := &types.Shape{ShapeId: nil}
	validations.ShapeIdValidation(shape, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing shape_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeIdValidation_EmptyShapeId(t *testing.T) {
	services.AppMessageService.Clear()
	empty := ""
	shape := &types.Shape{ShapeId: &empty}
	validations.ShapeIdValidation(shape, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Empty shape_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeIdValidation_ValidShapeId(t *testing.T) {
	services.AppMessageService.Clear()
	valid := "SHP1"
	shape := &types.Shape{ShapeId: &valid}
	validations.ShapeIdValidation(shape, 3)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid shape_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
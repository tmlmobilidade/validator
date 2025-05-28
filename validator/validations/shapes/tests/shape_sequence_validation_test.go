package shapes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/shapes/validations"
	"testing"
)

func TestShapeSequenceValidation_Valid(t *testing.T) {
	services.AppMessageService.Clear()
	
	shapes := []types.Shape{
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: lib.Ptr(0), ShapeDistTraveled: lib.Ptr(0.0) },
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: lib.Ptr(1), ShapeDistTraveled: lib.Ptr(1.0) },
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: lib.Ptr(2), ShapeDistTraveled: lib.Ptr(2.0) },
	}

	validations.ShapeSequenceValidation(shapes)
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid shape sequence should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeSequenceValidation_Valid_NoShapeDistTraveled(t *testing.T) {
	services.AppMessageService.Clear()
	
	shapes := []types.Shape{
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: lib.Ptr(0), ShapeDistTraveled: nil },
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: lib.Ptr(1), ShapeDistTraveled: nil },
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: lib.Ptr(2), ShapeDistTraveled: nil },
	}

	validations.ShapeSequenceValidation(shapes)
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid shape sequence should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeSequenceValidation_Invalid_ShapeDistTraveledDoesNotIncrease(t *testing.T) {
	services.AppMessageService.Clear()
	
	shapes := []types.Shape{
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: lib.Ptr(0), ShapeDistTraveled: lib.Ptr(0.0) },
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: lib.Ptr(1), ShapeDistTraveled: lib.Ptr(2.0) },
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: lib.Ptr(2), ShapeDistTraveled: lib.Ptr(1.0) },
	}

	validations.ShapeSequenceValidation(shapes)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "ShapeDistTraveled does not increase",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeSequenceValidation_Invalid_NoShapeId(t *testing.T) {
	services.AppMessageService.Clear()
	
	shapes := []types.Shape{
		{ ShapeId: nil, ShapePtSequence: lib.Ptr(0), ShapeDistTraveled: nil },
		{ ShapeId: nil, ShapePtSequence: lib.Ptr(1), ShapeDistTraveled: nil },
		{ ShapeId: nil, ShapePtSequence: lib.Ptr(2), ShapeDistTraveled: nil },
	}

	validations.ShapeSequenceValidation(shapes)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "ShapeId is required",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeSequenceValidation_Invalid_NoShapeDistTraveled(t *testing.T) {
	services.AppMessageService.Clear()
	
	shapes := []types.Shape{
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: lib.Ptr(0), ShapeDistTraveled: nil },
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: lib.Ptr(1), ShapeDistTraveled: nil },
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: lib.Ptr(1), ShapeDistTraveled: nil },
	}

	validations.ShapeSequenceValidation(shapes)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "shape_pt_sequence must increase along the trip",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeSequenceValidation_Invalid_NoShapePtSequence(t *testing.T) {
	services.AppMessageService.Clear()
	
	shapes := []types.Shape{
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: nil, ShapeDistTraveled: lib.Ptr(0.0) },
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: nil, ShapeDistTraveled: lib.Ptr(1.0) },
		{ ShapeId: lib.Ptr("A_shp"), ShapePtSequence: nil, ShapeDistTraveled: lib.Ptr(2.0) },
	}

	validations.ShapeSequenceValidation(shapes)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "shape_pt_sequence is required and must not be empty.",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
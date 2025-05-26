package shapes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/shapes/validations"
	"testing"
)

func TestShapeDistTraveledValidation_Missing(t *testing.T) {
	services.AppMessageService.Clear()
	shape := &types.Shape{ShapeDistTraveled: nil}
	validations.ShapeDistTraveledValidation(nil, shape, 1)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing shape_dist_traveled should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeDistTraveledValidation_Negative(t *testing.T) {
	services.AppMessageService.Clear()
	neg := -1.0
	shape := &types.Shape{ShapeDistTraveled: &neg}
	validations.ShapeDistTraveledValidation(nil, shape, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Negative shape_dist_traveled should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeDistTraveledValidation_Valid(t *testing.T) {
	services.AppMessageService.Clear()
	val := 6.8310
	shape := &types.Shape{ShapeDistTraveled: &val}
	validations.ShapeDistTraveledValidation(nil, shape, 3)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid shape_dist_traveled should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeDistTraveledValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	
	shape := &types.Shape{}
	severity := types.SEVERITY_ERROR
	validations.ShapeDistTraveledValidation(&severity, shape, 4)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Severity error should error",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeDistTraveledValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	
	shape := &types.Shape{}
	severity := types.SEVERITY_WARNING
	validations.ShapeDistTraveledValidation(&severity, shape, 4)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Severity warning should warn",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeDistTraveledGroupValidation_Increasing(t *testing.T) {
	services.AppMessageService.Clear()
	shapeId := "A_shp"
	shapes := []types.Shape{
		{ShapeId: &shapeId, ShapePtSequence: intPtr(0), ShapeDistTraveled: floatPtr(0)},
		{ShapeId: &shapeId, ShapePtSequence: intPtr(6), ShapeDistTraveled: floatPtr(6.8310)},
		{ShapeId: &shapeId, ShapePtSequence: intPtr(11), ShapeDistTraveled: floatPtr(15.8765)},
	}
	validations.ShapeDistTraveledGroupValidation(shapes)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Increasing shape_dist_traveled should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeDistTraveledGroupValidation_NonIncreasing(t *testing.T) {
	services.AppMessageService.Clear()
	shapeId := "A_shp"
	shapes := []types.Shape{
		{ShapeId: &shapeId, ShapePtSequence: intPtr(0), ShapeDistTraveled: floatPtr(0)},
		{ShapeId: &shapeId, ShapePtSequence: intPtr(6), ShapeDistTraveled: floatPtr(6.8310)},
		{ShapeId: &shapeId, ShapePtSequence: intPtr(11), ShapeDistTraveled: floatPtr(5.0)}, // Not increasing
	}
	validations.ShapeDistTraveledGroupValidation(shapes)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Non-increasing shape_dist_traveled should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func floatPtr(f float64) *float64 { return &f }
func intPtr(i int) *int { return &i } 
package shapes

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/shapes/validations"
	"testing"
)

func TestShapeDistTraveledValidation_Negative(t *testing.T) {
	services.AppMessageService.Clear()

	neg := -1.0
	shape := &types.Shape{ShapeDistTraveled: &neg}
	validations.ShapeDistTraveledValidation(shape, 2, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Negative shape_dist_traveled should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeDistTraveledValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()

	shape := &types.Shape{}
	severity := types.SEVERITY_ERROR
	validations.ShapeDistTraveledValidation(shape, 4, &types.ShapesRules{ShapeDistTraveled: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Severity error should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShapeDistTraveledValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()

	shape := &types.Shape{}
	severity := types.SEVERITY_WARNING
	validations.ShapeDistTraveledValidation(shape, 4, &types.ShapesRules{ShapeDistTraveled: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Severity warning should warn",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAllShapeDistTraveledValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetShapeDistTraveledValidOptions()
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("shape_dist_traveled") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var severity types.Severity
			if tc.ExpectedCode == "shape_dist_traveled_validation.required" {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}

			var shapeDistTraveled *float64
			if tc.Value != nil {
				shapeDistTraveled = &validOptions[tc.Row-1]
			} else {
				shapeDistTraveled = nil
			}
			validations.ShapeDistTraveledValidation(&types.Shape{ShapeDistTraveled: shapeDistTraveled}, tc.Row, &types.ShapesRules{ShapeDistTraveled: types.RuleConfig{Severity: severity}})
			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
		})
	}
}

package shapes

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/shapes/validations"
	"testing"
)

func TestAllShapeDistTraveledValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetShapeFloat64ValidOptions()
	negativeOptions := test_helpers.GetShapeFloat64InvalidOptions()
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("shape_dist_traveled") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var severity types.Severity
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			} else if tc.Name == "Recommended_Missing" {
				severity = types.SEVERITY_WARNING
			} else {
				severity = types.SEVERITY_ERROR
			}

			var shapeDistTraveled *float64

			if tc.Name == "Invalid_Value" {
				shapeDistTraveled = &negativeOptions[0]
			} else if tc.Value != nil {
				shapeDistTraveled = &validOptions[0]
			} else {
				shapeDistTraveled = nil
			}

			validations.ShapeDistTraveledValidation(&types.Shape{ShapeDistTraveled: shapeDistTraveled}, tc.Row, &types.ShapesRules{ShapeDistTraveled: types.RuleConfig{Severity: severity}})
			if tc.Name == "Recommended_Missing" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}
}

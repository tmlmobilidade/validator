package stop_times

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestAllShapeDistTraveledValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetShapeFloat64ValidOptions()
	negativeOptions := test_helpers.GetShapeFloat64InvalidOptions()
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("shape_dist_traveled") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var shapeDistTraveled *float64
			if tc.Name == "Invalid_Value" {
				shapeDistTraveled = &negativeOptions[0]
			} else if tc.Value != nil {
				shapeDistTraveled = &validOptions[0]
			} else {
				shapeDistTraveled = nil
			}

			var rules *types.StopTimesRules
			if tc.Name == "Recommended_Missing" {
				rules = &types.StopTimesRules{ShapeDistTraveled: types.RuleConfig{Severity: types.SEVERITY_WARNING}}
			} else {
				rules = &types.StopTimesRules{ShapeDistTraveled: types.RuleConfig{Severity: types.SEVERITY_ERROR}}
			}

			validations.ShapeDistTraveledValidation(&types.StopTime{ShapeDistTraveled: shapeDistTraveled}, tc.Row, rules)
			if tc.Name == "Recommended_Missing" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}

}

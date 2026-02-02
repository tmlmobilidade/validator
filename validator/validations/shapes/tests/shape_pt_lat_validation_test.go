package shapes

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/shapes/validations"
	"testing"
)

func TestShapePtLatValidation_InvalidLatitude(t *testing.T) {
	services.AppMessageService.Clear()
	invalid := float32(1000000.0)
	shape := &types.Shape{ShapePtLat: &invalid}
	validations.ShapePtLatValidation(shape, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid latitude should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAllShapePtLatValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetShapePtLatValidOptions()
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("shape_pt_lat") {
		if tc.Name == "Recommended_Missing" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var shapePtLat *float32
			if tc.Value != nil {
				shapePtLat = &validOptions[0]
			} else {
				shapePtLat = nil
			}
			validations.ShapePtLatValidation(&types.Shape{ShapePtLat: shapePtLat}, tc.Row)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
}

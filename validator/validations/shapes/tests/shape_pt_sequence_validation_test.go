package shapes

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/shapes/validations"
	"testing"
)

func TestAllShapePtSequenceValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetValidShapeOptions()
	negativeOption := -1
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("shape_pt_sequence") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var shapePtSequence *int
			if tc.Name == "Invalid_Value" {
				shapePtSequence = &negativeOption
			} else if tc.Value != nil {
				shapePtSequence = &validOptions[0]
			} else {
				shapePtSequence = nil
			}
			validations.ShapePtSequenceValidation(&types.Shape{ShapePtSequence: shapePtSequence}, tc.Row)
			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
		})
	}
}

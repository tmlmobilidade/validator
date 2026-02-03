package fare_attributes

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/fare_attributes/validations"
	"testing"
)

func TestAllPriceValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetShapeFloat64ValidOptions()
	negativeOptions := test_helpers.GetShapeFloat64InvalidOptions()
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("price") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var price *float64

			if tc.Name == "Invalid_Value" {
				price = &negativeOptions[0]
			} else if tc.Value != nil {
				price = &validOptions[0]
			} else {
				price = nil
			}
			validations.PriceValidation(&types.FareAttribute{Price: price}, tc.Row)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
}

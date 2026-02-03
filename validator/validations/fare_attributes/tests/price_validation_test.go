package fare_attributes

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/fare_attributes/validations"
	"testing"
)

func TestAllPriceValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetPriceValidOptions()
	negativePrice := -1.0
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("price") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var price *float64
			if tc.Value != nil {
				price = &validOptions[tc.Row-1]
			}

			if tc.Name == "Invalid_Value" {
				price = &negativePrice
			}

			validations.PriceValidation(&types.FareAttribute{Price: price}, tc.Row)
			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
		})
	}
	t.Run("NegativePrice", func(t *testing.T) {
		services.AppMessageService.Clear()
		validations.PriceValidation(&types.FareAttribute{Price: &negativePrice}, 1)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Negative price should error")
	})
}

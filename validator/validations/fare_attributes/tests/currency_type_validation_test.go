package fare_attributes

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/fare_attributes/validations"
	"testing"
)

func TestAllCurrencyTypeValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("currency_type") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var currencyType *string
			if tc.Value != nil {
				if tc.Name == "Valid_Present" {
					validCurrency := "USD"
					currencyType = &validCurrency
				} else {
					currencyType = tc.Value
				}
			}
			validations.CurrencyTypeValidation(&types.FareAttribute{CurrencyType: currencyType}, tc.Row)
			if tc.Name == "Recommended_Missing" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, 1, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}
}

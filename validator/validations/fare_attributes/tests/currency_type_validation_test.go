package fare_attributes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_attributes/validations"
	"testing"
)

func TestCurrencyTypeValidation_ValidCurrencyType(t *testing.T) {
	services.AppMessageService.Clear()
	
	currencyType := "EUR"
	fareAttribute := &types.FareAttribute{CurrencyType: &currencyType}
	validations.CurrencyTypeValidation(fareAttribute, 1)
	
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid currency_type should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestCurrencyTypeValidation_InvalidCurrencyType(t *testing.T) {
	services.AppMessageService.Clear()
	
	currencyType := "INVALID"
	fareAttribute := &types.FareAttribute{CurrencyType: &currencyType}
	validations.CurrencyTypeValidation(fareAttribute, 1)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid currency_type should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestCurrencyTypeValidation_MissingCurrencyType(t *testing.T) {
	services.AppMessageService.Clear()
	
	fareAttribute := &types.FareAttribute{}
	validations.CurrencyTypeValidation(fareAttribute, 1)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing currency_type should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

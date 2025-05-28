package fare_attributes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_attributes/validations"
	"testing"
)

func TestPriceValidation_MissingPrice(t *testing.T) {
	services.AppMessageService.Clear()
	fareAttribute := &types.FareAttribute{Price: nil}
	validations.PriceValidation(fareAttribute, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing price should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPriceValidation_NegativePrice(t *testing.T) {
	services.AppMessageService.Clear()
	price := -1.0
	fareAttribute := &types.FareAttribute{Price: &price}
	validations.PriceValidation(fareAttribute, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Negative price should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPriceValidation_ValidPrice(t *testing.T) {
	services.AppMessageService.Clear()
	price := 2.5
	fareAttribute := &types.FareAttribute{Price: &price}
	validations.PriceValidation(fareAttribute, 3)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid price should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
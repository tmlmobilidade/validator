package fare_attributes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_attributes/validations"
	"testing"
)

func TestPaymentMethodValidation_MissingPaymentMethod(t *testing.T) {
	services.AppMessageService.Clear()
	fareAttribute := &types.FareAttribute{PaymentMethod: nil}
	validations.PaymentMethodValidation(fareAttribute, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing payment_method should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPaymentMethodValidation_InvalidPaymentMethod(t *testing.T) {
	services.AppMessageService.Clear()
	invalid := 2
	fareAttribute := &types.FareAttribute{PaymentMethod: &invalid}
	validations.PaymentMethodValidation(fareAttribute, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid payment_method should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPaymentMethodValidation_ValidPaymentMethod0(t *testing.T) {
	services.AppMessageService.Clear()
	pm := 0
	fareAttribute := &types.FareAttribute{PaymentMethod: &pm}
	validations.PaymentMethodValidation(fareAttribute, 3)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid payment_method 0 should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPaymentMethodValidation_ValidPaymentMethod1(t *testing.T) {
	services.AppMessageService.Clear()
	pm := 1
	fareAttribute := &types.FareAttribute{PaymentMethod: &pm}
	validations.PaymentMethodValidation(fareAttribute, 4)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid payment_method 1 should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
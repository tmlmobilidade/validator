package fare_attributes

import (
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

	- File: [fare_attributes.txt]
	- Field: payment_method
	- Presence: Required
	- Type: Enum

# Description

Indicates when the fare must be paid.

Valid options are:

	- 0 - Fare is paid on board.
	- 1 - Fare must be paid before boarding.

[fare_attributes.txt]: https://gtfs.org/schedule/reference/#fare_attributestxt
*/
func PaymentMethodValidation(fareAttribute *types.FareAttribute, row int) {

	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "payment_method",
			FileName:     "fare_attributes.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "payment_method_validation",
		})
	}
	
	if fareAttribute.PaymentMethod == nil {
		addMessage("Payment method is required")
		return
	}
	
	validPaymentMethods := []int{0, 1}
	if !slices.Contains(validPaymentMethods, *fareAttribute.PaymentMethod) {
		addMessage("Invalid payment method. Valid options are 0 and 1.")
		return
	}
}

package fare_attributes

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

	- File: [fare_attributes.txt]
	- Field: currency_type
	- Presence: Required
	- Type: Currency code

# Description

Currency used to pay the fare.

[fare_attributes.txt]: https://gtfs.org/schedule/reference/#fare_attributestxt
*/
func CurrencyTypeValidation(fareAttribute *types.FareAttribute, row int) {

	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "currency_type",
			FileName:     "fare_attributes.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "currency_type_validation",
		})
	}
	
	if fareAttribute.CurrencyType == nil {
		addMessage("Currency type is required")
		return
	}
	
	if errMsg := lib.ValidateCurrencyType(*fareAttribute.CurrencyType); errMsg != "" {
		addMessage(errMsg)
		return
	}
}

package fare_attributes

import (
	"main/services"
	"main/types"
)

/*
# Attributes

	- File: [fare_attributes.txt]
	- Field: price
	- Presence: Required
	- Type: Non-negative float

# Description

Fare price, in the unit specified by currency_type.

[fare_attributes.txt]: https://gtfs.org/schedule/reference/#fare_attributestxt
*/
func PriceValidation(fareAttribute *types.FareAttribute, row int, gtfs *types.Gtfs) {

	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "price",
			FileName:     "fare_attributes.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "price_validation",
		})
	}
	
	if fareAttribute.Price == nil {
		addMessage("Price is required")
		return
	}
	
	if *fareAttribute.Price < 0 {
		addMessage("Price must be non-negative")
	}
}

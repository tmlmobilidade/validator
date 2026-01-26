package validations

import (
	"main/i18n"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [fare_products.txt]
  - Field: fare_product_name
  - Presence: Optional
  - Type: Text

# Description

The name of the fare product as displayed to riders.

[fare_products.txt]: https://gtfs.org/schedule/reference/#fare_productstxt
*/

func FareProductsNameValidation(fareProduct *types.FareProduct, row int, rules *types.FareProductRules) {
	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "fare_product_name",
			FileName:     "fare_products.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_WARNING,
			ValidationID: "fare_products_name_validation",
		})
	}

	if fareProduct.FareProductName == nil || *fareProduct.FareProductName == "" {
		addMessage(i18n.AppTranslator.Get("fare_products_name_validation.warning"))
		return
	}
}

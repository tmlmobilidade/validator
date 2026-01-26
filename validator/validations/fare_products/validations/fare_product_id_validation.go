package validations

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [fare_products.txt]
  - Field: fare_product_id
  - Presence: Required
  - Type: ID

# Description

Identifies a fare product or set of fare products.

Multiple records sharing the same fare_product_id are permitted as long as they contain different fare_media_ids or rider_category_ids.
Differing fare_media_ids would indicate various methods are available for employing the fare product, potentially at different prices.
Differing rider_category_ids would indicate multiple rider categories are eligible for the fare product, potentially at different prices.

[fare_products.txt]: https://gtfs.org/schedule/reference/#fare_productstxt
*/

func FareProductIdValidation(fareProduct *types.FareProduct, row int, gtfs *types.Gtfs, rules *types.FareProductRules) {
	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "fare_product_id",
			FileName:     "fare_products.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "fare_product_id_validation",
		})
	}

	// Validate presence
	if fareProduct.FareProductId == nil || *fareProduct.FareProductId == "" {
		addMessage(i18n.AppTranslator.Get("fare_product_id_validation.required"))
		return
	}

	// Validate that fareProduct.FareProductId exists in the fare_products.txt file
	if gtfs == nil || !lib.GtfsIdMapKeyExists(gtfs, "fare_products", *fareProduct.FareProductId) {
		addMessage(i18n.AppTranslator.Get("fare_product_id_validation.invalid"))
		return
	}
}

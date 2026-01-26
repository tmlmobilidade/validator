package validations

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [rider_categories.txt]
  - Field: rider_category_id
  - Presence: Required
  - Type: Unique ID

# Description

Identifies a rider category.

[rider_categories.txt]: https://gtfs.org/schedule/reference/#rider_categoriestxt
*/

func RiderCategoryIdValidation(riderCategory *types.RiderCategory, row int, gtfs *types.Gtfs, rules *types.RiderCategory) {
	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "rider_category_id",
			FileName:     "rider_categories.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "rider_category_id_validation",
		})
	}

	// Validate presence
	if riderCategory.RiderCategoryId == nil || *riderCategory.RiderCategoryId == "" {
		addMessage(i18n.AppTranslator.Get("rider_category_id_validation.required"))
		return
	}

	// Validate foreign key
	if gtfs != nil && !lib.GtfsIdMapKeyExists(gtfs, "rider_categories", *riderCategory.RiderCategoryId) {
		addMessage(i18n.AppTranslator.Get("rider_category_id_validation.invalid"))
		return
	}

}

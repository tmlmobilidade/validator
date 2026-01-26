package rider_categories

import (
	"main/i18n"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [rider_categories.txt]
  - Field: rider_category_name
  - Presence: required
  - Type: text

# Description

Rider category name as displayed to the rider.

[rider_categories.txt]: https://gtfs.org/schedule/reference/#rider_categoriestxt
*/

func RiderCategoryNameValidation(riderCategory *types.RiderCategory, row int, gtfs *types.Gtfs, rules *types.RiderCategory) {
	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "rider_category_name",
			FileName:     "rider_categories.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_WARNING,
			ValidationID: "rider_category_name_validation",
		})
	}

	if riderCategory.RiderCategoryName == nil || *riderCategory.RiderCategoryName == "" {
		addMessage(i18n.AppTranslator.Get("rider_category_name_validation.required"))
		return
	}

	// Rules validation can be added here if needed
	if rules != nil {
		// Rules validation logic can be implemented here
	}
}

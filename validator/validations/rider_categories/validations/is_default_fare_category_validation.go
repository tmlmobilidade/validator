package validations

import (
	"main/i18n"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

  - File: [rider_categories.txt]
  - Field: is_default_fare_category
  - Presence: required
  - Type: enum

# Description

Specifies if an entry in rider_categories.txt should be considered the default category (i.e. the main category that should be displayed to riders).

For example: Adult fare, Regular fare, etc. Valid options are:

	0 or empty - Category is not considered the default.
	1 - Category is considered the default one.

When multiple rider categories are eligible for a single fare product specified by a fare_product_id, there must be exactly one of these eligible rider categories indicated as the default rider category (is_default_fare_category = 1).

[rider_categories.txt]: https://gtfs.org/schedule/reference/#rider_categoriestxt
*/

func IsDefaultFareCategoryValidation(riderCategory *types.RiderCategory, row int, gtfs *types.Gtfs, rules *types.RiderCategory) {
	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "is_default_fare_category",
			FileName:     "rider_categories.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "is_default_fare_category_validation",
		})
	}

	// Validate presence
	if riderCategory.IsDefaultFareCategory == nil {
		return
	}

	// Validate enum
	validOptions := []int{0, 1}
	if !slices.Contains(validOptions, *riderCategory.IsDefaultFareCategory) {
		addMessage(i18n.AppTranslator.Get("is_default_fare_category_validation.invalid"))
		return
	}
}

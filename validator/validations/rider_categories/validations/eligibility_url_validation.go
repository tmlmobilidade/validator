package rider_categories

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [rider_categories.txt]
  - Field: eligibility_url
  - Presence: optional
  - Type: URL

# Description

URL of a web page, usually from the operating agency, that provides detailed information about a specific rider category and/or describes its eligibility criteria.

[rider_categories.txt]: https://gtfs.org/schedule/reference/#rider_categoriestxt
*/

func EligibilityUrlValidation(riderCategory *types.RiderCategory, row int, gtfs *types.Gtfs, rules *types.RiderCategory) {
	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "eligibility_url",
			FileName:     "rider_categories.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "eligibility_url_validation",
		})
	}

	// Validate presence - optional field, so nil or empty is valid
	if riderCategory.EligibilityUrl == nil || *riderCategory.EligibilityUrl == "" {
		return
	}

	// Validate URL
	if !lib.ValidateUrl(*riderCategory.EligibilityUrl) {
		addMessage(i18n.AppTranslator.Get("eligibility_url_validation.invalid"))
		return
	}
}

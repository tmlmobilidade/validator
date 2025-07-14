package agency

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

  - File: [agency.txt]
  - Field: agency_url
  - Presence: Required
  - Type: URL

# Description

URL of the transit agency.

[agency.txt]: https://gtfs.org/schedule/reference/#agencytxt
*/
func AgencyUrlValidation(agency *types.Agency, row int, rules *types.AgencyRules) {

	addMessage := func(message string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "agency_url",
			FileName:     "agency.txt",
			Message:      message,
			Rows:         []int{row},
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "agency_url_validation",
		})
	}

	if agency.AgencyUrl == nil {
		addMessage(i18n.AppTranslator.Get("agency_url_validation.required"))
		return
	}

	if !lib.ValidateUrl(*agency.AgencyUrl) {
		addMessage(i18n.AppTranslator.Get("agency_url_validation.invalid"))
		return
	}

	// Validate rules
	if rules != nil && rules.AgencyUrl.Options != nil {
		if slices.Contains(*rules.AgencyUrl.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.AgencyUrl.Options, *agency.AgencyUrl) {
			addMessage(i18n.AppTranslator.Get("agency_url_validation.not_allowed", *agency.AgencyUrl))
			return
		}
	}
}

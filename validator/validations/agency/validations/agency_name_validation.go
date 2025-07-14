package agency

import (
	"main/i18n"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

  - File: [agency.txt]
  - Field: agency_name
  - Presence: Required
  - Type: String

# Description

Full name of the transit agency.

[agency.txt]: https://gtfs.org/schedule/reference/#agencytxt
*/
func AgencyNameValidation(agency *types.Agency, row int, rules *types.AgencyRules) {
	s := types.SEVERITY_ERROR

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "agency_name",
			FileName:     "agency.txt",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
			ValidationID: "agency_name_validation",
		})
	}

	// Check if agency_email is required
	if agency.AgencyName == nil {
		addMessage(i18n.AppTranslator.Get("agency_name_validation.required"), s)
		return
	}

	// Validate rules
	if rules != nil && rules.AgencyName.Options != nil {
		if slices.Contains(*rules.AgencyName.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.AgencyName.Options, *agency.AgencyName) {
			addMessage(i18n.AppTranslator.Get("agency_name_validation.not_allowed", *agency.AgencyName), s)
			return
		}
	}
}

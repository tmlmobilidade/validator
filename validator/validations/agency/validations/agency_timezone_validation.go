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
  - Field: agency_timezone
  - Presence: Required
  - Type: Timezone

# Description

Timezone where the transit agency is located.
If multiple agencies are specified in the dataset, each must have the same 'agency_timezone'.

[agency.txt]: https://gtfs.org/schedule/reference/#agencytxt
*/
func AgencyTimezoneValidation(agency *types.Agency, row int, rules *types.AgencyRules) {
	addMessage := func(message string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "agency_timezone",
			FileName:     "agency.txt",
			Message:      message,
			Rows:         []int{row},
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "agency_timezone_validation",
		})
	}

	if agency.AgencyTimezone == nil {
		addMessage(i18n.AppTranslator.Get("agency_timezone_validation.required"))
		return
	}

	if !lib.ValidateTimezone(*agency.AgencyTimezone) {
		addMessage(i18n.AppTranslator.Get("agency_timezone_validation.invalid"))
		return
	}

	// Validate rules
	if rules != nil && rules.AgencyTimezone.Options != nil {
		if slices.Contains(*rules.AgencyTimezone.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.AgencyTimezone.Options, *agency.AgencyTimezone) {
			addMessage(i18n.AppTranslator.Get("agency_timezone_validation.not_allowed", *agency.AgencyTimezone))
			return
		}
	}
}

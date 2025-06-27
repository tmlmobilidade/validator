package agency

import (
	"fmt"
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
			Field: "agency_timezone",
			FileName: "agency.txt",
			Message: message,
			Rows: []int{row},
			Severity: types.SEVERITY_ERROR,
			ValidationID: "agency_timezone_validation",
		})
	}

	if agency.AgencyTimezone == nil {
		addMessage("Agency timezone is required")
		return
	}

	err := lib.ValidateTimezone(*agency.AgencyTimezone)
	if err != "" {
		addMessage(err)
		return
	}

	// Validate rules
	if rules != nil && rules.AgencyTimezone.Options != nil {
		if slices.Contains(*rules.AgencyTimezone.Options, types.ALL_OPTIONS) {
			return
		}

		if slices.Contains(*rules.AgencyTimezone.Options, *agency.AgencyTimezone) {
			return
		}
		
		addMessage(fmt.Sprintf("Agency timezone is not allowed: %s", *agency.AgencyTimezone))
		return
	}
}
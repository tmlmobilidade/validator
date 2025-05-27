package agency

import (
	"main/lib"
	"main/services"
	"main/types"
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
func AgencyTimezoneValidation(agency *types.Agency, row int) {
	s := types.SEVERITY_ERROR

	// Check if agency_timezone is required
	if agency.AgencyTimezone == "" {
		services.AppMessageService.AddMessage(types.Message{
			Field: "agency_timezone",
			FileName: "agency.txt",
			Message: "Agency timezone is required",
			Rows: []int{row},
			Severity: s,
			ValidationID: "agency_timezone_validation",
		})
	}

	// Check if agency_timezone is valid
	err := lib.ValidateTimezone(agency.AgencyTimezone)
	if err != "" {
		services.AppMessageService.AddMessage(types.Message{
			Field: "agency_timezone",
			FileName: "agency.txt",
			Message: err,
			Rows: []int{row},
			Severity: s,
			ValidationID: "agency_timezone_validation",
		})
	}
}
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
			Field: "agency_url",
			FileName: "agency.txt",
			Message: message,
			Rows: []int{row},
			Severity: types.SEVERITY_ERROR,
			ValidationID: "agency_url_validation",
		})
	}

	if agency.AgencyUrl == nil {
		addMessage("Agency URL is required")
		return
	}

	err := lib.ValidateUrl(*agency.AgencyUrl)
	if err != "" {
		addMessage(err)
		return
	}

	// Validate rules
	if rules != nil && rules.AgencyUrl.Options != nil {
		if slices.Contains(*rules.AgencyUrl.Options, types.ALL_OPTIONS) {
			return
		}

		if slices.Contains(*rules.AgencyUrl.Options, *agency.AgencyUrl) {
			return
		}

		addMessage(fmt.Sprintf("Agency URL is not allowed: %s", *agency.AgencyUrl))
		return
	}
}
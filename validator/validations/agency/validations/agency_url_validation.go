package agency

import (
	"main/lib"
	"main/services"
	"main/types"
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
func AgencyUrlValidation(agency *types.Agency, row int) {
	
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
}
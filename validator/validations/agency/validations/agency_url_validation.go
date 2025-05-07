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
func AgencyUrlValidation(agency *types.Agency, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_ERROR

	// Check if agency_url is required
	if agency.AgencyUrl == "" {
		services.AppMessageService.AddMessage(types.Message{
			Field: "agency_url",
			FileName: "agency.txt",
			Message: "Agency URL is required",
			Rows: []int{row},
			Severity: s,
			ValidationID: "agency_url_validation",
		})
	}

	// Check if agency_url is valid
	err := lib.ValidateUrl(agency.AgencyUrl)
	if err != "" {
		services.AppMessageService.AddMessage(types.Message{
			Field: "agency_url",
			FileName: "agency.txt",
			Message: err,
			Rows: []int{row},
			Severity: s,
			ValidationID: "agency_url_validation",
		})
	}
}
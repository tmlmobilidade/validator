package agency

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

	- File: [agency.txt]
	- Field: agency_phone
	- Presence: Optional
	- Type: URL

# Description

URL of a web page where a rider can purchase tickets or other fare instruments for that agency, or a web page containing information about that agency's fares.

[agency.txt]: https://gtfs.org/schedule/reference/#agencytxt
*/
func AgencyFareUrlValidation(severity *types.Severity, agency *types.Agency, row int) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	// Check if agency_fare_url is required
	if agency.AgencyFareUrl == nil && s != types.SEVERITY_IGNORE {
		services.AppMessageService.AddMessage(types.Message{
			Field: "agency_fare_url",
			FileName: "agency.txt",
			Message: lib.IfThenElse(s == types.SEVERITY_ERROR, "Agency fare URL is required", "Agency fare URL is recommended"),
			Rows: []int{row},
			Severity: s,
			ValidationID: "agency_fare_url_validation",
		})
	}

	// Check if agency_fare_url is valid
	if agency.AgencyFareUrl != nil {
		if urlErrors := lib.ValidateUrl(*agency.AgencyFareUrl); urlErrors != "" {
			services.AppMessageService.AddMessage(types.Message{
				Field: "agency_fare_url",
				FileName: "agency.txt",
				Message: urlErrors,
				Rows: []int{row},
				Severity: types.SEVERITY_ERROR,
				ValidationID: "agency_fare_url_validation",
			})
		}
	}
}

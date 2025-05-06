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
func AgencyFareUrlValidation(severity *types.Severity, agency *types.Agency, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	// Check if agency_phone is required
	if agency.AgencyPhone == nil && s != types.SEVERITY_IGNORE {
		services.AppMessageService.AddMessage(types.Message{
			Field: "agency_phone",
			FileName: "agency.txt",
			Message: lib.IfThenElse(s == types.SEVERITY_ERROR, "Agency phone is required", "Agency phone is recommended"),
			Rows: []int{row},
			Severity: s,
			ValidationID: "agency_phone_validation",
		})
	}

	// Check if agency_phone is valid
	if agency.AgencyPhone != nil {
		if phoneErrors := lib.ValidatePhone(*agency.AgencyPhone); phoneErrors != "" {
			services.AppMessageService.AddMessage(types.Message{
				Field: "agency_phone",
				FileName: "agency.txt",
				Message: phoneErrors,
				Rows: []int{row},
				Severity: types.SEVERITY_ERROR,
				ValidationID: "agency_phone_validation",
			})
		}
	}
}

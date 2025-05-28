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
	- Type: Phone Number

# Description

A voice telephone number for the specified agency.
This field is a string value that presents the telephone number as typical for the agency's service area.
It may contain punctuation marks to group the digits of the number.
Dialable text (for example, TriMet's "503-238-RIDE") is permitted, but the field must not contain any other descriptive text.

[agency.txt]: https://gtfs.org/schedule/reference/#agencytxt
*/
func AgencyPhoneValidation(severity *types.Severity, agency *types.Agency, row int) {
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
}

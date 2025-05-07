package agency

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

	- File: [agency.txt]
	- Field: agency_email
	- Presence: Optional
	- Type: Email

# Description

Email address actively monitored by the agency’s customer service department.
This email address should be a direct contact point where transit riders can reach a customer service representative at the agency.

[agency.txt]: https://gtfs.org/schedule/reference/#agencytxt
*/
func AgencyEmailValidation(severity *types.Severity, agency *types.Agency, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	// Check if agency_phone is required
	if agency.AgencyEmail == nil && s != types.SEVERITY_IGNORE {
		services.AppMessageService.AddMessage(types.Message{
			Field: "agency_email",
			FileName: "agency.txt",
			Message: lib.IfThenElse(s == types.SEVERITY_ERROR, "Agency email is required", "Agency email is recommended"),
			Rows: []int{row},
			Severity: s,
			ValidationID: "agency_email_validation",
		})
	}

	// Check if agency_phone is valid
	if agency.AgencyEmail != nil {
		if emailErrors := lib.ValidateEmail(*agency.AgencyEmail); emailErrors != "" {
			services.AppMessageService.AddMessage(types.Message{
				Field: "agency_email",
				FileName: "agency.txt",
				Message: emailErrors,
				Rows: []int{row},
				Severity: types.SEVERITY_ERROR,
				ValidationID: "agency_email_validation",
			})
		}
	}
}

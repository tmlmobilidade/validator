package agency

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
$validation
id: agency.agency_email_validation
severity_options: [error, warning, ignore]
description: Validates if the agency email is present and valid.
name: Agency Email Validation
*/

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
func AgencyEmailValidation(severity *types.Severity, agency *types.Agency, row int) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field: "agency_email",
			FileName: "agency.txt",
			Message: msg,
			Rows: []int{row},
			Severity: severity,
			ValidationID: "agency.agency_email_validation",
		})
	}

	// Check if agency_email is required
	if agency.AgencyEmail == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := lib.IfThenElse(s == types.SEVERITY_ERROR, "Agency email is required", "Agency email is recommended")
		addMessage(message, s)
		return
	}

	// Check if agency_email is valid
	if emailErrors := lib.ValidateEmail(*agency.AgencyEmail); emailErrors != "" {
		addMessage(emailErrors, types.SEVERITY_ERROR)
		return
	}
}

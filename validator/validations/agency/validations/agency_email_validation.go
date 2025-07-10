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
  - Field: agency_email
  - Presence: Optional
  - Type: Email

# Description

Email address actively monitored by the agency’s customer service department.
This email address should be a direct contact point where transit riders can reach a customer service representative at the agency.

[agency.txt]: https://gtfs.org/schedule/reference/#agencytxt
*/
func AgencyEmailValidation(agency *types.Agency, row int, rules *types.AgencyRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.AgencyEmail.Severity != "" {
		s = rules.AgencyEmail.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "agency_email",
			FileName:     "agency.txt",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
			ValidationID: "agency_email_validation",
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

	// Validate rules
	if rules != nil && rules.AgencyEmail.Options != nil {
		if slices.Contains(*rules.AgencyEmail.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.AgencyEmail.Options, *agency.AgencyEmail) {
			addMessage(fmt.Sprintf("Agency email is not allowed: %s", *agency.AgencyEmail), s)
			return
		}
	}
}

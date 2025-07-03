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
func AgencyPhoneValidation(agency *types.Agency, row int, rules *types.AgencyRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.AgencyPhone.Severity != "" {
		s = rules.AgencyPhone.Severity
	}

	addMessage := func(message string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "agency_phone",
			FileName:     "agency.txt",
			Message:      message,
			Rows:         []int{row},
			Severity:     severity,
			ValidationID: "agency_phone_validation",
		})
	}

	// Check if agency_phone is required
	if agency.AgencyPhone == nil && s != types.SEVERITY_IGNORE {
		addMessage(lib.IfThenElse(s == types.SEVERITY_ERROR, "Agency phone is required", "Agency phone is recommended"), s)
	}

	// Validate rules
	if rules != nil && rules.AgencyPhone.Options != nil {
		if slices.Contains(*rules.AgencyPhone.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.AgencyPhone.Options, *agency.AgencyPhone) {
			addMessage(fmt.Sprintf("Agency phone is not allowed: %s", *agency.AgencyPhone), s)
			return
		}
	}
}

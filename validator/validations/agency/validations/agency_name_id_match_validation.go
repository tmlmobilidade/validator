package agency

import (
	"fmt"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [agency.txt]
  - Field: agency_name
  - Presence: Required
  - Type: String

# Description

Full name of the transit agency.

[agency.txt]: https://gtfs.org/schedule/reference/#agencytxt
*/
func AgencyNameIdMatchValidation(agency *types.Agency, row int, rules *types.AgencyRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.AgencyNameIdMatch.Severity != types.SEVERITY_IGNORE {
		s = rules.AgencyNameIdMatch.Severity
	}

	if s == types.SEVERITY_IGNORE {
		return
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "agency_name_id_match",
			FileName:     "agency.txt",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
			ValidationID: "agency_name_id_match_validation",
		})
	}

	// Check if agency_id matches agency_name
	if agency.AgencyId == nil || agency.AgencyName == nil {
		addMessage("Agency ID and name are required", s)
		return
	}

	// Validate rules
	if rules != nil && rules.AgencyNameIdMatch.Compare != nil {
		// Find the matching key and value
		for _, compare := range *rules.AgencyNameIdMatch.Compare {
			if compare.Key == *agency.AgencyId && compare.Value == *agency.AgencyName {
				return
			}
		}

		addMessage(fmt.Sprintf("Agency ID %s does not match agency name %s", *agency.AgencyId, *agency.AgencyName), s)
		return
	}
}

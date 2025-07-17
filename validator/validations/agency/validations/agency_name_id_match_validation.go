package agency

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [agency.txt]
  - Field: agency_name_id_match
  - Presence: Optional
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

	if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
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
		addMessage(i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"agency_name_id_match_validation.required",
				"agency_name_id_match_validation.recommended",
			),
		), s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("agency_name_id_match_validation.forbidden"), s)
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

		addMessage(i18n.AppTranslator.Get("agency_name_id_match_validation.no_match", *agency.AgencyId, *agency.AgencyName), s)
		return
	}
}

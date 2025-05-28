package agency

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

	- File: [agency.txt]
	- Field: agency_lang
	- Presence: Optional
	- Type: Language Code

# Description

Primary language used by this transit agency.
Should be provided to help GTFS consumers choose capitalization rules and other language-specific settings for the dataset.

[agency.txt]: https://gtfs.org/schedule/reference/#agencytxt
*/
func AgencyLangValidation(severity *types.Severity, agency *types.Agency, row int) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	// Check if agency_lang is required
	if agency.AgencyLang == nil && s != types.SEVERITY_IGNORE {
		services.AppMessageService.AddMessage(types.Message{
			Field: "agency_lang",
			FileName: "agency.txt",
			Message: lib.IfThenElse(s == types.SEVERITY_ERROR, "Agency language is required", "Agency language is recommended"),
			Rows: []int{row},
			Severity: s,
			ValidationID: "agency_lang_validation",
		})
	}

	// Check if agency_lang is valid
	if agency.AgencyLang != nil {
		if langErrors := lib.ValidateLanguage(*agency.AgencyLang); langErrors != "" {
			services.AppMessageService.AddMessage(types.Message{
				Field: "agency_lang",
				FileName: "agency.txt",
				Message: langErrors,
				Rows: []int{row},
				Severity: types.SEVERITY_ERROR,
				ValidationID: "agency_lang_validation",
			})
		}
	}
}

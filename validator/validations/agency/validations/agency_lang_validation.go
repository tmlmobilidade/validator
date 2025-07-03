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
  - Field: agency_lang
  - Presence: Optional
  - Type: Language Code

# Description

Primary language used by this transit agency.
Should be provided to help GTFS consumers choose capitalization rules and other language-specific settings for the dataset.

[agency.txt]: https://gtfs.org/schedule/reference/#agencytxt
*/
func AgencyLangValidation(agency *types.Agency, row int, rules *types.AgencyRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.AgencyLang.Severity != "" {
		s = rules.AgencyLang.Severity
	}

	addMessage := func(message string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "agency_lang",
			FileName:     "agency.txt",
			Message:      message,
			Rows:         []int{row},
			Severity:     severity,
			ValidationID: "agency_lang_validation",
		})
	}

	// Check if agency_lang is required
	if agency.AgencyLang == nil && s != types.SEVERITY_IGNORE {
		addMessage(lib.IfThenElse(s == types.SEVERITY_ERROR, "Agency language is required", "Agency language is recommended"), s)
	}

	// Check if agency_lang is valid
	if agency.AgencyLang != nil {
		if langErrors := lib.ValidateLanguage(*agency.AgencyLang); langErrors != "" {
			addMessage(langErrors, types.SEVERITY_ERROR)
		}
	}

	// Validate rules
	if rules != nil && rules.AgencyLang.Options != nil {
		if slices.Contains(*rules.AgencyLang.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.AgencyLang.Options, *agency.AgencyLang) {
			addMessage(fmt.Sprintf("Agency language is not allowed: %s", *agency.AgencyLang), s)
			return
		}
	}
}

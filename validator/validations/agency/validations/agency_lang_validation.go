package agency

import (
	"main/i18n"
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
		addMessage(i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"agency_lang_validation.required",
				"agency_lang_validation.recommended",
			),
		), s)
	}

	// Check if agency_lang is valid
	if agency.AgencyLang != nil && !lib.ValidateLanguage(*agency.AgencyLang) {
		addMessage(i18n.AppTranslator.Get("agency_lang_validation.invalid"), types.SEVERITY_ERROR)
		return
	}

	// Validate rules
	if rules != nil && rules.AgencyLang.Options != nil {
		if slices.Contains(*rules.AgencyLang.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.AgencyLang.Options, *agency.AgencyLang) {
			addMessage(i18n.AppTranslator.Get("agency_lang_validation.not_allowed", *agency.AgencyLang), s)
			return
		}
	}
}

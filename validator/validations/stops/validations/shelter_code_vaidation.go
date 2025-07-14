package stops

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: shelter_code
  - Presence: Optional
  - Type: String

# Description

Shelter code for a stop.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func ShelterCodeValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.ShelterCode.Severity != "" {
		s = rules.ShelterCode.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "shelter_code",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "shelter_code_validation",
		})
	}

	if stop.ShelterCode == nil || *stop.ShelterCode == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"shelter_code_validation.required",
				"shelter_code_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	// Validate rules
	if rules != nil && rules.ShelterCode.Options != nil {
		if slices.Contains(*rules.ShelterCode.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.ShelterCode.Options, *stop.ShelterCode) {
			addMessage(i18n.AppTranslator.Get("shelter_code_validation.not_allowed", *stop.ShelterCode), s)
			return
		}
	}
}

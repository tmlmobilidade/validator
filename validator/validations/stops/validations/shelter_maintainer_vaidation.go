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
  - Field: shelter_maintainer
  - Presence: Optional
  - Type: String

# Description

Shelter code for a stop.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func ShelterMaintainerValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.ShelterMaintainer.Severity != "" {
		s = rules.ShelterMaintainer.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "shelter_maintainer",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "shelter_maintainer_validation",
		})
	}

	if stop.ShelterMaintainer == nil || *stop.ShelterMaintainer == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"shelter_maintainer_validation.required",
				"shelter_maintainer_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	// Validate rules
	if rules != nil && rules.ShelterMaintainer.Options != nil {
		if slices.Contains(*rules.ShelterMaintainer.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.ShelterMaintainer.Options, *stop.ShelterMaintainer) {
			addMessage(i18n.AppTranslator.Get("shelter_maintainer_validation.not_allowed", *stop.ShelterMaintainer), s)
			return
		}
	}
}

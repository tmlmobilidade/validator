package stops

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
	"strconv"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: has_shelter
  - Presence: Optional
  - Type: Boolean

# Description

Describes if the stop has a shelter.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func HasShelterValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.HasShelter.Severity != "" {
		s = rules.HasShelter.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "has_shelter",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "has_shelter_validation",
		})
	}

	if stop.HasShelter == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"has_shelter_validation.required",
				"has_shelter_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	// Validate value
	validValues := []int{0, 1}
	if !slices.Contains(validValues, *stop.HasShelter) {
		addMessage(i18n.AppTranslator.Get("has_shelter_validation.invalid", *stop.HasShelter), types.SEVERITY_ERROR)
		return
	}

	// Validate value based on rules
	if rules != nil && rules.HasShelter.Options != nil {
		if slices.Contains(*rules.HasShelter.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.HasShelter.Options, strconv.Itoa(*stop.HasShelter)) {
			addMessage(i18n.AppTranslator.Get("has_shelter_validation.not_allowed", *stop.HasShelter), s)
			return
		}
	}
}

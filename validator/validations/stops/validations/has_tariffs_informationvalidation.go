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
  - Field: has_tariffs_information
  - Presence: Optional
  - Type: Enum

# Description

Describes if the stop has tariffs information.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func HasTariffsInformationValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.HasTariffsInformation.Severity != "" {
		s = rules.HasTariffsInformation.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "has_tariffs_information",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "has_tariffs_information_validation",
		})
	}

	if stop.HasTariffsInformation == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"has_tariffs_information_validation.required",
				"has_tariffs_information_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("has_tariffs_information_validation.forbidden"), s)
		return
	}

	// Validate value
	validValues := []int{0, 1, 2, 3}
	if !slices.Contains(validValues, *stop.HasTariffsInformation) {
		addMessage(i18n.AppTranslator.Get("has_tariffs_information_validation.invalid", *stop.HasTariffsInformation), types.SEVERITY_ERROR)
		return
	}

	// Validate value based on rules
	if rules != nil && rules.HasTariffsInformation.Options != nil {
		if slices.Contains(*rules.HasTariffsInformation.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.HasTariffsInformation.Options, strconv.Itoa(*stop.HasTariffsInformation)) {
			addMessage(i18n.AppTranslator.Get("has_tariffs_information_validation.not_allowed", *stop.HasTariffsInformation), s)
			return
		}
	}
}

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
  - Field: has_stop_sign
  - Presence: Optional
  - Type: Boolean

# Description

Describes if the stop has a stop sign.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func HasStopSignValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.HasStopSign.Severity != "" {
		s = rules.HasStopSign.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "has_stop_sign",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "has_stop_sign_validation",
		})
	}

	if stop.HasStopSign == nil {
		if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"has_stop_sign_validation.required",
				"has_stop_sign_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("has_stop_sign_validation.forbidden"), s)
		return
	}

	// Validate value
	validValues := []int{0, 1, 2, 3}
	if !slices.Contains(validValues, *stop.HasStopSign) {
		addMessage(i18n.AppTranslator.Get("has_stop_sign_validation.invalid", *stop.HasStopSign), types.SEVERITY_ERROR)
		return
	}

	// Validate Rule options
	if rules != nil && rules.HasStopSign.Options != nil {
		if slices.Contains(*rules.HasStopSign.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.HasStopSign.Options, strconv.Itoa(*stop.HasStopSign)) {
			addMessage(i18n.AppTranslator.Get("has_stop_sign_validation.not_allowed", *stop.HasStopSign), s)
			return
		}
	}
}

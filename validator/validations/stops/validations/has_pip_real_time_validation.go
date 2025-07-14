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
  - Field: has_pip_real_time
  - Presence: Optional
  - Type: Boolean

# Description

Describes if the stop has a network map.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func HasPipRealTimeValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.HasPipRealTime.Severity != "" {
		s = rules.HasPipRealTime.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "has_pip_real_time",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "has_pip_real_time_validation",
		})
	}

	if stop.HasPipRealTime == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"has_pip_real_time_validation.required",
				"has_pip_real_time_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("has_pip_real_time_validation.forbidden"), s)
		return
	}

	// Validate value
	validValues := []int{0, 1, 2}
	if !slices.Contains(validValues, *stop.HasPipRealTime) {
		addMessage(i18n.AppTranslator.Get("has_pip_real_time_validation.invalid", *stop.HasPipRealTime), types.SEVERITY_ERROR)
		return
	}

	// Validate Rule options
	if rules != nil && rules.HasPipRealTime.Options != nil {
		if slices.Contains(*rules.HasPipRealTime.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.HasPipRealTime.Options, strconv.Itoa(*stop.HasPipRealTime)) {
			addMessage(i18n.AppTranslator.Get("has_pip_real_time_validation.not_allowed", *stop.HasPipRealTime), s)
			return
		}
	}
}

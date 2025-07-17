/*
# Attributes

 - File: [stops.txt]
 - Field: stop_short_name
 - Presence: Optional
 - Type: String

# Description

The stop_short_name is an optional field that can be used to provide a short name for the stop.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
[translations.txt]: https://gtfs.org/schedule/reference/#translationstxt
*/

package stops

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

// StopShortNameValidation validates the presence of stop_short_name in stops.txt according to location_type
func StopShortNameValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.StopShortName.Severity != "" {
		s = rules.StopShortName.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_short_name",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "stop_short_name_validation",
		})
	}

	// 1. Check presence of stop_short_name based on severity
	if stop.StopShortName == nil {
		if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"stop_short_name_validation.required",
				"stop_short_name_validation.recommended",
			),
		)

		addMessage(message, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("stop_short_name_validation.forbidden"), s)
		return
	}

	// 2. Validate rules
	if rules != nil && rules.StopShortName.Options != nil {
		if slices.Contains(*rules.StopShortName.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.StopShortName.Options, *stop.StopShortName) {
			addMessage(i18n.AppTranslator.Get("stop_short_name_validation.not_allowed", *stop.StopShortName), s)
			return
		}
	}
}

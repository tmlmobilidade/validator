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
  - Field: stop_desc
  - Presence: Optional
  - Type: String

# Description

Description of the location that provides useful, quality information. Should not be a duplicate of stop_name.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func StopDescValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.StopDesc.Severity != "" {
		s = rules.StopDesc.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_desc",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "stop_desc_validation",
		})
	}

	if stop.StopDesc == nil || *stop.StopDesc == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"stop_desc_validation.required",
				"stop_desc_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	if stop.StopName != nil && *stop.StopName == *stop.StopDesc {
		addMessage(i18n.AppTranslator.Get("stop_desc_validation.duplicate"), types.SEVERITY_WARNING)
		return
	}

	// Validate rules
	if rules != nil && rules.StopDesc.Options != nil {
		if slices.Contains(*rules.StopDesc.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.StopDesc.Options, *stop.StopDesc) {
			addMessage(i18n.AppTranslator.Get("stop_desc_validation.not_allowed", *stop.StopDesc), s)
			return
		}
	}

}

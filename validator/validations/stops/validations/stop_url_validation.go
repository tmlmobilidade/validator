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
  - Field: stop_url
  - Presence: Optional
  - Type: URL

# Description

URL of the transit stop.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func StopUrlValidation(stop *types.Stop, row int, rules *types.StopsRules) {

	s := types.SEVERITY_IGNORE
	if rules != nil {
		s = rules.StopUrl.Severity
	}

	addMessage := func(message string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_url",
			FileName:     "stops.txt",
			Message:      message,
			Rows:         []int{row},
			Severity:     severity,
			ValidationID: "stop_url_validation",
		})
	}

	if stop.StopUrl == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"stop_url_validation.required",
				"stop_url_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("stop_url_validation.forbidden"), s)
		return
	}

	if !lib.ValidateUrl(*stop.StopUrl) {
		addMessage(i18n.AppTranslator.Get("stop_url_validation.invalid", *stop.StopUrl), types.SEVERITY_ERROR)
		return
	}

	// Validate rules
	if rules != nil && rules.StopUrl.Options != nil {
		if slices.Contains(*rules.StopUrl.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.StopUrl.Options, *stop.StopUrl) {
			addMessage(i18n.AppTranslator.Get("stop_url_validation.not_allowed", *stop.StopUrl), s)
			return
		}
	}
}

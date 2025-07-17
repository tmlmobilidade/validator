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
  - Field: tts_stop_name
  - Presence: Optional
  - Type: String

# Description

Readable version of the stop_name. See "Text-to-speech field" in the [Term Definitions] for more.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
[Term Definitions]: https://gtfs.org/schedule/reference/#term-definitions
*/
func TtsStopNameValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.TtsStopName.Severity != "" {
		s = rules.TtsStopName.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "tts_stop_name",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "tts_stop_name_validation",
		})
	}

	if stop.TtsStopName == nil || *stop.TtsStopName == "" {
		if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"tts_stop_name_validation.required",
				"tts_stop_name_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("tts_stop_name_validation.forbidden"), s)
		return
	}

	// Validate rules
	if rules != nil && rules.TtsStopName.Options != nil {
		if slices.Contains(*rules.TtsStopName.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.TtsStopName.Options, *stop.TtsStopName) {
			addMessage(i18n.AppTranslator.Get("tts_stop_name_validation.not_allowed", *stop.TtsStopName), s)
			return
		}
	}
}

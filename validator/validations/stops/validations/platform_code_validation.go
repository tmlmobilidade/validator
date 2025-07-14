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
  - Field: platform_code
  - Presence: Optional
  - Type: String

# Description

Platform identifier for a platform stop (a stop belonging to a station).
This should be just the platform identifier (eg. "G" or "3").
Words like "platform" or "track" (or the feed's language-specific equivalent) should not be included.
This allows feed consumers to more easily internationalize and localize the platform identifier into other languages.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func PlatformCodeValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.PlatformCode.Severity != "" {
		s = rules.PlatformCode.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "platform_code",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "platform_code_validation",
		})
	}

	if stop.PlatformCode == nil || *stop.PlatformCode == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"platform_code_validation.required",
				"platform_code_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	// Validate rules
	if rules != nil && rules.PlatformCode.Options != nil {
		if slices.Contains(*rules.PlatformCode.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.PlatformCode.Options, *stop.PlatformCode) {
			addMessage(i18n.AppTranslator.Get("platform_code_validation.not_allowed", *stop.PlatformCode), s)
			return
		}
	}
}

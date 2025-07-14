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
  - Field: parish_id
  - Presence: Optional
  - Type: String

# Description

Parish identifier for a stop.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func ParishIdValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.ParishId.Severity != "" {
		s = rules.ParishId.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "parish_id",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "parish_id_validation",
		})
	}

	if stop.ParishId == nil || *stop.ParishId == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"parish_id_validation.required",
				"parish_id_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	// Validate rules
	if rules != nil && rules.ParishId.Options != nil {
		if slices.Contains(*rules.ParishId.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.ParishId.Options, *stop.ParishId) {
			addMessage(i18n.AppTranslator.Get("parish_id_validation.not_allowed", *stop.ParishId), s)
			return
		}
	}
}

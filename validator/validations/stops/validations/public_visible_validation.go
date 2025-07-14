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
  - Field: public_visible
  - Presence: Optional
  - Type: Boolean

# Description

Describes if the stop is visible to the public.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func PublicVisibleValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.PublicVisible.Severity != "" {
		s = rules.PublicVisible.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "public_visible",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "public_visible_validation",
		})
	}

	if stop.PublicVisible == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"public_visible_validation.required",
				"public_visible_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	// Validate value
	validValues := []int{0, 1}
	if !slices.Contains(validValues, *stop.PublicVisible) {
		addMessage(i18n.AppTranslator.Get("public_visible_validation.invalid", *stop.PublicVisible), types.SEVERITY_ERROR)
		return
	}

	// Validate Rule options
	if rules != nil && rules.PublicVisible.Options != nil {
		if slices.Contains(*rules.PublicVisible.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.PublicVisible.Options, strconv.Itoa(*stop.PublicVisible)) {
			addMessage(i18n.AppTranslator.Get("public_visible_validation.not_allowed", *stop.PublicVisible), s)
			return
		}
	}
}

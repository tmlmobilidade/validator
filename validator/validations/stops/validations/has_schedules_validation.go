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
  - Field: has_schedules
  - Presence: Optional
  - Type: Boolean

# Description

Describes if the stop has schedules.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func HasSchedulesValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.HasSchedules.Severity != "" {
		s = rules.HasSchedules.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "has_schedules",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "has_schedules_validation",
		})
	}

	if stop.HasSchedules == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"has_schedules_validation.required",
				"has_schedules_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	// Validate value
	validValues := []int{0, 1}
	if !slices.Contains(validValues, *stop.HasSchedules) {
		addMessage(i18n.AppTranslator.Get("has_schedules_validation.invalid", *stop.HasSchedules), types.SEVERITY_ERROR)
		return
	}

	// Validate Rule options
	if rules != nil && rules.HasSchedules.Options != nil {
		if slices.Contains(*rules.HasSchedules.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.HasSchedules.Options, strconv.Itoa(*stop.HasSchedules)) {
			addMessage(i18n.AppTranslator.Get("has_schedules_validation.not_allowed", *stop.HasSchedules), s)
			return
		}
	}
}

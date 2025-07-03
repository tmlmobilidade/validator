package stops

import (
	"fmt"
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

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "has_schedules is required", "has_schedules is recommended")
		addMessage(warn, s)
		return
	}

	// Validate value
	validValues := []int{0, 1, 2}
	if !slices.Contains(validValues, *stop.HasSchedules) {
		addMessage("has_schedules must be 0, 1, or 2", types.SEVERITY_ERROR)
		return
	}

	// Validate Rule options
	if rules != nil && rules.HasSchedules.Options != nil {
		if slices.Contains(*rules.HasSchedules.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.HasSchedules.Options, strconv.Itoa(*stop.HasSchedules)) {
			addMessage(fmt.Sprintf("has_schedules is not allowed: %d", *stop.HasSchedules), s)
			return
		}
	}
}

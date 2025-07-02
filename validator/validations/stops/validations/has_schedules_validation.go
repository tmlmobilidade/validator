package stops

import (
	"main/lib"
	"main/services"
	"main/types"
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
}

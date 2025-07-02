package stops

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: has_pip_real_time
  - Presence: Optional
  - Type: Boolean

# Description

Describes if the stop has a network map.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func HasPipRealTimeValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.HasPipRealTime.Severity != "" {
		s = rules.HasPipRealTime.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "has_pip_real_time",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "has_pip_real_time_validation",
		})
	}

	if stop.HasPipRealTime == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "has_pip_real_time is required", "has_pip_real_time is recommended")
		addMessage(warn, s)
		return
	}
}

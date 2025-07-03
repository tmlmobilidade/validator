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

	// Validate value
	validValues := []int{0, 1, 2}
	if !slices.Contains(validValues, *stop.HasPipRealTime) {
		addMessage("has_pip_real_time must be 0, 1, or 2", types.SEVERITY_ERROR)
		return
	}

	// Validate Rule options
	if rules != nil && rules.HasPipRealTime.Options != nil {
		if slices.Contains(*rules.HasPipRealTime.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.HasPipRealTime.Options, strconv.Itoa(*stop.HasPipRealTime)) {
			addMessage(fmt.Sprintf("has_pip_real_time is not allowed: %d", *stop.HasPipRealTime), s)
			return
		}
	}
}

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
  - Field: has_stop_sign
  - Presence: Optional
  - Type: Boolean

# Description

Describes if the stop has a stop sign.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func HasStopSignValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.HasStopSign.Severity != "" {
		s = rules.HasStopSign.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "has_stop_sign",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "has_stop_sign_validation",
		})
	}

	if stop.HasStopSign == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "has_stop_sign is required", "has_stop_sign is recommended")
		addMessage(warn, s)
		return
	}

	// Validate value
	validValues := []int{0, 1, 2, 3}
	if !slices.Contains(validValues, *stop.HasStopSign) {
		addMessage("has_stop_sign must be 0, 1, 2, or 3", types.SEVERITY_ERROR)
		return
	}

	// Validate Rule options
	if rules != nil && rules.HasStopSign.Options != nil {
		if slices.Contains(*rules.HasStopSign.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.HasStopSign.Options, strconv.Itoa(*stop.HasStopSign)) {
			addMessage(fmt.Sprintf("has_stop_sign is not allowed: %d", *stop.HasStopSign), s)
			return
		}
	}
}

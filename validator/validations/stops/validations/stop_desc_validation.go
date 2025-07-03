package stops

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: stop_desc
  - Presence: Optional
  - Type: String

# Description

Description of the location that provides useful, quality information. Should not be a duplicate of stop_name.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func StopDescValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.StopDesc.Severity != "" {
		s = rules.StopDesc.Severity
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

	if stop.StopDesc == nil || *stop.StopDesc == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "tts_stop_name is required", "tts_stop_name is recommended")
		addMessage(warn, s)
		return
	}

	if stop.StopName != nil && *stop.StopName == *stop.StopDesc {
		addMessage("stop_desc should not be a duplicate of stop_name", types.SEVERITY_WARNING)
		return
	}

	// Validate rules
	if rules != nil && rules.StopDesc.Options != nil {
		if slices.Contains(*rules.StopDesc.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.StopDesc.Options, *stop.StopDesc) {
			addMessage(fmt.Sprintf("stop_desc is not allowed: %s", *stop.StopDesc), s)
			return
		}
	}

}

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
  - Field: stop_code
  - Presence: Optional
  - Type: String

# Description

Short text or a number that identifies the location for riders.

These codes are often used in phone-based transit information systems or printed on signage to make it easier for riders to get information for a particular location.

The `stop_code` may be the same as `stop_id` if it is public facing.

This field should be left empty for locations without a code presented to riders.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func StopCodeValidation(stop *types.Stop, row int, gtfs *types.Gtfs, rules *types.StopsRules) {

	s := types.SEVERITY_IGNORE
	if rules != nil && rules.StopCode.Severity != "" {
		s = rules.StopCode.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_code",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "stop_code_validation",
		})
	}

	if stop.StopCode == nil || *stop.StopCode == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "stop_code is required", "stop_code is recommended")
		addMessage(warn, s)
		return
	}

	// Check if stop_code is unique
	if stop.StopCode != nil {
		count := len(lib.RemoveDuplicates(gtfs.IdMap["stops"][*stop.StopCode]))

		if count > 1 {
			addMessage("Duplicate stop_code found: "+*stop.StopCode, types.SEVERITY_WARNING)
			return
		}
	}

	// Validate rules
	if rules != nil && rules.StopCode.Options != nil {
		if slices.Contains(*rules.StopCode.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.StopCode.Options, *stop.StopCode) {
			addMessage(fmt.Sprintf("stop_code is not allowed: %s", *stop.StopCode), s)
			return
		}
	}

}

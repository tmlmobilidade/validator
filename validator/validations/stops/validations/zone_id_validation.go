/*
# Attributes

 - File: [stops.txt]
 - Field: zone_id
 - Presence: Optional
 - Type: ID

# Description

Identifies the fare zone for a stop. If this record represents a station or station entrance, the `zone_id` is ignored.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/

package stops

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

// ZoneIdValidation validates the zone_id field in stops.txt
func ZoneIdValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.ZoneId.Severity != "" {
		s = rules.ZoneId.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "zone_id",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "zone_id_validation",
		})
	}

	if stop.ZoneId == nil || *stop.ZoneId == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "zone_id is required", "zone_id is recommended")
		addMessage(warn, s)
		return
	}

	// Validate rules
	if rules != nil && rules.ZoneId.Options != nil {
		if slices.Contains(*rules.ZoneId.Options, types.ALL_OPTIONS) {
			return
		}

		if slices.Contains(*rules.ZoneId.Options, *stop.ZoneId) {
			return
		}

		addMessage(fmt.Sprintf("zone_id is not allowed: %s", *stop.ZoneId), s)
		return
	}
}

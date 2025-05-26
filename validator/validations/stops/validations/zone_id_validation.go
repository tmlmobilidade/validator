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
	"main/lib"
	"main/services"
	"main/types"
)

// ZoneIdValidation validates the zone_id field in stops.txt
func ZoneIdValidation(severity *types.Severity, stop *types.Stop, row int) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
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
} 
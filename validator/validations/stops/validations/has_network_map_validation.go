package stops

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: has_network_map
  - Presence: Optional
  - Type: Boolean

# Description

Describes if the stop has a network map.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func HasNetworkMapValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.HasNetworkMap.Severity != "" {
		s = rules.HasNetworkMap.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "has_network_map",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "has_network_map_validation",
		})
	}

	if stop.HasNetworkMap == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "has_network_map is required", "has_network_map is recommended")
		addMessage(warn, s)
		return
	}
}

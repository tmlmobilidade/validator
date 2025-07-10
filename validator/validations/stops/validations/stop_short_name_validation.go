/*
# Attributes

 - File: [stops.txt]
 - Field: stop_name
 - Presence: Conditionally Required
 - Type: String

# Description

Name of the location. The stop_name should match the agency's rider-facing name for the location as printed on a timetable, published online, or represented on signage. For translations into other languages, use [translations.txt].

When the location is a boarding area (location_type=4), the stop_name should contains the name of the boarding area as displayed by the agency. It could be just one letter (like on some European intercity railway stations), or text like "Wheelchair boarding area" (NYC's Subway) or "Head of short trains" (Paris' RER).

Conditionally Required:

  - Required for locations which are stops (location_type=0), stations (location_type=1) or entrances/exits (location_type=2).
  - Optional for locations which are generic nodes (location_type=3) or boarding areas (location_type=4).

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
[translations.txt]: https://gtfs.org/schedule/reference/#translationstxt
*/

package stops

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

// StopNameValidation validates the presence of stop_name in stops.txt according to location_type
func StopShortNameValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.StopShortName.Severity != "" {
		s = rules.StopShortName.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_short_name",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "stop_short_name_validation",
		})
	}

	// If stop_short_name is present, return
	if stop.StopShortName != nil && *stop.StopShortName != "" {
		return
	}

	// Check presence of stop_short_name based on location_type
	locationType := -1
	if stop.LocationType != nil {
		locationType = *stop.LocationType
	}

	if locationType == 0 || locationType == 1 || locationType == 2 {
		addMessage("stop_short_name is required when location_type is 0, 1, or 2", s)
		return
	}

	// Check presence of stop_short_name based on severity
	if s != types.SEVERITY_IGNORE {
		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "stop_short_name is required", "stop_short_name is recommended")
		addMessage(warn, s)
		return
	}

	// Validate rules
	if rules != nil && rules.StopShortName.Options != nil {
		if slices.Contains(*rules.StopShortName.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.StopShortName.Options, *stop.StopShortName) {
			addMessage(fmt.Sprintf("stop_short_name is not allowed: %s", *stop.StopShortName), s)
			return
		}
	}
}

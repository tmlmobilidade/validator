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
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

// StopNameValidation validates the presence of stop_name in stops.txt according to location_type
func StopNameValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.StopName.Severity != "" {
		s = rules.StopName.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_name",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "stop_name_validation",
		})
	}

	// If stop_name is present, return
	if stop.StopName != nil && *stop.StopName != "" {
		return
	}

	// Check presence of stop_name based on location_type
	locationType := -1
	if stop.LocationType != nil {
		locationType = *stop.LocationType
	}

	if locationType == 0 || locationType == 1 || locationType == 2 {
		addMessage(i18n.AppTranslator.Get("stop_short_name_validation.required_location_type"), s)
		return
	}

	// Check presence of stop_name based on severity
	if s != types.SEVERITY_IGNORE {
		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"stop_short_name_validation.required",
				"stop_short_name_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	// Validate rules
	if rules != nil && rules.StopName.Options != nil {
		if slices.Contains(*rules.StopName.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.StopName.Options, *stop.StopName) {
			addMessage(i18n.AppTranslator.Get("stop_short_name_validation.not_allowed", *stop.StopName), s)
			return
		}
	}
}

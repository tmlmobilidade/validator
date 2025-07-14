package stops

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: stop_lat
  - Presence: Conditionally Required
  - Type: Latitude

# Description

Latitude of the location.

For stops/platforms (location_type=0) and boarding area (location_type=4), the coordinates must be the ones of the bus pole — if exists — and otherwise of where the travelers are boarding the vehicle (on the sidewalk or the platform, and not on the roadway or the track where the vehicle stops).

Conditionally Required:
  - Required for locations which are stops (location_type=0), stations (location_type=1) or entrances/exits (location_type=2).
  - Optional for locations which are generic nodes (location_type=3) or boarding areas (location_type=4).

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func StopLatValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.StopLat.Severity != "" {
		s = rules.StopLat.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_lat",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "stop_lat_validation",
		})
	}

	locationType := -1
	if stop.LocationType != nil {
		locationType = *stop.LocationType
	}

	isRequired := locationType == 0 || locationType == 1 || locationType == 2

	if stop.StopLat == nil {
		if s == types.SEVERITY_IGNORE && !isRequired {
			return
		}

		if isRequired {
			addMessage(i18n.AppTranslator.Get("stop_lat_validation.required_location_type"), types.SEVERITY_ERROR)
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"stop_lat_validation.required",
				"stop_lat_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	if !lib.ValidateLatitude(*stop.StopLat) {
		addMessage(i18n.AppTranslator.Get("stop_lat_validation.invalid", *stop.StopLat), types.SEVERITY_ERROR)
		return
	}
}

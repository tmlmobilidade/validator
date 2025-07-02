package stops

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

 - File: [stops.txt]
 - Field: stop_lon
 - Presence: Conditionally Required
 - Type: Lonitude

# Description

Longitude of the location.

For stops/platforms (location_type=0) and boarding area (location_type=4), the coordinates must be the ones of the bus pole — if exists — and otherwise of where the travelers are boarding the vehicle (on the sidewalk or the platform, and not on the roadway or the track where the vehicle stops).

Conditionally Required:
  - Required for locations which are stops (location_type=0), stations (location_type=1) or entrances/exits (location_type=2).
  - Optional for locations which are generic nodes (location_type=3) or boarding areas (location_type=4).

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func StopLonValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.StopLon.Severity != "" {
		s = rules.StopLon.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_lon",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "stop_lon_validation",
		})
	}

	locationType := -1
	if stop.LocationType != nil {
		locationType = *stop.LocationType
	}

	isRequired := locationType == 0 || locationType == 1 || locationType == 2

	if stop.StopLon == nil {
		if s == types.SEVERITY_IGNORE && !isRequired {
			return
		}

		if isRequired {
			addMessage("stop_lon is required when location_type is 0, 1, or 2", types.SEVERITY_ERROR)
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "stop_lon is required", "stop_lon is recommended")
		addMessage(warn, s)
		return
	}
	
	err := lib.ValidateLongitude(*stop.StopLon)
	if err != "" {
		addMessage(err, types.SEVERITY_ERROR)
		return
	}
} 
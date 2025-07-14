package stops

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: parent_station
  - Presence: Optional
  - Type: Foreign ID referencing stops.stop_id

# Description

Defines hierarchy between the different locations defined in stops.txt.

It contains the ID of the parent location, as follows:

  - Stop/platform (location_type=0): the parent_station field contains the ID of a station.
  - Station (location_type=1): this field must be empty.
  - Entrance/exit (location_type=2) or generic node (location_type=3): the parent_station field contains the ID of a station (location_type=1)
  - Boarding Area (location_type=4): the parent_station field contains ID of a platform.

Conditionally Required:

  - Required for locations which are entrances (location_type=2), generic nodes (location_type=3) or boarding areas (location_type=4).
  - Optional for stops/platforms (location_type=0).
  - Forbidden for stations (location_type=1).

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func ParentStationValidation(stop *types.Stop, row int, gtfs types.Gtfs, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.ParentStation.Severity != "" {
		s = rules.ParentStation.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "parent_station",
			FileName:     "stops.txt",
			ValidationID: "parent_station_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	locationType := 0
	if stop.LocationType != nil {
		locationType = *stop.LocationType
	}

	// Handle Nil Parent Station
	if stop.ParentStation == nil {

		// Handle Severity
		if s != types.SEVERITY_IGNORE {
			message := i18n.AppTranslator.Get(
				lib.IfThenElse(s == types.SEVERITY_ERROR,
					"parent_station_validation.required",
					"parent_station_validation.recommended",
				),
			)
			addMessage(message, s)
			return
		}

		// Allow nil parent_station for location_type=0 (Stop/Platform)
		if locationType == 0 || locationType == 1 {
			return
		}
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("parent_station_validation.forbidden"), s)
		return
	}

	// Validate Parent Station for Location Type 1 (Station)
	if locationType == 1 && stop.ParentStation != nil {
		addMessage(i18n.AppTranslator.Get("parent_station_validation.forbidden"), types.SEVERITY_ERROR)
		return
	}

	// Validate Parent Station for Location Type 2 (Entrance/Exit), 3 (Generic Node), or 4 (Boarding Area)
	if (locationType == 2 || locationType == 3 || locationType == 4) && stop.ParentStation == nil {
		addMessage(i18n.AppTranslator.Get("parent_station_validation.required_location_type"), types.SEVERITY_ERROR)
		return
	}

	// Validate Foreign Key
	if !lib.GtfsIdMapKeyExists(&gtfs, "stops", *stop.ParentStation) {
		addMessage(i18n.AppTranslator.Get("parent_station_validation.not_found", *stop.ParentStation), types.SEVERITY_ERROR)
		return
	}

	// Validate rules
	if rules != nil && rules.ParentStation.Options != nil {
		if slices.Contains(*rules.ParentStation.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.ParentStation.Options, *stop.ParentStation) {
			addMessage(i18n.AppTranslator.Get("parent_station_validation.not_allowed", *stop.ParentStation), s)
			return
		}
	}
}

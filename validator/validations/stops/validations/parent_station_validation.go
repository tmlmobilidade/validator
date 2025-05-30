package stops

import (
	"main/lib"
	"main/services"
	"main/types"
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
func ParentStationValidation(severity *types.Severity, stop *types.Stop, row int, gtfs types.Gtfs) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
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
			warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "parent_station is required", "parent_station is recommended")
			addMessage(warn, s)
			return
		}

		// Allow nil parent_station for location_type=0 (Stop/Platform)
		if locationType == 0 || locationType == 1 {
			return
		}
	}

	// Validate Parent Station for Location Type 1 (Station)
	if locationType == 1 && stop.ParentStation != nil {
		addMessage("parent_station is forbidden for stations", types.SEVERITY_ERROR)
		return
	}

	// Validate Parent Station for Location Type 2 (Entrance/Exit), 3 (Generic Node), or 4 (Boarding Area)
	if (locationType == 2 || locationType == 3 || locationType == 4) && stop.ParentStation == nil {
		addMessage("parent_station is required for location_type=2 (Entrance/Exit), 3 (Generic Node), or 4 (Boarding Area)", types.SEVERITY_ERROR)
		return
	}

	// Validate Foreign Key
	if !lib.GtfsIdMapKeyExists(&gtfs, "stops", *stop.ParentStation) {
		addMessage("parent_station '"+ *stop.ParentStation + "' does not exist in stops.txt", types.SEVERITY_ERROR)
		return
	}
}
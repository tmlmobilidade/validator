package stops

import (
	"main/lib"
	"main/services"
	"main/types"
)

func ParentStationValidation(severity *types.Severity, stop *types.Stop, row int) {
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
		if locationType == 0 {
			return
		}
	}

	if locationType == 1 && stop.ParentStation != nil {
		addMessage("parent_station is forbidden for stations", types.SEVERITY_ERROR)
		return
	}

	if (locationType == 2 || locationType == 3 || locationType == 4) && stop.ParentStation == nil {
		addMessage("parent_station is required for location_type=2 (Entrance/Exit), 3 (Generic Node), or 4 (Boarding Area)", types.SEVERITY_ERROR)
		return
	}
}
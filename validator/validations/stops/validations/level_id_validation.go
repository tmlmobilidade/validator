package stops

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

 - File: [stops.txt]
 - Field: level_id
 - Presence: Optional
 - Type: Foreign ID referencing levels.level_id

# Description

Level of the location. The same level may be used by multiple unlinked stations.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func LevelIdValidation(severity *types.Severity, stop *types.Stop, row int, gtfs types.Gtfs) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "level_id",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "level_id_validation",
		})
	}

	if stop.LevelId == nil || *stop.LevelId == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "level_id is required", "level_id is recommended")
		addMessage(warn, s)
		return
	}

	// Check Foreign Key
	if !lib.GtfsIdMapKeyExists(&gtfs, "levels", *stop.LevelId) {
		addMessage("level_id '"+ *stop.LevelId + "' does not exist in levels.txt", types.SEVERITY_ERROR)
		return
	}
}
package stops

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

 - File: [stops.txt]
 - Field: stop_desc
 - Presence: Optional
 - Type: String

# Description

Description of the location that provides useful, quality information. Should not be a duplicate of stop_name.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func StopDescValidation(severity *types.Severity, stop *types.Stop, row int) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "tts_stop_name",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "tts_stop_name_validation",
		})
	}


	if stop.StopDesc == nil || *stop.StopDesc == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "tts_stop_name is required", "tts_stop_name is recommended")
		addMessage(warn, s)
		return
	}

	if stop.StopName != nil && *stop.StopName == *stop.StopDesc {
		addMessage("stop_desc should not be a duplicate of stop_name", types.SEVERITY_WARNING)
		return
	}
	
}
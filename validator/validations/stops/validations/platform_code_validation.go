package stops

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

 - File: [stops.txt]
 - Field: platform_code
 - Presence: Optional
 - Type: String

# Description

Platform identifier for a platform stop (a stop belonging to a station).
This should be just the platform identifier (eg. "G" or "3").
Words like “platform” or "track" (or the feed’s language-specific equivalent) should not be included.
This allows feed consumers to more easily internationalize and localize the platform identifier into other languages.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func PlatformCodeValidation(severity *types.Severity, stop *types.Stop, row int) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "platform_code",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "platform_code_validation",
		})
	}

	if stop.PlatformCode == nil || *stop.PlatformCode == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "platform_code is required", "platform_code is recommended")
		addMessage(warn, s)
		return
	}
}
package stops

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

  - File: [stop.txt]
  - Field: stop_url
  - Presence: Optional
  - Type: URL

# Description

URL of the transit stop.

[stop.txt]: https://gtfs.org/schedule/reference/#stoptxt
*/
func StopUrlValidation(stop *types.Stop, row int, rules *types.StopsRules) {

	s := types.SEVERITY_IGNORE
	if rules != nil {
		s = rules.StopUrl.Severity
	}

	addMessage := func(message string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_url",
			FileName:     "stop.txt",
			Message:      message,
			Rows:         []int{row},
			Severity:     severity,
			ValidationID: "stop_url_validation",
		})
	}

	if stop.StopUrl == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "stop_url is required", "stop_url is recommended")
		addMessage(warn, s)
		return
	}

	err := lib.ValidateUrl(*stop.StopUrl)
	if err != "" {
		addMessage(err, types.SEVERITY_ERROR)
		return
	}

	// Validate rules
	if rules != nil && rules.StopUrl.Options != nil {
		if slices.Contains(*rules.StopUrl.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.StopUrl.Options, *stop.StopUrl) {
			addMessage(fmt.Sprintf("stop_url is not allowed: %s", *stop.StopUrl), s)
			return
		}
	}
}

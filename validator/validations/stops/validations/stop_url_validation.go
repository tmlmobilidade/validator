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
  - Presence: Required
  - Type: URL

# Description

URL of the transit stop.

[stop.txt]: https://gtfs.org/schedule/reference/#stoptxt
*/
func StopUrlValidation(stop *types.Stop, row int, rules *types.StopsRules) {

	addMessage := func(message string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_url",
			FileName:     "stop.txt",
			Message:      message,
			Rows:         []int{row},
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "stop_url_validation",
		})
	}

	if stop.StopUrl == nil {
		addMessage("Stop URL is required")
		return
	}

	err := lib.ValidateUrl(*stop.StopUrl)
	if err != "" {
		addMessage(err)
		return
	}

	// Validate rules
	if rules != nil && rules.StopUrl.Options != nil {
		if slices.Contains(*rules.StopUrl.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.StopUrl.Options, *stop.StopUrl) {
			addMessage(fmt.Sprintf("stop_url is not allowed: %s", *stop.StopUrl))
			return
		}
	}
}

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

  - File: [stops.txt]
  - Field: shelter_maintainer
  - Presence: Optional
  - Type: String

# Description

Shelter code for a stop.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func ShelterMaintainerValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.ShelterMaintainer.Severity != "" {
		s = rules.ShelterMaintainer.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "shelter_maintainer",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "shelter_maintainer_validation",
		})
	}

	if stop.ShelterMaintainer == nil || *stop.ShelterMaintainer == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "shelter_maintainer is required", "shelter_maintainer is recommended")
		addMessage(warn, s)
		return
	}

	// Validate rules
	if rules != nil && rules.ShelterMaintainer.Options != nil {
		if slices.Contains(*rules.ShelterMaintainer.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.ShelterMaintainer.Options, *stop.ShelterMaintainer) {
			addMessage(fmt.Sprintf("shelter_maintainer is not allowed: %s", *stop.ShelterMaintainer), s)
			return
		}
	}
}

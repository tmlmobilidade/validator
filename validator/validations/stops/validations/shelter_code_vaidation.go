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
  - Field: shelter_code
  - Presence: Optional
  - Type: String

# Description

Shelter code for a stop.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func ShelterCodeValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.ShelterCode.Severity != "" {
		s = rules.ShelterCode.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "shelter_code",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "shelter_code_validation",
		})
	}

	if stop.ShelterCode == nil || *stop.ShelterCode == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "shelter_code is required", "shelter_code is recommended")
		addMessage(warn, s)
		return
	}

	// Validate rules
	if rules != nil && rules.ShelterCode.Options != nil {
		if slices.Contains(*rules.ShelterCode.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.ShelterCode.Options, *stop.ShelterCode) {
			addMessage(fmt.Sprintf("shelter_code is not allowed: %s", *stop.ShelterCode), s)
			return
		}
	}
}

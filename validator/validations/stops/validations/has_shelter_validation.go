package stops

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
	"strconv"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: has_shelter
  - Presence: Optional
  - Type: Boolean

# Description

Describes if the stop has a shelter.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func HasShelterValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.HasShelter.Severity != "" {
		s = rules.HasShelter.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "has_shelter",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "has_shelter_validation",
		})
	}

	if stop.HasShelter == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "has_shelter is required", "has_shelter is recommended")
		addMessage(warn, s)
		return
	}

	// Validate value
	validValues := []int{0, 1, 2, 3}
	if !slices.Contains(validValues, *stop.HasShelter) {
		addMessage("has_shelter must be 0, 1, 2, or 3", types.SEVERITY_ERROR)
		return
	}

	// Validate value based on rules
	if rules != nil && rules.HasShelter.Options != nil {
		if slices.Contains(*rules.HasShelter.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.HasShelter.Options, strconv.Itoa(*stop.HasShelter)) {
			addMessage(fmt.Sprintf("has_shelter is not allowed: %d", *stop.HasShelter), types.SEVERITY_ERROR)
			return
		}
	}
}

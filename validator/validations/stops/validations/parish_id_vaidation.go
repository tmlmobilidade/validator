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
  - Field: parish_id
  - Presence: Optional
  - Type: String

# Description

Parish identifier for a stop.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func ParishIdValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.ParishId.Severity != "" {
		s = rules.ParishId.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "parish_id",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "parish_id_validation",
		})
	}

	if stop.ParishId == nil || *stop.ParishId == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "parish_id is required", "parish_id is recommended")
		addMessage(warn, s)
		return
	}

	// Validate rules
	if rules != nil && rules.ParishId.Options != nil {
		if slices.Contains(*rules.ParishId.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.ParishId.Options, *stop.ParishId) {
			addMessage(fmt.Sprintf("parish_id is not allowed: %s", *stop.ParishId), s)
			return
		}
	}
}

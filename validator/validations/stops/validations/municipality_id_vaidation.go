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
  - Field: municipality_id
  - Presence: Optional
  - Type: String

# Description

Municipality identifier for a stop.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func MunicipalityIdValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.MunicipalityId.Severity != "" {
		s = rules.MunicipalityId.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "municipality_id",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "municipality_id_validation",
		})
	}

	if stop.MunicipalityId == nil || *stop.MunicipalityId == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "municipality_id is required", "municipality_id is recommended")
		addMessage(warn, s)
		return
	}

	// Validate rules
	if rules != nil && rules.MunicipalityId.Options != nil {
		if slices.Contains(*rules.MunicipalityId.Options, types.ALL_OPTIONS) {
			return
		}

		if slices.Contains(*rules.MunicipalityId.Options, *stop.MunicipalityId) {
			return
		}

		addMessage(fmt.Sprintf("This municipality_id is not allowed: %s", *stop.MunicipalityId), types.SEVERITY_ERROR)
		return
	}
}

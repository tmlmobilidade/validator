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
  - Field: region_id
  - Presence: Optional
  - Type: String

# Description

Region identifier for a stop.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func RegionIdValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.RegionId.Severity != "" {
		s = rules.RegionId.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "region_id",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "region_id_validation",
		})
	}

	if stop.RegionId == nil || *stop.RegionId == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "region_id is required", "region_id is recommended")
		addMessage(warn, s)
		return
	}

	// Validate rules
	if rules != nil && rules.RegionId.Options != nil {
		if slices.Contains(*rules.RegionId.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.RegionId.Options, *stop.RegionId) {
			addMessage(fmt.Sprintf("region_id is not allowed: %s", *stop.RegionId), s)
			return
		}
	}
}

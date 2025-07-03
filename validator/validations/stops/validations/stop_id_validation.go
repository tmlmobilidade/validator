/*
# Attributes

 - File: [stops.txt]
 - Field: stop_id
 - Presence: Required
 - Type: Unique ID

# Description

Identifies a location: stop/platform, station, entrance/exit, generic node or boarding area (see location_type).

ID must be unique across all stops.stop_id, locations.geojson id, and location_groups.location_group_id values.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/

package stops

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

// StopIdValidation validates the presence and uniqueness of stop_id in stops.txt
func StopIdValidation(stop *types.Stop, row int, gtfs *types.Gtfs, rules *types.StopsRules) {
	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_id",
			FileName:     "stops.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "stop_id_validation",
		})
	}

	// Check if stop_id is missing
	if stop.StopId == nil || *stop.StopId == "" {
		addMessage("Missing required stop_id.")
	}

	// Check if stop_id is unique
	if stop.StopId != nil {
		count := len(lib.RemoveDuplicates(gtfs.IdMap["stops"][*stop.StopId]))

		if count > 1 {
			addMessage("Duplicate stop_id found: " + *stop.StopId)
			return
		}
	}

	// Validate rules
	if rules != nil && rules.StopId.Options != nil {
		if slices.Contains(*rules.StopId.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.StopId.Options, *stop.StopId) {
			addMessage(fmt.Sprintf("stop_id is not allowed: %s", *stop.StopId))
			return
		}
	}
}

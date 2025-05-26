package stop_times

import (
	"main/services"
	"main/types"
)

/*
# Attributes

 - File: [stop_times.txt]
 - Field: stop_id
 - Presence: Conditionally Required
 - Type: Foreign ID referencing stops.stop_id

# Description

Identifies the serviced stop. All stops serviced during a trip must have a record in [stop_times.txt].

Referenced locations must be stops/platforms, i.e. their `stops.location_type` value must be `0` or `empty`. A stop may be serviced multiple times in the same trip, and multiple trips and routes may service the same stop.

On-demand service using stops should be referenced in the sequence in which service is available at those stops.

A data consumer should assume that travel is possible from one stop or location to any stop or location later in the trip, provided that the `pickup/drop_off_type` of each `stop_time` and the time constraints of each `start/end_pickup_drop_off_window` do not forbid it.

Conditionally Required:

  - Required if stop_times.location_group_id AND stop_times.location_id are NOT defined.
  - Forbidden if stop_times.location_group_id or stop_times.location_id are defined.

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func StopIdValidation(stopTime *types.StopTime, row int, gtfs *types.Gtfs) {
	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "stop_id",
			FileName:     "stop_times.txt",
			ValidationID: "stop_id_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	// Forbidden if location_group_id or location_id are defined
	if (stopTime.LocationGroupId != nil && *stopTime.LocationGroupId != "") || (stopTime.LocationId != nil && *stopTime.LocationId != "") {
		if stopTime.StopId != nil && *stopTime.StopId != "" {
			addMessage("stop_id is forbidden if location_group_id or location_id are defined.", types.SEVERITY_ERROR)
		}
		return
	}

	// Required if location_group_id AND location_id are NOT defined
	if (stopTime.LocationGroupId == nil || *stopTime.LocationGroupId == "") && (stopTime.LocationId == nil || *stopTime.LocationId == "") {
		if stopTime.StopId == nil || *stopTime.StopId == "" {
			addMessage("stop_id is required if location_group_id and location_id are not defined.", types.SEVERITY_ERROR)
			return
		}

		// Foreign key check: must reference a valid stop_id from stops.txt
		stopsMap, ok := gtfs.IdMap["stops"]
		if !ok || stopsMap == nil {
			addMessage("stops.txt is missing or not indexed.", types.SEVERITY_ERROR)
			return
		}
		rows, ok := stopsMap[*stopTime.StopId]
		if !ok || len(rows) == 0 {
			addMessage("stop_id must reference a valid stop_id from stops.txt.", types.SEVERITY_ERROR)
			return
		}

		// Check location_type is 0 or empty
		rowIdx := rows[0]
		stopRow := gtfs.Files["stops"][rowIdx]
		locationTypeStr := stopRow["location_type"]
		if locationTypeStr != "" && locationTypeStr != "0" {
			addMessage("stop_id must reference a stop with location_type 0 or empty (platform/stop).", types.SEVERITY_ERROR)
			return
		}
	}
} 
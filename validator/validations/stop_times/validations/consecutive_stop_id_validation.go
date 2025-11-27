package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [stop_times.txt]
  - Field: stop_id
  - Rule: ISO 1083

# Description

Não deve ser possível ter a mesma paragem, diretamente uma depois da outra.

The same stop_id cannot appear consecutively (one directly after another) in stop_times.txt within the same trip.

# Example

Invalid:
stop_id,stop_sequence
111, 1
222, 2
333, 3
333, 4  <- Error: same stop_id (333) appears consecutively
444, 5

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func ConsecutiveStopIdValidation(stopTime *types.StopTime, row int, previousStopIdByTrip map[string]*string) {
	ctx := lib.NewValidationContext("stop_id", "stop_times.txt", "consecutive_stop_id_validation", row, services.AppMessageService)

	// Skip if stop_id is not defined (using location_group_id or location_id instead)
	if stopTime.StopId == nil || *stopTime.StopId == "" {
		return
	}

	// Skip if trip_id is not defined
	if stopTime.TripId == nil || *stopTime.TripId == "" {
		return
	}

	tripId := *stopTime.TripId
	currentStopId := *stopTime.StopId

	// Check if previous stop_id exists for this trip
	if previousStopId, exists := previousStopIdByTrip[tripId]; exists && previousStopId != nil {
		if *previousStopId == currentStopId {
			ctx.AddError(ctx.GetTranslatedMessage("consecutive_stop_id_validation.consecutive_stop_ids", currentStopId, tripId))
			// Continue to update map so next validation can check against this stop_id
		}
	}

	// Update the previous stop_id for this trip (allocate new string to avoid pointer reuse issues)
	stopIdCopy := new(string)
	*stopIdCopy = currentStopId
	previousStopIdByTrip[tripId] = stopIdCopy
}

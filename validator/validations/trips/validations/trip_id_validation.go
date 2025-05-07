package trips

import (
	"main/services"
	"main/types"
)

/*
# Attributes

	- File: [trips.txt]
	- Field: trip_id
	- Presence: Required
	- Type: Unique ID

# Description

Identifies a trip.

[trips.txt]: https://gtfs.org/schedule/reference/#trips
*/
func TripIdValidation(trip *types.Trip, row int, gtfs *types.Gtfs) {

	message := types.Message{
		Field: "trip_id",
		FileName: "trips.txt",
		Message: "trip_id is required",
		Rows: []int{row},
		Severity: types.SEVERITY_ERROR,
		ValidationID: "trip_id_validation",
	}

	if trip.TripId != "" {
		// Check if trip_id is Unique ID
		if gtfs.IdMap["trips"] != nil && len(gtfs.IdMap["trips"][trip.TripId]) > 1 {
			message.Message = "Duplicate trip_id found. Trip IDs must be unique."
			message.Severity = types.SEVERITY_ERROR
			services.AppMessageService.AddMessage(message)
		}
		
		return;
	}
	
	services.AppMessageService.AddMessage(message)
}
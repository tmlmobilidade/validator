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

	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "trip_id",
			FileName:     "trips.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "trip_id_validation",
		})
	}
	
	if trip.TripId == nil {
		addMessage("trip_id is required")
		return
	}
	
	if gtfs.IdMap["trips"] != nil && len(gtfs.IdMap["trips"][*trip.TripId]) > 1 {
		addMessage("Duplicate trip_id found. Trip IDs must be unique.")
		return
	}
}
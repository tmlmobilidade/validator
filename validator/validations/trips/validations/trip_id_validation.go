package trips

import (
	"main/i18n"
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
		addMessage(i18n.AppTranslator.Get("trip_id_validation.required"))
		return
	}

	rows, err := gtfs.GetRowsById("trips", *trip.TripId)
	if err == nil && len(rows) > 1 {
		addMessage(i18n.AppTranslator.Get("trip_id_validation.duplicate", map[string]interface{}{"trip_id": *trip.TripId}))
		return
	}
}

package stop_times

import (
	"main/i18n"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [stop_times.txt]
- Field: trip_id
- Presence: Required
- Type: Foreign ID referencing trips.trip_id

# Description

Identifies a trip.

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func TripIdValidation(stopTime *types.StopTime, row int, gtfs *types.Gtfs) {
	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "trip_id",
			FileName:     "stop_times.txt",
			ValidationID: "trip_id_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	if stopTime.TripId == nil || *stopTime.TripId == "" {
		addMessage(i18n.AppTranslator.Get("trip_id_validation.required"), types.SEVERITY_ERROR)
		return
	}

	if _, ok := gtfs.IdMap["trips"][*stopTime.TripId]; !ok {
		addMessage(i18n.AppTranslator.Get("trip_id_validation.not_found", *stopTime.TripId), types.SEVERITY_ERROR)
		return
	}
}

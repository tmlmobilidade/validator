package trips

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [trips.txt]
  - Field: service_id
  - Presence: Required
  - Type: Foreign Key referencing calendar.service_id or calendar_dates.service_id

# Description

Identifies a service.

[trips.txt]: https://gtfs.org/schedule/reference/#trips
*/
func ServiceIdValidation(trip *types.Trip, row int, gtfs *types.Gtfs) {
	message := types.Message{
		Field:        "service_id",
		FileName:     "trips.txt",
		Message:      i18n.AppTranslator.Get("service_id_validation.required"),
		Rows:         []int{row},
		Severity:     types.SEVERITY_ERROR,
		ValidationID: "service_id_validation",
	}

	if trip.ServiceId == nil {
		message.Message = i18n.AppTranslator.Get("service_id_validation.required")
		services.AppMessageService.AddMessage(message)
		return
	}

	// Merge calendar and calendar_dates id maps
	serviceIds := lib.MergeMaps(gtfs.IdMap["calendar"], gtfs.IdMap["calendar_dates"])
	if _, ok := serviceIds[*trip.ServiceId]; !ok {
		message.Message = i18n.AppTranslator.Get("service_id_validation.not_found", map[string]interface{}{"service_id": *trip.ServiceId})
		services.AppMessageService.AddMessage(message)
	}
}

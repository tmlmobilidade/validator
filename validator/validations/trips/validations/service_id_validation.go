package trips

import (
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
		Field: "service_id",
		FileName: "trips.txt",
		Message: "Service ID is required",
		Rows: []int{row},
		Severity: types.SEVERITY_ERROR,
		ValidationID: "service_id_validation",
	}

	if trip.ServiceId == nil {
		message.Message = "Service ID is required"
		services.AppMessageService.AddMessage(message)
		return;
	}

	// Merge calendar and calendar_dates id maps
	serviceIds := lib.MergeMaps(gtfs.IdMap["calendar"], gtfs.IdMap["calendar_dates"])
	if _, ok := serviceIds[*trip.ServiceId]; !ok {
		message.Message = "Service ID must reference a valid service_id from calendar.txt or calendar_dates.txt."
		services.AppMessageService.AddMessage(message)
	}
}
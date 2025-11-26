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
	ctx := lib.NewValidationContext("service_id", "trips.txt", "service_id_validation", row, services.AppMessageService)

	if trip.ServiceId == nil {
		ctx.AddError(ctx.GetTranslatedMessage("service_id_validation.required"))
		return
	}

	// Check in calendar or calendar_dates without merging the maps
	calendarRows, err := gtfs.GetRowsById("calendar", *trip.ServiceId)
	if err == nil && len(calendarRows) > 0 {
		return
	}
	calendarDatesRows, err := gtfs.GetRowsById("calendar_dates", *trip.ServiceId)
	if err == nil && len(calendarDatesRows) > 0 {
		return
	}
	ctx.AddError(ctx.GetTranslatedMessage("service_id_validation.not_found", map[string]interface{}{"service_id": *trip.ServiceId}))
}

package trips

import (
	"main/i18n"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [trips.txt]
  - Field: route_id
  - Presence: Required
  - Type: Foreign Key referencing routes.route_id

# Description

Identifies a route.

[trips.txt]: https://gtfs.org/schedule/reference/#trips
*/
func RouteIdValidation(trip *types.Trip, row int, gtfs *types.Gtfs) {
	message := types.Message{
		Field:        "route_id",
		FileName:     "trips.txt",
		Message:      i18n.AppTranslator.Get("route_id_validation.required"),
		Rows:         []int{row},
		Severity:     types.SEVERITY_ERROR,
		ValidationID: "route_id_validation",
	}

	if trip.RouteId == nil {
		message.Message = i18n.AppTranslator.Get("route_id_validation.required")
		services.AppMessageService.AddMessage(message)
		return
	}

	// Check if route_id is Foreign Key referencing routes.route_id
	if _, ok := gtfs.IdMap["routes"][*trip.RouteId]; !ok {
		message.Message = i18n.AppTranslator.Get("route_id_validation.not_found", map[string]interface{}{"route_id": *trip.RouteId})
		services.AppMessageService.AddMessage(message)
	}
}

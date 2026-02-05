package trips

import (
	"main/lib"
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
	ctx := lib.NewValidationContext("route_id", "trips.txt", "route_id_validation", row, services.AppMessageService)

	if trip.RouteId == nil {
		ctx.AddError(ctx.GetTranslatedMessage("route_id_validation.required"))
		return
	}

	// Check if route_id is Foreign Key referencing routes.route_id
	rows, err := gtfs.GetRowsById("routes", *trip.RouteId)
	if err != nil || len(rows) == 0 {
		ctx.AddError(ctx.GetTranslatedMessage("route_id_validation.not_found", map[string]interface{}{"route_id": *trip.RouteId}))
	}

	// Check if route_id is Foreign Key referencing routes.route_id
	// if !lib.GtfsIdMapKeyExists(gtfs, "routes", *trip.RouteId) {
	// 	ctx.AddError(ctx.GetTranslatedMessage("route_id_validation.not_found", map[string]interface{}{"route_id": *trip.RouteId}))
	// }
}

package trips

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes
	- File: [trips.txt]
	- Field: shape_id
	- Presence: Conditionally Required
	- Type: Foreign Key referencing shapes.shape_id

# Description

Identifies a geospatial shape describing the vehicle travel path for a trip.

Conditionally Required:
	- Required if the trip has a continuous pickup or drop-off behavior defined either in routes.txt or in stop_times.txt.
	- Optional otherwise.

[trips.txt]: https://gtfs.org/schedule/reference/#tripstxt
*/
func ShapeIdValidation(severity *types.Severity, trip *types.Trip, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(message string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field: "shape_id",
			FileName: "trips.txt",
			Rows: []int{row},
			Severity: severity,
			Message: message,
			ValidationID: "shape_id_validation",
		})
	}

	hasContinuousPickupDropoff := false
	
	// Check if the route has continuous pickup/dropoff behavior
	routeRow := gtfs.IdMap["routes"][*trip.RouteId]
	if gtfs.IdMap["routes"] != nil && gtfs.Route[routeRow[0]].ContinuousPickup != "" {
		hasContinuousPickupDropoff = true
	}

	// Check if the stop_times have continuous pickup/dropoff behavior
	if gtfs.IdMap["stop_times"] != nil && len(gtfs.IdMap["stop_times"][*trip.TripId]) > 0 && !hasContinuousPickupDropoff {

		for _, rowIndex := range gtfs.IdMap["stop_times"][*trip.TripId] {
			if continuousPickup := gtfs.StopTime[rowIndex].ContinuousPickup; continuousPickup != "" {
				hasContinuousPickupDropoff = true
				break // Exit early once we find a continuous pickup
			}
		}
	}

	if hasContinuousPickupDropoff && trip.ShapeId == nil {
		addMessage("shape_id is required when a continuous pickup or drop-off behavior is defined either in routes.txt or in stop_times.txt.", types.SEVERITY_ERROR)
		return
	}

	if trip.ShapeId == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := lib.IfThenElse(s == types.SEVERITY_ERROR, "shape_id is required", "shape_id is recommended")
		addMessage(message, s)
		return
	}	

	// Check Foreign Key
	if !lib.GtfsIdMapKeyExists(gtfs, "shapes", *trip.ShapeId) {
		addMessage("shape_id '"+ *trip.ShapeId + "' does not exist in shapes.txt", types.SEVERITY_ERROR)
		return
	}
}

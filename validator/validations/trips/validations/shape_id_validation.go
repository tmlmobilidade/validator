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

	message := types.Message{
		Field: "shape_id",
		FileName: "trips.txt",
		Rows: []int{row},
		Severity: s,
		ValidationID: "shape_id_validation",
	}

	hasContinuousPickupDropoff := false
	
	// Check if the route has continuous pickup/dropoff behavior
	routeRow := gtfs.IdMap["routes"][trip.RouteId]
	if gtfs.IdMap["routes"] != nil && gtfs.Files["routes"][routeRow[0]]["continuous_pickup"] != "" {
		hasContinuousPickupDropoff = true
	}

	// Check if the stop_times have continuous pickup/dropoff behavior
	if gtfs.IdMap["stop_times"] != nil && len(gtfs.IdMap["stop_times"][trip.TripId]) > 0 && !hasContinuousPickupDropoff {
		startRow := gtfs.IdMap["stop_times"][trip.TripId][0]
		endRow := gtfs.IdMap["stop_times"][trip.TripId][len(gtfs.IdMap["stop_times"][trip.TripId])-1]

		for i := startRow; i <= endRow; i++ {
			if gtfs.Files["stop_times"][i]["continuous_pickup"] != "" {
				hasContinuousPickupDropoff = true
			}
		}
	}

	if hasContinuousPickupDropoff && trip.ShapeId == nil {
		message.Message = "shape_id is required when a continuous pickup or drop-off behavior is defined either in routes.txt or in stop_times.txt."
		message.Severity = types.SEVERITY_ERROR

		services.AppMessageService.AddMessage(message)
		return
	}

	if trip.ShapeId != nil && s != types.SEVERITY_IGNORE {
		message.Message = lib.IfThenElse(s == types.SEVERITY_ERROR, "shape_id is required", "shape_id is recommended")
		services.AppMessageService.AddMessage(message)
		return
	}
	
	
}

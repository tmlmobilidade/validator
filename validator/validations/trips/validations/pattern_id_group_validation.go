package trips

import (
	"fmt"
	"main/services"
	"main/types"
)

/*
# Attributes

	- File: trips.txt
	- Field: pattern_id
	- Presence: Optional (Required for "Transportes Metropolitanos de Lisboa")
	- Type: ID

# Description

Validates if trips with the same pattern_id have the same route_id, trip_headsign, direction_id, shape_id and the same stop sequence.
*/
func PatternIdGroupValidation(tripsGroupedByPattern types.TripGroupedByPattern, gtfs *types.Gtfs) {
	addMessage := func(msg string, row int, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field: "pattern_id",
			FileName: "trips.txt",
			Message: msg,
			Rows: []int{row},
			Severity: severity,
			ValidationID: "pattern_id_validation",
		})
	}

	for patternId, group := range tripsGroupedByPattern {
		if len(group.Trips) == 0 {
			panic("trips is empty")
		}

		routeId := group.Trips[0].RouteId
		directionId := group.Trips[0].DirectionId
		shapeId := group.Trips[0].ShapeId

		for _, trip := range group.Trips {
			//check if route_id is the same
			if *trip.RouteId != *routeId {
				addMessage(fmt.Sprintf("For pattern_id %s, route_id %s is not the same as %s found in row %d", patternId, *trip.RouteId, *routeId, trip.Row), trip.Row, types.SEVERITY_ERROR)
				continue
			}

			//check if direction_id is the same
			if *trip.DirectionId != *directionId {
				addMessage(fmt.Sprintf("For pattern_id %s, direction_id %v is not the same as %v found in row %d", patternId, *trip.DirectionId, *directionId, trip.Row), trip.Row, types.SEVERITY_ERROR)
				continue
			}

			//check if shape_id is the same
			if *trip.ShapeId != *shapeId {
				addMessage(fmt.Sprintf("For pattern_id %s, shape_id %v is not the same as %v found in row %d", patternId, *trip.ShapeId, *shapeId, trip.Row), trip.Row, types.SEVERITY_ERROR)
				continue
			}
		}

		if len(group.Hash) > 1 {
			addMessage(fmt.Sprintf("For pattern_id %s, there are trips with multiple stop sequence variations", patternId), group.Trips[0].Row, types.SEVERITY_ERROR)
		}
	}
	
}
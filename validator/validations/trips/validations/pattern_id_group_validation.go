package trips

import (
	"main/lib"
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
	for patternId, group := range tripsGroupedByPattern {
		if len(group.Trips) == 0 {
			panic("trips is empty")
		}

		routeId := group.Trips[0].RouteId
		directionId := group.Trips[0].DirectionId
		shapeId := group.Trips[0].ShapeId

		for _, trip := range group.Trips {
			ctx := lib.NewValidationContext("pattern_id", "trips.txt", "pattern_id_validation", trip.Row, services.AppMessageService)

			//check if route_id is the same
			if *trip.RouteId != *routeId {
				ctx.AddError(ctx.GetTranslatedMessage("pattern_id_validation.different_route_id", patternId))
				continue
			}

			//check if direction_id is the same
			if *trip.DirectionId != *directionId {
				ctx.AddError(ctx.GetTranslatedMessage("pattern_id_validation.different_direction_id", patternId))
				continue
			}

			//check if shape_id is the same
			if *trip.ShapeId != *shapeId {
				ctx.AddError(ctx.GetTranslatedMessage("pattern_id_validation.different_shape_id", patternId))
				continue
			}
		}

		if len(group.Hash) > 1 {
			ctx := lib.NewValidationContext("pattern_id", "trips.txt", "pattern_id_validation", group.Trips[0].Row, services.AppMessageService)
			ctx.AddError(ctx.GetTranslatedMessage("pattern_id_validation.multiple_stop_sequence_variations", patternId))
		}
	}

}

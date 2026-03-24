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
	// Group trips by pattern_id and validate the group
	for patternId, group := range tripsGroupedByPattern {
		if len(group.Trips) == 0 {
			panic("trips is empty")
		}

		// 1. Validate shape_id, route_id and direction_id
		for _, trip := range group.Trips {
			ctx := lib.NewValidationContext("pattern_id", "trips.txt", "pattern_id_validation", trip.Row, services.AppMessageService)

			if trip.ShapeId == nil {
				ctx.AddError(ctx.GetTranslatedMessage("pattern_id_group_validation.shape_id_required"))
				continue
			}

			if trip.RouteId == nil {
				ctx.AddError(ctx.GetTranslatedMessage("pattern_id_group_validation.route_id_required"))
				continue
			}

			if trip.DirectionId == nil {
				ctx.AddError(ctx.GetTranslatedMessage("pattern_id_group_validation.direction_id_required"))
				continue
			}
		}

		// 2. Validate multiple stop sequence variations
		if len(group.Hash) > 1 {
			ctx := lib.NewValidationContext("pattern_id", "trips.txt", "pattern_id_validation", group.Trips[0].Row, services.AppMessageService)
			ctx.AddError(ctx.GetTranslatedMessage("pattern_id_group_validation.multiple_stop_sequence_variations", patternId))
		}
	}
}

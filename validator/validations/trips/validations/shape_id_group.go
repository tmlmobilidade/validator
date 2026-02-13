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
  - Presence: Optional (Required if pattern_id and shape_id are present)
  - Type: ID

# Description

Validates if the shape_id is in a unique pattern_id.

A shape_id must not be reused across different pattern_id values.
If a shape_id is already associated with one pattern_id, it cannot appear in a trip with a different pattern_id.
*/

func ShapeIdGroupValidation(tripsGroupedByShapeId types.TripGroupedByShapeId, gtfs *types.Gtfs) {
	for shapeId, group := range tripsGroupedByShapeId {
		if len(group.Trips) == 0 {
			panic("trips is empty")
		}

		patternId := group.Trips[0].PatternId
		shapeIdValue := group.Trips[0].ShapeId
		if shapeIdValue == nil {
			return
		}

		for _, trip := range group.Trips {
			ctx := lib.NewValidationContext("shape_id", "trips.txt", "shape_id_in_unique_pattern_id_validation", trip.Row, services.AppMessageService)
			// check if pattern_id is present
			if trip.PatternId == nil {
				lib.AppLogger.Accent("Pattern ID not found for shape_id: ")
				ctx.AddError(ctx.GetTranslatedMessage("shape_id_in_unique_pattern_id_validation.pattern_id_not_found"))
				return
			}

			// check if pattern_id is different and shape_id is the same
			if *trip.PatternId != *patternId {
				lib.AppLogger.Accent("Different pattern ID found for shape_id: ", shapeId)
				ctx.AddError(ctx.GetTranslatedMessage("shape_id_in_unique_pattern_id_validation.different_pattern_id", shapeId, *patternId, *trip.PatternId))
				return
			}
		}

		if len(group.Hash) > 1 {
			lib.AppLogger.Accent("Multiple pattern IDs found for shape_id: ", shapeId)
			ctx := lib.NewValidationContext("shape_id", "trips.txt", "shape_id_in_unique_pattern_id_validation", group.Trips[0].Row, services.AppMessageService)
			ctx.AddError(ctx.GetTranslatedMessage("shape_id_in_unique_pattern_id_validation.multiple_shape_id_in_unique_pattern_id", shapeId))
		}
	}
}

package trips

import (
	"main/types"
)

/*
Orchestrates pattern group validations to avoid duplicate errors.
Runs validations in optimal order and coordinates reporting between
route_id, direction_id, pattern_id, and shape_id group validators.
*/

func ValidatePatternGroups(
	tripsGroupedByPattern types.TripGroupedByPattern,
	tripsGroupedByShapeId types.TripGroupedByShapeId,
	gtfs *types.Gtfs,
) {
	PatternIdGroupValidation(tripsGroupedByPattern, gtfs)
	RouteIdGroupValidation(tripsGroupedByPattern, gtfs)
	DirectionIdGroupValidation(tripsGroupedByPattern, gtfs)
	ShapeIdGroupValidation(tripsGroupedByPattern, tripsGroupedByShapeId, gtfs)
}

package stops

import (
	"main/lib"
	"main/services"
	shapes_coordinates "main/services/geo/shapes"
	"main/types"
)

/*
# Attributes
  - File: [stops.txt]
  - Field: coordinates
  - Presence: optional
  - Type: coordinates

# Description
Validate if the stop_lat and stop_lon are valid.
*/

func StopCoordinatesValidation(stop *types.Stop, row int, stopClosestShapeInfo map[string]shapes_coordinates.StopClosestShapeInfo, rules *types.StopsRules) {
	ctx := lib.NewValidationContext("coordinates", "stops.txt", "coordinates_validation", row, services.AppMessageService)
	if rules != nil && rules.StopCoordinates.Severity != "" {
		ctx.WithSeverity(rules.StopCoordinates.Severity)
	}

	// Other validations already handle mandatory presence and format checks.
	if stop.StopLat == nil || stop.StopLon == nil {
		return
	}

	if stop.StopId == nil || *stop.StopId == "" || stopClosestShapeInfo == nil {
		return
	}

	info, ok := stopClosestShapeInfo[*stop.StopId]
	if !ok {
		return
	}

	if info.DistanceMeters > shapes_coordinates.MaxStopDistanceToClosestShapeMeters {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("coordinates_validation.invalid_distance_to_shape", *stop.StopLat, *stop.StopLon, info.ShapeID, info.ClosestShapePtSeq, info.ClosestShapePtLat, info.ClosestShapePtLon, info.DistanceMeters))
	}
}

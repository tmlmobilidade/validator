package stops

import (
	"main/lib"
	"main/services"
	municipality_coordinates "main/services/geo/municipalities"
	shapes_coordinates "main/services/geo/shapes"
	"main/types"
)

/*
# Attributes
  - File: [stops.txt]
  - Field: coordenates
  - Presence: optional
  - Type: coordenates

# Description
Validate if the stop_lat and stop_lon are valid.
*/

func CoordenatesValidation(stop *types.Stop, row int, rules *types.StopsRules, stopClosestShapeDistance map[string]float64) {
	ctx := lib.NewValidationContext("coordenates", "stops.txt", "coordenates_validation", row, services.AppMessageService)

	// Other validations already handle mandatory presence and format checks.
	if stop.StopLat == nil || stop.StopLon == nil {
		return
	}

	if municipality_coordinates.MunicipalityCoordinatesEnabled() && stop.MunicipalityId != nil && *stop.MunicipalityId != "" {
		expectedMunicipalityID, found, _ := municipality_coordinates.ResolveMunicipalityByCoordinates(*stop.StopLat, *stop.StopLon)
		if !found {
			ctx.AddError(ctx.GetTranslatedMessage("coordenates_validation.not_mapped", *stop.StopLat, *stop.StopLon))
			return
		}

		if expectedMunicipalityID != *stop.MunicipalityId {
			ctx.AddError(ctx.GetTranslatedMessage("coordenates_validation.invalid_municipality_id", *stop.StopLat, *stop.StopLon, expectedMunicipalityID, *stop.MunicipalityId))
		}
	}

	if stop.StopId == nil || *stop.StopId == "" || stopClosestShapeDistance == nil {
		return
	}

	minDistanceMeters, ok := stopClosestShapeDistance[*stop.StopId]
	if !ok {
		return
	}

	if minDistanceMeters > shapes_coordinates.MaxStopDistanceToClosestShapeMeters {
		ctx.AddError(ctx.GetTranslatedMessage("coordenates_validation.invalid_distance_to_shape", *stop.StopLat, *stop.StopLon))
	}
}

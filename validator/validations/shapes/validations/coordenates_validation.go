package shapes

import (
	"main/lib"
	"main/services"
	municipality_coordinates "main/services/geo/municipalities"
	"main/types"
)

/*
# Attributes
  - File: [shapes.txt]
  - Field: coordenates
  - Presence: Required
  - Type: Foreign Key

# Description

Foreign key to check if shape_lat and shape_lon are valid.
*/

func CoordenatesValidation(shape *types.Shape, row int, rules *types.ShapesRules) {
	ctx := lib.NewValidationContext("coordenates", "shapes.txt", "coordenates_validation", row, services.AppMessageService)

	// This validation is only meaningful when the coordinates map is loaded.
	if !municipality_coordinates.MunicipalityCoordinatesEnabled() {
		return
	}

	// Other validations already handle mandatory presence and format checks.
	if shape.ShapePtLat == nil || shape.ShapePtLon == nil {
		return
	}

	if _, found, _ := municipality_coordinates.ResolveMunicipalityByCoordinates(*shape.ShapePtLat, *shape.ShapePtLon); !found {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("coordenates_validation.not_mapped", *shape.ShapePtLat, *shape.ShapePtLon))
		return
	}
}

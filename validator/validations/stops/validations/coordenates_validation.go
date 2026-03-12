package stops

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes
  - File: [stops.txt]
  - Field: coordenates
  - Presence: Required
  - Type: Foreign Key

# Description

Foreign key to check if stop_lat and stop_lon are valid.
*/

func CoordenatesValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	ctx := lib.NewValidationContext("coordenates", "stops.txt", "coordenates_validation", row, services.AppMessageService)

	// This validation is only meaningful when the coordinates map is loaded.
	if !services.MunicipalityCoordinatesEnabled() {
		return
	}

	// Other validations already handle mandatory presence and format checks.
	if stop.StopLat == nil || stop.StopLon == nil || stop.MunicipalityId == nil || *stop.MunicipalityId == "" {
		return
	}

	expectedMunicipalityID, found, _ := services.ResolveMunicipalityByCoordinates(*stop.StopLat, *stop.StopLon)
	if !found {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("coordenates_validation.not_mapped", *stop.StopLat, *stop.StopLon))
		return
	}

	if expectedMunicipalityID != *stop.MunicipalityId {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("coordenates_validation.invalid_municipality_id", *stop.StopLat, *stop.StopLon, expectedMunicipalityID, *stop.MunicipalityId))
	}
}

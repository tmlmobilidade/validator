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

	lib.AppLogger.Accent("checking if stop is mapped to a municipality")
	// Other validations already handle mandatory presence and format checks.
	if stop.StopLat == nil || stop.StopLon == nil || stop.MunicipalityId == nil || *stop.MunicipalityId == "" {
		return
	}

	lib.AppLogger.Accent("checking if stop is mapped to a municipality")
	expectedMunicipalityID, found, _ := services.ResolveMunicipalityByCoordinates(*stop.StopLat, *stop.StopLon)
	if !found {
		ctx.AddError(ctx.GetTranslatedMessage("coordenates_validation.not_mapped", *stop.StopLat, *stop.StopLon))
		return
	}

	if expectedMunicipalityID != *stop.MunicipalityId {
		ctx.AddError(ctx.GetTranslatedMessage("coordenates_validation.invalid_municipality_id", *stop.StopLat, *stop.StopLon, expectedMunicipalityID, *stop.MunicipalityId))
	}
}

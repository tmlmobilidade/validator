package stops

import (
	"encoding/json"
	"main/lib"
	"main/services"
	municipality_coordinates "main/services/geo/municipalities"
	"main/types"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: municipality_id
  - Presence: Optional
  - Type: String

# Description

Municipality identifier for a stop.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func MunicipalityIdValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	ctx := lib.NewValidationContext("municipality_id", "stops.txt", "municipality_id_validation", row, services.AppMessageService)
	if rules != nil && rules.MunicipalityId.Severity != "" {
		ctx.WithSeverity(rules.MunicipalityId.Severity)
	}

	if stop.MunicipalityId == nil || *stop.MunicipalityId == "" {
		if ctx.ShouldSkip() {
			return
		}

		message := ctx.GetRequiredMessage("municipality_id_validation.required", "municipality_id_validation.recommended")
		ctx.AddMessageWithSeverity(message)
		return
	}

	if ctx.IsForbidden() {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("municipality_id_validation.forbidden"))
		return
	}

	if municipality_coordinates.MunicipalityIDFromRaw(json.RawMessage(*stop.MunicipalityId)) != *stop.MunicipalityId {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("municipality_id_validation.not_found", *stop.MunicipalityId))
		return
	}
}

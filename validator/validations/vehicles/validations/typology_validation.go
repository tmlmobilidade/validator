package vehicles

import (
	"main/lib"
	"main/services"
	"main/types"
	"slices"
	"strconv"
)

/*
# Attributes
  - File: [vehicles.txt]
  - Field: typology
  - Presence: Required
  - Type: Enum

# Description

The typology of the vehicle.
*/

func TypologyValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("typology", "vehicles.txt", "typology_validation", row, services.AppMessageService)
	if rules != nil && rules.Typology.Severity != "" {
		ctx.WithSeverity(rules.Typology.Severity)
	}

	if vehicle.Typology == nil {
		ctx.AddError(ctx.GetTranslatedMessage("typology_validation.required"))
		return
	}

	validOptions := []float32{0.1, 0.2, 0.3, 1.1, 1.2, 1.3, 2.1, 2.2, 2.3, 3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 4.1, 4.2, 4.3, 7.1, 7.2, 7.3}
	typology, err := strconv.ParseFloat(*vehicle.Typology, 32)
	if err != nil {
		ctx.AddError(ctx.GetTranslatedMessage("typology_validation.invalid", *vehicle.Typology))
		return
	}
	if !slices.Contains(validOptions, float32(typology)) {
		ctx.AddError(ctx.GetTranslatedMessage("typology_validation.invalid", float32(typology)))
		return
	}
}

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
  - Field: bicycles
  - Presence: Required
  - Type: Enum

# Description

Permission to transport non-folding bicycles (folding bicycles are always allowed).

Valid options are:

  - 0 - No
  - 1 - Yes
*/
func BicyclesValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("bicycles", "vehicles.txt", "bicycles_validation", row, services.AppMessageService)
	if rules != nil && rules.Bicycles.Severity != "" {
		ctx.WithSeverity(rules.Bicycles.Severity)
	}

	if vehicle.Bicycles == nil {
		ctx.AddError(ctx.GetTranslatedMessage("bicycles_validation.required"))
		return
	}

	validOptions := []int{0, 1}
	if !slices.Contains(validOptions, *vehicle.Bicycles) {
		ctx.AddError(ctx.GetTranslatedMessage("bicycles_validation.invalid", strconv.Itoa(*vehicle.Bicycles)))
		return
	}
}

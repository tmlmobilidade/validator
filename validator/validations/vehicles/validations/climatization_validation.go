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
  - Field: climatization
  - Presence: Required
  - Type: Enum

# Description

The climatization of the vehicle.

Valid options are:

  - 0 - No
  - 1 - Yes
*/

func ClimatizationValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("climatization", "vehicles.txt", "climatization_validation", row, services.AppMessageService)
	if rules != nil && rules.Climatization.Severity != "" {
		ctx.WithSeverity(rules.Climatization.Severity)
	}

	if vehicle.Climatization == nil {
		ctx.AddError(ctx.GetTranslatedMessage("climatization_validation.required"))
		return
	}

	validOptions := []int{0, 1}
	if !slices.Contains(validOptions, *vehicle.Climatization) {
		ctx.AddError(ctx.GetTranslatedMessage("climatization_validation.invalid", strconv.Itoa(*vehicle.Climatization)))
		return
	}
}

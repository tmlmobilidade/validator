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
  - Field: ramp
  - Presence: Required
  - Type: Enum


# Description

Ramp for people with reduced mobility.

Valid options are:

  - 0 - No
  - 1 - Manual ramp
  - 2 - Eletric folding system
  - 3 - Not Applicable
*/

func RampValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("ramp", "vehicles.txt", "ramp_validation", row, services.AppMessageService)
	if rules != nil && rules.Ramp.Severity != "" {
		ctx.WithSeverity(rules.Ramp.Severity)
	}

	if vehicle.Ramp == nil {
		ctx.AddError(ctx.GetTranslatedMessage("ramp_validation.required"))
		return
	}

	validOptions := []int{0, 1, 2, 3}
	if !slices.Contains(validOptions, *vehicle.Ramp) {
		ctx.AddError(ctx.GetTranslatedMessage("ramp_validation.invalid", strconv.Itoa(*vehicle.Ramp)))
		return
	}
}

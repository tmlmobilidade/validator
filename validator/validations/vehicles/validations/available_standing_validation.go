package vehicles

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes
  - File: [vehicles.txt]
  - Field: available_standing
  - Presence: Required
  - Type: Number

# Description

The number of standing available on the vehicle.
*/

func AvailableStandingValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("available_standing", "vehicles.txt", "available_standing_validation", row, services.AppMessageService)
	if rules != nil && rules.AvailableStanding.Severity != "" {
		ctx.WithSeverity(rules.AvailableStanding.Severity)
	}

	if vehicle.AvailableStanding == nil {
		ctx.AddError(ctx.GetTranslatedMessage("available_standing_validation.required"))
		return
	}

	if *vehicle.AvailableStanding <= 0 {
		ctx.AddError(ctx.GetTranslatedMessage("available_standing_validation.invalid", *vehicle.AvailableStanding))
		return
	}
}

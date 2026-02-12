package vehicles

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [vehicles.txt]
- Field: owner
- Presence: Required
- Type: String

# Description

The owner of the vehicle.
*/
func OwnerValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("owner", "vehicles.txt", "owner_validation", row, services.AppMessageService)
	if rules != nil && rules.Owner.Severity != "" {
		ctx.WithSeverity(rules.Owner.Severity)
	}

	if vehicle.Owner == nil {
		ctx.AddError(ctx.GetTranslatedMessage("owner_validation.required"))
		return
	}
}

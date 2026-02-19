package vehicles

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes
  - File: [vehicles.txt]
  - Field: registration_date
  - Presence: Required
  - Type: Date

# Description

Date of the first registration.
*/

func RegistrationDateValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("registration_date", "vehicles.txt", "registration_date_validation", row, services.AppMessageService)
	if rules != nil && rules.RegistrationDate.Severity != "" {
		ctx.WithSeverity(rules.RegistrationDate.Severity)
	}

	if vehicle.RegistrationDate == nil {
		ctx.AddError(ctx.GetTranslatedMessage("registration_date_validation.required"))
		return
	}

	if !lib.IsValidServiceDate(*vehicle.RegistrationDate) {
		ctx.AddError(ctx.GetTranslatedMessage("registration_date_validation.invalid", &vehicle.RegistrationDate))
		return
	}
}

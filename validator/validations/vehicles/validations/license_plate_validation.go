package vehicles

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [vehicles.txt]
- Field: license_plate
- Presence: Required
- Type: String

# Description

The license plate of the vehicle.

The license plate must be in the format XX-XX-XX.
*/
func LicensePlateValidation(vehicle *types.Vehicle, row int, gtfs *types.Gtfs, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("license_plate", "vehicles.txt", "license_plate_validation", row, services.AppMessageService)
	if rules != nil && rules.LicensePlate.Severity != "" {
		ctx.WithSeverity(rules.LicensePlate.Severity)
	}

	if vehicle.LicensePlate == nil {
		ctx.AddError(ctx.GetTranslatedMessage("license_plate_validation.required"))
		return
	}

	if !lib.ValidateLicensePlate(*vehicle.LicensePlate) {
		ctx.AddError(ctx.GetTranslatedMessage("license_plate_validation.invalid", *vehicle.LicensePlate))
		return
	}

	if gtfs != nil {
		rows, err := gtfs.GetRowsByField("vehicles", "license_plate", *vehicle.LicensePlate)
		if err == nil && len(rows) > 1 {
			ctx.AddError(ctx.GetTranslatedMessage("license_plate_validation.duplicate", *vehicle.LicensePlate))
			return
		}
	}
}

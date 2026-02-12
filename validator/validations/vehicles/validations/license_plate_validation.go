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
		lib.AppLogger.Accent("Invalid license plate format", "license_plate", *vehicle.LicensePlate)
		ctx.AddError(ctx.GetTranslatedMessage("license_plate_validation.invalid", map[string]interface{}{"license_plate": *vehicle.LicensePlate}))
		return
	}

	if gtfs != nil {
		rows, err := gtfs.GetRowsByField("vehicles", "license_plate", *vehicle.LicensePlate)
		if err == nil && len(rows) > 1 {
			ctx.AddError(ctx.GetTranslatedMessage("license_plate_validation.duplicate", map[string]interface{}{"license_plate": *vehicle.LicensePlate}))
			return
		}
	}
}

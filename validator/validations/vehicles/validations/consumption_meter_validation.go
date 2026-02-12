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
  - Field: consumption_meter
  - Presence: Required
  - Type: Enum

# Description

Consumption measurement.

Valid options are:

  - 0 - No
  - 1 - Yes
*/
func ConsumptionMeterValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("consumption_meter", "vehicles.txt", "consumption_meter_validation", row, services.AppMessageService)
	if rules != nil && rules.ConsumptionMeter.Severity != "" {
		ctx.WithSeverity(rules.ConsumptionMeter.Severity)
	}

	if vehicle.ConsumptionMeter == nil {
		ctx.AddError(ctx.GetTranslatedMessage("consumption_meter_validation.required"))
		return
	}

	validOptions := []int{0, 1}
	if !slices.Contains(validOptions, *vehicle.ConsumptionMeter) {
		ctx.AddError(ctx.GetTranslatedMessage("consumption_meter_validation.invalid", strconv.Itoa(*vehicle.ConsumptionMeter)))
		return
	}
}

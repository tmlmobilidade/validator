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
  - Field: emission
  - Presence: Required
  - Type: Enum

# Description

The emission of the vehicle.

Valid options are:

  - 1 - Euro I
  - 2 - Euro II
  - 3 - Euro III
  - 4 - Euro IV
  - 5 - Euro V
  - 6 - Euro VI
*/

func EmissionValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("emission", "vehicles.txt", "emission_validation", row, services.AppMessageService)
	ctx.Severity = types.SEVERITY_ERROR
	if rules != nil && rules.Emission.Severity != "" {
		ctx.WithSeverity(rules.Emission.Severity)
	}

	if vehicle.Emission == nil {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("emission_validation.required"))
		return
	}

	validOptions := []int{1, 2, 3, 4, 5, 6}
	if !slices.Contains(validOptions, *vehicle.Emission) {
		ctx.AddError(ctx.GetTranslatedMessage("emission_validation.invalid", *vehicle.Emission))
		return
	}

	// Validate rules
	if rules != nil && rules.Emission.Options != nil {
		if slices.Contains(*rules.Emission.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.Emission.Options, strconv.Itoa(*vehicle.Emission)) {
			ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("emission_validation.not_allowed", *vehicle.Emission))
			return
		}
	}
}

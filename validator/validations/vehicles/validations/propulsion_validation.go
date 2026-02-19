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
  - Field: propulsion
  - Presence: Required
  - Type: Enum

# Description

The propulsion of the vehicle.

Valid options are:

  - 1 - Gasoline
  - 2 - Diesel
  - 3 - LPG auto
  - 4 - Mixture
  - 5 - Biodiesel
  - 6 - Electricity
  - 7 - Hybrid
  - 8 - Natural Gas
*/

func PropulsionValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("propulsion", "vehicles.txt", "propulsion_validation", row, services.AppMessageService)
	ctx.Severity = types.SEVERITY_ERROR
	if rules != nil && rules.Propulsion.Severity != "" {
		ctx.WithSeverity(rules.Propulsion.Severity)
	}

	if vehicle.Propulsion == nil {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("propulsion_validation.required"))
		return
	}

	validOptions := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !slices.Contains(validOptions, *vehicle.Propulsion) {
		ctx.AddError(ctx.GetTranslatedMessage("propulsion_validation.invalid", strconv.Itoa(*vehicle.Propulsion)))
		return
	}

	// Validate rules
	if rules != nil && rules.Propulsion.Options != nil {
		if slices.Contains(*rules.Propulsion.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.Propulsion.Options, strconv.Itoa(*vehicle.Propulsion)) {
			ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("propulsion_validation.not_allowed", *vehicle.Propulsion))
			return
		}
	}
}

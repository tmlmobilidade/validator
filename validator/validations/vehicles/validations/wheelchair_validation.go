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
  - Field: wheelchair
  - Presence: Required
  - Type: Enum

# Description

Wheelchair accessibility.

Valid options are:

  - 0 - No
  - 1 - Yes
*/

func WheelchairValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("wheelchair", "vehicles.txt", "wheelchair_validation", row, services.AppMessageService)
	ctx.Severity = types.SEVERITY_ERROR
	ctx.Severity = types.SEVERITY_ERROR
	if rules != nil && rules.Wheelchair.Severity != "" {
		ctx.WithSeverity(rules.Wheelchair.Severity)
	}

	if vehicle.Wheelchair == nil {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("wheelchair_validation.required"))
		return
	}

	validOptions := []int{0, 1}
	if !slices.Contains(validOptions, *vehicle.Wheelchair) {
		ctx.AddError(ctx.GetTranslatedMessage("wheelchair_validation.invalid", *vehicle.Wheelchair))
		return
	}

	// Validate rules
	ctx.Severity = types.SEVERITY_ERROR
	if rules != nil && rules.Wheelchair.Options != nil {
		if slices.Contains(*rules.Wheelchair.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.Wheelchair.Options, strconv.Itoa(*vehicle.Wheelchair)) {
			ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("wheelchair_validation.not_allowed", *vehicle.Wheelchair))
			return
		}
	}
}

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
  - Field: external_sound
  - Presence: Required
  - Type: Enum

# Description

External sound information (external Public Address).

Valid options are:

  - 0 - No
  - 1 - Yes
*/

func ExternalSoundValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("external_sound", "vehicles.txt", "external_sound_validation", row, services.AppMessageService)
	if rules != nil && rules.ExternalSound.Severity != "" {
		ctx.WithSeverity(rules.ExternalSound.Severity)
	}

	if vehicle.ExternalSound == nil {
		ctx.AddError(ctx.GetTranslatedMessage("external_sound_validation.required"))
		return
	}

	validOptions := []int{0, 1}
	if !slices.Contains(validOptions, *vehicle.ExternalSound) {
		ctx.AddError(ctx.GetTranslatedMessage("external_sound_validation.invalid", strconv.Itoa(*vehicle.ExternalSound)))
		return
	}
}

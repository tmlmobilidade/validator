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
  - Field: onboard_monitor
  - Presence: Required
  - Type: Enum


# Description

On-board monitor.

Valid options are:

  - 0 - No
  - 1 - Yes
*/

func OnboardMonitorValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("onboard_monitor", "vehicles.txt", "onboard_monitor_validation", row, services.AppMessageService)
	if rules != nil && rules.OnboardMonitor.Severity != "" {
		ctx.WithSeverity(rules.OnboardMonitor.Severity)
	}

	if vehicle.OnboardMonitor == nil {
		ctx.AddError(ctx.GetTranslatedMessage("onboard_monitor_validation.required"))
		return
	}

	validOptions := []int{0, 1}
	if !slices.Contains(validOptions, *vehicle.OnboardMonitor) {
		ctx.AddError(ctx.GetTranslatedMessage("onboard_monitor_validation.invalid", strconv.Itoa(*vehicle.OnboardMonitor)))
		return
	}
}

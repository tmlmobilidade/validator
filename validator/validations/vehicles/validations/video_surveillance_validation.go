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
  - Field: video_surveillance
  - Presence: Required
  - Type: Enum

# Description

Vehicle has a video surveillance system.

Valid options are:

  - 0 - No
  - 1 - Yes
*/
func VideoSurveillanceValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("video_surveillance", "vehicles.txt", "video_surveillance_validation", row, services.AppMessageService)
	if rules != nil && rules.VideoSurveillance.Severity != "" {
		ctx.WithSeverity(rules.VideoSurveillance.Severity)
	}

	if vehicle.VideoSurveillance == nil {
		ctx.AddError(ctx.GetTranslatedMessage("video_surveillance_validation.required"))
		return
	}

	validOptions := []int{0, 1}
	if !slices.Contains(validOptions, *vehicle.VideoSurveillance) {
		ctx.AddError(ctx.GetTranslatedMessage("video_surveillance_validation.invalid", strconv.Itoa(*vehicle.VideoSurveillance)))
		return
	}
}

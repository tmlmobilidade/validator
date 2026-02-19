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
	ctx.Severity = types.SEVERITY_ERROR
	if rules != nil && rules.VideoSurveillance.Severity != "" {
		ctx.WithSeverity(rules.VideoSurveillance.Severity)
	}

	if vehicle.VideoSurveillance == nil {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("video_surveillance_validation.required"))
		return
	}

	validOptions := []int{0, 1}
	if !slices.Contains(validOptions, *vehicle.VideoSurveillance) {
		ctx.AddError(ctx.GetTranslatedMessage("video_surveillance_validation.invalid", *vehicle.VideoSurveillance))
		return
	}

	// Validate rules
	if rules != nil && rules.VideoSurveillance.Options != nil {
		if slices.Contains(*rules.VideoSurveillance.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.VideoSurveillance.Options, strconv.Itoa(*vehicle.VideoSurveillance)) {
			ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("video_surveillance_validation.not_allowed", *vehicle.VideoSurveillance))
			return
		}
	}
}

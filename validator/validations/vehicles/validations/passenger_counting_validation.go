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
  - Field: passenger_counting
  - Presence: Required
  - Type: Enum

# Description

Vehicle has a passenger counting system.

Valid options are:

  - 0 - No
  - 1 - Yes
*/

func PassengerCountingValidation(vehicle *types.Vehicle, row int, rules *types.VehiclesRules) {
	ctx := lib.NewValidationContext("passenger_counting", "vehicles.txt", "passenger_counting_validation", row, services.AppMessageService)
	ctx.Severity = types.SEVERITY_ERROR
	if rules != nil && rules.PassengerCounting.Severity != "" {
		ctx.WithSeverity(rules.PassengerCounting.Severity)
	}

	if vehicle.PassengerCounting == nil {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("passenger_counting_validation.required"))
		return
	}

	validOptions := []int{0, 1}
	if !slices.Contains(validOptions, *vehicle.PassengerCounting) {
		ctx.AddError(ctx.GetTranslatedMessage("passenger_counting_validation.invalid", strconv.Itoa(*vehicle.PassengerCounting)))
		return
	}

	// Validate rules
	ctx.Severity = types.SEVERITY_ERROR
	if rules != nil && rules.PassengerCounting.Options != nil {
		if slices.Contains(*rules.PassengerCounting.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.PassengerCounting.Options, strconv.Itoa(*vehicle.PassengerCounting)) {
			ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("passenger_counting_validation.not_allowed", *vehicle.PassengerCounting))
			return
		}
	}
}

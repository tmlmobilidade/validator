package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	"slices"
	"strconv"
)

/*
# Attributes

  - File: [stops.txt]
  - Field: has_shelter
  - Presence: Optional
  - Type: Boolean

# Description

Describes if the stop has a shelter.

[stops.txt]: https://gtfs.org/schedule/reference/#stopstxt
*/
func HasShelterValidation(stop *types.Stop, row int, rules *types.StopsRules) {
	ctx := lib.NewValidationContext("has_shelter", "stops.txt", "has_shelter_validation", "has_shelter_valid", row, services.AppMessageService)
	if rules != nil && rules.HasShelter.Severity != "" {
		ctx.WithSeverity(rules.HasShelter.Severity)
	}

	if stop.HasShelter == nil {
		if ctx.ShouldSkip() {
			return
		}

		message := ctx.GetRequiredMessage("has_shelter_validation.required", "has_shelter_validation.recommended")
		ctx.AddMessageWithSeverity(message)
		return
	}

	if ctx.IsForbidden() {
		ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("has_shelter_validation.forbidden"))
		return
	}

	// Validate value
	validValues := []int{0, 1}
	if !slices.Contains(validValues, *stop.HasShelter) {
		ctx.AddError(ctx.GetTranslatedMessage("has_shelter_validation.invalid", *stop.HasShelter))
		return
	}

	// Validate value based on rules
	if rules != nil && rules.HasShelter.Options != nil {
		if slices.Contains(*rules.HasShelter.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.HasShelter.Options, strconv.Itoa(*stop.HasShelter)) {
			ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("has_shelter_validation.not_allowed", *stop.HasShelter))
			return
		}
	}
}

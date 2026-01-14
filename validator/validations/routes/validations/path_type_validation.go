package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

- File: [routes.txt]
- Field: path_type
- Presence: Optional (TML-specific)
- Type: Enum

# Description

TML-specific field indicating the path type of a route.

Valid options are:

  - 1: First path type
  - 2: Second path type
  - 3: Third path type

This is a TML-specific extension to the GTFS standard.
*/
func PathTypeValidation(route *types.Route, row int, rules *types.RoutesRules) {
	ctx := lib.NewValidationContext("path_type", "routes.txt", "path_type_validation", row, services.AppMessageService)
	if rules != nil && rules.PathType.Severity != "" {
		ctx.WithSeverity(rules.PathType.Severity)
	}

	validTypes := []string{"1", "2", "3"}

	// path_type is optional, so if it's nil or empty, we don't validate
	if route.PathType == nil || *route.PathType == "" {
		if ctx.IsForbidden() {
			ctx.AddError(ctx.GetTranslatedMessage("path_type_validation.forbidden"))
		}
		return
	}

	// Check if field is forbidden - if present, it's an error
	if ctx.IsForbidden() {
		ctx.AddError(ctx.GetTranslatedMessage("path_type_validation.forbidden"))
		return
	}

	// Validate against valid options
	if !slices.Contains(validTypes, *route.PathType) {
		ctx.AddError(ctx.GetTranslatedMessage("path_type_validation.invalid", *route.PathType))
		return
	}

	// Validate rules - check if the value is in the allowed options
	if rules != nil && rules.PathType.Options != nil {
		if slices.Contains(*rules.PathType.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.PathType.Options, *route.PathType) {
			ctx.AddError(ctx.GetTranslatedMessage("path_type_validation.not_allowed", *route.PathType))
			return
		}
	}
}

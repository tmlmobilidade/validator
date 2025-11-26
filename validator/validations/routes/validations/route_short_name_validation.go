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
- Field: route_short_name
- Presence: Conditionally Required
- Type: String

# Description

Short name of a route. Often a short, abstract identifier (e.g., "32", "100X", "Green") that riders use to identify a route.
Both route_short_name and route_long_name may be defined.

Conditionally Required:
  - Required if routes.route_long_name is empty.
  - Recommended if there is a brief service designation. This should be the commonly-known passenger name of the service, and should be no longer than 12 characters.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
*/
func RouteShortNameValidation(route *types.Route, row int, rules *types.RoutesRules) {
	ctx := lib.NewValidationContext("route_short_name", "routes.txt", "route_short_name_validation", row, services.AppMessageService)
	if rules != nil && rules.RouteShortName.Severity != "" {
		ctx.WithSeverity(rules.RouteShortName.Severity)
	} else {
		ctx.WithSeverity(types.SEVERITY_WARNING)
	}

	// Extract values with nil checks
	if route.RouteShortName == nil && route.RouteLongName == nil {
		ctx.AddError(ctx.GetTranslatedMessage("route_short_name_validation.required_if_long_name_empty"))
		return
	}

	if route.RouteShortName == nil {
		if !ctx.ShouldIgnore() {
			message := ctx.GetRequiredMessage("route_short_name_validation.required", "route_short_name_validation.recommended")
			ctx.AddMessageWithSeverity(message)
		}
		return
	}

	// Validate length
	if len(*route.RouteShortName) > 12 {
		ctx.AddWarning(ctx.GetTranslatedMessage("route_short_name_validation.too_long"))
	}

	// Validate rules
	if rules != nil && rules.RouteShortName.Options != nil {
		if slices.Contains(*rules.RouteShortName.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.RouteShortName.Options, *route.RouteShortName) {
			ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("route_short_name_validation.not_allowed", map[string]interface{}{"value": *route.RouteShortName}))
			return
		}
	}
}

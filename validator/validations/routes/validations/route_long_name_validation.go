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
- Field: route_long_name
- Presence: Conditionally Required
- Type: String

# Description

Full name of a route. This name is generally more descriptive than the route_short_name and often includes the route's destination or stop.

Both route_short_name and route_long_name may be defined.

Conditionally Required:
  - Required if routes.route_short_name is empty.
  - Optional otherwise.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
*/
func RouteLongNameValidation(route *types.Route, row int, rules *types.RoutesRules) {
	ctx := lib.NewValidationContext("route_long_name", "routes.txt", "route_long_name_validation", row, services.AppMessageService)
	if rules != nil && rules.RouteLongName.Severity != "" {
		ctx.WithSeverity(rules.RouteLongName.Severity)
	}

	if (route.RouteLongName == nil || *route.RouteLongName == "") && (route.RouteShortName == nil || *route.RouteShortName == "") {
		ctx.AddError(ctx.GetTranslatedMessage("route_long_name_validation.required_if_short_name_empty"))
		return
	}

	if route.RouteLongName == nil || *route.RouteLongName == "" {
		if ctx.ShouldSkip() {
			return
		}

		message := ctx.GetRequiredMessage("route_long_name_validation.required", "route_long_name_validation.recommended")
		ctx.AddMessageWithSeverity(message)
		return
	}

	// Validate rules
	if rules != nil && rules.RouteLongName.Options != nil {
		if slices.Contains(*rules.RouteLongName.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.RouteLongName.Options, *route.RouteLongName) {
			ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("route_long_name_validation.not_allowed", map[string]interface{}{"value": *route.RouteLongName}))
			return
		}
	}
}

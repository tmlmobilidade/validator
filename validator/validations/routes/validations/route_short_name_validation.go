package routes

import (
	"main/i18n"
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
	s := types.SEVERITY_WARNING
	if rules != nil && rules.RouteShortName.Severity != "" {
		s = rules.RouteShortName.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "route_short_name",
			FileName:     "routes.txt",
			ValidationID: "route_short_name_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	// Extract values with nil checks
	if route.RouteShortName == nil && route.RouteLongName == nil {
		addMessage(i18n.AppTranslator.Get("route_short_name_validation.required_if_long_name_empty"), types.SEVERITY_ERROR)
		return
	}

	if route.RouteShortName == nil {
		if s != types.SEVERITY_IGNORE {
			warn := lib.IfThenElse(s == types.SEVERITY_WARNING, i18n.AppTranslator.Get("route_short_name_validation.recommended"), i18n.AppTranslator.Get("route_short_name_validation.required"))
			addMessage(warn, s)
		}
		return
	}

	// Validate length
	if len(*route.RouteShortName) > 12 {
		addMessage(i18n.AppTranslator.Get("route_short_name_validation.too_long"), types.SEVERITY_WARNING)
	}

	// Validate rules
	if rules != nil && rules.RouteShortName.Options != nil {
		if slices.Contains(*rules.RouteShortName.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.RouteShortName.Options, *route.RouteShortName) {
			addMessage(i18n.AppTranslator.Get("route_short_name_validation.not_allowed", map[string]interface{}{"value": *route.RouteShortName}), s)
			return
		}
	}
}

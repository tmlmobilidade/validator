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
- Field: route_desc
- Presence: Optional
- Type: String

# Description

Description of a route that provides useful, quality information. Should not be a duplicate of route_short_name or route_long_name.

# Example

"A" trains operate between Inwood-207 St, Manhattan and Far Rockaway-Mott Avenue, Queens at all times. Also from about 6AM until about midnight, additional "A" trains operate between Inwood-207 St and Lefferts Boulevard (trains typically alternate between Lefferts Blvd and Far Rockaway).

Conditionally Required:
  - Required if routes.route_short_name is empty.
  - Optional otherwise.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
*/
func RouteDescValidation(route *types.Route, row int, rules *types.RoutesRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.RouteDesc.Severity != "" {
		s = rules.RouteDesc.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "route_desc",
			FileName:     "routes.txt",
			ValidationID: "route_desc_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	if route.RouteDesc == nil || *route.RouteDesc == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, i18n.AppTranslator.Get("route_desc_validation.recommended"), i18n.AppTranslator.Get("route_desc_validation.required"))
		addMessage(warn, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("route_desc_validation.forbidden"), s)
		return
	}

	if route.RouteShortName != nil && *route.RouteDesc == *route.RouteShortName {
		addMessage(i18n.AppTranslator.Get("route_desc_validation.duplicate_short_name"), types.SEVERITY_WARNING)
	}
	if route.RouteLongName != nil && *route.RouteDesc == *route.RouteLongName {
		addMessage(i18n.AppTranslator.Get("route_desc_validation.duplicate_long_name"), types.SEVERITY_WARNING)
	}

	// Validate rules
	if rules != nil && rules.RouteDesc.Options != nil {
		if slices.Contains(*rules.RouteDesc.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.RouteDesc.Options, *route.RouteDesc) {
			addMessage(i18n.AppTranslator.Get("route_desc_validation.not_allowed", map[string]interface{}{"value": *route.RouteDesc}), s)
			return
		}
	}
}

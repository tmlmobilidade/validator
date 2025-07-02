package routes

import (
	"fmt"
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

	s := types.SEVERITY_IGNORE
	if rules != nil && rules.RouteLongName.Severity != "" {
		s = rules.RouteLongName.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "route_long_name",
			FileName:     "routes.txt",
			ValidationID: "route_long_name_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	if (route.RouteLongName == nil || *route.RouteLongName == "") && (route.RouteShortName == nil || *route.RouteShortName == "") {
		addMessage("route_long_name is required if route_short_name is empty.", types.SEVERITY_ERROR)
		return
	}

	if route.RouteLongName == nil || *route.RouteLongName == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "route_long_name is recommended.", "route_long_name is required.")
		addMessage(warn, s)
	}

	// Validate rules
	if rules != nil && rules.RouteLongName.Options != nil {
		if slices.Contains(*rules.RouteLongName.Options, types.ALL_OPTIONS) {
			return
		}

		if slices.Contains(*rules.RouteLongName.Options, *route.RouteLongName) {
			return
		}

		addMessage(fmt.Sprintf("route_long_name is not allowed: %s", *route.RouteLongName), s)
		return
	}
}

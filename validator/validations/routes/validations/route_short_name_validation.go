package routes

import (
	"main/lib"
	"main/services"
	"main/types"
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
func RouteShortNameValidation(severity *types.Severity, route *types.Route, row int) {
	s := types.SEVERITY_WARNING
	if severity != nil {
		s = *severity
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
		addMessage("route_short_name is required if route_long_name is empty.", types.SEVERITY_ERROR)
		return
	}

	if route.RouteShortName == nil {
		if s != types.SEVERITY_IGNORE {
			warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "route_short_name is recommended.", "route_short_name is required.")
			addMessage(warn, s)
		}
		return
	}

	// Validate length
	if len(*route.RouteShortName) > 12 {
		addMessage("route_short_name should be no longer than 12 characters.", types.SEVERITY_WARNING)
	}
} 
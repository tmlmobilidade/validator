package routes

import (
	"main/lib"
	"main/services"
	"main/types"
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
func RouteLongNameValidation(severity *types.Severity, route *types.Route, row int) {

	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
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
} 
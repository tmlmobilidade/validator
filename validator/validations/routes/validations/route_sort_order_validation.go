package routes

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [routes.txt]
- Field: route_sort_order
- Presence: Optional
- Type: Non-negative integer

# Description

Orders the routes in a way which is ideal for presentation to customers. Routes with smaller route_sort_order values should be displayed first.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
*/
func RouteSortOrderValidation(route *types.Route, row int, rules *types.RoutesRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.RouteSortOrder.Severity != "" {
		s = rules.RouteSortOrder.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "route_sort_order",
			FileName:     "routes.txt",
			ValidationID: "route_sort_order_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	if route.RouteSortOrder == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "route_sort_order is recommended", "route_sort_order is required")
		addMessage(warn, s)
		return
	}

	if *route.RouteSortOrder < 0 {
		addMessage("route_sort_order must be a non-negative integer.", types.SEVERITY_ERROR)
		return
	}
}

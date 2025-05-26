package routes

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [routes.txt]
- Field: network_id
- Presence: Conditionally Forbidden
- Type: ID

# Description

Identifies a group of routes. Multiple rows in [routes.txt] may have the same network_id.

Conditionally Forbidden:
- Forbidden if the [route_networks.txt] file exists.
- Optional otherwise.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
[route_networks.txt]: https://gtfs.org/schedule/reference/#routenetworkstxt
*/
func NetworkIdValidation(severity *types.Severity, route *types.Route, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "network_id",
			FileName:     "routes.txt",
			ValidationID: "network_id_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	if route.NetworkId == nil && s != types.SEVERITY_IGNORE {
		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "network_id is recommended", "network_id is required")
		addMessage(warn, s)
		return
	}

	if _, exists := gtfs.Files["route_networks"]; exists && route.NetworkId != nil {
		addMessage("network_id is forbidden if route_networks.txt exists", types.SEVERITY_ERROR)
		return
	}
}
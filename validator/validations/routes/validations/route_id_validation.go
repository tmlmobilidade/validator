package routes

import (
	"main/i18n"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [routes.txt]
- Field: route_id
- Presence: Required
- Type: Unique ID

# Description

Identifies a route.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
*/
func RouteIdValidation(route *types.Route, row int, gtfs *types.Gtfs) {

	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "route_id",
			FileName:     "routes.txt",
			ValidationID: "route_id_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     types.SEVERITY_ERROR,
		})
	}

	if route.RouteId == nil || *route.RouteId == "" {
		addMessage(i18n.AppTranslator.Get("route_id_validation.required"))
		return
	}

	// Check if route_id is Unique ID
	rows, err := gtfs.GetRowsById("routes", *route.RouteId)
	if err == nil && len(rows) > 1 {
		addMessage(i18n.AppTranslator.Get("route_id_validation.duplicate", *route.RouteId))
		return
	}
}

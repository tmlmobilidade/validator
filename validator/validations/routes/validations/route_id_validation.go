package routes

import (
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
			Field: "route_id",
			FileName: "routes.txt",
			ValidationID: "route_id_validation",
			Message: msg,
			Rows: []int{row},
			Severity: types.SEVERITY_ERROR,
		})
	}

	if route.RouteId == nil || *route.RouteId == "" {
		addMessage("route_id is required.")
		return
	}

	
	// Check if route_id is Unique ID
	if _, ok := gtfs.IdMap["routes"][*route.RouteId]; ok && len(gtfs.IdMap["routes"][*route.RouteId]) > 1 {
		addMessage("Duplicate route_id found. Route IDs must be unique.")
		return
	}
} 
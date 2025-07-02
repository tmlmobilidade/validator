package routes

import (
	"fmt"
	"main/services"
	"main/types"
	"slices"
	"strconv"
)

/*
# Attributes

- File: [routes.txt]
- Field: route_type
- Presence: Required
- Type: Enum

# Description

Indicates the type of transportation used on a route.

Valid options are:

  - 0: Tram, Streetcar, Light rail. Any light rail or street level system within a metropolitan area.
  - 1: Subway, Metro. Any underground rail system within a metropolitan area.
  - 2: Rail. Used for intercity or long-distance travel.
  - 3: Bus. Used for short- and long-distance bus routes.
  - 4: Ferry. Used for short- and long-distance boat service.
  - 5: Cable tram. Used for street-level rail cars where the cable runs beneath the vehicle (e.g., cable car in San Francisco).
  - 6: Aerial lift, suspended cable car (e.g., gondola lift, aerial tramway). Cable transport where cabins, cars, gondolas or open chairs are suspended by means of one or more cables.
  - 7: Funicular. Any rail system designed for steep inclines.
  - 11: Trolleybus. Electric buses that draw power from overhead wires using poles.
  - 12: Monorail. Railway in which the track consists of a single rail or a beam.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
*/
func RouteTypeValidation(route *types.Route, row int, rules *types.RoutesRules) {
	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "route_type",
			FileName:     "routes.txt",
			ValidationID: "route_type_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     types.SEVERITY_ERROR,
		})
	}

	validTypes := map[int]struct{}{
		0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 11: {}, 12: {},
	}

	if route.RouteType == nil {
		addMessage("route_type is required.")
		return
	}

	if _, ok := validTypes[*route.RouteType]; !ok {
		addMessage("route_type must be one of the valid GTFS enum values (0,1,2,3,4,5,6,7,11,12).")
	}

	// Validate rules
	if rules != nil && rules.RouteType.Options != nil {
		if slices.Contains(*rules.RouteType.Options, types.ALL_OPTIONS) {
			return
		}

		if slices.Contains(*rules.RouteType.Options, strconv.Itoa(*route.RouteType)) {
			return
		}

		addMessage(fmt.Sprintf("route_type is not allowed: %d", *route.RouteType))
		return
	}
}

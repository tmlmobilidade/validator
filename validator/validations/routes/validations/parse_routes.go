package routes

import (
	"main/lib"
	"main/services"
	"main/types"
)

func ParseRoutes(rawRoute map[string]string, row int, gtfs *types.Gtfs) types.Route {
	var (
		route types.Route = types.Route{}
		routeId string
		routeType int
		agencyId, continuousDropOff, continuousPickup, routeColor, routeDesc, routeLongName, routeShortName, routeTextColor, routeUrl string
		routeSortOrder int
		messages []types.Message
	)

	stringFields := map[string]*string{
		"route_id": &routeId,
		"agency_id": &agencyId,
		"continuous_drop_off": &continuousDropOff,
		"continuous_pickup": &continuousPickup,
		"route_color": &routeColor,
		"route_desc": &routeDesc,
		"route_long_name": &routeLongName,
		"route_short_name": &routeShortName,
		"route_text_color": &routeTextColor,
		"route_url": &routeUrl,
	}

	intFields := map[string]*int{
		"route_type": &routeType,
		"route_sort_order": &routeSortOrder,
	}

	addMessage := func(field, msg string) {
		messages = append(messages, types.Message{
			Field:        field,
			FileName:     "routes.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "routes_parse",
		})
	}

	for field, target := range stringFields {
		if errMsg := lib.ParseStringToPrimitive(rawRoute[field], target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}

	for field, target := range intFields {
		if errMsg := lib.ParseStringToPrimitive(rawRoute[field], target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}

	if len(messages) > 0 {
		services.AppMessageService.AddMessages(messages)
		return route
	}

	route.RouteId = lib.IfThenElse(rawRoute["route_id"] != "", &routeId, nil)
	route.RouteType = lib.IfThenElse(rawRoute["route_type"] != "", &routeType, nil)
	route.AgencyId = lib.IfThenElse(rawRoute["agency_id"] != "", &agencyId, nil)
	route.ContinuousDropOff = lib.IfThenElse(rawRoute["continuous_drop_off"] != "", &continuousDropOff, nil)
	route.ContinuousPickup = lib.IfThenElse(rawRoute["continuous_pickup"] != "", &continuousPickup, nil)
	route.RouteColor = lib.IfThenElse(rawRoute["route_color"] != "", &routeColor, nil)
	route.RouteDesc = lib.IfThenElse(rawRoute["route_desc"] != "", &routeDesc, nil)
	route.RouteLongName = lib.IfThenElse(rawRoute["route_long_name"] != "", &routeLongName, nil)
	route.RouteShortName = lib.IfThenElse(rawRoute["route_short_name"] != "", &routeShortName, nil)
	route.RouteSortOrder = lib.IfThenElse(rawRoute["route_sort_order"] != "", &routeSortOrder, nil)
	route.RouteTextColor = lib.IfThenElse(rawRoute["route_text_color"] != "", &routeTextColor, nil)
	route.RouteUrl = lib.IfThenElse(rawRoute["route_url"] != "", &routeUrl, nil)

	return route
}

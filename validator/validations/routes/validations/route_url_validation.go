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
- Field: route_url
- Presence: Optional
- Type: URL

# Description

URL of a web page about the particular route. Should be different from the agency.agency_url value.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
*/
func RouteUrlValidation(route *types.Route, row int, gtfs *types.Gtfs, rules *types.RoutesRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.RouteUrl.Severity != "" {
		s = rules.RouteUrl.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "route_url",
			FileName:     "routes.txt",
			ValidationID: "route_url_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	if route.RouteUrl == nil || *route.RouteUrl == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := lib.IfThenElse(s == types.SEVERITY_WARNING, "route_url is recommended.", "route_url is required.")
		addMessage(message, s)
		return
	}

	if err := lib.ValidateUrl(*route.RouteUrl); err != "" {
		addMessage(err, types.SEVERITY_ERROR)
		return
	}

	// Check if route_url is the same as agency_url
	if route.AgencyId != nil {
		agencyId := *route.AgencyId
		agencyRow := gtfs.IdMap["agency"][agencyId]
		agencyUrl := gtfs.Agency[agencyRow[0]].AgencyUrl

		if agencyUrl != "" && *route.RouteUrl == agencyUrl {
			addMessage("route_url should be different from agency_url", types.SEVERITY_WARNING)
		}
	}

	// Validate rules
	if rules != nil && rules.RouteUrl.Options != nil {
		if slices.Contains(*rules.RouteUrl.Options, types.ALL_OPTIONS) {
			return
		}

		if slices.Contains(*rules.RouteUrl.Options, *route.RouteUrl) {
			return
		}

		addMessage(fmt.Sprintf("route_url is not allowed: %s", *route.RouteUrl), s)
		return
	}
}

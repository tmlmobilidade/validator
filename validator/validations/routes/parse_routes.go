package routes

import (
	"main/lib"
	"main/types"
)

type parseRouteValidation struct {
	*types.Validation
}

func NewParseRouteValidation(severity *types.Severity) *parseRouteValidation {

	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	return &parseRouteValidation{
		Validation: &types.Validation{
			ID:          "parse_route",
			Description: "Validate route data",
			Severity:    s,
		},
	}
}

func (v *parseRouteValidation) Validate(gtfs types.Gtfs) (routes []types.Route, messages []types.Message) {
	routeIds := make(map[string]bool)

	// Check if multiple agencies exist by checking length of agencies file
	multipleAgencies := len(gtfs.Files["agency"]) > 1

	// Check if any stop_times have pickup/dropoff window fields with non-zero counts
	hasPickupDropoffWindows := false
	stopTimesFields := gtfs.FieldCounter["stop_times"]
	if stopTimesFields != nil {
		startWindowCount := stopTimesFields["start_pickup_drop_off_window"]
		endWindowCount := stopTimesFields["end_pickup_drop_off_window"]
		hasPickupDropoffWindows = startWindowCount > 0 || endWindowCount > 0
	}

	for i, route := range gtfs.Files["routes"] {
		route, routeMessages := parseRoute(route, multipleAgencies, hasPickupDropoffWindows)
		routes = append(routes, route)

		// Check for duplicate route IDs
		if route.RouteId != "" {
			if routeIds[route.RouteId] {
				messages = append(messages, types.Message{
					Field:        "route_id",
					FileName:     "routes.txt",
					Message:      "Duplicate route_id found. Route IDs must be unique.",
					Rows:         []int{i + 1},
					Severity:     v.Severity,
					ValidationID: v.ID,
				})
			}
			routeIds[route.RouteId] = true
		}

		// Update row number and other fields for each message
		for _, msg := range routeMessages {
			msg.Rows = []int{i + 1}
			msg.FileName = "routes.txt"
			msg.Severity = v.Severity
			msg.ValidationID = v.ID
			messages = append(messages, msg)
		}
	}
	return routes, messages
}

func parseRoute(m map[string]string, multipleAgencies bool, hasPickupDropoffWindows bool) (route types.Route, messages []types.Message) {
	var parsingErrors []string

	// Convert Optional Primitive Values
	var agencyId, routeColor, routeDesc, routeLongName, routeShortName, routeTextColor, routeUrl, continuousPickup, continuousDropOff string
	var routeType, routeSortOrder int

	lib.ParseStringToPrimitive(m["agency_id"], &agencyId)
	lib.ParseStringToPrimitive(m["route_color"], &routeColor)
	lib.ParseStringToPrimitive(m["route_desc"], &routeDesc)
	lib.ParseStringToPrimitive(m["route_long_name"], &routeLongName)
	lib.ParseStringToPrimitive(m["route_short_name"], &routeShortName)
	lib.ParseStringToPrimitive(m["route_text_color"], &routeTextColor)
	lib.ParseStringToPrimitive(m["route_url"], &routeUrl)
	lib.ParseStringToPrimitive(m["continuous_pickup"], &continuousPickup)
	lib.ParseStringToPrimitive(m["continuous_drop_off"], &continuousDropOff)
	lib.ParseStringToPrimitive(m["route_type"], &routeType)
	lib.ParseStringToPrimitive(m["route_sort_order"], &routeSortOrder)

	route.AgencyId = lib.IfThenElse(m["agency_id"] != "", &agencyId, nil)
	route.RouteColor = lib.IfThenElse(m["route_color"] != "", &routeColor, nil)
	route.RouteDesc = lib.IfThenElse(m["route_desc"] != "", &routeDesc, nil)
	route.RouteLongName = lib.IfThenElse(m["route_long_name"] != "", &routeLongName, nil)
	route.RouteShortName = lib.IfThenElse(m["route_short_name"] != "", &routeShortName, nil)
	route.RouteTextColor = lib.IfThenElse(m["route_text_color"] != "", &routeTextColor, nil)
	route.RouteUrl = lib.IfThenElse(m["route_url"] != "", &routeUrl, nil)
	route.ContinuousPickup = lib.IfThenElse(m["continuous_pickup"] != "", &continuousPickup, nil)
	route.ContinuousDropOff = lib.IfThenElse(m["continuous_drop_off"] != "", &continuousDropOff, nil)
	route.RouteSortOrder = lib.IfThenElse(m["route_sort_order"] != "", &routeSortOrder, nil)

	// Convert Required Values
	lib.ParseStringToPrimitive(m["route_id"], &route.RouteId)
	lib.ParseStringToPrimitive(m["route_type"], &route.RouteType)

	if len(parsingErrors) > 0 {
		for _, err := range parsingErrors {
			messages = append(messages, types.Message{
				Field:   "N/A", //TODO: Add field name
				Message: err,
			})
		}
	}

	// Validate Values
	// Validate Required route_id
	if route.RouteId == "" {
		messages = append(messages, types.Message{
			Field:   "route_id",
			Message: "Route ID is required and must be unique.",
		})
	}

	// Validate Required route_type
	if route.RouteType == 0 && m["route_type"] == "" {
		messages = append(messages, types.Message{
			Field:   "route_type",
			Message: "Route type is required.",
		})
	}

	// Validate route_type enum values
	validRouteTypes := map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 11: true, 12: true}
	if !validRouteTypes[route.RouteType] {
		messages = append(messages, types.Message{
			Field:   "route_type",
			Message: "Invalid route_type. Valid values are 0, 1, 2, 3, 4, 5, 6, 7, 11, 12.",
		})
	}

	// Validate agency_id based on multiple agencies
	if multipleAgencies && (route.AgencyId == nil || *route.AgencyId == "") {
		messages = append(messages, types.Message{
			Field:   "agency_id",
			Message: "Agency ID is required when multiple agencies are defined in agency.txt.",
		})
	}

	// Validate route_short_name and route_long_name
	if (route.RouteShortName == nil || *route.RouteShortName == "") && (route.RouteLongName == nil || *route.RouteLongName == "") {
		messages = append(messages, types.Message{
			Field:   "route_short_name/route_long_name",
			Message: "At least one of route_short_name or route_long_name must be provided.",
		})
	}

	// Validate route_short_name length
	if route.RouteShortName != nil && *route.RouteShortName != "" && len(*route.RouteShortName) > 12 {
		messages = append(messages, types.Message{
			Field:   "route_short_name",
			Message: "Route short name should be no longer than 12 characters.",
		})
	}

	// Validate URLs if provided
	if route.RouteUrl != nil && *route.RouteUrl != "" {
		if urlErrors := lib.ValidateUrl(*route.RouteUrl); urlErrors != "" {
			messages = append(messages, types.Message{
				Field:   "route_url",
				Message: urlErrors,
			})
		}
	}

	// Validate continuous_pickup enum values
	if route.ContinuousPickup != nil && *route.ContinuousPickup != "" {
		validContinuousPickup := map[string]bool{"0": true, "1": true, "2": true, "3": true}
		if !validContinuousPickup[*route.ContinuousPickup] {
			messages = append(messages, types.Message{
				Field:   "continuous_pickup",
				Message: "Invalid continuous_pickup value. Valid values are 0, 1, 2, 3.",
			})
		}
	}

	// Validate continuous_pickup is forbidden if stop_times have pickup/dropoff windows
	if hasPickupDropoffWindows && route.ContinuousPickup != nil && *route.ContinuousPickup != "" {
		messages = append(messages, types.Message{
			Field:   "continuous_pickup",
			Message: "continuous_pickup is forbidden when stop_times.start_pickup_drop_off_window or stop_times.end_pickup_drop_off_window are defined for any trip of this route.",
		})
	}

	// Validate continuous_drop_off enum values
	if route.ContinuousDropOff != nil && *route.ContinuousDropOff != "" {
		validContinuousDropOff := map[string]bool{"0": true, "1": true, "2": true, "3": true}
		if !validContinuousDropOff[*route.ContinuousDropOff] {
			messages = append(messages, types.Message{
				Field:   "continuous_drop_off",
				Message: "Invalid continuous_drop_off value. Valid values are 0, 1, 2, 3.",
			})
		}
	}

	// Validate continuous_drop_off is forbidden if stop_times have pickup/dropoff windows
	if hasPickupDropoffWindows && route.ContinuousDropOff != nil && *route.ContinuousDropOff != "" {
		messages = append(messages, types.Message{
			Field:   "continuous_drop_off",
			Message: "continuous_drop_off is forbidden when stop_times.start_pickup_drop_off_window or stop_times.end_pickup_drop_off_window are defined for any trip of this route.",
		})
	}

	// Validate route_sort_order is non-negative
	if route.RouteSortOrder != nil && *route.RouteSortOrder < 0 {
		messages = append(messages, types.Message{
			Field:   "route_sort_order",
			Message: "Route sort order must be a non-negative integer.",
		})
	}

	return route, messages
}

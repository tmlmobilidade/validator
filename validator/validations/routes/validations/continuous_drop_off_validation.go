package routes

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

- File: [routes.txt]
- Field: continuous_drop_off
- Presence: Conditionally Forbidden
- Type: Enum

# Description

Indicates that the rider can board the transit vehicle at any point along the vehicle’s travel path as described by shapes.txt, on every trip of the route.

Valid options are:

  - 0 - Continuous stopping drop off.
  - 1 or empty - No continuous stopping drop off.
  - 2 - Must phone agency to arrange continuous stopping drop off.
  - 3 - Must coordinate with driver to arrange continuous stopping drop off.

Values for `routes.continuous_drop_off` may be overridden by defining values in `stop_times.continuous_drop_off` for specific `stop_times` along the route.

Conditionally Forbidden:
- Any value other than `1` or `empty` is Forbidden if `stop_times.start_pickup_drop_off_window` or `stop_times.end_pickup_drop_off_window` are defined for any trip of this route.
- Optional otherwise.

[routes.txt]: https://gtfs.org/schedule/reference/#routestxt
*/
func ContinuousDropOffValidation(severity *types.Severity, route *types.Route, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "continuous_drop_off",
			FileName:     "routes.txt",
			ValidationID: "continuous_drop_off_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	// If continuous_drop_off is "1", it's valid and we can return early
	if route.ContinuousDropOff != nil && *route.ContinuousDropOff == "1" {
		return
	}

	if route.ContinuousDropOff == nil || *route.ContinuousDropOff == "" {
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "continuous_drop_off is recommended", "continuous_drop_off is required")
		addMessage(warn, s)
		return
	}
	
	// Get all trip IDs for this route
	tripIds := make([]string, 0)
	if trips, exists := gtfs.IdMap["trips"][*route.RouteId]; exists {
		for _, row := range trips {
			if tripId := gtfs.Files["trips"][row]["trip_id"]; tripId != "" {
				tripIds = append(tripIds, tripId)
			}
		}
	}

	// Check each trip's stop times for pickup/dropoff windows
	for _, tripId := range tripIds {
		if stopTimes, exists := gtfs.IdMap["stop_times"][tripId]; exists {
			for _, row := range stopTimes {
				startWindow := gtfs.Files["stop_times"][row]["start_pickup_drop_off_window"]
				endWindow := gtfs.Files["stop_times"][row]["end_pickup_drop_off_window"]
				
				if startWindow != "" || endWindow != "" {
					lib.AppLogger.Accent("route.ContinuousDropOff", *route.ContinuousDropOff)
						addMessage("continuous_drop_off must be 1 or empty if start_pickup_drop_off_window or end_pickup_drop_off_window is defined", types.SEVERITY_ERROR)
					return
				}
			}
		}
	}
}
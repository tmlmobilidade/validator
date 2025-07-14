package routes

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
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
func ContinuousDropOffValidation(route *types.Route, row int, gtfs *types.Gtfs, rules *types.RoutesRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.ContinuousDropOff.Severity != "" {
		s = rules.ContinuousDropOff.Severity
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

		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, i18n.AppTranslator.Get("continuous_drop_off_validation.recommended"), i18n.AppTranslator.Get("continuous_drop_off_validation.required"))
		addMessage(warn, s)
		return
	}

	// Get all trip IDs for this route
	tripIds := make([]string, 0)
	if trips, exists := gtfs.IdMap["trips"][*route.RouteId]; exists {
		for _, row := range trips {
			if tripId := gtfs.Trip[row].TripId; tripId != "" {
				tripIds = append(tripIds, tripId)
			}
		}
	}

	// Check each trip's stop times for pickup/dropoff windows
	for _, tripId := range tripIds {
		if stopTimes, exists := gtfs.IdMap["stop_times"][tripId]; exists {
			for _, row := range stopTimes {
				startWindow := gtfs.StopTime[row].StartPickupDropOffWindow
				endWindow := gtfs.StopTime[row].EndPickupDropOffWindow

				if startWindow != "" || endWindow != "" {
					lib.AppLogger.Accent("route.ContinuousDropOff", *route.ContinuousDropOff)
					addMessage(i18n.AppTranslator.Get("continuous_drop_off_validation.forbidden_with_window"), types.SEVERITY_ERROR)
					return
				}
			}
		}
	}

	// Validate rules
	if rules != nil && rules.ContinuousDropOff.Options != nil {
		if slices.Contains(*rules.ContinuousDropOff.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.ContinuousDropOff.Options, *route.ContinuousDropOff) {
			addMessage(i18n.AppTranslator.Get("continuous_drop_off_validation.not_allowed", map[string]interface{}{"value": *route.ContinuousDropOff}), s)
			return
		}
	}
}

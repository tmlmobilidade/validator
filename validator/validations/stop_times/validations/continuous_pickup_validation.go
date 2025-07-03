package stop_times

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

  - File: [stop_times.txt]
  - Field: continuous_pickup
  - Presence: Conditionally Forbidden
  - Type: Enum

# Description

Indicates that the rider can board the transit vehicle at any point along the vehicle's travel path as described by shapes.txt, from this stop_time to the next stop_time in the trip's stop_sequence.

Valid options are:

  - 0 - Continuous stopping pickup.
  - 1 or empty - No continuous stopping pickup.
  - 2 - Must phone agency to arrange continuous stopping pickup.
  - 3 - Must coordinate with driver to arrange continuous stopping pickup.

If this field is populated, it overrides any continuous pickup behavior defined in routes.txt. If this field is empty, the stop_time inherits any continuous pickup behavior defined in routes.txt.

Conditionally Forbidden:

  - Any value other than 1 or empty is Forbidden if start_pickup_drop_off_window or end_pickup_drop_off_window are defined.
  - Optional otherwise.

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func ContinuousPickupValidation(stopTime *types.StopTime, row int, rules *types.StopTimesRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.ContinuousPickup.Severity != "" {
		s = rules.ContinuousPickup.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "continuous_pickup",
			FileName:     "stop_times.txt",
			ValidationID: "continuous_pickup_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	// If not present, it's optional unless severity is set
	if stopTime.ContinuousPickup == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}
		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "continuous_pickup is recommended", "continuous_pickup is required")
		addMessage(warn, s)
		return
	}

	cp := *stopTime.ContinuousPickup
	if cp < 0 || cp > 3 {
		addMessage("continuous_pickup must be 0, 1, 2, or 3.", types.SEVERITY_ERROR)
		return
	}

	// Forbidden: Any value other than 1 or empty if start/end_pickup_drop_off_window are defined
	if ((stopTime.StartPickupDropOffWindow != nil && *stopTime.StartPickupDropOffWindow != "") || (stopTime.EndPickupDropOffWindow != nil && *stopTime.EndPickupDropOffWindow != "")) && (cp != 1) {
		addMessage("continuous_pickup must be 1 or empty if start_pickup_drop_off_window or end_pickup_drop_off_window are defined.", types.SEVERITY_ERROR)
		return
	}

	// Validate Rule Options
	if rules != nil && rules.ContinuousPickup.Options != nil {
		if slices.Contains(*rules.ContinuousPickup.Options, types.ALL_OPTIONS) {
			return
		}

		if slices.Contains(*rules.ContinuousPickup.Options, fmt.Sprintf("%d", cp)) {
			return
		}

		addMessage(fmt.Sprintf("continuous_pickup is not allowed: %d", cp), s)
		return
	}
}

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
  - Field: pickup_booking_rule_id
  - Presence: Optional
  - Type: Foreign ID referencing booking_rules.booking_rule_id

# Description

Identifies the boarding booking rule at this stop time.

Recommended when pickup_type=2.

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func PickupBookingRuleIdValidation(stopTime *types.StopTime, row int, gtfs *types.Gtfs, rules *types.StopTimesRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.PickupBookingRuleId.Severity != "" {
		s = rules.PickupBookingRuleId.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "pickup_booking_rule_id",
			FileName:     "stop_times.txt",
			ValidationID: "pickup_booking_rule_id_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	if stopTime.PickupBookingRuleId == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}
		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "pickup_booking_rule_id is recommended", "pickup_booking_rule_id is required")
		addMessage(warn, s)
		return
	}

	// Foreign key check: must reference a valid booking_rule_id from booking_rules.txt
	if !lib.GtfsIdMapKeyExists(gtfs, "booking_rules", *stopTime.PickupBookingRuleId) {
		addMessage("pickup_booking_rule_id '"+*stopTime.PickupBookingRuleId+"' does not exist in booking_rules.txt", types.SEVERITY_ERROR)
		return
	}

	// Validate Rule Options
	if rules != nil && rules.PickupBookingRuleId.Options != nil {
		if slices.Contains(*rules.PickupBookingRuleId.Options, types.ALL_OPTIONS) {
			return
		}

		if slices.Contains(*rules.PickupBookingRuleId.Options, *stopTime.PickupBookingRuleId) {
			return
		}

		addMessage(fmt.Sprintf("pickup_booking_rule_id is not allowed: %s", *stopTime.PickupBookingRuleId), s)
		return
	}
}

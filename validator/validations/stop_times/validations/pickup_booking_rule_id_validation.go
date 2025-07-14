package stop_times

import (
	"main/i18n"
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
		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, i18n.AppTranslator.Get("pickup_booking_rule_id_validation.recommended"), i18n.AppTranslator.Get("pickup_booking_rule_id_validation.required"))
		addMessage(warn, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("pickup_booking_rule_id_validation.forbidden"), s)
		return
	}

	// Foreign key check: must reference a valid booking_rule_id from booking_rules.txt
	if !lib.GtfsIdMapKeyExists(gtfs, "booking_rules", *stopTime.PickupBookingRuleId) {
		addMessage(i18n.AppTranslator.Get("pickup_booking_rule_id_validation.not_found", *stopTime.PickupBookingRuleId), types.SEVERITY_ERROR)
		return
	}

	// Validate Rule Options
	if rules != nil && rules.PickupBookingRuleId.Options != nil {
		if slices.Contains(*rules.PickupBookingRuleId.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.PickupBookingRuleId.Options, *stopTime.PickupBookingRuleId) {
			addMessage(i18n.AppTranslator.Get("pickup_booking_rule_id_validation.not_allowed", *stopTime.PickupBookingRuleId), s)
			return
		}
	}
}

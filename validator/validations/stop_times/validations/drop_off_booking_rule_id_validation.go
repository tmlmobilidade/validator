package stop_times

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [stop_times.txt]
  - Field: drop_off_booking_rule_id
  - Presence: Optional
  - Type: Foreign ID referencing booking_rules.booking_rule_id

# Description

Identifies the alighting booking rule at this stop time.

Recommended when drop_off_type=2.

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func DropOffBookingRuleIdValidation(stopTime *types.StopTime, row int, gtfs *types.Gtfs, rules *types.StopTimesRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.DropOffBookingRuleId.Severity != "" {
		s = rules.DropOffBookingRuleId.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "drop_off_booking_rule_id",
			FileName:     "stop_times.txt",
			ValidationID: "drop_off_booking_rule_id_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	if stopTime.DropOffBookingRuleId == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}
		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, i18n.AppTranslator.Get("drop_off_booking_rule_id_validation.recommended"), i18n.AppTranslator.Get("drop_off_booking_rule_id_validation.required"))
		addMessage(warn, s)
		return
	}

	// Foreign key check: must reference a valid booking_rule_id from booking_rules.txt
	if !lib.GtfsIdMapKeyExists(gtfs, "booking_rules", *stopTime.DropOffBookingRuleId) {
		addMessage(i18n.AppTranslator.Get("drop_off_booking_rule_id_validation.not_found", *stopTime.DropOffBookingRuleId), types.SEVERITY_ERROR)
		return
	}
}

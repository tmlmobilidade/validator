package stop_times

import (
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
func DropOffBookingRuleIdValidation(severity *types.Severity, stopTime *types.StopTime, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
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
		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, "drop_off_booking_rule_id is recommended", "drop_off_booking_rule_id is required")
		addMessage(warn, s)
		return
	}

	// Foreign key check: must reference a valid booking_rule_id from booking_rules.txt
	if gtfs.IdMap != nil {
		bookingRulesMap, ok := gtfs.IdMap["booking_rules"]
		if !ok || bookingRulesMap == nil {
			addMessage("booking_rules.txt is missing or not indexed.", types.SEVERITY_ERROR)
			return
		}
		rows, ok := bookingRulesMap[*stopTime.DropOffBookingRuleId]
		if !ok || len(rows) == 0 {
			addMessage("drop_off_booking_rule_id must reference a valid booking_rule_id from booking_rules.txt.", types.SEVERITY_ERROR)
			return
		}
	}
} 
package fare_rules

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

	- File: [fare_rules.txt]
	- Field: origin_id
	- Presence: Optional
	- Type: Foreign ID referencing [stops.zone_id]

# Description

Identifies an origin zone. If a fare class has multiple origin zones, create a record in fare_rules.txt for each origin_id.

# Example

If fare class "b" is valid for all travel originating from either zone "2" or zone "8", the fare_rules.txt file would contain these records for the fare class:

	fare_id,...,origin_id
	b,...,2
	b,...,8

[fare_rules.txt]: https://gtfs.org/schedule/reference/#fare_rulestxt
[stops.zone_id]: https://gtfs.org/schedule/reference/#stopstxt
*/
func OriginIdValidation(fareRule *types.FareRule, row int, gtfs *types.Gtfs, severity *types.Severity) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "origin_id",
			FileName:     "fare_rules.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "fare_rules_parse",
		})
	}

	if fareRule.OriginId == nil {
		
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "required", "recommended")
		addMessage(fmt.Sprintf("Origin ID is %s", warn), s)
		return
	}

	if _, ok := gtfs.IdMap["stops"][*fareRule.OriginId]; !ok {
		addMessage("Origin ID does not exist in stops.txt", types.SEVERITY_ERROR)
		return
	}
} 
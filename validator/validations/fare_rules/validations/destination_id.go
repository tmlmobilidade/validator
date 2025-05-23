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
	- Field: destination_id
	- Presence: Optional
	- Type: Foreign ID referencing [stops.zone_id]

# Description

Identifies a destination zone. If a fare class has multiple destination zones, create a record in fare_rules.txt for each destination_id.

# Example

The destination_id and destination_id fields could be used together to specify that fare class "b" is valid for travel between zones 3 and 4, and for travel between zones 3 and 5, the fare_rules.txt file would contain these records for the fare class:

	fare_id,...,destination_id,destination_id
	b,...,3,4
	b,...,3,5

[fare_rules.txt]: https://gtfs.org/schedule/reference/#fare_rulestxt
[stops.zone_id]: https://gtfs.org/schedule/reference/#stopstxt
*/
func DestinationIdValidation(fareRule *types.FareRule, row int, gtfs *types.Gtfs, severity *types.Severity) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "destination_id",
			FileName:     "fare_rules.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "fare_rules_parse",
		})
	}

	if fareRule.DestinationId == nil {
		
		if s == types.SEVERITY_IGNORE {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, "required", "recommended")
		addMessage(fmt.Sprintf("Destination ID is %s", warn), s)
		return
	}

	if _, ok := gtfs.IdMap["stops"][*fareRule.DestinationId]; !ok {
		addMessage("Destination ID does not exist in stops.txt", types.SEVERITY_ERROR)
		return
	}
} 
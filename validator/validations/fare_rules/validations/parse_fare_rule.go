package fare_rules

import (
	"main/lib"
	"main/services"
	"main/types"
)

func ParseFareRule(rawFareRule map[string]string, row int, gtfs *types.Gtfs) types.FareRule {
	var (
		fareRule types.FareRule = types.FareRule{}
		fareId, routeId, originId, destinationId, containsId string
		messages []types.Message
	)

	stringFields := map[string]*string{
		"fare_id": &fareId,
		"route_id": &routeId,
		"origin_id": &originId,
		"destination_id": &destinationId,
		"contains_id": &containsId,
	}

	// Helper to collect error messages
	addMessage := func(field, msg string) {
		messages = append(messages, types.Message{
			Field:        field,
			FileName:     "fare_rules.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "fare_rules_parse",
		})
	}

	// Parse string fields
	for field, target := range stringFields {
		if errMsg := lib.ParseStringToPrimitive(rawFareRule[field], target); errMsg != "" {
			addMessage(field, errMsg)
		}
	}

	if len(messages) > 0 {
		services.AppMessageService.AddMessages(messages)
		return fareRule
	}
	
	fareRule.FareId = lib.IfThenElse(rawFareRule["fare_id"] != "", &fareId, nil)
	fareRule.RouteId = lib.IfThenElse(rawFareRule["route_id"] != "", &routeId, nil)
	fareRule.OriginId = lib.IfThenElse(rawFareRule["origin_id"] != "", &originId, nil)
	fareRule.DestinationId = lib.IfThenElse(rawFareRule["destination_id"] != "", &destinationId, nil)
	fareRule.ContainsId = lib.IfThenElse(rawFareRule["contains_id"] != "", &containsId, nil)
	
	return fareRule
}
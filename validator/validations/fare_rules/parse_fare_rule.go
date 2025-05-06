package fare_rules

import (
	"main/lib"
	"main/types"
)

type parseFareRuleValidation struct {
	*types.Validation
}

func NewParseFareRuleValidation(severity *types.Severity) *parseFareRuleValidation {
	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	return &parseFareRuleValidation{
		Validation: &types.Validation{
			ID:          "parse_fare_rule",
			Description: "Validate fare rule data",
			Severity:    s,
		},
	}
}

func (v *parseFareRuleValidation) Validate(gtfs types.Gtfs) (fareRules []types.FareRule, messages []types.Message) {
	// Get reference maps for foreign key validation
	fareAttributeIds := gtfs.IdMap["fare_attributes"]
	routeIds := gtfs.IdMap["routes"]
	zoneIds := make(map[string]bool)

	// Build zone IDs map from stops.txt
	for _, stop := range gtfs.Files["stops"] {
		if zoneId := stop["zone_id"]; zoneId != "" {
			zoneIds[zoneId] = true
		}
	}

	for i, fareRule := range gtfs.Files["fare_rules"] {
		fareRule, fareRuleMessages := parseFareRule(fareRule, fareAttributeIds, routeIds, zoneIds)
		fareRules = append(fareRules, fareRule)

		// Update row number and other fields for each message
		for _, msg := range fareRuleMessages {
			msg.Rows = []int{i + 1}
			msg.FileName = "fare_rules.txt"
			msg.Severity = v.Severity
			msg.ValidationID = v.ID
			messages = append(messages, msg)
		}
	}

	return fareRules, messages
}

func parseFareRule(m map[string]string, fareAttributeIds map[string]int, routeIds map[string]int, zoneIds map[string]bool) (fareRule types.FareRule, messages []types.Message) {
	var parsingErrors []string

	// Parse required field
	lib.ParseStringToPrimitive(m["fare_id"], &fareRule.FareId)

	// Parse optional fields
	var routeId, originId, destinationId, containsId string
	lib.ParseStringToPrimitive(m["route_id"], &routeId)
	lib.ParseStringToPrimitive(m["origin_id"], &originId)
	lib.ParseStringToPrimitive(m["destination_id"], &destinationId)
	lib.ParseStringToPrimitive(m["contains_id"], &containsId)

	fareRule.RouteId = lib.IfThenElse(m["route_id"] != "", &routeId, nil)
	fareRule.OriginId = lib.IfThenElse(m["origin_id"] != "", &originId, nil)
	fareRule.DestinationId = lib.IfThenElse(m["destination_id"] != "", &destinationId, nil)
	fareRule.ContainsId = lib.IfThenElse(m["contains_id"] != "", &containsId, nil)

	if len(parsingErrors) > 0 {
		for _, err := range parsingErrors {
			messages = append(messages, types.Message{
				Field:   "N/A",
				Message: err,
			})
		}
	}

	// Validate required fare_id
	if fareRule.FareId == "" {
		messages = append(messages, types.Message{
			Field:   "fare_id",
			Message: "Fare ID is required.",
		})
	} else {
		// Validate fare_id references a valid fare_attributes.fare_id
		if _, ok := fareAttributeIds[fareRule.FareId]; !ok {
			messages = append(messages, types.Message{
				Field:   "fare_id",
				Message: "Fare ID must reference a valid fare_id from fare_attributes.txt.",
			})
		}
	}

	// Validate optional route_id if provided
	if fareRule.RouteId != nil && *fareRule.RouteId != "" {
		if _, ok := routeIds[*fareRule.RouteId]; !ok {
			messages = append(messages, types.Message{
				Field:   "route_id",
				Message: "Route ID must reference a valid route_id from routes.txt.",
			})
		}
	}

	// Validate optional origin_id if provided
	if fareRule.OriginId != nil && *fareRule.OriginId != "" {
		if !zoneIds[*fareRule.OriginId] {
			messages = append(messages, types.Message{
				Field:   "origin_id",
				Message: "Origin ID must reference a valid zone_id from stops.txt.",
			})
		}
	}

	// Validate optional destination_id if provided
	if fareRule.DestinationId != nil && *fareRule.DestinationId != "" {
		if !zoneIds[*fareRule.DestinationId] {
			messages = append(messages, types.Message{
				Field:   "destination_id",
				Message: "Destination ID must reference a valid zone_id from stops.txt.",
			})
		}
	}

	// Validate optional contains_id if provided
	if fareRule.ContainsId != nil && *fareRule.ContainsId != "" {
		if !zoneIds[*fareRule.ContainsId] {
			messages = append(messages, types.Message{
				Field:   "contains_id",
				Message: "Contains ID must reference a valid zone_id from stops.txt.",
			})
		}
	}

	return fareRule, messages
}

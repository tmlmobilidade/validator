package fare_rules

import (
	"testing"
)

func TestParseFareRule_ValidFareRule(t *testing.T) {
	// Test a valid fare rule with all fields
	input := map[string]string{
		"fare_id":        "fare1",
		"route_id":       "route1",
		"origin_id":      "zone1",
		"destination_id": "zone2",
		"contains_id":    "zone3",
	}

	fareAttributeIds := map[string]int{"fare1": 1}
	routeIds := map[string]int{"route1": 1}
	zoneIds := map[string]bool{"zone1": true, "zone2": true, "zone3": true}

	fareRule, messages := parseFareRule(input, fareAttributeIds, routeIds, zoneIds)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the fare rule was parsed correctly
	if fareRule.FareId != "fare1" {
		t.Errorf("Expected fare_id to be 'fare1', got '%s'", fareRule.FareId)
	}
	if *fareRule.RouteId != "route1" {
		t.Errorf("Expected route_id to be 'route1', got '%s'", *fareRule.RouteId)
	}
	if *fareRule.OriginId != "zone1" {
		t.Errorf("Expected origin_id to be 'zone1', got '%s'", *fareRule.OriginId)
	}
	if *fareRule.DestinationId != "zone2" {
		t.Errorf("Expected destination_id to be 'zone2', got '%s'", *fareRule.DestinationId)
	}
	if *fareRule.ContainsId != "zone3" {
		t.Errorf("Expected contains_id to be 'zone3', got '%s'", *fareRule.ContainsId)
	}
}

func TestParseFareRule_MinimalValidFareRule(t *testing.T) {
	// Test a minimal valid fare rule with only required fields
	input := map[string]string{
		"fare_id": "fare1",
	}

	fareAttributeIds := map[string]int{"fare1": 1}
	routeIds := map[string]int{}
	zoneIds := map[string]bool{}

	fareRule, messages := parseFareRule(input, fareAttributeIds, routeIds, zoneIds)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the fare rule was parsed correctly
	if fareRule.FareId != "fare1" {
		t.Errorf("Expected fare_id to be 'fare1', got '%s'", fareRule.FareId)
	}
	if fareRule.RouteId != nil {
		t.Error("Expected route_id to be nil")
	}
	if fareRule.OriginId != nil {
		t.Error("Expected origin_id to be nil")
	}
	if fareRule.DestinationId != nil {
		t.Error("Expected destination_id to be nil")
	}
	if fareRule.ContainsId != nil {
		t.Error("Expected contains_id to be nil")
	}
}

func TestParseFareRule_MissingRequiredFields(t *testing.T) {
	// Test a fare rule with missing required fields
	input := map[string]string{
		"route_id": "route1",
	}

	fareAttributeIds := map[string]int{}
	routeIds := map[string]int{}
	zoneIds := map[string]bool{}

	_, messages := parseFareRule(input, fareAttributeIds, routeIds, zoneIds)

	// Check for validation message for missing fare_id
	found := false
	for _, msg := range messages {
		if msg.Field == "fare_id" && msg.Message == "Fare ID is required." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for missing fare_id not found")
	}
}

func TestParseFareRule_InvalidFareId(t *testing.T) {
	// Test a fare rule with an invalid fare_id
	input := map[string]string{
		"fare_id": "invalid_fare",
	}

	fareAttributeIds := map[string]int{"fare1": 1}
	routeIds := map[string]int{}
	zoneIds := map[string]bool{}

	_, messages := parseFareRule(input, fareAttributeIds, routeIds, zoneIds)

	// Check for validation message for invalid fare_id
	found := false
	for _, msg := range messages {
		if msg.Field == "fare_id" && msg.Message == "Fare ID must reference a valid fare_id from fare_attributes.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid fare_id not found")
	}
}

func TestParseFareRule_InvalidRouteId(t *testing.T) {
	// Test a fare rule with an invalid route_id
	input := map[string]string{
		"fare_id":  "fare1",
		"route_id": "invalid_route",
	}

	fareAttributeIds := map[string]int{"fare1": 1}
	routeIds := map[string]int{"route1": 1}
	zoneIds := map[string]bool{}

	_, messages := parseFareRule(input, fareAttributeIds, routeIds, zoneIds)

	// Check for validation message for invalid route_id
	found := false
	for _, msg := range messages {
		if msg.Field == "route_id" && msg.Message == "Route ID must reference a valid route_id from routes.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid route_id not found")
	}
}

func TestParseFareRule_InvalidZoneIds(t *testing.T) {
	// Test a fare rule with invalid zone IDs
	input := map[string]string{
		"fare_id":        "fare1",
		"origin_id":      "invalid_zone1",
		"destination_id": "invalid_zone2",
		"contains_id":    "invalid_zone3",
	}

	fareAttributeIds := map[string]int{"fare1": 1}
	routeIds := map[string]int{}
	zoneIds := map[string]bool{"zone1": true, "zone2": true, "zone3": true}

	_, messages := parseFareRule(input, fareAttributeIds, routeIds, zoneIds)

	// Check for validation messages for invalid zone IDs
	expectedErrors := map[string]bool{
		"Origin ID must reference a valid zone_id from stops.txt.":      false,
		"Destination ID must reference a valid zone_id from stops.txt.": false,
		"Contains ID must reference a valid zone_id from stops.txt.":    false,
	}

	for _, msg := range messages {
		expectedErrors[msg.Message] = true
	}

	for errMsg, found := range expectedErrors {
		if !found {
			t.Errorf("Expected error message not found: '%s'", errMsg)
		}
	}
}

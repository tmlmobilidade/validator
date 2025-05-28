package fare_rules

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_rules/validations"
	"testing"
)

func TestParseFareRule_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	rawFareRule := map[string]string{
		"fare_id": "FARE1",
		"route_id": "ROUTE1",
		"origin_id": "ORIGIN1",
		"destination_id": "DEST1",
		"contains_id": "CONTAIN1",
	}
	gtfs := &types.Gtfs{}
	fareRule := validations.ParseFareRule(rawFareRule, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid input should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	if fareRule.FareId == nil || *fareRule.FareId != "FARE1" {
		t.Errorf("Expected FareId to be 'FARE1', got %v", fareRule.FareId)
	}
	if fareRule.RouteId == nil || *fareRule.RouteId != "ROUTE1" {
		t.Errorf("Expected RouteId to be 'ROUTE1', got %v", fareRule.RouteId)
	}
	if fareRule.OriginId == nil || *fareRule.OriginId != "ORIGIN1" {
		t.Errorf("Expected OriginId to be 'ORIGIN1', got %v", fareRule.OriginId)
	}
	if fareRule.DestinationId == nil || *fareRule.DestinationId != "DEST1" {
		t.Errorf("Expected DestinationId to be 'DEST1', got %v", fareRule.DestinationId)
	}
	if fareRule.ContainsId == nil || *fareRule.ContainsId != "CONTAIN1" {
		t.Errorf("Expected ContainsId to be 'CONTAIN1', got %v", fareRule.ContainsId)
	}
} 
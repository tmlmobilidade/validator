package trips

import (
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestParseTrips_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	raw := map[string]string{
		"trip_id":               "T1",
		"route_id":              "R1",
		"service_id":            "S1",
		"trip_headsign":         "Headsign",
		"trip_short_name":       "Short",
		"direction_id":          "1",
		"block_id":              "B1",
		"shape_id":              "SH1",
		"wheelchair_accessible": "2",
		"bikes_allowed":         "1",
	}
	trip := validations.ParseTrips(raw, row)

	if *trip.TripId != "T1" {
		t.Errorf("Expected TripId 'T1', got '%s'", *trip.TripId)
	}
	if *trip.RouteId != "R1" {
		t.Errorf("Expected RouteId 'R1', got '%s'", *trip.RouteId)
	}
	if *trip.ServiceId != "S1" {
		t.Errorf("Expected ServiceId 'S1', got '%s'", *trip.ServiceId)
	}
	if trip.TripHeadsign == nil || *trip.TripHeadsign != "Headsign" {
		t.Errorf("Expected TripHeadsign 'Headsign', got '%v'", trip.TripHeadsign)
	}
	if trip.TripShortName == nil || *trip.TripShortName != "Short" {
		t.Errorf("Expected TripShortName 'Short', got '%v'", trip.TripShortName)
	}
	if trip.BlockId == nil || *trip.BlockId != "B1" {
		t.Errorf("Expected BlockId 'B1', got '%v'", trip.BlockId)
	}
	if trip.ShapeId == nil || *trip.ShapeId != "SH1" {
		t.Errorf("Expected ShapeId 'SH1', got '%v'", trip.ShapeId)
	}
	if trip.DirectionId == nil || *trip.DirectionId != 1 {
		t.Errorf("Expected DirectionId 1, got '%v'", trip.DirectionId)
	}
	if trip.WheelchairAccessible == nil || *trip.WheelchairAccessible != 2 {
		t.Errorf("Expected WheelchairAccessible 2, got '%v'", trip.WheelchairAccessible)
	}
	if trip.BikesAllowed == nil || *trip.BikesAllowed != 1 {
		t.Errorf("Expected BikesAllowed 1, got '%v'", trip.BikesAllowed)
	}
}

func TestParseTrips_ValidIntField(t *testing.T) {
	services.AppMessageService.Clear()
	row := 2
	raw := map[string]string{
		"trip_id":               "T1",
		"route_id":              "R1",
		"service_id":            "S1",
		"trip_headsign":         "Headsign",
		"trip_short_name":       "Short",
		"direction_id":          "1",
		"block_id":              "B1",
		"shape_id":              "SH1",
		"wheelchair_accessible": "2",
		"bikes_allowed":         "1",
	}
	trip := validations.ParseTrips(raw, row)

	if trip.DirectionId == nil || *trip.DirectionId != 1 {
		t.Errorf("Expected DirectionId 1, got '%v'", trip.DirectionId)
	}
	
	if trip.WheelchairAccessible == nil || *trip.WheelchairAccessible != 2 {
		t.Errorf("Expected WheelchairAccessible 2, got '%v'", trip.WheelchairAccessible)
	}

	if trip.BikesAllowed == nil || *trip.BikesAllowed != 1 {
		t.Errorf("Expected BikesAllowed 1, got '%v'", trip.BikesAllowed)
	}
}

func TestParseTrips_InvalidIntField(t *testing.T) {
	services.AppMessageService.Clear()
	row := 3
	raw := map[string]string{
		"trip_id":               "T1",
		"route_id":              "R1",
		"service_id":            "S1",
		"trip_headsign":         "Headsign",
		"trip_short_name":       "Short",
		"direction_id":          "not_an_int",
		"block_id":              "B1",
		"shape_id":              "SH1",
		"wheelchair_accessible": "2",
		"bikes_allowed":         "1",
	}
	trip := validations.ParseTrips(raw, row)

	if trip != (types.Trip{}) {
		t.Errorf("Expected empty Trip struct when int field is invalid, got '%+v'", trip)
	}
	if services.AppMessageService.GetSummary().TotalErrors != 1 {
		t.Errorf("Expected 1 error message, got %d", services.AppMessageService.GetSummary().TotalErrors)
	}
}

func TestParseTrips_OptionalFieldsEmpty(t *testing.T) {
	services.AppMessageService.Clear()
	row := 4
	raw := map[string]string{
		"trip_id":               "T2",
		"route_id":              "R2",
		"service_id":            "S2",
		"trip_headsign":         "",
		"trip_short_name":       "",
		"direction_id":          "",
		"block_id":              "",
		"shape_id":              "",
		"wheelchair_accessible": "",
		"bikes_allowed":         "",
	}
	trip := validations.ParseTrips(raw, row)

	if *trip.TripId != "T2" {
		t.Errorf("Expected TripId 'T2', got '%s'", *trip.TripId)
	}
	if *trip.RouteId != "R2" {
		t.Errorf("Expected RouteId 'R2', got '%s'", *trip.RouteId)
	}
	if *trip.ServiceId != "S2" {
		t.Errorf("Expected ServiceId 'S2', got '%s'", *trip.ServiceId)
	}
	if trip.TripHeadsign != nil {
		t.Errorf("Expected TripHeadsign nil, got '%v'", trip.TripHeadsign)
	}
	if trip.TripShortName != nil {
		t.Errorf("Expected TripShortName nil, got '%v'", trip.TripShortName)
	}
	if trip.BlockId != nil {
		t.Errorf("Expected BlockId nil, got '%v'", trip.BlockId)
	}
	if trip.ShapeId != nil {
		t.Errorf("Expected ShapeId nil, got '%v'", trip.ShapeId)
	}
	if trip.DirectionId != nil {
		t.Errorf("Expected DirectionId nil, got '%v'", trip.DirectionId)
	}
	if trip.WheelchairAccessible != nil {
		t.Errorf("Expected WheelchairAccessible nil, got '%v'", trip.WheelchairAccessible)
	}
	if trip.BikesAllowed != nil {
		t.Errorf("Expected BikesAllowed nil, got '%v'", trip.BikesAllowed)
	}
}

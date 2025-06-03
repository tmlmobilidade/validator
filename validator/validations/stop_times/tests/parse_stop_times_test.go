package stop_times

import (
	"main/lib"
	"main/services"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestParseStopTimes_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	raw := map[string]string{
		"trip_id": "T1",
		"arrival_time": "08:00:00",
		"departure_time": "09:00:00",
		"stop_id": "S1",
		"stop_sequence": "1",
		"stop_headsign": "Headsign",
		"start_pickup_drop_off_window": "07:00:00",
		"end_pickup_drop_off_window": "09:00:00",
		"pickup_type": "1",
		"drop_off_type": "1",
		"continuous_pickup": "1",
		"continuous_drop_off": "1",
		"shape_dist_traveled": "100.0",
		"timepoint": "1",
		"pickup_booking_rule_id": "R1",
		"drop_off_booking_rule_id": "R2",
		"location_group_id": "L1",
		"location_id": "L2",
	}
	stopTime := validations.ParseStopTimes(raw, row)

	if *stopTime.TripId != "T1" {
		t.Errorf("Expected TripId 'T1', got '%s'", *stopTime.TripId)
	}

	if *stopTime.ArrivalTime != "08:00:00" {
		t.Errorf("Expected ArrivalTime '08:00:00', got '%s'", *stopTime.ArrivalTime)
	}

	if *stopTime.DepartureTime != "09:00:00" {
		t.Errorf("Expected DepartureTime '09:00:00', got '%s'", *stopTime.DepartureTime)
	}

	if *stopTime.StopId != "S1" {
		t.Errorf("Expected StopId 'S1', got '%s'", *stopTime.StopId)
	}

	if *stopTime.StopSequence != 1 {
		t.Errorf("Expected StopSequence '1', got '%d'", *stopTime.StopSequence)
	}

	if *stopTime.StopHeadsign != "Headsign" {
		t.Errorf("Expected StopHeadsign 'Headsign', got '%s'", *stopTime.StopHeadsign)
	}

	if *stopTime.StartPickupDropOffWindow != "07:00:00" {
		t.Errorf("Expected StartPickupDropOffWindow '07:00:00', got '%s'", *stopTime.StartPickupDropOffWindow)
	}

	if *stopTime.EndPickupDropOffWindow != "09:00:00" {
		t.Errorf("Expected EndPickupDropOffWindow '09:00:00', got '%s'", *stopTime.EndPickupDropOffWindow)
	}

	if *stopTime.PickupType != 1 {
		t.Errorf("Expected PickupType '1', got '%d'", *stopTime.PickupType)
	}

	if *stopTime.DropOffType != 1 {
		t.Errorf("Expected DropOffType '1', got '%d'", *stopTime.DropOffType)
	}

	if *stopTime.ContinuousPickup != 1 {
		t.Errorf("Expected ContinuousPickup '1', got '%d'", *stopTime.ContinuousPickup)
	}

	if *stopTime.ContinuousDropOff != 1 {
		t.Errorf("Expected ContinuousDropOff '1', got '%d'", *stopTime.ContinuousDropOff)
	}

	if *stopTime.ShapeDistTraveled != 100.0 {
		t.Errorf("Expected ShapeDistTraveled '100.0', got '%f'", *stopTime.ShapeDistTraveled)
	}

	if *stopTime.Timepoint != 1 {
		t.Errorf("Expected Timepoint '1', got '%d'", *stopTime.Timepoint)
	}

	if *stopTime.PickupBookingRuleId != "R1" {
		t.Errorf("Expected PickupBookingRuleId 'R1', got '%s'", *stopTime.PickupBookingRuleId)
	}

	if *stopTime.DropOffBookingRuleId != "R2" {
		t.Errorf("Expected DropOffBookingRuleId 'R2', got '%s'", *stopTime.DropOffBookingRuleId)
	}

	if *stopTime.LocationGroupId != "L1" {
		t.Errorf("Expected LocationGroupId 'L1', got '%s'", *stopTime.LocationGroupId)
	}

	if *stopTime.LocationId != "L2" {
		t.Errorf("Expected LocationId 'L2', got '%s'", *stopTime.LocationId)
	}

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid stop_times should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestParseStopTimes_InvalidInts(t *testing.T) {
	services.AppMessageService.Clear()
	raw := map[string]string{
		"stop_sequence": "INVALID",
		"pickup_type": "INVALID",
		"drop_off_type": "INVALID",
		"continuous_pickup": "INVALID",
		"continuous_drop_off": "INVALID",
		"timepoint": "INVALID",
	}
	
	validations.ParseStopTimes(raw, 1)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid ints should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
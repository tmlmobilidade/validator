package stop_times

import (
	"main/src/types"
	"testing"
)

func TestParseStopTime_ValidStopTime(t *testing.T) {
	// Test a valid stop time with all fields
	input := map[string]string{
		"trip_id":                  "trip1",
		"arrival_time":             "08:00:00",
		"departure_time":           "08:05:00",
		"stop_id":                  "stop1",
		"stop_sequence":            "1",
		"stop_headsign":            "Downtown",
		"pickup_type":              "0",
		"drop_off_type":            "0",
		"continuous_pickup":        "0",
		"continuous_drop_off":      "0",
		"shape_dist_traveled":      "1.5",
		"timepoint":                "1",
		"pickup_booking_rule_id":   "booking1",
		"drop_off_booking_rule_id": "booking2",
	}

	tripIds := map[string]int{"trip1": 1}
	stopIds := map[string]int{"stop1": 1}
	locationGroupIds := map[string]int{"location_group1": 1}
	hasContinuousPickupDropoff := false

	stopTime, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the stop time was parsed correctly
	if stopTime.TripId != "trip1" {
		t.Errorf("Expected trip_id to be 'trip1', got '%s'", stopTime.TripId)
	}
	if *stopTime.ArrivalTime != "08:00:00" {
		t.Errorf("Expected arrival_time to be '08:00:00', got '%s'", *stopTime.ArrivalTime)
	}
	if *stopTime.DepartureTime != "08:05:00" {
		t.Errorf("Expected departure_time to be '08:05:00', got '%s'", *stopTime.DepartureTime)
	}
	if stopTime.StopId != "stop1" {
		t.Errorf("Expected stop_id to be 'stop1', got '%s'", stopTime.StopId)
	}
	if stopTime.StopSequence != 1 {
		t.Errorf("Expected stop_sequence to be 1, got %d", stopTime.StopSequence)
	}
	if *stopTime.StopHeadsign != "Downtown" {
		t.Errorf("Expected stop_headsign to be 'Downtown', got '%s'", *stopTime.StopHeadsign)
	}
	if *stopTime.PickupType != 0 {
		t.Errorf("Expected pickup_type to be 0, got %d", *stopTime.PickupType)
	}
	if *stopTime.DropOffType != 0 {
		t.Errorf("Expected drop_off_type to be 0, got %d", *stopTime.DropOffType)
	}
	if *stopTime.ContinuousPickup != 0 {
		t.Errorf("Expected continuous_pickup to be 0, got %d", *stopTime.ContinuousPickup)
	}
	if *stopTime.ContinuousDropOff != 0 {
		t.Errorf("Expected continuous_drop_off to be 0, got %d", *stopTime.ContinuousDropOff)
	}
	if *stopTime.ShapeDistTraveled != 1.5 {
		t.Errorf("Expected shape_dist_traveled to be 1.5, got %f", *stopTime.ShapeDistTraveled)
	}
	if *stopTime.Timepoint != true {
		t.Errorf("Expected timepoint to be true, got %v", *stopTime.Timepoint)
	}
	if *stopTime.PickupBookingRuleId != "booking1" {
		t.Errorf("Expected pickup_booking_rule_id to be 'booking1', got '%s'", *stopTime.PickupBookingRuleId)
	}
	if *stopTime.DropOffBookingRuleId != "booking2" {
		t.Errorf("Expected drop_off_booking_rule_id to be 'booking2', got '%s'", *stopTime.DropOffBookingRuleId)
	}
}

func TestParseStopTime_MinimalValidStopTime(t *testing.T) {
	// Test a minimal valid stop time with only required fields
	input := map[string]string{
		"trip_id":       "trip1",
		"stop_id":       "stop1",
		"stop_sequence": "1",
	}

	tripIds := map[string]int{"trip1": 1}
	stopIds := map[string]int{"stop1": 1}
	locationGroupIds := map[string]int{}
	hasContinuousPickupDropoff := false

	stopTime, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the stop time was parsed correctly
	if stopTime.TripId != "trip1" {
		t.Errorf("Expected trip_id to be 'trip1', got '%s'", stopTime.TripId)
	}
	if stopTime.StopId != "stop1" {
		t.Errorf("Expected stop_id to be 'stop1', got '%s'", stopTime.StopId)
	}
	if stopTime.StopSequence != 1 {
		t.Errorf("Expected stop_sequence to be 1, got %d", stopTime.StopSequence)
	}
}

func TestParseStopTime_MissingRequiredFields(t *testing.T) {
	// Test a stop time with missing required fields
	input := map[string]string{
		"stop_headsign": "Downtown",
	}

	tripIds := map[string]int{}
	stopIds := map[string]int{}
	locationGroupIds := map[string]int{}
	hasContinuousPickupDropoff := false

	_, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check for validation messages for missing required fields
	expectedErrors := map[string]bool{
		"Trip ID is required.": false,
		"Stop sequence is required and must be a non-negative integer.":               false,
		"At least one of Stop ID, Location Group ID, or Location ID must be defined.": false,
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

func TestParseStopTime_InvalidTripId(t *testing.T) {
	// Test a stop time with an invalid trip_id
	input := map[string]string{
		"trip_id":       "invalid_trip",
		"stop_id":       "stop1",
		"stop_sequence": "1",
	}

	tripIds := map[string]int{"trip1": 1}
	stopIds := map[string]int{"stop1": 1}
	locationGroupIds := map[string]int{}
	hasContinuousPickupDropoff := false

	_, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check for validation message for invalid trip_id
	found := false
	for _, msg := range messages {
		if msg.Field == "trip_id" && msg.Message == "Trip ID must reference a valid trip_id from trips.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid trip_id not found")
	}
}

func TestParseStopTime_InvalidStopId(t *testing.T) {
	// Test a stop time with an invalid stop_id
	input := map[string]string{
		"trip_id":       "trip1",
		"stop_id":       "invalid_stop",
		"stop_sequence": "1",
	}

	tripIds := map[string]int{"trip1": 1}
	stopIds := map[string]int{"stop1": 1}
	locationGroupIds := map[string]int{}
	hasContinuousPickupDropoff := false

	_, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check for validation message for invalid stop_id
	found := false
	for _, msg := range messages {
		if msg.Field == "stop_id" && msg.Message == "Stop ID must reference a valid stop_id from stops.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid stop_id not found")
	}
}

func TestParseStopTime_InvalidLocationGroupId(t *testing.T) {
	// Test a stop time with an invalid location_group_id
	input := map[string]string{
		"trip_id":           "trip1",
		"location_group_id": "invalid_location_group",
		"stop_sequence":     "1",
	}

	tripIds := map[string]int{"trip1": 1}
	stopIds := map[string]int{}
	locationGroupIds := map[string]int{"location_group1": 1}
	hasContinuousPickupDropoff := false

	_, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check for validation message for invalid location_group_id
	found := false
	for _, msg := range messages {
		if msg.Field == "location_group_id" && msg.Message == "Location Group ID must reference a valid location_group_id from location_groups.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid location_group_id not found")
	}
}

func TestParseStopTime_MutuallyExclusiveFields(t *testing.T) {
	// Test a stop time with mutually exclusive fields
	input := map[string]string{
		"trip_id":           "trip1",
		"stop_id":           "stop1",
		"location_group_id": "location_group1",
		"stop_sequence":     "1",
	}

	tripIds := map[string]int{"trip1": 1}
	stopIds := map[string]int{"stop1": 1}
	locationGroupIds := map[string]int{"location_group1": 1}
	hasContinuousPickupDropoff := false

	_, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check for validation message for mutually exclusive fields
	found := false
	for _, msg := range messages {
		if msg.Field == "stop_id" && msg.Message == "Stop ID and Location Group ID cannot both be defined." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for mutually exclusive fields not found")
	}
}

func TestParseStopTime_TimepointValidation(t *testing.T) {
	// Test a stop time with timepoint=1 but missing arrival_time and departure_time
	input := map[string]string{
		"trip_id":       "trip1",
		"stop_id":       "stop1",
		"stop_sequence": "1",
		"timepoint":     "1",
	}

	tripIds := map[string]int{"trip1": 1}
	stopIds := map[string]int{"stop1": 1}
	locationGroupIds := map[string]int{}
	hasContinuousPickupDropoff := false

	_, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check for validation message for timepoint=1 without arrival_time and departure_time
	expectedErrors := map[string]bool{
		"Arrival time is required when timepoint=1.":   false,
		"Departure time is required when timepoint=1.": false,
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

func TestParseStopTime_PickupDropOffWindowValidation(t *testing.T) {
	// Test a stop time with start_pickup_drop_off_window but missing end_pickup_drop_off_window
	input := map[string]string{
		"trip_id":                      "trip1",
		"stop_id":                      "stop1",
		"stop_sequence":                "1",
		"start_pickup_drop_off_window": "08:00:00",
	}

	tripIds := map[string]int{"trip1": 1}
	stopIds := map[string]int{"stop1": 1}
	locationGroupIds := map[string]int{}
	hasContinuousPickupDropoff := false

	_, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check for validation message for start_pickup_drop_off_window without end_pickup_drop_off_window
	found := false
	for _, msg := range messages {
		if msg.Field == "start_pickup_drop_off_window" && msg.Message == "Start pickup drop off window and end pickup drop off window must be defined together." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for start_pickup_drop_off_window without end_pickup_drop_off_window not found")
	}
}

func TestParseStopTime_LocationGroupIdRequiresPickupDropOffWindow(t *testing.T) {
	// Test a stop time with location_group_id but missing start_pickup_drop_off_window and end_pickup_drop_off_window
	input := map[string]string{
		"trip_id":           "trip1",
		"location_group_id": "location_group1",
		"stop_sequence":     "1",
	}

	tripIds := map[string]int{"trip1": 1}
	stopIds := map[string]int{}
	locationGroupIds := map[string]int{"location_group1": 1}
	hasContinuousPickupDropoff := false

	_, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check for validation message for location_group_id without start_pickup_drop_off_window and end_pickup_drop_off_window
	found := false
	for _, msg := range messages {
		if msg.Field == "start_pickup_drop_off_window" && msg.Message == "Start pickup drop off window and end pickup drop off window are required when location_group_id or location_id is defined." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for location_group_id without start_pickup_drop_off_window and end_pickup_drop_off_window not found")
	}
}

func TestParseStopTime_InvalidPickupType(t *testing.T) {
	// Test a stop time with an invalid pickup_type
	input := map[string]string{
		"trip_id":       "trip1",
		"stop_id":       "stop1",
		"stop_sequence": "1",
		"pickup_type":   "5", // Invalid value
	}

	tripIds := map[string]int{"trip1": 1}
	stopIds := map[string]int{"stop1": 1}
	locationGroupIds := map[string]int{}
	hasContinuousPickupDropoff := false

	_, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check for validation message for invalid pickup_type
	found := false
	for _, msg := range messages {
		if msg.Field == "pickup_type" && msg.Message == "Invalid pickup_type value. Valid values are 0, 1, 2, 3." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid pickup_type not found")
	}
}

func TestParseStopTime_InvalidDropOffType(t *testing.T) {
	// Test a stop time with an invalid drop_off_type
	input := map[string]string{
		"trip_id":       "trip1",
		"stop_id":       "stop1",
		"stop_sequence": "1",
		"drop_off_type": "5", // Invalid value
	}

	tripIds := map[string]int{"trip1": 1}
	stopIds := map[string]int{"stop1": 1}
	locationGroupIds := map[string]int{}
	hasContinuousPickupDropoff := false

	_, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check for validation message for invalid drop_off_type
	found := false
	for _, msg := range messages {
		if msg.Field == "drop_off_type" && msg.Message == "Invalid drop_off_type value. Valid values are 0, 1, 2, 3." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid drop_off_type not found")
	}
}

func TestParseStopTime_InvalidContinuousPickup(t *testing.T) {
	// Test a stop time with an invalid continuous_pickup
	input := map[string]string{
		"trip_id":           "trip1",
		"stop_id":           "stop1",
		"stop_sequence":     "1",
		"continuous_pickup": "5", // Invalid value
	}

	tripIds := map[string]int{"trip1": 1}
	stopIds := map[string]int{"stop1": 1}
	locationGroupIds := map[string]int{}
	hasContinuousPickupDropoff := false

	_, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check for validation message for invalid continuous_pickup
	found := false
	for _, msg := range messages {
		if msg.Field == "continuous_pickup" && msg.Message == "Invalid continuous_pickup value. Valid values are 0, 1, 2, 3." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid continuous_pickup not found")
	}
}

func TestParseStopTime_InvalidContinuousDropOff(t *testing.T) {
	// Test a stop time with an invalid continuous_drop_off
	input := map[string]string{
		"trip_id":             "trip1",
		"stop_id":             "stop1",
		"stop_sequence":       "1",
		"continuous_drop_off": "5", // Invalid value
	}

	tripIds := map[string]int{"trip1": 1}
	stopIds := map[string]int{"stop1": 1}
	locationGroupIds := map[string]int{}
	hasContinuousPickupDropoff := false

	_, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check for validation message for invalid continuous_drop_off
	found := false
	for _, msg := range messages {
		if msg.Field == "continuous_drop_off" && msg.Message == "Invalid continuous_drop_off value. Valid values are 0, 1, 2, 3." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid continuous_drop_off not found")
	}
}

func TestParseStopTime_BookingRuleRecommendation(t *testing.T) {
	// Test a stop time with pickup_type=2 but missing pickup_booking_rule_id
	input := map[string]string{
		"trip_id":       "trip1",
		"stop_id":       "stop1",
		"stop_sequence": "1",
		"pickup_type":   "2",
	}

	tripIds := map[string]int{"trip1": 1}
	stopIds := map[string]int{"stop1": 1}
	locationGroupIds := map[string]int{}
	hasContinuousPickupDropoff := false

	_, messages := parseStopTime(input, tripIds, stopIds, locationGroupIds, hasContinuousPickupDropoff)

	// Check for validation message for pickup_type=2 without pickup_booking_rule_id
	found := false
	for _, msg := range messages {
		if msg.Field == "pickup_booking_rule_id" && msg.Message == "Pickup booking rule ID is recommended when pickup_type=2." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for pickup_type=2 without pickup_booking_rule_id not found")
	}
}

func TestParseStopTimeValidation_Validate(t *testing.T) {
	// Test the Validate method with a valid stop time
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"stop_times": {
				{
					"trip_id":       "trip1",
					"stop_id":       "stop1",
					"stop_sequence": "1",
				},
			},
			"trips": {
				{
					"trip_id": "trip1",
				},
			},
			"stops": {
				{
					"stop_id": "stop1",
				},
			},
		},
		IdMap: map[string]map[string]int{
			"stop_times": {
				"trip1_1": 1,
			},
			"trips": {
				"trip1": 1,
			},
			"stops": {
				"stop1": 1,
			},
		},
	}

	validator := NewParseStopTimeValidation(nil)
	stopTimes, messages := validator.Validate(gtfs)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the stop time was parsed correctly
	if len(stopTimes) != 1 {
		t.Errorf("Expected 1 stop time, got %d", len(stopTimes))
	} else {
		stopTime := stopTimes[0]
		if stopTime.TripId != "trip1" {
			t.Errorf("Expected trip_id to be 'trip1', got '%s'", stopTime.TripId)
		}
		if stopTime.StopId != "stop1" {
			t.Errorf("Expected stop_id to be 'stop1', got '%s'", stopTime.StopId)
		}
		if stopTime.StopSequence != 1 {
			t.Errorf("Expected stop_sequence to be 1, got %d", stopTime.StopSequence)
		}
	}
}

func TestParseStopTimeValidation_ValidateDuplicateStopSequence(t *testing.T) {
	// Test the Validate method with duplicate stop_sequence for the same trip_id
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"stop_times": {
				{
					"trip_id":       "trip1",
					"stop_id":       "stop1",
					"stop_sequence": "1",
				},
				{
					"trip_id":       "trip1",
					"stop_id":       "stop2",
					"stop_sequence": "1", // Duplicate stop_sequence for trip1
				},
			},
			"trips": {
				{
					"trip_id": "trip1",
				},
			},
			"stops": {
				{
					"stop_id": "stop1",
				},
				{
					"stop_id": "stop2",
				},
			},
		},
		IdMap: map[string]map[string]int{
			"stop_times": {
				"trip1_1": 1,
			},
			"trips": {
				"trip1": 1,
			},
			"stops": {
				"stop1": 1,
				"stop2": 2,
			},
		},
	}

	validator := NewParseStopTimeValidation(nil)
	_, messages := validator.Validate(gtfs)

	// Check for validation message for duplicate stop_sequence
	found := false
	for _, msg := range messages {
		if msg.Field == "stop_sequence" && msg.Message == "Duplicate stop_sequence found for trip_id. Each stop_sequence must be unique within a trip." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for duplicate stop_sequence not found")
	}
}

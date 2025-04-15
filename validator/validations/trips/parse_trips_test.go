package trips

import (
	"main/validator/types"
	"testing"
)

func TestParseTrip_ValidTrip(t *testing.T) {
	// Test a valid trip with all fields
	input := map[string]string{
		"trip_id":               "trip1",
		"route_id":              "route1",
		"service_id":            "service1",
		"trip_headsign":         "Downtown",
		"trip_short_name":       "T1",
		"direction_id":          "0",
		"block_id":              "block1",
		"shape_id":              "shape1",
		"wheelchair_accessible": "1",
		"bikes_allowed":         "1",
	}

	routeIds := map[string]int{"route1": 1}
	serviceIds := map[string]int{"service1": 1}
	shapeIds := map[string]int{"shape1": 1}

	trip, messages := parseTrip(input, routeIds, serviceIds, shapeIds, false, false)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the trip was parsed correctly
	if trip.TripId != "trip1" {
		t.Errorf("Expected trip_id to be 'trip1', got '%s'", trip.TripId)
	}
	if trip.RouteId != "route1" {
		t.Errorf("Expected route_id to be 'route1', got '%s'", trip.RouteId)
	}
	if trip.ServiceId != "service1" {
		t.Errorf("Expected service_id to be 'service1', got '%s'", trip.ServiceId)
	}
	if *trip.TripHeadsign != "Downtown" {
		t.Errorf("Expected trip_headsign to be 'Downtown', got '%s'", *trip.TripHeadsign)
	}
	if *trip.TripShortName != "T1" {
		t.Errorf("Expected trip_short_name to be 'T1', got '%s'", *trip.TripShortName)
	}
	if *trip.DirectionId != false {
		t.Errorf("Expected direction_id to be false, got %v", *trip.DirectionId)
	}
	if *trip.BlockId != "block1" {
		t.Errorf("Expected block_id to be 'block1', got '%s'", *trip.BlockId)
	}
	if *trip.ShapeId != "shape1" {
		t.Errorf("Expected shape_id to be 'shape1', got '%s'", *trip.ShapeId)
	}
	if *trip.WheelchairAccessible != "1" {
		t.Errorf("Expected wheelchair_accessible to be '1', got '%s'", *trip.WheelchairAccessible)
	}
	if *trip.BikesAllowed != 1 {
		t.Errorf("Expected bikes_allowed to be 1, got %d", *trip.BikesAllowed)
	}
}

func TestParseTrip_MinimalValidTrip(t *testing.T) {
	// Test a minimal valid trip with only required fields
	input := map[string]string{
		"trip_id":    "trip1",
		"route_id":   "route1",
		"service_id": "service1",
	}

	routeIds := map[string]int{"route1": 1}
	serviceIds := map[string]int{"service1": 1}
	shapeIds := map[string]int{"shape1": 1}

	trip, messages := parseTrip(input, routeIds, serviceIds, shapeIds, false, false)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the trip was parsed correctly
	if trip.TripId != "trip1" {
		t.Errorf("Expected trip_id to be 'trip1', got '%s'", trip.TripId)
	}
	if trip.RouteId != "route1" {
		t.Errorf("Expected route_id to be 'route1', got '%s'", trip.RouteId)
	}
	if trip.ServiceId != "service1" {
		t.Errorf("Expected service_id to be 'service1', got '%s'", trip.ServiceId)
	}
}

func TestParseTrip_MissingRequiredFields(t *testing.T) {
	// Test a trip with missing required fields
	input := map[string]string{
		"trip_headsign": "Downtown",
	}

	routeIds := map[string]int{}
	serviceIds := map[string]int{}
	shapeIds := map[string]int{}

	_, messages := parseTrip(input, routeIds, serviceIds, shapeIds, false, false)

	// Check for validation messages for missing required fields
	expectedErrors := map[string]bool{
		"Trip ID is required and must be unique.": false,
		"Route ID is required.":                   false,
		"Service ID is required.":                 false,
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

func TestParseTrip_InvalidRouteId(t *testing.T) {
	// Test a trip with an invalid route_id
	input := map[string]string{
		"trip_id":    "trip1",
		"route_id":   "invalid_route",
		"service_id": "service1",
	}

	routeIds := map[string]int{"route1": 1}
	serviceIds := map[string]int{"service1": 1}
	shapeIds := map[string]int{}

	_, messages := parseTrip(input, routeIds, serviceIds, shapeIds, false, false)

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

func TestParseTrip_InvalidServiceId(t *testing.T) {
	// Test a trip with an invalid service_id
	input := map[string]string{
		"trip_id":    "trip1",
		"route_id":   "route1",
		"service_id": "invalid_service",
	}

	routeIds := map[string]int{"route1": 1}
	serviceIds := map[string]int{"service1": 1}
	shapeIds := map[string]int{}

	_, messages := parseTrip(input, routeIds, serviceIds, shapeIds, false, false)

	// Check for validation message for invalid service_id
	found := false
	for _, msg := range messages {
		if msg.Field == "service_id" && msg.Message == "Service ID "+input["service_id"]+" must reference a valid service_id from calendar.txt or calendar_dates.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid service_id not found")
	}
}

func TestParseTrip_InvalidShapeId(t *testing.T) {
	// Test a trip with an invalid shape_id
	input := map[string]string{
		"trip_id":    "trip1",
		"route_id":   "route1",
		"service_id": "service1",
		"shape_id":   "invalid_shape",
	}

	routeIds := map[string]int{"route1": 1}
	serviceIds := map[string]int{"service1": 1}
	shapeIds := map[string]int{"shape1": 1}

	_, messages := parseTrip(input, routeIds, serviceIds, shapeIds, false, false)

	// Check for validation message for invalid shape_id
	found := false
	for _, msg := range messages {
		if msg.Field == "shape_id" && msg.Message == "Shape ID must reference a valid shape_id from shapes.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid shape_id not found")
	}
}

func TestParseTrip_MissingShapeIdWithContinuousPickupDropoff(t *testing.T) {
	// Test a trip with missing shape_id when continuous pickup/dropoff behavior is defined
	input := map[string]string{
		"trip_id":    "trip1",
		"route_id":   "route1",
		"service_id": "service1",
	}

	routeIds := map[string]int{"route1": 1}
	serviceIds := map[string]int{"service1": 1}
	shapeIds := map[string]int{}

	_, messages := parseTrip(input, routeIds, serviceIds, shapeIds, true, false)

	// Check for validation message for missing shape_id with continuous pickup/dropoff
	found := false
	for _, msg := range messages {
		if msg.Field == "shape_id" && msg.Message == "Shape ID is required when continuous pickup or drop-off behavior is defined in routes.txt or stop_times.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for missing shape_id with continuous pickup/dropoff not found")
	}
}

func TestParseTrip_MissingShapeIdWithStopTimesContinuousPickupDropoff(t *testing.T) {
	// Test a trip with missing shape_id when stop_times have continuous pickup/dropoff behavior
	input := map[string]string{
		"trip_id":    "trip1",
		"route_id":   "route1",
		"service_id": "service1",
	}

	routeIds := map[string]int{"route1": 1}
	serviceIds := map[string]int{"service1": 1}
	shapeIds := map[string]int{}

	_, messages := parseTrip(input, routeIds, serviceIds, shapeIds, false, true)

	// Check for validation message for missing shape_id with stop_times continuous pickup/dropoff
	found := false
	for _, msg := range messages {
		if msg.Field == "shape_id" && msg.Message == "Shape ID is required when continuous pickup or drop-off behavior is defined in routes.txt or stop_times.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for missing shape_id with stop_times continuous pickup/dropoff not found")
	}
}

func TestParseTrip_InvalidBikesAllowed(t *testing.T) {
	// Test a trip with an invalid bikes_allowed value
	input := map[string]string{
		"trip_id":       "trip1",
		"route_id":      "route1",
		"service_id":    "service1",
		"bikes_allowed": "5", // Invalid value
	}

	routeIds := map[string]int{"route1": 1}
	serviceIds := map[string]int{"service1": 1}
	shapeIds := map[string]int{}

	_, messages := parseTrip(input, routeIds, serviceIds, shapeIds, false, false)

	// Check for validation message for invalid bikes_allowed
	found := false
	for _, msg := range messages {
		if msg.Field == "bikes_allowed" && msg.Message == "Invalid bikes_allowed value. Valid values are 0, 1, 2." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid bikes_allowed not found")
	}
}

func TestParseTrip_InvalidWheelchairAccessible(t *testing.T) {
	// Test a trip with an invalid wheelchair_accessible value
	input := map[string]string{
		"trip_id":               "trip1",
		"route_id":              "route1",
		"service_id":            "service1",
		"wheelchair_accessible": "5", // Invalid value
	}

	routeIds := map[string]int{"route1": 1}
	serviceIds := map[string]int{"service1": 1}
	shapeIds := map[string]int{}

	_, messages := parseTrip(input, routeIds, serviceIds, shapeIds, false, false)

	// Check for validation message for invalid wheelchair_accessible
	found := false
	for _, msg := range messages {
		if msg.Field == "wheelchair_accessible" && msg.Message == "Invalid wheelchair_accessible value. Valid values are 0, 1, 2." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid wheelchair_accessible not found")
	}
}

func TestParseTripValidation_Validate(t *testing.T) {
	// Test the Validate method with a valid trip
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"trips": {
				{
					"trip_id":    "trip1",
					"route_id":   "route1",
					"service_id": "service1",
				},
			},
			"routes": {
				{
					"route_id": "route1",
				},
			},
			"calendar": {
				{
					"service_id": "service1",
				},
			},
		},
		IdMap: map[string]map[string]int{
			"trips": {
				"trip1": 1,
			},
			"routes": {
				"route1": 1,
			},
			"calendar": {
				"service1": 1,
			},
		},
	}

	validator := NewParseTripValidation(nil)
	trips, messages := validator.Validate(gtfs)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the trip was parsed correctly
	if len(trips) != 1 {
		t.Errorf("Expected 1 trip, got %d", len(trips))
	} else {
		trip := trips[0]
		if trip.TripId != "trip1" {
			t.Errorf("Expected trip_id to be 'trip1', got '%s'", trip.TripId)
		}
		if trip.RouteId != "route1" {
			t.Errorf("Expected route_id to be 'route1', got '%s'", trip.RouteId)
		}
		if trip.ServiceId != "service1" {
			t.Errorf("Expected service_id to be 'service1', got '%s'", trip.ServiceId)
		}
	}
}

func TestParseTripValidation_ValidateDuplicateTripId(t *testing.T) {
	// Test the Validate method with duplicate trip IDs
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"trips": {
				{
					"trip_id":    "trip1",
					"route_id":   "route1",
					"service_id": "service1",
				},
				{
					"trip_id":    "trip1", // Duplicate trip_id
					"route_id":   "route2",
					"service_id": "service2",
				},
			},
			"routes": {
				{
					"route_id": "route1",
				},
				{
					"route_id": "route2",
				},
			},
			"calendar": {
				{
					"service_id": "service1",
				},
				{
					"service_id": "service2",
				},
			},
		},
	}

	validator := NewParseTripValidation(nil)
	_, messages := validator.Validate(gtfs)

	// Check for validation message for duplicate trip_id
	found := false
	for _, msg := range messages {
		if msg.Field == "trip_id" && msg.Message == "Duplicate trip_id found. Trip IDs must be unique." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for duplicate trip_id not found")
	}
}

func TestParseTripValidation_ValidateInvalidRouteId(t *testing.T) {
	// Test the Validate method with an invalid route_id
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"trips": {
				{
					"trip_id":    "trip1",
					"route_id":   "invalid_route",
					"service_id": "service1",
				},
			},
			"routes": {
				{
					"route_id": "route1",
				},
			},
			"calendar": {
				{
					"service_id": "service1",
				},
			},
		},
	}

	validator := NewParseTripValidation(nil)
	_, messages := validator.Validate(gtfs)

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

func TestParseTripValidation_ValidateInvalidServiceId(t *testing.T) {
	// Test the Validate method with an invalid service_id
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"trips": {
				{
					"trip_id":    "trip1",
					"route_id":   "route1",
					"service_id": "invalid_service",
				},
			},
			"routes": {
				{
					"route_id": "route1",
				},
			},
			"calendar": {
				{
					"service_id": "service1",
				},
			},
		},
	}

	validator := NewParseTripValidation(nil)
	_, messages := validator.Validate(gtfs)

	// Check for validation message for invalid service_id
	found := false
	for _, msg := range messages {
		if msg.Field == "service_id" && msg.Message == "Service ID "+gtfs.Files["trips"][0]["service_id"]+" must reference a valid service_id from calendar.txt or calendar_dates.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid service_id not found")
	}
}

func TestParseTripValidation_ValidateMissingShapeIdWithContinuousPickupDropoff(t *testing.T) {
	// Test the Validate method with missing shape_id when continuous pickup/dropoff behavior is defined
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"trips": {
				{
					"trip_id":    "trip1",
					"route_id":   "route1",
					"service_id": "service1",
				},
			},
			"routes": {
				{
					"route_id":            "route1",
					"continuous_pickup":   "1",
					"continuous_drop_off": "1",
				},
			},
			"calendar": {
				{
					"service_id": "service1",
				},
			},
		},
	}

	validator := NewParseTripValidation(nil)
	_, messages := validator.Validate(gtfs)

	// Check for validation message for missing shape_id with continuous pickup/dropoff
	found := false
	for _, msg := range messages {
		if msg.Field == "shape_id" && msg.Message == "Shape ID is required when continuous pickup or drop-off behavior is defined in routes.txt or stop_times.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for missing shape_id with continuous pickup/dropoff not found")
	}
}

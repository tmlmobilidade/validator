package routes

import (
	"main/validator/types"
	"testing"
)

func TestParseRoute_ValidRoute(t *testing.T) {
	// Test a valid route with all fields
	input := map[string]string{
		"route_id":            "route1",
		"agency_id":           "agency1",
		"route_short_name":    "R1",
		"route_long_name":     "Route 1",
		"route_desc":          "Test route description",
		"route_type":          "3",
		"route_url":           "http://example.com/route1",
		"route_color":         "FF0000",
		"route_text_color":    "FFFFFF",
		"route_sort_order":    "1",
		"continuous_pickup":   "1",
		"continuous_drop_off": "1",
	}

	route, messages := parseRoute(input, false, false)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the route was parsed correctly
	if route.RouteId != "route1" {
		t.Errorf("Expected route_id to be 'route1', got '%s'", route.RouteId)
	}
	if *route.AgencyId != "agency1" {
		t.Errorf("Expected agency_id to be 'agency1', got '%s'", *route.AgencyId)
	}
	if *route.RouteShortName != "R1" {
		t.Errorf("Expected route_short_name to be 'R1', got '%s'", *route.RouteShortName)
	}
	if *route.RouteLongName != "Route 1" {
		t.Errorf("Expected route_long_name to be 'Route 1', got '%s'", *route.RouteLongName)
	}
	if *route.RouteDesc != "Test route description" {
		t.Errorf("Expected route_desc to be 'Test route description', got '%s'", *route.RouteDesc)
	}
	if route.RouteType != 3 {
		t.Errorf("Expected route_type to be 3, got %d", route.RouteType)
	}
	if *route.RouteUrl != "http://example.com/route1" {
		t.Errorf("Expected route_url to be 'http://example.com/route1', got '%s'", *route.RouteUrl)
	}
	if *route.RouteColor != "FF0000" {
		t.Errorf("Expected route_color to be 'FF0000', got '%s'", *route.RouteColor)
	}
	if *route.RouteTextColor != "FFFFFF" {
		t.Errorf("Expected route_text_color to be 'FFFFFF', got '%s'", *route.RouteTextColor)
	}
	if *route.RouteSortOrder != 1 {
		t.Errorf("Expected route_sort_order to be 1, got %d", *route.RouteSortOrder)
	}
	if *route.ContinuousPickup != "1" {
		t.Errorf("Expected continuous_pickup to be '1', got '%s'", *route.ContinuousPickup)
	}
	if *route.ContinuousDropOff != "1" {
		t.Errorf("Expected continuous_drop_off to be '1', got '%s'", *route.ContinuousDropOff)
	}
}

func TestParseRoute_MinimalValidRoute(t *testing.T) {
	// Test a minimal valid route with only required fields
	input := map[string]string{
		"route_id":         "route1",
		"route_type":       "3",
		"route_short_name": "R1",
	}

	route, messages := parseRoute(input, false, false)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the route was parsed correctly
	if route.RouteId != "route1" {
		t.Errorf("Expected route_id to be 'route1', got '%s'", route.RouteId)
	}
	if route.RouteType != 3 {
		t.Errorf("Expected route_type to be 3, got %d", route.RouteType)
	}
	if *route.RouteShortName != "R1" {
		t.Errorf("Expected route_short_name to be 'R1', got '%s'", *route.RouteShortName)
	}
}

func TestParseRoute_MissingRequiredFields(t *testing.T) {
	// Test a route with missing required fields
	input := map[string]string{
		"route_short_name": "R1",
	}

	_, messages := parseRoute(input, false, false)

	// Check for validation messages for missing required fields
	expectedErrors := map[string]bool{
		"Route ID is required and must be unique.": false,
		"Route type is required.":                  false,
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

func TestParseRoute_InvalidRouteType(t *testing.T) {
	// Test a route with an invalid route_type
	input := map[string]string{
		"route_id":         "route1",
		"route_type":       "99", // Invalid route type
		"route_short_name": "R1",
	}

	_, messages := parseRoute(input, false, false)

	// Check for validation message for invalid route_type
	found := false
	for _, msg := range messages {
		if msg.Field == "route_type" && msg.Message == "Invalid route_type. Valid values are 0, 1, 2, 3, 4, 5, 6, 7, 11, 12." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid route_type not found")
	}
}

func TestParseRoute_MissingAgencyIdWithMultipleAgencies(t *testing.T) {
	// Test a route with missing agency_id when multiple agencies exist
	input := map[string]string{
		"route_id":         "route1",
		"route_type":       "3",
		"route_short_name": "R1",
	}

	_, messages := parseRoute(input, true, false)

	// Check for validation message for missing agency_id
	found := false
	for _, msg := range messages {
		if msg.Field == "agency_id" && msg.Message == "Agency ID is required when multiple agencies are defined in agency.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for missing agency_id not found")
	}
}

func TestParseRoute_MissingNameFields(t *testing.T) {
	// Test a route with missing both route_short_name and route_long_name
	input := map[string]string{
		"route_id":   "route1",
		"route_type": "3",
	}

	_, messages := parseRoute(input, false, false)

	// Check for validation message for missing name fields
	found := false
	for _, msg := range messages {
		if msg.Field == "route_short_name/route_long_name" && msg.Message == "At least one of route_short_name or route_long_name must be provided." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for missing name fields not found")
	}
}

func TestParseRoute_ShortNameTooLong(t *testing.T) {
	// Test a route with a route_short_name that's too long
	input := map[string]string{
		"route_id":         "route1",
		"route_type":       "3",
		"route_short_name": "ThisIsTooLong123",
	}

	_, messages := parseRoute(input, false, false)

	// Check for validation message for short name too long
	found := false
	for _, msg := range messages {
		if msg.Field == "route_short_name" && msg.Message == "Route short name should be no longer than 12 characters." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for short name too long not found")
	}
}

func TestParseRoute_InvalidUrl(t *testing.T) {
	// Test a route with an invalid URL
	input := map[string]string{
		"route_id":         "route1",
		"route_type":       "3",
		"route_short_name": "R1",
		"route_url":        "not-a-url",
	}

	_, messages := parseRoute(input, false, false)

	// Check for validation message for invalid URL
	found := false
	for _, msg := range messages {
		if msg.Field == "route_url" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid URL not found")
	}
}

func TestParseRoute_InvalidContinuousPickup(t *testing.T) {
	// Test a route with an invalid continuous_pickup value
	input := map[string]string{
		"route_id":          "route1",
		"route_type":        "3",
		"route_short_name":  "R1",
		"continuous_pickup": "5", // Invalid value
	}

	_, messages := parseRoute(input, false, false)

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

func TestParseRoute_ContinuousPickupWithWindows(t *testing.T) {
	// Test a route with continuous_pickup when pickup/dropoff windows are defined
	input := map[string]string{
		"route_id":          "route1",
		"route_type":        "3",
		"route_short_name":  "R1",
		"continuous_pickup": "1",
	}

	_, messages := parseRoute(input, false, true)

	// Check for validation message for continuous_pickup with windows
	found := false
	for _, msg := range messages {
		if msg.Field == "continuous_pickup" && msg.Message == "continuous_pickup is forbidden when stop_times.start_pickup_drop_off_window or stop_times.end_pickup_drop_off_window are defined for any trip of this route." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for continuous_pickup with windows not found")
	}
}

func TestParseRoute_InvalidContinuousDropOff(t *testing.T) {
	// Test a route with an invalid continuous_drop_off value
	input := map[string]string{
		"route_id":            "route1",
		"route_type":          "3",
		"route_short_name":    "R1",
		"continuous_drop_off": "5", // Invalid value
	}

	_, messages := parseRoute(input, false, false)

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

func TestParseRoute_ContinuousDropOffWithWindows(t *testing.T) {
	// Test a route with continuous_drop_off when pickup/dropoff windows are defined
	input := map[string]string{
		"route_id":            "route1",
		"route_type":          "3",
		"route_short_name":    "R1",
		"continuous_drop_off": "1",
	}

	_, messages := parseRoute(input, false, true)

	// Check for validation message for continuous_drop_off with windows
	found := false
	for _, msg := range messages {
		if msg.Field == "continuous_drop_off" && msg.Message == "continuous_drop_off is forbidden when stop_times.start_pickup_drop_off_window or stop_times.end_pickup_drop_off_window are defined for any trip of this route." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for continuous_drop_off with windows not found")
	}
}

func TestParseRoute_NegativeSortOrder(t *testing.T) {
	// Test a route with a negative route_sort_order
	input := map[string]string{
		"route_id":         "route1",
		"route_type":       "3",
		"route_short_name": "R1",
		"route_sort_order": "-1",
	}

	_, messages := parseRoute(input, false, false)

	// Check for validation message for negative route_sort_order
	found := false
	for _, msg := range messages {
		if msg.Field == "route_sort_order" && msg.Message == "Route sort order must be a non-negative integer." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for negative route_sort_order not found")
	}
}

func TestParseRouteValidation_Validate(t *testing.T) {
	// Test the Validate method with a valid route
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"routes": {
				{
					"route_id":         "route1",
					"route_type":       "3",
					"route_short_name": "R1",
				},
			},
			"agency": {
				{
					"agency_id":   "agency1",
					"agency_name": "Test Agency",
				},
			},
		},
		FieldCounter: map[string]map[string]int{
			"stop_times": {
				"start_pickup_drop_off_window": 0,
				"end_pickup_drop_off_window":   0,
			},
		},
	}

	validator := NewParseRouteValidation(nil)
	routes, messages := validator.Validate(gtfs)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the route was parsed correctly
	if len(routes) != 1 {
		t.Errorf("Expected 1 route, got %d", len(routes))
	} else {
		route := routes[0]
		if route.RouteId != "route1" {
			t.Errorf("Expected route_id to be 'route1', got '%s'", route.RouteId)
		}
		if route.RouteType != 3 {
			t.Errorf("Expected route_type to be 3, got %d", route.RouteType)
		}
		if *route.RouteShortName != "R1" {
			t.Errorf("Expected route_short_name to be 'R1', got '%s'", *route.RouteShortName)
		}
	}
}

func TestParseRouteValidation_ValidateDuplicateRouteId(t *testing.T) {
	// Test the Validate method with duplicate route IDs
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"routes": {
				{
					"route_id":         "route1",
					"route_type":       "3",
					"route_short_name": "R1",
				},
				{
					"route_id":         "route1", // Duplicate route_id
					"route_type":       "3",
					"route_short_name": "R2",
				},
			},
			"agency": {
				{
					"agency_id":   "agency1",
					"agency_name": "Test Agency",
				},
			},
		},
		FieldCounter: map[string]map[string]int{
			"stop_times": {
				"start_pickup_drop_off_window": 0,
				"end_pickup_drop_off_window":   0,
			},
		},
	}

	validator := NewParseRouteValidation(nil)
	_, messages := validator.Validate(gtfs)

	// Check for validation message for duplicate route_id
	found := false
	for _, msg := range messages {
		if msg.Field == "route_id" && msg.Message == "Duplicate route_id found. Route IDs must be unique." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for duplicate route_id not found")
	}
}

func TestParseRouteValidation_ValidateMultipleAgencies(t *testing.T) {
	// Test the Validate method with multiple agencies
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"routes": {
				{
					"route_id":         "route1",
					"route_type":       "3",
					"route_short_name": "R1",
					// Missing agency_id
				},
			},
			"agency": {
				{
					"agency_id":   "agency1",
					"agency_name": "Test Agency 1",
				},
				{
					"agency_id":   "agency2",
					"agency_name": "Test Agency 2",
				},
			},
		},
		FieldCounter: map[string]map[string]int{
			"stop_times": {
				"start_pickup_drop_off_window": 0,
				"end_pickup_drop_off_window":   0,
			},
		},
	}

	validator := NewParseRouteValidation(nil)
	_, messages := validator.Validate(gtfs)

	// Check for validation message for missing agency_id with multiple agencies
	found := false
	for _, msg := range messages {
		if msg.Field == "agency_id" && msg.Message == "Agency ID is required when multiple agencies are defined in agency.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for missing agency_id with multiple agencies not found")
	}
}

func TestParseRouteValidation_ValidatePickupDropoffWindows(t *testing.T) {
	// Test the Validate method with pickup/dropoff windows
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"routes": {
				{
					"route_id":            "route1",
					"route_type":          "3",
					"route_short_name":    "R1",
					"continuous_pickup":   "1",
					"continuous_drop_off": "1",
				},
			},
			"agency": {
				{
					"agency_id":   "agency1",
					"agency_name": "Test Agency",
				},
			},
		},
		FieldCounter: map[string]map[string]int{
			"stop_times": {
				"start_pickup_drop_off_window": 1, // Has pickup/dropoff windows
				"end_pickup_drop_off_window":   0,
			},
		},
	}

	validator := NewParseRouteValidation(nil)
	_, messages := validator.Validate(gtfs)

	// Check for validation messages for continuous_pickup and continuous_drop_off with windows
	foundPickup := false
	foundDropoff := false
	for _, msg := range messages {
		if msg.Field == "continuous_pickup" && msg.Message == "continuous_pickup is forbidden when stop_times.start_pickup_drop_off_window or stop_times.end_pickup_drop_off_window are defined for any trip of this route." {
			foundPickup = true
		}
		if msg.Field == "continuous_drop_off" && msg.Message == "continuous_drop_off is forbidden when stop_times.start_pickup_drop_off_window or stop_times.end_pickup_drop_off_window are defined for any trip of this route." {
			foundDropoff = true
		}
	}

	if !foundPickup {
		t.Error("Expected validation message for continuous_pickup with windows not found")
	}
	if !foundDropoff {
		t.Error("Expected validation message for continuous_drop_off with windows not found")
	}
}

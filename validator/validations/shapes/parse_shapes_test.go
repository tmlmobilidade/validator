package shapes

import (
	"fmt"
	"main/types"
	"strconv"
	"testing"
)

func TestParseShape_ValidShape(t *testing.T) {
	// Test a valid shape with all fields
	input := map[string]string{
		"shape_id":            "shape1",
		"shape_pt_lat":        "37.61956",
		"shape_pt_lon":        "-122.48161",
		"shape_pt_sequence":   "0",
		"shape_dist_traveled": "0",
	}

	shape, messages := parseShape(input)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the shape was parsed correctly
	if shape.ShapeId != "shape1" {
		t.Errorf("Expected shape_id to be 'shape1', got '%s'", shape.ShapeId)
	}
	if *shape.ShapePtLat != 37.61956 {
		t.Errorf("Expected shape_pt_lat to be 37.61956, got '%f'", *shape.ShapePtLat)
	}
	if *shape.ShapePtLon != -122.48161 {
		t.Errorf("Expected shape_pt_lon to be -122.48161, got '%f'", *shape.ShapePtLon)
	}
	if *shape.ShapePtSequence != 0 {
		t.Errorf("Expected shape_pt_sequence to be 0, got '%d'", *shape.ShapePtSequence)
	}
	if *shape.ShapeDistTraveled != 0 {
		t.Errorf("Expected shape_dist_traveled to be 0, got '%f'", *shape.ShapeDistTraveled)
	}
}

func TestParseShape_MissingRequiredFields(t *testing.T) {
	// Test a shape with missing required fields
	input := map[string]string{}

	_, messages := parseShape(input)

	// Check for validation messages for missing required fields
	expectedErrors := map[string]bool{
		"Shape ID is required.":                              false,
		"Shape point latitude is required.":                  false,
		"Shape point longitude is required.":                 false,
		"Shape point sequence must be non-negative integer.": false,
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

func TestParseShape_InvalidLatitude(t *testing.T) {
	// Test a shape with invalid latitude
	input := map[string]string{
		"shape_id":            "shape1",
		"shape_pt_lat":        "91.0", // Invalid latitude (> 90)
		"shape_pt_lon":        "-122.48161",
		"shape_pt_sequence":   "0",
		"shape_dist_traveled": "0",
	}

	_, messages := parseShape(input)

	// Check for validation message for invalid latitude
	found := false
	for _, msg := range messages {
		lat, _ := strconv.ParseFloat(input["shape_pt_lat"], 32)
		if msg.Field == "shape_pt_lat" && msg.Message == fmt.Sprintf("Invalid latitude, expected range: -90 to 90, got: %f", lat) {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid latitude not found")
	}
}

func TestParseShape_InvalidLongitude(t *testing.T) {
	// Test a shape with invalid longitude
	input := map[string]string{
		"shape_id":            "shape1",
		"shape_pt_lat":        "37.61956",
		"shape_pt_lon":        "-181.0", // Invalid longitude (< -180)
		"shape_pt_sequence":   "0",
		"shape_dist_traveled": "0",
	}

	_, messages := parseShape(input)

	// Check for validation message for invalid longitude
	found := false
	for _, msg := range messages {
		lon, _ := strconv.ParseFloat(input["shape_pt_lon"], 32)
		if msg.Field == "shape_pt_lon" && msg.Message == fmt.Sprintf("Invalid longitude, expected range: -180 to 180, got: %f", lon) {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid longitude not found")
	}
}

func TestParseShape_InvalidSequence(t *testing.T) {
	// Test a shape with invalid sequence number
	input := map[string]string{
		"shape_id":            "shape1",
		"shape_pt_lat":        "37.61956",
		"shape_pt_lon":        "-122.48161",
		"shape_pt_sequence":   "-1", // Invalid sequence (negative)
		"shape_dist_traveled": "0",
	}

	_, messages := parseShape(input)

	// Check for validation message for invalid sequence
	found := false
	for _, msg := range messages {
		if msg.Field == "shape_pt_sequence" && msg.Message == "Shape point sequence must be non-negative integer." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid sequence not found")
	}
}

func TestParseShape_InvalidDistanceTraveled(t *testing.T) {
	// Test a shape with invalid distance traveled
	input := map[string]string{
		"shape_id":            "shape1",
		"shape_pt_lat":        "37.61956",
		"shape_pt_lon":        "-122.48161",
		"shape_pt_sequence":   "0",
		"shape_dist_traveled": "-1.0", // Invalid distance (negative)
	}

	_, messages := parseShape(input)

	// Check for validation message for invalid distance
	found := false
	for _, msg := range messages {
		if msg.Field == "shape_dist_traveled" && msg.Message == "Shape distance traveled must be non-negative float." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid distance not found")
	}
}

func TestParseShapeValidation_Validate(t *testing.T) {
	// Test the Validate method with valid shapes
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"shapes": {
				{
					"shape_id":            "shape1",
					"shape_pt_lat":        "37.61956",
					"shape_pt_lon":        "-122.48161",
					"shape_pt_sequence":   "0",
					"shape_dist_traveled": "0",
				},
				{
					"shape_id":            "shape1",
					"shape_pt_lat":        "37.64430",
					"shape_pt_lon":        "-122.41070",
					"shape_pt_sequence":   "6",
					"shape_dist_traveled": "6.8310",
				},
			},
		},
	}

	validator := NewParseShapeValidation(nil)
	shapes, messages := validator.Validate(gtfs)

	// Check that shapes were parsed correctly
	if len(shapes) != 2 {
		t.Errorf("Expected 2 shapes, got %d", len(shapes))
	}

	// Check for duplicate shape_id message
	found := false
	for _, msg := range messages {
		if msg.Field == "shape_id" && msg.Message == "Duplicate shape_id found. Shape IDs must be unique." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for duplicate shape_id not found")
	}
}

package stops

import (
	"main/src/types"
	"testing"
)

func TestParseStopValidation_SingleValidStop(t *testing.T) {
	gtfsData := types.Gtfs{
		Files: map[string][]map[string]string{
			"stops": {
				{
					"stop_id":   "stop1",
					"stop_name": "Test Stop",
					"stop_lat":  "40.7128",
					"stop_lon":  "-74.0060",
				},
			},
		},
	}

	validator := NewParseStopValidation(nil)
	_, messages := validator.Validate(gtfsData)

	if len(messages) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Error message: %s", msg.Message)
		}
	}
}

func TestParseStopValidation_DuplicateStopID(t *testing.T) {
	gtfsData := types.Gtfs{
		Files: map[string][]map[string]string{
			"stops": {
				{
					"stop_id":   "stop1",
					"stop_name": "Test Stop 1",
					"stop_lat":  "40.7128",
					"stop_lon":  "-74.0060",
				},
				{
					"stop_id":   "stop1",
					"stop_name": "Test Stop 2",
					"stop_lat":  "40.7129",
					"stop_lon":  "-74.0061",
				},
			},
		},
	}

	validator := NewParseStopValidation(nil)
	_, messages := validator.Validate(gtfsData)

	if len(messages) != 1 {
		t.Errorf("Expected 1 error for duplicate stop_id, got %d", len(messages))
	}

	if len(messages) > 0 && messages[0].Field != "stop_id" {
		t.Errorf("Expected error for field 'stop_id', got %s", messages[0].Field)
	}
}

func TestParseStop_CompleteValidStop(t *testing.T) {
	input := map[string]string{
		"stop_id":             "stop1",
		"stop_code":           "S1",
		"stop_name":           "Test Stop",
		"stop_desc":           "A test stop description",
		"stop_lat":            "40.7128",
		"stop_lon":            "-74.0060",
		"zone_id":             "zone1",
		"stop_url":            "http://example.com/stop",
		"location_type":       "0",
		"parent_station":      "station1",
		"stop_timezone":       "America/New_York",
		"wheelchair_boarding": "1",
		"level_id":            "level1",
		"platform_code":       "A",
	}

	_, messages := parseStop(input)
	if len(messages) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}
}

func TestParseStop_MinimalValidStop(t *testing.T) {
	input := map[string]string{
		"stop_id": "stop1",
	}

	_, messages := parseStop(input)
	if len(messages) != 0 {
		t.Errorf("Expected 0 errors for minimal valid stop, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}
}

func TestParseStop_InvalidCoordinates(t *testing.T) {
	input := map[string]string{
		"stop_id":  "stop1",
		"stop_lat": "invalid",
		"stop_lon": "invalid",
	}

	_, messages := parseStop(input)
	if len(messages) != 2 {
		t.Errorf("Expected 2 errors for invalid coordinates, got %d", len(messages))
	}
}

func TestParseStop_InvalidEnumValues(t *testing.T) {
	input := map[string]string{
		"stop_id":             "stop1",
		"location_type":       "5", // Invalid location type (valid: 0-4)
		"wheelchair_boarding": "3", // Invalid wheelchair boarding (valid: 0-2)
	}

	_, messages := parseStop(input)
	if len(messages) != 2 {
		t.Errorf("Expected 2 errors for invalid enum values, got %d", len(messages))
	}
}

func TestParseStop_InvalidURL(t *testing.T) {
	input := map[string]string{
		"stop_id":  "stop1",
		"stop_url": "not-a-url",
	}

	_, messages := parseStop(input)
	if len(messages) != 1 {
		t.Errorf("Expected 1 error for invalid URL, got %d", len(messages))
	}
}

package agency

import (
	"main/src/types"
	"testing"
)

// --- ParseAgencyValidation Tests ---

func TestParseAgencyValidation_SingleValidAgency(t *testing.T) {
	gtfsData := types.Gtfs{
		"agency": []map[string]string{
			{
				"agency_name":     "Test Agency",
				"agency_url":      "http://www.testagency.com",
				"agency_timezone": "America/New_York",
			},
		},
	}

	validator := NewParseAgencyValidation(nil)
	messages := validator.Validate(gtfsData)

	if len(messages) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Error message: %s", msg.Message)
		}
	}
}

func TestParseAgencyValidation_MultipleValidAgenciesWithIDs(t *testing.T) {
	gtfsData := types.Gtfs{
		"agency": []map[string]string{
			{
				"agency_id":       "agency1",
				"agency_name":     "Test Agency 1",
				"agency_url":      "http://www.testagency1.com",
				"agency_timezone": "America/New_York",
			},
			{
				"agency_id":       "agency2",
				"agency_name":     "Test Agency 2",
				"agency_url":      "http://www.testagency2.com",
				"agency_timezone": "America/Los_Angeles",
				"agency_phone":    "+351210410400",
			},
		},
	}

	validator := NewParseAgencyValidation(nil)
	messages := validator.Validate(gtfsData)

	if len(messages) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Error message: %s", msg.Message)
		}
	}
}

func TestParseAgencyValidation_MissingRequiredFields(t *testing.T) {
	gtfsData := types.Gtfs{
		"agency": []map[string]string{
			{
				"agency_id": "agency1",
			},
		},
	}

	validator := NewParseAgencyValidation(nil)
	messages := validator.Validate(gtfsData)

	if len(messages) != 3 {
		t.Errorf("Expected 3 errors, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Error message: %s", msg.Message)
		}
	}
}

func TestParseAgencyValidation_InvalidOptionalFields(t *testing.T) {
	gtfsData := types.Gtfs{
		"agency": []map[string]string{
			{
				"agency_name":     "Test Agency",
				"agency_url":      "http://www.testagency.com",
				"agency_timezone": "America/New_York",
				"agency_email":    "invalid-email",
				"agency_phone":    "123",
				"agency_lang":     "invalid-lang",
				"agency_fare_url": "invalid-url",
			},
		},
	}

	validator := NewParseAgencyValidation(nil)
	messages := validator.Validate(gtfsData)

	if len(messages) != 4 {
		t.Errorf("Expected 4 errors, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Error message: %s", msg.Message)
		}
	}
}

func TestParseAgencyValidation_MultipleAgenciesWithoutID(t *testing.T) {
	gtfsData := types.Gtfs{
		"agency": []map[string]string{
			{
				"agency_name":     "Test Agency 1",
				"agency_url":      "http://www.testagency1.com",
				"agency_timezone": "America/New_York",
			},
			{
				"agency_name":     "Test Agency 2",
				"agency_url":      "http://www.testagency2.com",
				"agency_timezone": "America/Los_Angeles",
			},
		},
	}

	validator := NewParseAgencyValidation(nil)
	messages := validator.Validate(gtfsData)

	if len(messages) != 2 {
		t.Errorf("Expected 2 errors, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Error message: %s", msg.Message)
		}
	}
}

// --- parseAgency Tests ---

func TestParseAgency_ValidAgencyAllFields(t *testing.T) {
	input := map[string]string{
		"agency_id":       "agency1",
		"agency_name":     "Test Agency",
		"agency_url":      "http://www.testagency.com",
		"agency_timezone": "America/New_York",
		"agency_lang":     "en",
		"agency_phone":    "1234567890",
		"agency_fare_url": "http://www.testagency.com/fares",
		"agency_email":    "contact@testagency.com",
	}

	_, errors := parseAgency(input, 1)
	if len(errors) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(errors))
		for _, err := range errors {
			t.Logf("Error: %s", err)
		}
	}
}

func TestParseAgency_ValidAgencyMinimumFields(t *testing.T) {
	input := map[string]string{
		"agency_name":     "Test Agency",
		"agency_url":      "http://www.testagency.com",
		"agency_timezone": "America/New_York",
	}

	_, errors := parseAgency(input, 1)
	if len(errors) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(errors))
		for _, err := range errors {
			t.Logf("Error: %s", err)
		}
	}
}

func TestParseAgency_InvalidTimezone(t *testing.T) {
	input := map[string]string{
		"agency_name":     "Test Agency",
		"agency_url":      "http://www.testagency.com",
		"agency_timezone": "Invalid/Timezone",
	}

	_, errors := parseAgency(input, 1)
	if len(errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(errors))
		for _, err := range errors {
			t.Logf("Error: %s", err)
		}
	}
}

package calendar

import (
	"main/types"
	"testing"
)

func TestParseCalendar_ValidCalendar(t *testing.T) {
	// Test a valid calendar with all fields
	input := map[string]string{
		"service_id": "service1",
		"monday":     "1",
		"tuesday":    "1",
		"wednesday":  "1",
		"thursday":   "1",
		"friday":     "1",
		"saturday":   "0",
		"sunday":     "0",
		"start_date": "20240101",
		"end_date":   "20241231",
	}

	calendar, messages := parseCalendar(input)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the calendar was parsed correctly
	if calendar.ServiceId != "service1" {
		t.Errorf("Expected service_id to be 'service1', got '%s'", calendar.ServiceId)
	}
	if !calendar.Monday {
		t.Error("Expected monday to be true")
	}
	if !calendar.Tuesday {
		t.Error("Expected tuesday to be true")
	}
	if !calendar.Wednesday {
		t.Error("Expected wednesday to be true")
	}
	if !calendar.Thursday {
		t.Error("Expected thursday to be true")
	}
	if !calendar.Friday {
		t.Error("Expected friday to be true")
	}
	if calendar.Saturday {
		t.Error("Expected saturday to be false")
	}
	if calendar.Sunday {
		t.Error("Expected sunday to be false")
	}
	if calendar.StartDate != "20240101" {
		t.Errorf("Expected start_date to be '20240101', got '%s'", calendar.StartDate)
	}
	if calendar.EndDate != "20241231" {
		t.Errorf("Expected end_date to be '20241231', got '%s'", calendar.EndDate)
	}
}

func TestParseCalendar_MinimalValidCalendar(t *testing.T) {
	// Test a minimal valid calendar with only required fields
	input := map[string]string{
		"service_id": "service1",
		"monday":     "0",
		"tuesday":    "0",
		"wednesday":  "0",
		"thursday":   "0",
		"friday":     "0",
		"saturday":   "0",
		"sunday":     "0",
		"start_date": "20240101",
		"end_date":   "20240101",
	}

	calendar, messages := parseCalendar(input)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the calendar was parsed correctly
	if calendar.ServiceId != "service1" {
		t.Errorf("Expected service_id to be 'service1', got '%s'", calendar.ServiceId)
	}
	if calendar.Monday {
		t.Error("Expected monday to be false")
	}
	if calendar.Tuesday {
		t.Error("Expected tuesday to be false")
	}
	if calendar.Wednesday {
		t.Error("Expected wednesday to be false")
	}
	if calendar.Thursday {
		t.Error("Expected thursday to be false")
	}
	if calendar.Friday {
		t.Error("Expected friday to be false")
	}
	if calendar.Saturday {
		t.Error("Expected saturday to be false")
	}
	if calendar.Sunday {
		t.Error("Expected sunday to be false")
	}
	if calendar.StartDate != "20240101" {
		t.Errorf("Expected start_date to be '20240101', got '%s'", calendar.StartDate)
	}
	if calendar.EndDate != "20240101" {
		t.Errorf("Expected end_date to be '20240101', got '%s'", calendar.EndDate)
	}
}

func TestParseCalendar_MissingRequiredFields(t *testing.T) {
	// Test a calendar with missing required fields
	input := map[string]string{
		"monday": "1",
	}

	_, messages := parseCalendar(input)

	// Check for validation messages for missing required fields
	expectedErrors := map[string]bool{
		"Service ID is required.": false,
		"Start date is required.": false,
		"End date is required.":   false,
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

func TestParseCalendar_InvalidDayValues(t *testing.T) {
	// Test a calendar with invalid day values
	input := map[string]string{
		"service_id": "service1",
		"monday":     "2", // Invalid value
		"tuesday":    "1",
		"wednesday":  "1",
		"thursday":   "1",
		"friday":     "1",
		"saturday":   "1",
		"sunday":     "1",
		"start_date": "20240101",
		"end_date":   "20241231",
	}

	_, messages := parseCalendar(input)

	// Check for validation message for invalid monday value
	found := false
	for _, msg := range messages {
		if msg.Field == "monday" && msg.Message == "Monday value must be either 0 or 1." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid monday value not found")
	}
}

func TestParseCalendar_InvalidDateFormat(t *testing.T) {
	// Test a calendar with invalid date format
	input := map[string]string{
		"service_id": "service1",
		"monday":     "1",
		"tuesday":    "1",
		"wednesday":  "1",
		"thursday":   "1",
		"friday":     "1",
		"saturday":   "1",
		"sunday":     "1",
		"start_date": "2024-01-01", // Invalid format
		"end_date":   "20241231",
	}

	_, messages := parseCalendar(input)

	// Check for validation message for invalid start_date format
	found := false
	for _, msg := range messages {
		if msg.Field == "start_date" && msg.Message == "Invalid start date format. Expected YYYYMMDD." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid start_date format not found")
	}
}

func TestParseCalendar_InvalidDateRange(t *testing.T) {
	// Test a calendar with invalid date range (end_date before start_date)
	input := map[string]string{
		"service_id": "service1",
		"monday":     "1",
		"tuesday":    "1",
		"wednesday":  "1",
		"thursday":   "1",
		"friday":     "1",
		"saturday":   "1",
		"sunday":     "1",
		"start_date": "20241231",
		"end_date":   "20240101", // Before start_date
	}

	_, messages := parseCalendar(input)

	// Check for validation message for invalid date range
	found := false
	for _, msg := range messages {
		if msg.Field == "end_date" && msg.Message == "End date must be after or equal to start date." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid date range not found")
	}
}

func TestParseCalendarValidation_Validate(t *testing.T) {
	// Test the Validate method with a valid calendar
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"calendar": {
				{
					"service_id": "service1",
					"monday":     "1",
					"tuesday":    "1",
					"wednesday":  "1",
					"thursday":   "1",
					"friday":     "1",
					"saturday":   "0",
					"sunday":     "0",
					"start_date": "20240101",
					"end_date":   "20241231",
				},
			},
		},
	}

	validator := NewParseCalendarValidation(nil)
	calendars, messages := validator.Validate(gtfs)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the calendar was parsed correctly
	if len(calendars) != 1 {
		t.Errorf("Expected 1 calendar, got %d", len(calendars))
	} else {
		calendar := calendars[0]
		if calendar.ServiceId != "service1" {
			t.Errorf("Expected service_id to be 'service1', got '%s'", calendar.ServiceId)
		}
		if !calendar.Monday {
			t.Error("Expected monday to be true")
		}
		if !calendar.Tuesday {
			t.Error("Expected tuesday to be true")
		}
		if !calendar.Wednesday {
			t.Error("Expected wednesday to be true")
		}
		if !calendar.Thursday {
			t.Error("Expected thursday to be true")
		}
		if !calendar.Friday {
			t.Error("Expected friday to be true")
		}
		if calendar.Saturday {
			t.Error("Expected saturday to be false")
		}
		if calendar.Sunday {
			t.Error("Expected sunday to be false")
		}
		if calendar.StartDate != "20240101" {
			t.Errorf("Expected start_date to be '20240101', got '%s'", calendar.StartDate)
		}
		if calendar.EndDate != "20241231" {
			t.Errorf("Expected end_date to be '20241231', got '%s'", calendar.EndDate)
		}
	}
}

func TestParseCalendarValidation_ValidateDuplicateServiceId(t *testing.T) {
	// Test the Validate method with duplicate service IDs
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"calendar": {
				{
					"service_id": "service1",
					"monday":     "1",
					"tuesday":    "1",
					"wednesday":  "1",
					"thursday":   "1",
					"friday":     "1",
					"saturday":   "0",
					"sunday":     "0",
					"start_date": "20240101",
					"end_date":   "20241231",
				},
				{
					"service_id": "service1", // Duplicate service_id
					"monday":     "0",
					"tuesday":    "0",
					"wednesday":  "0",
					"thursday":   "0",
					"friday":     "0",
					"saturday":   "1",
					"sunday":     "1",
					"start_date": "20240101",
					"end_date":   "20241231",
				},
			},
		},
	}

	validator := NewParseCalendarValidation(nil)
	_, messages := validator.Validate(gtfs)

	// Check for validation message for duplicate service_id
	found := false
	for _, msg := range messages {
		if msg.Field == "service_id" && msg.Message == "Duplicate service_id found. Service IDs must be unique." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for duplicate service_id not found")
	}
}

package calendar_dates

import (
	"main/src/types"
	"testing"
)

func TestParseCalendarDate_ValidCalendarDate(t *testing.T) {
	// Test a valid calendar date entry
	input := map[string]string{
		"service_id":     "service1",
		"date":           "20240101",
		"exception_type": "1",
	}

	calendarDate, messages := parseCalendarDate(input, true)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the calendar date was parsed correctly
	if calendarDate.ServiceId != "service1" {
		t.Errorf("Expected service_id to be 'service1', got '%s'", calendarDate.ServiceId)
	}
	if calendarDate.Date != "20240101" {
		t.Errorf("Expected date to be '20240101', got '%s'", calendarDate.Date)
	}
	if calendarDate.ExceptionType != 1 {
		t.Errorf("Expected exception_type to be 1, got %d", calendarDate.ExceptionType)
	}
}

func TestParseCalendarDate_MissingRequiredFields(t *testing.T) {
	// Test a calendar date with missing required fields
	input := map[string]string{}

	_, messages := parseCalendarDate(input, true)

	// Check for validation messages for missing required fields
	expectedErrors := map[string]bool{
		"Service ID is required.": false,
		"Date is required.":       false,
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

func TestParseCalendarDate_InvalidDate(t *testing.T) {
	// Test a calendar date with invalid date format
	input := map[string]string{
		"service_id":     "service1",
		"date":           "2024-01-01", // Invalid format, should be YYYYMMDD
		"exception_type": "1",
	}

	_, messages := parseCalendarDate(input, true)

	// Check for validation message for invalid date format
	found := false
	for _, msg := range messages {
		if msg.Field == "date" && msg.Message == "Invalid date format. Date must be in YYYYMMDD format." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid date format not found")
	}
}

func TestParseCalendarDate_InvalidExceptionType(t *testing.T) {
	// Test a calendar date with invalid exception_type
	input := map[string]string{
		"service_id":     "service1",
		"date":           "20240101",
		"exception_type": "3", // Invalid value, should be 1 or 2
	}

	_, messages := parseCalendarDate(input, true)

	// Check for validation message for invalid exception_type
	found := false
	for _, msg := range messages {
		if msg.Field == "exception_type" && msg.Message == "Invalid exception_type value. Valid values are 1 (service added) or 2 (service removed)." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid exception_type not found")
	}
}

func TestParseCalendarDatesValidation_Validate(t *testing.T) {
	// Test the Validate method with valid calendar dates
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"calendar_dates": {
				{
					"service_id":     "service1",
					"date":           "20240101",
					"exception_type": "1",
				},
				{
					"service_id":     "service1",
					"date":           "20240102",
					"exception_type": "2",
				},
			},
		},
	}

	validator := NewParseCalendarDatesValidation(nil)
	calendarDates, messages := validator.Validate(gtfs)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the calendar dates were parsed correctly
	if len(calendarDates) != 2 {
		t.Errorf("Expected 2 calendar dates, got %d", len(calendarDates))
	}
}

func TestParseCalendarDatesValidation_DuplicatePair(t *testing.T) {
	// Test the Validate method with duplicate service_id + date pairs
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"calendar_dates": {
				{
					"service_id":     "service1",
					"date":           "20240101",
					"exception_type": "1",
				},
				{
					"service_id":     "service1",
					"date":           "20240101", // Duplicate pair
					"exception_type": "2",
				},
			},
		},
	}

	validator := NewParseCalendarDatesValidation(nil)
	_, messages := validator.Validate(gtfs)

	// Check for validation message for duplicate pair
	found := false
	for _, msg := range messages {
		if msg.Field == "service_id,date" && msg.Message == "Duplicate (service_id, date) pair found. Each pair may only appear once in calendar_dates.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for duplicate (service_id, date) pair not found")
	}
}

func TestParseCalendarDatesValidation_RequiredWhenCalendarOmitted(t *testing.T) {
	// Test that calendar_dates.txt is required when calendar.txt is omitted
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"calendar_dates": {}, // Empty calendar_dates.txt
		},
	}

	validator := NewParseCalendarDatesValidation(nil)
	_, messages := validator.Validate(gtfs)

	// Check for validation message requiring entries when calendar.txt is omitted
	found := false
	for _, msg := range messages {
		if msg.Message == "calendar_dates.txt must contain at least one entry when calendar.txt is omitted" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message requiring entries when calendar.txt is omitted not found")
	}
}

func TestParseCalendarDate_ValidExceptionTypes(t *testing.T) {
	// Test both valid exception types
	testCases := []struct {
		exceptionType string
		expected      int
	}{
		{"1", 1}, // Service added
		{"2", 2}, // Service removed
	}

	for _, tc := range testCases {
		input := map[string]string{
			"service_id":     "service1",
			"date":           "20240101",
			"exception_type": tc.exceptionType,
		}

		calendarDate, messages := parseCalendarDate(input, true)

		if len(messages) != 0 {
			t.Errorf("Expected 0 validation messages for exception_type %s, got %d", tc.exceptionType, len(messages))
		}

		if calendarDate.ExceptionType != tc.expected {
			t.Errorf("Expected exception_type to be %d, got %d", tc.expected, calendarDate.ExceptionType)
		}
	}
}

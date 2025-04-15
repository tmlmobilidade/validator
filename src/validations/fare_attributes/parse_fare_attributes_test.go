package fare_attributes

import (
	"testing"
)

func TestParseFareAttribute_ValidFare(t *testing.T) {
	// Test a valid fare attribute with all fields
	input := map[string]string{
		"fare_id":           "fare1",
		"price":             "2.50",
		"currency_type":     "USD",
		"payment_method":    "1",
		"transfers":         "1",
		"agency_id":         "agency1",
		"transfer_duration": "3600",
	}

	agencyIds := map[string]int{"agency1": 1}

	fare, messages := parseFareAttribute(input, agencyIds)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the fare was parsed correctly
	if fare.FareId != "fare1" {
		t.Errorf("Expected fare_id to be 'fare1', got '%s'", fare.FareId)
	}
	if *fare.Price != 2.50 {
		t.Errorf("Expected price to be 2.50, got %f", *fare.Price)
	}
	if fare.CurrencyType != "USD" {
		t.Errorf("Expected currency_type to be 'USD', got '%s'", fare.CurrencyType)
	}
	if *fare.PaymentMethod != 1 {
		t.Errorf("Expected payment_method to be 1, got %d", *fare.PaymentMethod)
	}
	if *fare.Transfers != 1 {
		t.Errorf("Expected transfers to be 1, got %d", *fare.Transfers)
	}
	if *fare.AgencyId != "agency1" {
		t.Errorf("Expected agency_id to be 'agency1', got '%s'", *fare.AgencyId)
	}
	if *fare.TransferDuration != 3600 {
		t.Errorf("Expected transfer_duration to be 3600, got %d", *fare.TransferDuration)
	}
}

func TestParseFareAttribute_MinimalValidFare(t *testing.T) {
	// Test a minimal valid fare with only required fields
	input := map[string]string{
		"fare_id":        "fare1",
		"price":          "2.50",
		"currency_type":  "USD",
		"payment_method": "1",
	}

	agencyIds := map[string]int{}

	fare, messages := parseFareAttribute(input, agencyIds)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the fare was parsed correctly
	if fare.FareId != "fare1" {
		t.Errorf("Expected fare_id to be 'fare1', got '%s'", fare.FareId)
	}
	if *fare.Price != 2.50 {
		t.Errorf("Expected price to be 2.50, got %f", *fare.Price)
	}
	if fare.CurrencyType != "USD" {
		t.Errorf("Expected currency_type to be 'USD', got '%s'", fare.CurrencyType)
	}
	if *fare.PaymentMethod != 1 {
		t.Errorf("Expected payment_method to be 1, got %d", *fare.PaymentMethod)
	}
}

func TestParseFareAttribute_MissingRequiredFields(t *testing.T) {
	// Test a fare with missing required fields
	input := map[string]string{
		"transfer_duration": "3600",
	}

	agencyIds := map[string]int{}

	_, messages := parseFareAttribute(input, agencyIds)

	// Check for validation messages for missing required fields
	expectedErrors := map[string]bool{
		"Fare ID is required and must be unique.": false,
		"Price is required.":                      false,
		"Currency type is required.":              false,
		"Payment method is required.":             false,
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

func TestParseFareAttribute_InvalidPrice(t *testing.T) {
	// Test cases for invalid price values
	testCases := []struct {
		price    string
		expected string
	}{
		{"invalid", "Price must be a valid non-negative float value."},
		{"-1.50", "Price must be non-negative."},
	}

	for _, tc := range testCases {
		input := map[string]string{
			"fare_id":        "fare1",
			"price":          tc.price,
			"currency_type":  "USD",
			"payment_method": "1",
		}

		_, messages := parseFareAttribute(input, map[string]int{})

		found := false
		for _, msg := range messages {
			if msg.Field == "price" && msg.Message == tc.expected {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expected validation message '%s' for price '%s' not found", tc.expected, tc.price)
		}
	}
}

func TestParseFareAttribute_InvalidPaymentMethod(t *testing.T) {
	// Test a fare with invalid payment_method
	input := map[string]string{
		"fare_id":        "fare1",
		"price":          "2.50",
		"currency_type":  "USD",
		"payment_method": "2", // Invalid value
	}

	_, messages := parseFareAttribute(input, map[string]int{})

	found := false
	for _, msg := range messages {
		if msg.Field == "payment_method" && msg.Message == "Invalid payment_method value. Valid values are 0 (paid on board) or 1 (paid before boarding)." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid payment_method not found")
	}
}

func TestParseFareAttribute_InvalidTransfers(t *testing.T) {
	// Test a fare with invalid transfers value
	input := map[string]string{
		"fare_id":        "fare1",
		"price":          "2.50",
		"currency_type":  "USD",
		"payment_method": "1",
		"transfers":      "3", // Invalid value
	}

	_, messages := parseFareAttribute(input, map[string]int{})

	found := false
	for _, msg := range messages {
		if msg.Field == "transfers" && msg.Message == "Invalid transfers value. Valid values are 0 (no transfers), 1 (one transfer), 2 (two transfers), or empty (unlimited)." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid transfers not found")
	}
}

func TestParseFareAttribute_InvalidAgencyId(t *testing.T) {
	// Test a fare with invalid agency_id
	input := map[string]string{
		"fare_id":        "fare1",
		"price":          "2.50",
		"currency_type":  "USD",
		"payment_method": "1",
		"agency_id":      "invalid_agency",
	}

	agencyIds := map[string]int{"agency1": 1}

	_, messages := parseFareAttribute(input, agencyIds)

	found := false
	for _, msg := range messages {
		if msg.Field == "agency_id" && msg.Message == "Agency ID must reference a valid agency_id from agency.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid agency_id not found")
	}
}

func TestParseFareAttribute_RequiredAgencyId(t *testing.T) {
	// Test a fare without agency_id when multiple agencies exist
	input := map[string]string{
		"fare_id":        "fare1",
		"price":          "2.50",
		"currency_type":  "USD",
		"payment_method": "1",
	}

	agencyIds := map[string]int{"agency1": 1, "agency2": 2}

	_, messages := parseFareAttribute(input, agencyIds)

	found := false
	for _, msg := range messages {
		if msg.Field == "agency_id" && msg.Message == "Agency ID is required when multiple agencies are defined in agency.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for required agency_id not found")
	}
}

func TestParseFareAttribute_InvalidTransferDuration(t *testing.T) {
	// Test cases for invalid transfer_duration values
	testCases := []struct {
		duration string
		expected string
	}{
		{"invalid", "Transfer duration must be a valid non-negative integer."},
		{"-3600", "Transfer duration must be non-negative."},
	}

	for _, tc := range testCases {
		input := map[string]string{
			"fare_id":           "fare1",
			"price":             "2.50",
			"currency_type":     "USD",
			"payment_method":    "1",
			"transfer_duration": tc.duration,
		}

		_, messages := parseFareAttribute(input, map[string]int{})

		found := false
		for _, msg := range messages {
			if msg.Field == "transfer_duration" && msg.Message == tc.expected {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expected validation message '%s' for transfer_duration '%s' not found", tc.expected, tc.duration)
		}
	}
}

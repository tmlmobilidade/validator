// Package fare_attributes provides validation for fare attributes in GTFS feeds
package fare_attributes

import (
	"main/types"
	"testing"
)

// TestParseFareAttributeValidation_SingleValidFareAttribute tests validation of a single valid fare attribute
func TestParseFareAttributeValidation_SingleValidFareAttribute(t *testing.T) {
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"fare_attributes": {
				{
					"fare_id":        "fare1",
					"price":          "2.50",
					"currency_type":  "USD",
					"payment_method": "1",
					"transfers":      "1",
				},
			},
			"agency": {
				{
					"agency_id": "agency1",
				},
			},
		},
		IdMap: map[string]map[string]int{
			"agency": {
				"agency1": 1,
			},
		},
	}

	validator := NewParseFareAttributeValidation(nil)
	fareAttributes, messages := validator.Validate(gtfs)

	if len(messages) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Error message: %s", msg.Message)
		}
	}

	if len(fareAttributes) != 1 {
		t.Errorf("Expected 1 fare attribute, got %d", len(fareAttributes))
	} else {
		fareAttr := fareAttributes[0]
		if fareAttr.FareId != "fare1" {
			t.Errorf("Expected fare_id to be 'fare1', got '%s'", fareAttr.FareId)
		}
		if *fareAttr.Price != 2.50 {
			t.Errorf("Expected price to be 2.50, got %f", *fareAttr.Price)
		}
		if fareAttr.CurrencyType != "USD" {
			t.Errorf("Expected currency_type to be 'USD', got '%s'", fareAttr.CurrencyType)
		}
		if *fareAttr.PaymentMethod != 1 {
			t.Errorf("Expected payment_method to be 1, got %d", *fareAttr.PaymentMethod)
		}
		if *fareAttr.Transfers != 1 {
			t.Errorf("Expected transfers to be 1, got %d", *fareAttr.Transfers)
		}
	}
}

// TestParseFareAttributeValidation_DuplicateFareId tests validation of duplicate fare IDs
func TestParseFareAttributeValidation_DuplicateFareId(t *testing.T) {
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"fare_attributes": {
				{
					"fare_id":        "fare1",
					"price":          "2.50",
					"currency_type":  "USD",
					"payment_method": "1",
				},
				{
					"fare_id":        "fare1",
					"price":          "3.00",
					"currency_type":  "USD",
					"payment_method": "0",
				},
			},
		},
	}

	validator := NewParseFareAttributeValidation(nil)
	_, messages := validator.Validate(gtfs)

	if len(messages) != 1 {
		t.Errorf("Expected 1 error for duplicate fare_id, got %d", len(messages))
	}

	if len(messages) > 0 && messages[0].Field != "fare_id" {
		t.Errorf("Expected error for field 'fare_id', got %s", messages[0].Field)
	}
}

// TestParseFareAttributeValidation_MissingRequiredFields tests validation of missing required fields
func TestParseFareAttributeValidation_MissingRequiredFields(t *testing.T) {
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"fare_attributes": {
				{
					"fare_id": "fare1",
				},
			},
		},
	}

	validator := NewParseFareAttributeValidation(nil)
	_, messages := validator.Validate(gtfs)

	if len(messages) != 3 { // Missing price, currency_type, and payment_method
		t.Errorf("Expected 3 errors for missing required fields, got %d", len(messages))
	}

	expectedFields := map[string]bool{
		"currency_type":  false,
		"payment_method": false,
		"price":          false,
	}

	for _, msg := range messages {
		expectedFields[msg.Field] = true
	}

	for field, found := range expectedFields {
		if !found {
			t.Errorf("Expected error for field '%s' not found", field)
		}
	}
}

// TestParseFareAttributeValidation_InvalidValues tests validation of invalid field values
func TestParseFareAttributeValidation_InvalidValues(t *testing.T) {
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"fare_attributes": {
				{
					"fare_id":           "fare1",
					"price":             "-1.00",
					"currency_type":     "USD",
					"payment_method":    "2",   // Invalid payment method
					"transfers":         "3",   // Invalid transfers value
					"transfer_duration": "-60", // Invalid negative duration
				},
			},
		},
	}

	validator := NewParseFareAttributeValidation(nil)
	_, messages := validator.Validate(gtfs)

	if len(messages) != 4 { // Invalid price, payment_method, transfers, and transfer_duration
		t.Errorf("Expected 4 validation errors, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	expectedFields := map[string]bool{
		"price":             false,
		"payment_method":    false,
		"transfers":         false,
		"transfer_duration": false,
	}

	for _, msg := range messages {
		expectedFields[msg.Field] = true
	}

	for field, found := range expectedFields {
		if !found {
			t.Errorf("Expected error for field '%s' not found", field)
		}
	}
}

// TestParseFareAttributeValidation_MultipleAgenciesWithoutAgencyId tests validation when multiple agencies exist but agency_id is missing
func TestParseFareAttributeValidation_MultipleAgenciesWithoutAgencyId(t *testing.T) {
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"fare_attributes": {
				{
					"fare_id":        "fare1",
					"price":          "2.50",
					"currency_type":  "USD",
					"payment_method": "1",
				},
			},
			"agency": {
				{
					"agency_id": "agency1",
				},
				{
					"agency_id": "agency2",
				},
			},
		},
	}

	validator := NewParseFareAttributeValidation(nil)
	_, messages := validator.Validate(gtfs)

	found := false
	for _, msg := range messages {
		if msg.Field == "agency_id" && msg.Message == "Agency ID is required when the dataset contains multiple agencies." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for missing agency_id with multiple agencies not found")
	}
}

// TestParseFareAttributeValidation_InvalidAgencyId tests validation of invalid agency ID references
func TestParseFareAttributeValidation_InvalidAgencyId(t *testing.T) {
	gtfs := types.Gtfs{
		Files: map[string][]map[string]string{
			"fare_attributes": {
				{
					"fare_id":        "fare1",
					"price":          "2.50",
					"currency_type":  "USD",
					"payment_method": "1",
					"agency_id":      "invalid_agency",
				},
			},
			"agency": {
				{
					"agency_id": "agency1",
				},
				{
					"agency_id": "agency2",
				},
			},
		},
		IdMap: map[string]map[string]int{
			"agency": {
				"agency1": 1,
				"agency2": 2,
			},
		},
	}

	validator := NewParseFareAttributeValidation(nil)
	_, messages := validator.Validate(gtfs)

	found := false
	for _, msg := range messages {
		if msg.Field == "agency_id" && msg.Message == "Agency ID must reference a valid agency_id from agency.txt." {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid agency_id reference not found")
	}
}

// TestParseFareAttribute_AllValidOptionalFields tests parsing of a fare attribute with all valid optional fields
func TestParseFareAttribute_AllValidOptionalFields(t *testing.T) {
	input := map[string]string{
		"fare_id":           "fare1",
		"price":             "2.50",
		"currency_type":     "USD",
		"payment_method":    "1",
		"transfers":         "2",
		"transfer_duration": "7200",
		"agency_id":         "agency1",
	}

	fareAttr, messages := parseFareAttribute(input, false, map[string]int{"agency1": 1})

	if len(messages) != 0 {
		t.Errorf("Expected 0 errors, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Verify all fields are set correctly
	if fareAttr.FareId != "fare1" {
		t.Errorf("Expected fare_id to be 'fare1', got '%s'", fareAttr.FareId)
	}
	if *fareAttr.Price != 2.50 {
		t.Errorf("Expected price to be 2.50, got %f", *fareAttr.Price)
	}
	if fareAttr.CurrencyType != "USD" {
		t.Errorf("Expected currency_type to be 'USD', got '%s'", fareAttr.CurrencyType)
	}
	if *fareAttr.PaymentMethod != 1 {
		t.Errorf("Expected payment_method to be 1, got %d", *fareAttr.PaymentMethod)
	}
	if *fareAttr.Transfers != 2 {
		t.Errorf("Expected transfers to be 2, got %d", *fareAttr.Transfers)
	}
	if *fareAttr.TransferDuration != 7200 {
		t.Errorf("Expected transfer_duration to be 7200, got %d", *fareAttr.TransferDuration)
	}
	if *fareAttr.AgencyId != "agency1" {
		t.Errorf("Expected agency_id to be 'agency1', got '%s'", *fareAttr.AgencyId)
	}
}

package agency

import (
	"main/services"
	validations "main/validations/agency/validations"
	"testing"
)

func TestParseAgency_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	raw := map[string]string{
		"agency_name":     "Agency Name",
		"agency_url":      "https://agency.com",
		"agency_timezone": "America/New_York",
		"agency_id":       "1",
		"agency_lang":     "en",
		"agency_phone":    "1234567890",
		"agency_fare_url": "https://agency.com/fares",
		"agency_email":    "agency@example.com",
	}
	agency := validations.ParseAgency(raw, row)

	if *agency.AgencyName != "Agency Name" {
		t.Errorf("Expected AgencyName 'Agency Name', got '%s'", *agency.AgencyName)
	}
	if *agency.AgencyUrl != "https://agency.com" {
		t.Errorf("Expected AgencyUrl 'https://agency.com', got '%s'", *agency.AgencyUrl)
	}
	if *agency.AgencyTimezone != "America/New_York" {
		t.Errorf("Expected AgencyTimezone 'America/New_York', got '%s'", *agency.AgencyTimezone)
	}
	if *agency.AgencyId != "1" {
		t.Errorf("Expected AgencyId '1', got '%s'", *agency.AgencyId)
	}
	if *agency.AgencyLang != "en" {
		t.Errorf("Expected AgencyLang 'en', got '%s'", *agency.AgencyLang)
	}
	if *agency.AgencyPhone != "1234567890" {
		t.Errorf("Expected AgencyPhone '1234567890', got '%s'", *agency.AgencyPhone)
	}
	if *agency.AgencyFareUrl != "https://agency.com/fares" {
		t.Errorf("Expected AgencyFareUrl 'https://agency.com/fares', got '%s'", *agency.AgencyFareUrl)
	}
	if *agency.AgencyEmail != "agency@example.com" {
		t.Errorf("Expected AgencyEmail 'agency@example.com', got '%s'", *agency.AgencyEmail)
	}
}

func TestParseAgency_OptionalFieldsEmpty(t *testing.T) {
	services.AppMessageService.Clear()
	row := 4
	raw := map[string]string{
		"agency_name":     "Agency Name",
		"agency_url":      "https://agency.com",
		"agency_timezone": "America/New_York",
		"agency_id":       "1",
		"agency_lang":     "en",
		"agency_phone":    "1234567890",
		"agency_fare_url": "https://agency.com/fares",
		"agency_email":    "agency@example.com",
	}
	agency := validations.ParseAgency(raw, row)

	if *agency.AgencyId != "1" {
		t.Errorf("Expected AgencyId '1', got '%s'", *agency.AgencyId)
	}
	if *agency.AgencyLang != "en" {
		t.Errorf("Expected AgencyLang 'en', got '%s'", *agency.AgencyLang)
	}
	if *agency.AgencyPhone != "1234567890" {
		t.Errorf("Expected AgencyPhone '1234567890', got '%s'", *agency.AgencyPhone)
	}
	if *agency.AgencyFareUrl != "https://agency.com/fares" {
		t.Errorf("Expected AgencyFareUrl 'https://agency.com/fares', got '%s'", *agency.AgencyFareUrl)
	}
	if *agency.AgencyEmail != "agency@example.com" {
		t.Errorf("Expected AgencyEmail 'agency@example.com', got '%s'", *agency.AgencyEmail)
	}
}

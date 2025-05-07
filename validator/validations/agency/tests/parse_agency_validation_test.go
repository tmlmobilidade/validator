package agency

import (
	"main/services"
	validations "main/validations/agency/validations"
	"testing"
)

func TestParseAgencyValidation_ValidFields(t *testing.T) {
	raw := map[string]string{
		"agency_name": "Test Agency",
		"agency_url": "https://example.com",
		"agency_timezone": "America/New_York",
		"agency_id": "A1",
		"agency_lang": "en",
		"agency_phone": "503-238-RIDE",
		"agency_fare_url": "https://example.com/fare",
		"agency_email": "test@example.com",
	}
	agency := validations.ParseAgencyValidation(raw, 1, nil)
	if agency.AgencyName != "Test Agency" || agency.AgencyTimezone != "America/New_York" || agency.AgencyUrl != "https://example.com" {
		t.Errorf("Agency fields not parsed correctly")
	}
	services.AppMessageService.Clear()
}

func TestParseAgencyValidation_InvalidFields(t *testing.T) {
	raw := map[string]string{
		"agency_name": "Test Agency",
		"agency_url": "invalid-url",
		"agency_timezone": "Invalid/Timezone",
		"agency_id": "A1",
		"agency_lang": "invalid-lang",
		"agency_phone": "invalid-phone",
		"agency_fare_url": "invalid-url",
		"agency_email": "invalid-email",
	}
	_ = validations.ParseAgencyValidation(raw, 2, nil)
	if services.AppMessageService.GetSummary().TotalErrors == 0 {
		t.Errorf("Expected errors for invalid fields, got none")
	}
	services.AppMessageService.Clear()
} 
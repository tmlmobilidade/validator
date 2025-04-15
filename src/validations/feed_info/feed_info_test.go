package feed_info

import (
	"main/src/types"
	"testing"
)

func TestParseFeedInfo_ValidFeedInfo(t *testing.T) {
	// Test a valid feed info with all required fields and some optional fields
	input := map[string]string{
		"feed_publisher_name": "Transit Agency",
		"feed_publisher_url":  "https://example.com",
		"feed_lang":           "en",
		"feed_contact_email":  "contact@example.com",
		"feed_version":        "1.0",
	}

	info, messages := parseFeedInfo(input)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the feed info was parsed correctly
	if info.FeedPublisherName != "Transit Agency" {
		t.Errorf("Expected feed_publisher_name to be 'Transit Agency', got '%s'", info.FeedPublisherName)
	}
	if info.FeedPublisherUrl != "https://example.com" {
		t.Errorf("Expected feed_publisher_url to be 'https://example.com', got '%s'", info.FeedPublisherUrl)
	}
	if info.FeedLang != "en" {
		t.Errorf("Expected feed_lang to be 'en', got '%s'", info.FeedLang)
	}
	if *info.FeedContactEmail != "contact@example.com" {
		t.Errorf("Expected feed_contact_email to be 'contact@example.com', got '%s'", *info.FeedContactEmail)
	}
	if *info.FeedVersion != "1.0" {
		t.Errorf("Expected feed_version to be '1.0', got '%s'", *info.FeedVersion)
	}
}

func TestParseFeedInfo_ValidFeedInfoWithMultiLanguage(t *testing.T) {
	// Test a valid feed info with multilingual setting
	input := map[string]string{
		"feed_publisher_name": "Transit Agency",
		"feed_publisher_url":  "https://example.com",
		"feed_lang":           "mul",
		"default_lang":        "en",
		"feed_contact_email":  "contact@example.com",
	}

	info, messages := parseFeedInfo(input)

	// Check that no validation messages were generated
	if len(messages) != 0 {
		t.Errorf("Expected 0 validation messages, got %d", len(messages))
		for _, msg := range messages {
			t.Logf("Message: %s (Field: %s)", msg.Message, msg.Field)
		}
	}

	// Check that the feed info was parsed correctly
	if info.FeedLang != "mul" {
		t.Errorf("Expected feed_lang to be 'mul', got '%s'", info.FeedLang)
	}
	if *info.DefaultLang != "en" {
		t.Errorf("Expected default_lang to be 'en', got '%s'", *info.DefaultLang)
	}
}

func TestParseFeedInfo_MissingRequiredFields(t *testing.T) {
	// Test feed info with missing required fields
	input := map[string]string{
		"feed_version": "1.0",
	}

	_, messages := parseFeedInfo(input)

	// Check for validation messages for missing required fields
	expectedErrors := map[string]bool{
		"feed_publisher_name is required": false,
		"feed_publisher_url is required":  false,
		"feed_lang is required":           false,
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

func TestParseFeedInfo_InvalidUrls(t *testing.T) {
	// Test feed info with invalid URLs
	input := map[string]string{
		"feed_publisher_name": "Transit Agency",
		"feed_publisher_url":  "not-a-url",
		"feed_lang":           "en",
		"feed_contact_url":    "also-not-a-url",
		"feed_contact_email":  "contact@example.com",
	}

	_, messages := parseFeedInfo(input)

	// Check for validation messages for invalid URLs
	urlErrors := 0
	for _, msg := range messages {
		if msg.Field == "feed_publisher_url" || msg.Field == "feed_contact_url" {
			urlErrors++
		}
	}

	if urlErrors != 2 {
		t.Errorf("Expected 2 URL validation errors, got %d", urlErrors)
	}
}

func TestParseFeedInfo_InvalidLanguageCodes(t *testing.T) {
	// Test feed info with invalid language codes
	input := map[string]string{
		"feed_publisher_name": "Transit Agency",
		"feed_publisher_url":  "https://example.com",
		"feed_lang":           "invalid",
		"default_lang":        "also-invalid",
		"feed_contact_email":  "contact@example.com",
	}

	_, messages := parseFeedInfo(input)

	// Check for validation messages for invalid language codes
	langErrors := 0
	for _, msg := range messages {
		if msg.Field == "feed_lang" || msg.Field == "default_lang" {
			langErrors++
		}
	}

	if langErrors != 2 {
		t.Errorf("Expected 2 language code validation errors, got %d", langErrors)
	}
}

func TestParseFeedInfo_InvalidDates(t *testing.T) {
	// Test feed info with invalid dates
	input := map[string]string{
		"feed_publisher_name": "Transit Agency",
		"feed_publisher_url":  "https://example.com",
		"feed_lang":           "en",
		"feed_start_date":     "2024-01-01", // Wrong format
		"feed_end_date":       "2023-12-31", // Wrong format
		"feed_contact_email":  "contact@example.com",
	}

	_, messages := parseFeedInfo(input)

	// Check for validation messages for invalid dates
	dateErrors := 0
	for _, msg := range messages {
		if msg.Field == "feed_start_date" || msg.Field == "feed_end_date" {
			dateErrors++
		}
	}

	if dateErrors != 2 {
		t.Errorf("Expected 2 date validation errors, got %d", dateErrors)
	}
}

func TestParseFeedInfo_EndDateBeforeStartDate(t *testing.T) {
	// Test feed info with end date before start date
	input := map[string]string{
		"feed_publisher_name": "Transit Agency",
		"feed_publisher_url":  "https://example.com",
		"feed_lang":           "en",
		"feed_start_date":     "20240101",
		"feed_end_date":       "20231231",
		"feed_contact_email":  "contact@example.com",
	}

	_, messages := parseFeedInfo(input)

	// Check for validation message about end date being before start date
	found := false
	for _, msg := range messages {
		if msg.Field == "feed_end_date" && msg.Message == "feed_end_date cannot be earlier than feed_start_date" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for end date being before start date not found")
	}
}

func TestParseFeedInfo_InvalidEmail(t *testing.T) {
	// Test feed info with invalid email
	input := map[string]string{
		"feed_publisher_name": "Transit Agency",
		"feed_publisher_url":  "https://example.com",
		"feed_lang":           "en",
		"feed_contact_email":  "not-an-email",
	}

	_, messages := parseFeedInfo(input)

	// Check for validation message about invalid email
	found := false
	for _, msg := range messages {
		if msg.Field == "feed_contact_email" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for invalid email not found")
	}
}

func TestParseFeedInfo_NoContactInfo(t *testing.T) {
	// Test feed info with no contact information
	input := map[string]string{
		"feed_publisher_name": "Transit Agency",
		"feed_publisher_url":  "https://example.com",
		"feed_lang":           "en",
	}

	_, messages := parseFeedInfo(input)

	// Check for validation message about missing contact information
	found := false
	for _, msg := range messages {
		if msg.Message == "It's recommended to provide at least one of feed_contact_email or feed_contact_url" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for missing contact information not found")
	}
}

func TestValidate_RequiredWithTranslations(t *testing.T) {
	validator := NewParseFeedInfoValidation(nil)
	gtfs := types.Gtfs{
		Files: types.GtfsFiles{
			"translations": []map[string]string{{"field": "value"}},
		},
	}

	_, messages := validator.Validate(gtfs)

	// Check for validation message about feed_info.txt being required
	found := false
	for _, msg := range messages {
		if msg.Message == "feed_info.txt is required when translations.txt is present" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected validation message for missing feed_info.txt when translations.txt is present")
	}
}

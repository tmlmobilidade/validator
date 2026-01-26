package rider_categories_test

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/rider_categories/validations"
	"testing"
)

func TestEligibilityUrlValidation_NilEligibilityUrl(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC1"),
		EligibilityUrl:  nil,
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Nil eligibility_url should not error (optional field)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_EmptyEligibilityUrl(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC1"),
		EligibilityUrl:  lib.Ptr(""),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Empty eligibility_url should not error (optional field)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_ValidHttpUrl(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC1"),
		EligibilityUrl:  lib.Ptr("http://example.com"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid HTTP URL should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_ValidHttpsUrl(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC2"),
		EligibilityUrl:  lib.Ptr("https://example.com"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid HTTPS URL should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_ValidUrlWithPath(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC3"),
		EligibilityUrl:  lib.Ptr("https://example.com/eligibility/student"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid URL with path should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_ValidUrlWithQueryParams(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC4"),
		EligibilityUrl:  lib.Ptr("https://example.com/eligibility?category=student&lang=en"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid URL with query parameters should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_ValidUrlWithPort(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC5"),
		EligibilityUrl:  lib.Ptr("https://example.com:8080/eligibility"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid URL with port should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_InvalidUrl_NoProtocol(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC6"),
		EligibilityUrl:  lib.Ptr("example.com"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid URL without protocol should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_InvalidUrl_NotAUrl(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC7"),
		EligibilityUrl:  lib.Ptr("THIS IS NOT A URL"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid URL (not a URL) should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_InvalidUrl_Malformed(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC8"),
		EligibilityUrl:  lib.Ptr("https://"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Malformed URL should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_InvalidUrl_InvalidProtocol(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC9"),
		EligibilityUrl:  lib.Ptr("ftp://example.com"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	// FTP might be considered invalid depending on ValidateUrl implementation
	// This test checks if it errors
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "URL with invalid protocol (ftp) should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_ValidUrlWithFragment(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC10"),
		EligibilityUrl:  lib.Ptr("https://example.com/eligibility#student"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid URL with fragment should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_ValidUrlWithSpecialCharacters(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC11"),
		EligibilityUrl:  lib.Ptr("https://example.com/eligibility/student-18%2B"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid URL with special characters should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_InvalidUrl_Spaces(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC12"),
		EligibilityUrl:  lib.Ptr("https://example.com/eligibility page"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "URL with spaces should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_InvalidUrl_JustText(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC13"),
		EligibilityUrl:  lib.Ptr("invalid-url"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid URL (just text) should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_WithNilGtfs(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC14"),
		EligibilityUrl:  lib.Ptr("https://example.com"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid eligibility_url with nil gtfs should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_WithGtfs(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC15"),
		EligibilityUrl:  lib.Ptr("https://example.com"),
	}
	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{},
	}
	validations.EligibilityUrlValidation(riderCategory, row, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid eligibility_url with gtfs should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_WithNilRules(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC16"),
		EligibilityUrl:  lib.Ptr("https://example.com"),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid eligibility_url with nil rules should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_WithRules(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC17"),
		EligibilityUrl:  lib.Ptr("https://example.com"),
	}
	rules := &types.RiderCategory{}
	validations.EligibilityUrlValidation(riderCategory, row, nil, rules)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid eligibility_url with rules should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_ValidUrlWithOtherFields(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:       lib.Ptr("RC18"),
		RiderCategoryName:     lib.Ptr("Student"),
		EligibilityUrl:        lib.Ptr("https://example.com/eligibility"),
		IsDefaultFareCategory: lib.Ptr(0),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid eligibility_url with other fields populated should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_MultipleValidUrls(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	validUrls := []string{
		"https://example.com",
		"http://example.com",
		"https://example.com/eligibility",
		"https://example.com/eligibility?category=student",
		"https://subdomain.example.com/eligibility",
		"https://example.com:443/eligibility",
	}

	for i, url := range validUrls {
		services.AppMessageService.Clear()
		riderCategory := &types.RiderCategory{
			RiderCategoryId: lib.Ptr("RC19"),
			EligibilityUrl:  lib.Ptr(url),
		}
		validations.EligibilityUrlValidation(riderCategory, row+i, nil, nil)
		assertion := lib.AssertionMessage{
			Expected: 0,
			Actual:   services.AppMessageService.GetSummary().TotalErrors,
			Message:  "Valid eligibility_url '" + url + "' should not error",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Error(assert)
		}
	}
}

func TestEligibilityUrlValidation_MultipleInvalidUrls(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	invalidUrls := []string{
		"not-a-url",
		"example.com",
		"https://",
		"http://",
		"://example.com",
		"https:// example.com",
	}

	for i, url := range invalidUrls {
		services.AppMessageService.Clear()
		riderCategory := &types.RiderCategory{
			RiderCategoryId: lib.Ptr("RC20"),
			EligibilityUrl:  lib.Ptr(url),
		}
		validations.EligibilityUrlValidation(riderCategory, row+i, nil, nil)
		assertion := lib.AssertionMessage{
			Expected: 1,
			Actual:   services.AppMessageService.GetSummary().TotalErrors,
			Message:  "Invalid eligibility_url '" + url + "' should error",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Error(assert)
		}
	}
}

func TestEligibilityUrlValidation_NilVsEmptyString(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1

	// Test with nil
	riderCategoryNil := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC21"),
		EligibilityUrl:  nil,
	}
	validations.EligibilityUrlValidation(riderCategoryNil, row, nil, nil)
	assertionNil := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Nil eligibility_url should not error (optional field)",
	}
	if assert := lib.Assert(assertionNil); assert != "" {
		t.Error(assert)
	}

	// Test with empty string
	services.AppMessageService.Clear()
	riderCategoryEmpty := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC22"),
		EligibilityUrl:  lib.Ptr(""),
	}
	validations.EligibilityUrlValidation(riderCategoryEmpty, row+1, nil, nil)
	assertionEmpty := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Empty string eligibility_url should not error (optional field)",
	}
	if assert := lib.Assert(assertionEmpty); assert != "" {
		t.Error(assert)
	}
}

func TestEligibilityUrlValidation_LongValidUrl(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	longUrl := "https://example.com/very/long/path/to/eligibility/information/for/rider/categories/with/many/segments"
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC23"),
		EligibilityUrl:  lib.Ptr(longUrl),
	}
	validations.EligibilityUrlValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Long valid URL should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

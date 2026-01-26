package rider_categories_test

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/rider_categories/validations"
	"testing"
)

func TestRiderCategoryIdValidation_MissingRiderCategoryId(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId: nil,
	}

	gtfs := &types.Gtfs{}

	validations.RiderCategoryIdValidation(riderCategory, 1, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing rider_category_id should produce required error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryIdValidation_EmptyRiderCategoryId(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr(""),
	}

	gtfs := &types.Gtfs{}

	validations.RiderCategoryIdValidation(riderCategory, 2, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Empty rider_category_id should produce required error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryIdValidation_InvalidRiderCategoryId(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("INVALID_ID"),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"rider_categories": {},
		},
	}

	validations.RiderCategoryIdValidation(riderCategory, 3, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid rider_category_id should produce invalid error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryIdValidation_ValidRiderCategoryId(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("VALID_ID"),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"rider_categories": {
				"VALID_ID": {1},
			},
		},
	}

	validations.RiderCategoryIdValidation(riderCategory, 4, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid rider_category_id should not produce errors",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryIdValidation_ValidRiderCategoryId_MultipleRows(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("VALID_ID"),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"rider_categories": {
				"VALID_ID": {1, 2, 3},
			},
		},
	}

	validations.RiderCategoryIdValidation(riderCategory, 5, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid rider_category_id with multiple rows should not produce errors",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryIdValidation_EmptyGtfsIdMap(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("SOME_ID"),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{},
	}

	validations.RiderCategoryIdValidation(riderCategory, 6, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Rider category ID not in empty IdMap should produce invalid error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryIdValidation_NilGtfs(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("SOME_ID"),
	}

	validations.RiderCategoryIdValidation(riderCategory, 7, nil, nil)

	// When gtfs is nil, the validation should not error (it checks gtfs != nil first)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Rider category ID with nil gtfs should not error (validation skips IdMap check)",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryIdValidation_DifferentRiderCategoryIds(t *testing.T) {
	services.AppMessageService.Clear()

	testCases := []struct {
		name            string
		riderCategoryId string
		shouldError     bool
	}{
		{"Valid ID 1", "RC1", false},
		{"Valid ID 2", "RC2", false},
		{"Valid ID 3", "RC3", false},
		{"Invalid ID", "INVALID", true},
		{"Empty ID", "", true},
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"rider_categories": {
				"RC1": {1},
				"RC2": {2},
				"RC3": {3},
			},
		},
	}

	for i, tc := range testCases {
		services.AppMessageService.Clear()
		riderCategory := &types.RiderCategory{
			RiderCategoryId: lib.Ptr(tc.riderCategoryId),
		}

		validations.RiderCategoryIdValidation(riderCategory, i+10, gtfs, nil)

		expectedErrors := 0
		if tc.shouldError {
			expectedErrors = 1
		}

		assertion := lib.AssertionMessage{
			Expected: expectedErrors,
			Actual:   services.AppMessageService.GetSummary().TotalErrors,
			Message:  fmt.Sprintf("%s should have %d errors", tc.name, expectedErrors),
		}

		if assert := lib.Assert(assertion); assert != "" {
			t.Error(assert)
		}
	}
}

func TestRiderCategoryIdValidation_SpecialCharactersInId(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("RC-123_Test"),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"rider_categories": {
				"RC-123_Test": {1},
			},
		},
	}

	validations.RiderCategoryIdValidation(riderCategory, 8, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Rider category ID with special characters should not error if it exists",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryIdValidation_LongRiderCategoryId(t *testing.T) {
	services.AppMessageService.Clear()

	longId := "VERY_LONG_RIDER_CATEGORY_ID_THAT_SHOULD_STILL_BE_VALID_IF_IT_EXISTS_IN_THE_ID_MAP"
	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr(longId),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"rider_categories": {
				longId: {1},
			},
		},
	}

	validations.RiderCategoryIdValidation(riderCategory, 9, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Long rider category ID should not error if it exists",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryIdValidation_CaseSensitive(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("rc1"),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"rider_categories": {
				"RC1": {1}, // Different case
			},
		},
	}

	validations.RiderCategoryIdValidation(riderCategory, 10, gtfs, nil)

	// IDs are case-sensitive, so "rc1" != "RC1"
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Rider category ID should be case-sensitive",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryIdValidation_WhitespaceInId(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("  RC1  "),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"rider_categories": {
				"RC1": {1}, // Without whitespace
			},
		},
	}

	validations.RiderCategoryIdValidation(riderCategory, 11, gtfs, nil)

	// IDs with whitespace are different from IDs without
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Rider category ID with whitespace should be treated as different",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryIdValidation_WhitespaceInId_ExactMatch(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId: lib.Ptr("  RC1  "),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"rider_categories": {
				"  RC1  ": {1}, // With whitespace
			},
		},
	}

	validations.RiderCategoryIdValidation(riderCategory, 12, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Rider category ID with whitespace should match if exact match exists",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryIdValidation_NilVsEmptyString(t *testing.T) {
	services.AppMessageService.Clear()

	// Test with nil
	riderCategoryNil := &types.RiderCategory{
		RiderCategoryId: nil,
	}
	validations.RiderCategoryIdValidation(riderCategoryNil, 13, &types.Gtfs{}, nil)
	assertionNil := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Nil rider_category_id should error",
	}
	if assert := lib.Assert(assertionNil); assert != "" {
		t.Error(assert)
	}

	// Test with empty string
	services.AppMessageService.Clear()
	riderCategoryEmpty := &types.RiderCategory{
		RiderCategoryId: lib.Ptr(""),
	}
	validations.RiderCategoryIdValidation(riderCategoryEmpty, 14, &types.Gtfs{}, nil)
	assertionEmpty := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Empty string rider_category_id should error",
	}
	if assert := lib.Assert(assertionEmpty); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryIdValidation_ValidIdWithOtherFields(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId:       lib.Ptr("RC1"),
		RiderCategoryName:     lib.Ptr("Student"),
		EligibilityUrl:        lib.Ptr("https://example.com"),
		IsDefaultFareCategory: lib.Ptr(1),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"rider_categories": {
				"RC1": {1},
			},
		},
	}

	validations.RiderCategoryIdValidation(riderCategory, 15, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid rider_category_id with other fields populated should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

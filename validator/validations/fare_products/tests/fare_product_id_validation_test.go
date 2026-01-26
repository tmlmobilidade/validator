package fare_products

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_products/validations"
	"testing"
)

func TestFareProductIdValidation_MissingFareProductId(t *testing.T) {
	services.AppMessageService.Clear()

	fareProduct := &types.FareProduct{
		FareProductId: nil,
	}

	gtfs := &types.Gtfs{}

	validations.FareProductIdValidation(fareProduct, 1, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing fare_product_id should produce required error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductIdValidation_InvalidFareProductId(t *testing.T) {
	services.AppMessageService.Clear()

	fareProduct := &types.FareProduct{
		FareProductId: lib.Ptr("INVALID_ID"),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_products": {},
		},
	}

	validations.FareProductIdValidation(fareProduct, 2, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid fare_product_id should produce invalid error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductIdValidation_ValidFareProductId(t *testing.T) {
	services.AppMessageService.Clear()

	fareProduct := &types.FareProduct{
		FareProductId: lib.Ptr("VALID_ID"),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_products": {
				"VALID_ID": {1},
			},
		},
	}

	validations.FareProductIdValidation(fareProduct, 3, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid fare_product_id should not produce errors",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductIdValidation_ValidFareProductId_MultipleRows(t *testing.T) {
	services.AppMessageService.Clear()

	fareProduct := &types.FareProduct{
		FareProductId: lib.Ptr("VALID_ID"),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_products": {
				"VALID_ID": {1, 2, 3},
			},
		},
	}

	validations.FareProductIdValidation(fareProduct, 4, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid fare_product_id with multiple rows should not produce errors",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductIdValidation_EmptyGtfsIdMap(t *testing.T) {
	services.AppMessageService.Clear()

	fareProduct := &types.FareProduct{
		FareProductId: lib.Ptr("SOME_ID"),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{},
	}

	validations.FareProductIdValidation(fareProduct, 5, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Fare product ID not in empty IdMap should produce invalid error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductIdValidation_NilGtfs(t *testing.T) {
	services.AppMessageService.Clear()

	fareProduct := &types.FareProduct{
		FareProductId: lib.Ptr("SOME_ID"),
	}

	validations.FareProductIdValidation(fareProduct, 6, nil, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Fare product ID with nil gtfs should produce invalid error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductIdValidation_DifferentFareProductIds(t *testing.T) {
	services.AppMessageService.Clear()

	testCases := []struct {
		name          string
		fareProductId string
		shouldError   bool
	}{
		{"Valid ID 1", "FP1", false},
		{"Valid ID 2", "FP2", false},
		{"Valid ID 3", "FP3", false},
		{"Invalid ID", "INVALID", true},
		{"Empty ID", "", true},
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_products": {
				"FP1": {1},
				"FP2": {2},
				"FP3": {3},
			},
		},
	}

	for i, tc := range testCases {
		services.AppMessageService.Clear()
		fareProduct := &types.FareProduct{
			FareProductId: lib.Ptr(tc.fareProductId),
		}

		validations.FareProductIdValidation(fareProduct, i+10, gtfs, nil)

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

func TestFareProductIdValidation_SpecialCharactersInId(t *testing.T) {
	services.AppMessageService.Clear()

	fareProduct := &types.FareProduct{
		FareProductId: lib.Ptr("FP-123_Test"),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_products": {
				"FP-123_Test": {1},
			},
		},
	}

	validations.FareProductIdValidation(fareProduct, 7, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Fare product ID with special characters should not error if it exists",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductIdValidation_LongFareProductId(t *testing.T) {
	services.AppMessageService.Clear()

	longId := "VERY_LONG_FARE_PRODUCT_ID_THAT_SHOULD_STILL_BE_VALID_IF_IT_EXISTS_IN_THE_ID_MAP"
	fareProduct := &types.FareProduct{
		FareProductId: lib.Ptr(longId),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_products": {
				longId: {1},
			},
		},
	}

	validations.FareProductIdValidation(fareProduct, 8, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Long fare product ID should not error if it exists",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductIdValidation_CaseSensitive(t *testing.T) {
	services.AppMessageService.Clear()

	fareProduct := &types.FareProduct{
		FareProductId: lib.Ptr("fp1"),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_products": {
				"FP1": {1}, // Different case
			},
		},
	}

	validations.FareProductIdValidation(fareProduct, 9, gtfs, nil)

	// IDs are case-sensitive, so "fp1" != "FP1"
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Fare product ID should be case-sensitive",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductIdValidation_WhitespaceInId(t *testing.T) {
	services.AppMessageService.Clear()

	fareProduct := &types.FareProduct{
		FareProductId: lib.Ptr("  FP1  "),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_products": {
				"FP1": {1}, // Without whitespace
			},
		},
	}

	validations.FareProductIdValidation(fareProduct, 10, gtfs, nil)

	// IDs with whitespace are different from IDs without
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Fare product ID with whitespace should be treated as different",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductIdValidation_WhitespaceInId_ExactMatch(t *testing.T) {
	services.AppMessageService.Clear()

	fareProduct := &types.FareProduct{
		FareProductId: lib.Ptr("  FP1  "),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_products": {
				"  FP1  ": {1}, // With whitespace
			},
		},
	}

	validations.FareProductIdValidation(fareProduct, 11, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Fare product ID with whitespace should match if exact match exists",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

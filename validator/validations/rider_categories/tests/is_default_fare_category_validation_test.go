package rider_categories_test

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/rider_categories/validations"
	"testing"
)

func TestIsDefaultFareCategoryValidation_MissingIsDefaultFareCategory(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId:       lib.Ptr("RC1"),
		IsDefaultFareCategory: nil,
	}

	gtfs := &types.Gtfs{}

	validations.IsDefaultFareCategoryValidation(riderCategory, 1, gtfs, nil)

	// Nil is valid (treated as 0/empty)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing is_default_fare_category (nil) should be valid (treated as 0)",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestIsDefaultFareCategoryValidation_ValidFalse(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId:       lib.Ptr("RC1"),
		IsDefaultFareCategory: lib.Ptr(0),
	}

	gtfs := &types.Gtfs{}

	validations.IsDefaultFareCategoryValidation(riderCategory, 2, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "is_default_fare_category = false (0) should be valid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestIsDefaultFareCategoryValidation_ValidTrue(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId:       lib.Ptr("RC2"),
		IsDefaultFareCategory: lib.Ptr(1),
	}

	gtfs := &types.Gtfs{}

	validations.IsDefaultFareCategoryValidation(riderCategory, 3, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "is_default_fare_category = true (1) should be valid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestIsDefaultFareCategoryValidation_WithNilGtfs(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId:       lib.Ptr("RC3"),
		IsDefaultFareCategory: lib.Ptr(1),
	}

	validations.IsDefaultFareCategoryValidation(riderCategory, 4, nil, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid is_default_fare_category with nil gtfs should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestIsDefaultFareCategoryValidation_WithGtfs(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId:       lib.Ptr("RC4"),
		IsDefaultFareCategory: lib.Ptr(0),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{},
	}

	validations.IsDefaultFareCategoryValidation(riderCategory, 5, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid is_default_fare_category with gtfs should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestIsDefaultFareCategoryValidation_WithRules(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId:       lib.Ptr("RC5"),
		IsDefaultFareCategory: lib.Ptr(1),
	}

	rules := &types.RiderCategory{}

	validations.IsDefaultFareCategoryValidation(riderCategory, 6, nil, rules)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid is_default_fare_category with rules should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestIsDefaultFareCategoryValidation_WithNilRules(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId:       lib.Ptr("RC6"),
		IsDefaultFareCategory: lib.Ptr(0),
	}

	validations.IsDefaultFareCategoryValidation(riderCategory, 7, nil, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid is_default_fare_category with nil rules should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestIsDefaultFareCategoryValidation_WithOtherFields(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId:       lib.Ptr("RC7"),
		RiderCategoryName:     lib.Ptr("Student"),
		IsDefaultFareCategory: lib.Ptr(1),
		EligibilityUrl:        lib.Ptr("https://example.com"),
	}

	validations.IsDefaultFareCategoryValidation(riderCategory, 8, nil, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid is_default_fare_category with other fields populated should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestIsDefaultFareCategoryValidation_MultipleValidValues(t *testing.T) {
	services.AppMessageService.Clear()

	testCases := []struct {
		name                  string
		isDefaultFareCategory int
		shouldError           bool
	}{
		{"Zero (0)", 0, false},
		{"One (1)", 1, false},
	}

	for i, tc := range testCases {
		services.AppMessageService.Clear()
		riderCategory := &types.RiderCategory{
			RiderCategoryId:       lib.Ptr("RC8"),
			IsDefaultFareCategory: lib.Ptr(tc.isDefaultFareCategory),
		}

		validations.IsDefaultFareCategoryValidation(riderCategory, i+10, nil, nil)

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

func TestIsDefaultFareCategoryValidation_AllFieldsNil(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId:       nil,
		RiderCategoryName:     nil,
		IsDefaultFareCategory: nil,
		EligibilityUrl:        nil,
	}

	validations.IsDefaultFareCategoryValidation(riderCategory, 9, nil, nil)

	// Nil is valid (treated as 0/empty)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Nil is_default_fare_category should be valid (treated as 0)",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestIsDefaultFareCategoryValidation_InvalidValue(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId:       lib.Ptr("RC9"),
		IsDefaultFareCategory: lib.Ptr(2),
	}

	validations.IsDefaultFareCategoryValidation(riderCategory, 10, nil, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid is_default_fare_category value (2) should produce error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestIsDefaultFareCategoryValidation_InvalidNegativeValue(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId:       lib.Ptr("RC10"),
		IsDefaultFareCategory: lib.Ptr(-1),
	}

	validations.IsDefaultFareCategoryValidation(riderCategory, 11, nil, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid is_default_fare_category value (-1) should produce error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestIsDefaultFareCategoryValidation_InvalidLargeValue(t *testing.T) {
	services.AppMessageService.Clear()

	riderCategory := &types.RiderCategory{
		RiderCategoryId:       lib.Ptr("RC11"),
		IsDefaultFareCategory: lib.Ptr(99),
	}

	validations.IsDefaultFareCategoryValidation(riderCategory, 12, nil, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid is_default_fare_category value (99) should produce error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

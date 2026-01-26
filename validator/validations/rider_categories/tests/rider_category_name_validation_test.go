package RiderCategories_test

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/rider_categories/validations"
	"testing"
)

func TestRiderCategoryNameValidation_NilRiderCategoryName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC1"),
		RiderCategoryName: nil,
	}
	validations.RiderCategoryNameValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Nil rider_category_name should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryNameValidation_EmptyRiderCategoryName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC1"),
		RiderCategoryName: lib.Ptr(""),
	}
	validations.RiderCategoryNameValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Empty rider_category_name should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryNameValidation_ValidRiderCategoryName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC1"),
		RiderCategoryName: lib.Ptr("Student"),
	}
	validations.RiderCategoryNameValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Valid rider_category_name should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryNameValidation_LongRiderCategoryName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC2"),
		RiderCategoryName: lib.Ptr("This is a very long rider category name that should still be valid and not trigger any warnings"),
	}
	validations.RiderCategoryNameValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Long rider_category_name should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryNameValidation_ShortRiderCategoryName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC3"),
		RiderCategoryName: lib.Ptr("A"),
	}
	validations.RiderCategoryNameValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Short rider_category_name should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryNameValidation_SpecialCharactersInName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC4"),
		RiderCategoryName: lib.Ptr("Student & Senior (2024)"),
	}
	validations.RiderCategoryNameValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Rider category name with special characters should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryNameValidation_WhitespaceOnlyName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC5"),
		RiderCategoryName: lib.Ptr("   "),
	}
	validations.RiderCategoryNameValidation(riderCategory, row, nil, nil)
	// Whitespace-only string is not empty, so should not warn
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Whitespace-only rider_category_name should not warn (not empty)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryNameValidation_WhitespaceInName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC6"),
		RiderCategoryName: lib.Ptr("  Student  "),
	}
	validations.RiderCategoryNameValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Rider category name with whitespace should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryNameValidation_UnicodeCharactersInName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC7"),
		RiderCategoryName: lib.Ptr("Estudante"),
	}
	validations.RiderCategoryNameValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Rider category name with unicode characters should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryNameValidation_NumbersInName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC8"),
		RiderCategoryName: lib.Ptr("Student 18+"),
	}
	validations.RiderCategoryNameValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Rider category name with numbers should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryNameValidation_MultipleValidNames(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	validNames := []string{
		"Student",
		"Senior",
		"Child",
		"Adult",
		"Disabled",
		"Veteran",
	}

	for i, name := range validNames {
		services.AppMessageService.Clear()
		riderCategory := &types.RiderCategory{
			RiderCategoryId:   lib.Ptr("RC9"),
			RiderCategoryName: lib.Ptr(name),
		}
		validations.RiderCategoryNameValidation(riderCategory, row+i, nil, nil)
		assertion := lib.AssertionMessage{
			Expected: 0,
			Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
			Message:  "Valid rider_category_name '" + name + "' should not warn",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Error(assert)
		}
	}
}

func TestRiderCategoryNameValidation_EmptyStringVsNil(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1

	// Test with nil
	riderCategoryNil := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC10"),
		RiderCategoryName: nil,
	}
	validations.RiderCategoryNameValidation(riderCategoryNil, row, nil, nil)
	assertionNil := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Nil rider_category_name should warn",
	}
	if assert := lib.Assert(assertionNil); assert != "" {
		t.Error(assert)
	}

	// Test with empty string
	services.AppMessageService.Clear()
	riderCategoryEmpty := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC11"),
		RiderCategoryName: lib.Ptr(""),
	}
	validations.RiderCategoryNameValidation(riderCategoryEmpty, row+1, nil, nil)
	assertionEmpty := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Empty string rider_category_name should warn",
	}
	if assert := lib.Assert(assertionEmpty); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryNameValidation_WithNilRules(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC12"),
		RiderCategoryName: lib.Ptr("Student"),
	}

	validations.RiderCategoryNameValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Valid rider_category_name with nil rules should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryNameValidation_WithRules(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC13"),
		RiderCategoryName: lib.Ptr("Student"),
	}
	rules := &types.RiderCategory{}

	validations.RiderCategoryNameValidation(riderCategory, row, nil, rules)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Valid rider_category_name with rules should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryNameValidation_WithGtfs(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:   lib.Ptr("RC15"),
		RiderCategoryName: lib.Ptr("Student"),
	}
	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{},
	}

	validations.RiderCategoryNameValidation(riderCategory, row, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Valid rider_category_name with gtfs should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRiderCategoryNameValidation_ValidNameWithOtherFields(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	riderCategory := &types.RiderCategory{
		RiderCategoryId:       lib.Ptr("RC16"),
		RiderCategoryName:     lib.Ptr("Student"),
		EligibilityUrl:        lib.Ptr("https://example.com"),
		IsDefaultFareCategory: lib.Ptr(true),
	}

	validations.RiderCategoryNameValidation(riderCategory, row, nil, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Valid rider_category_name with other fields populated should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

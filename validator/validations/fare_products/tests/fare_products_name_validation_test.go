package fare_products

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_products/validations"
	"testing"
)

func TestFareProductsNameValidation_NilFareProductName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareProduct := &types.FareProduct{
		FareProductId:   lib.Ptr("FP1"),
		FareProductName: nil,
	}
	validations.FareProductsNameValidation(fareProduct, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Nil fare_product_name should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductsNameValidation_EmptyFareProductName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareProduct := &types.FareProduct{
		FareProductId:   lib.Ptr("FP1"),
		FareProductName: lib.Ptr(""),
	}
	validations.FareProductsNameValidation(fareProduct, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Empty fare_product_name should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductsNameValidation_ValidFareProductName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareProduct := &types.FareProduct{
		FareProductId:   lib.Ptr("FP1"),
		FareProductName: lib.Ptr("Single Ride Ticket"),
	}
	validations.FareProductsNameValidation(fareProduct, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Valid fare_product_name should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductsNameValidation_LongFareProductName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareProduct := &types.FareProduct{
		FareProductId:   lib.Ptr("FP2"),
		FareProductName: lib.Ptr("This is a very long fare product name that should still be valid and not trigger any warnings"),
	}
	validations.FareProductsNameValidation(fareProduct, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Long fare_product_name should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductsNameValidation_ShortFareProductName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareProduct := &types.FareProduct{
		FareProductId:   lib.Ptr("FP3"),
		FareProductName: lib.Ptr("A"),
	}
	validations.FareProductsNameValidation(fareProduct, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Short fare_product_name should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductsNameValidation_SpecialCharactersInName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareProduct := &types.FareProduct{
		FareProductId:   lib.Ptr("FP4"),
		FareProductName: lib.Ptr("Ticket & Pass (2024)"),
	}
	validations.FareProductsNameValidation(fareProduct, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Fare product name with special characters should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductsNameValidation_WhitespaceOnlyName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareProduct := &types.FareProduct{
		FareProductId:   lib.Ptr("FP5"),
		FareProductName: lib.Ptr("   "),
	}
	validations.FareProductsNameValidation(fareProduct, row, nil)
	// Whitespace-only string is not empty, so should not warn
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Whitespace-only fare_product_name should not warn (not empty)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductsNameValidation_WhitespaceInName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareProduct := &types.FareProduct{
		FareProductId:   lib.Ptr("FP6"),
		FareProductName: lib.Ptr("  Single Ride  "),
	}
	validations.FareProductsNameValidation(fareProduct, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Fare product name with whitespace should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductsNameValidation_UnicodeCharactersInName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareProduct := &types.FareProduct{
		FareProductId:   lib.Ptr("FP7"),
		FareProductName: lib.Ptr("Bilhete Único"),
	}
	validations.FareProductsNameValidation(fareProduct, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Fare product name with unicode characters should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductsNameValidation_NumbersInName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareProduct := &types.FareProduct{
		FareProductId:   lib.Ptr("FP8"),
		FareProductName: lib.Ptr("24-Hour Pass"),
	}
	validations.FareProductsNameValidation(fareProduct, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Fare product name with numbers should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductsNameValidation_MultipleValidNames(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	validNames := []string{
		"Single Ride",
		"Day Pass",
		"Monthly Pass",
		"Student Discount",
		"Senior Citizen Pass",
		"7-Day Unlimited",
	}

	for i, name := range validNames {
		services.AppMessageService.Clear()
		fareProduct := &types.FareProduct{
			FareProductId:   lib.Ptr("FP9"),
			FareProductName: lib.Ptr(name),
		}
		validations.FareProductsNameValidation(fareProduct, row+i, nil)
		assertion := lib.AssertionMessage{
			Expected: 0,
			Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
			Message:  "Valid fare_product_name '" + name + "' should not warn",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Error(assert)
		}
	}
}

func TestFareProductsNameValidation_EmptyStringVsNil(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1

	// Test with nil
	fareProductNil := &types.FareProduct{
		FareProductId:   lib.Ptr("FP10"),
		FareProductName: nil,
	}
	validations.FareProductsNameValidation(fareProductNil, row, nil)
	assertionNil := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Nil fare_product_name should warn",
	}
	if assert := lib.Assert(assertionNil); assert != "" {
		t.Error(assert)
	}

	// Test with empty string
	services.AppMessageService.Clear()
	fareProductEmpty := &types.FareProduct{
		FareProductId:   lib.Ptr("FP11"),
		FareProductName: lib.Ptr(""),
	}
	validations.FareProductsNameValidation(fareProductEmpty, row+1, nil)
	assertionEmpty := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Empty string fare_product_name should warn",
	}
	if assert := lib.Assert(assertionEmpty); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductsNameValidation_WithRules(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareProduct := &types.FareProduct{
		FareProductId:   lib.Ptr("FP12"),
		FareProductName: lib.Ptr("Valid Name"),
	}
	rules := &types.FareProductRules{}

	validations.FareProductsNameValidation(fareProduct, row, rules)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Valid fare_product_name with rules should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareProductsNameValidation_WithNilRules(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareProduct := &types.FareProduct{
		FareProductId:   lib.Ptr("FP13"),
		FareProductName: nil,
	}

	validations.FareProductsNameValidation(fareProduct, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Nil fare_product_name with nil rules should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

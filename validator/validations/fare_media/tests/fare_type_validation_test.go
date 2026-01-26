package fare_media

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_media/validations"
	"testing"
)

func TestFareTypeValidation_MissingFareMediaType(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   "FM1",
		FareMediaType: "",
	}
	validOptions := []string{"0", "1", "2", "3", "4"}
	rules := &types.FareMediaRules{
		FareType: types.RuleConfig{
			Options: &validOptions,
		},
	}
	validations.FareTypeValidation(fareMedia, row, rules)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing fare_media_type should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareTypeValidation_ValidTypes(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	validTypes := []string{"0", "1", "2", "3", "4"}
	validOptions := []string{"0", "1", "2", "3", "4"}
	rules := &types.FareMediaRules{
		FareType: types.RuleConfig{
			Options: &validOptions,
		},
	}
	for _, fareMediaType := range validTypes {
		fareMedia := &types.FareMedia{
			FareMediaId:   "FM1",
			FareMediaType: fareMediaType,
		}
		validations.FareTypeValidation(fareMedia, row, rules)
		assertion := lib.AssertionMessage{
			Expected: 0,
			Actual:   services.AppMessageService.GetSummary().TotalErrors,
			Message:  "Valid fare_media_type should not error",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Error(assert)
		}
		services.AppMessageService.Clear()
	}
}

func TestFareTypeValidation_InvalidType(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   "FM1",
		FareMediaType: "99",
	}
	validOptions := []string{"0", "1", "2", "3", "4"}
	rules := &types.FareMediaRules{
		FareType: types.RuleConfig{
			Options: &validOptions,
		},
	}
	validations.FareTypeValidation(fareMedia, row, rules)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid fare_media_type should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareTypeValidation_WithRestrictedOptions_Allowed(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   "FM1",
		FareMediaType: "2",
	}
	// Restricted to only allow "2" and "4"
	restrictedOptions := []string{"2", "4"}
	rules := &types.FareMediaRules{
		FareType: types.RuleConfig{
			Options: &restrictedOptions,
		},
	}
	validations.FareTypeValidation(fareMedia, row, rules)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Fare media type in restricted options should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareTypeValidation_WithRestrictedOptions_NotAllowed(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   "FM1",
		FareMediaType: "1",
	}
	// Restricted to only allow "2" and "4"
	restrictedOptions := []string{"2", "4"}
	rules := &types.FareMediaRules{
		FareType: types.RuleConfig{
			Options: &restrictedOptions,
		},
	}
	validations.FareTypeValidation(fareMedia, row, rules)
	// "1" is not in restrictedOptions, so it should error at step 2 (not in Options)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Fare media type not in restricted options should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareTypeValidation_WithAllOptions_InvalidType(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   "FM1",
		FareMediaType: "99",
	}
	// Options contains valid types plus ALL_OPTIONS
	allOptions := []string{"0", "1", "2", "3", "4", types.ALL_OPTIONS}
	rules := &types.FareMediaRules{
		FareType: types.RuleConfig{
			Options: &allOptions,
		},
	}
	validations.FareTypeValidation(fareMedia, row, rules)
	// "99" is not in the valid options list, so it should error at step 2
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid fare_media_type should error even with ALL_OPTIONS (must be in valid options first)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareTypeValidation_WithAllOptions_ValidType(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   "FM1",
		FareMediaType: "2",
	}
	// Options contains valid types plus ALL_OPTIONS
	allOptions := []string{"0", "1", "2", "3", "4", types.ALL_OPTIONS}
	rules := &types.FareMediaRules{
		FareType: types.RuleConfig{
			Options: &allOptions,
		},
	}
	validations.FareTypeValidation(fareMedia, row, rules)
	// "2" is in valid options, and ALL_OPTIONS allows it, so should not error
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid fare_media_type with ALL_OPTIONS should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareTypeValidation_WithNilRules(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   "FM1",
		FareMediaType: "2",
	}
	// This might cause a panic, but testing what happens
	validations.FareTypeValidation(fareMedia, row, nil)
	// If it doesn't panic, check if it errors (since rules.FareType.Options would be nil)
	summary := services.AppMessageService.GetSummary()
	if summary.TotalErrors == 0 {
		t.Log("Validation with nil rules completed without error (may be expected behavior)")
	}
}

func TestFareTypeValidation_WithNilOptions(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   "FM1",
		FareMediaType: "2",
	}
	rules := &types.FareMediaRules{
		FareType: types.RuleConfig{
			Options: nil,
		},
	}
	// This might cause a panic, but testing what happens
	validations.FareTypeValidation(fareMedia, row, rules)
	// If it doesn't panic, check if it errors
	summary := services.AppMessageService.GetSummary()
	if summary.TotalErrors == 0 {
		t.Log("Validation with nil options completed without error (may be expected behavior)")
	}
}

func TestFareTypeValidation_MultipleValidTypes(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	validOptions := []string{"0", "1", "2", "3", "4"}
	rules := &types.FareMediaRules{
		FareType: types.RuleConfig{
			Options: &validOptions,
		},
	}

	testCases := []struct {
		name          string
		fareMediaType string
		shouldError   bool
	}{
		{"Type 0 - None", "0", false},
		{"Type 1 - Physical paper ticket", "1", false},
		{"Type 2 - Physical transit card", "2", false},
		{"Type 3 - cEMV", "3", false},
		{"Type 4 - Mobile app", "4", false},
		{"Invalid type", "5", true},
		{"Invalid type", "99", true},
		{"Invalid type", "abc", true},
	}

	for _, tc := range testCases {
		services.AppMessageService.Clear()
		fareMedia := &types.FareMedia{
			FareMediaId:   "FM1",
			FareMediaType: tc.fareMediaType,
		}
		validations.FareTypeValidation(fareMedia, row, rules)
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

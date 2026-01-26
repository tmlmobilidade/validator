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
		FareMediaId:   lib.Ptr("FM1"),
		FareMediaType: nil,
	}
	// Rules are optional now - validation uses hardcoded valid options
	validations.FareTypeValidation(fareMedia, row, nil)
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
	validTypes := []int{0, 1, 2, 3, 4}
	// Rules are optional - validation uses hardcoded valid options
	for _, fareMediaType := range validTypes {
		fareMedia := &types.FareMedia{
			FareMediaId:   lib.Ptr("FM1"),
			FareMediaType: lib.Ptr(fareMediaType),
		}
		validations.FareTypeValidation(fareMedia, row, nil)
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
		FareMediaId:   lib.Ptr("FM1"),
		FareMediaType: lib.Ptr(99),
	}
	// Rules are optional - validation uses hardcoded valid options
	validations.FareTypeValidation(fareMedia, row, nil)
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
		FareMediaId:   lib.Ptr("FM1"),
		FareMediaType: lib.Ptr(2),
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
		FareMediaId:   lib.Ptr("FM1"),
		FareMediaType: lib.Ptr(1),
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
		FareMediaId:   lib.Ptr("FM1"),
		FareMediaType: lib.Ptr(99),
	}
	// ALL_OPTIONS in rules allows all valid types, but "99" is not a valid type
	allOptions := []string{types.ALL_OPTIONS}
	rules := &types.FareMediaRules{
		FareType: types.RuleConfig{
			Options: &allOptions,
		},
	}
	validations.FareTypeValidation(fareMedia, row, rules)
	// "99" is not in the hardcoded valid options list, so it should error
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid fare_media_type should error even with ALL_OPTIONS (must be in hardcoded valid options first)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareTypeValidation_WithAllOptions_ValidType(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   lib.Ptr("FM1"),
		FareMediaType: lib.Ptr(2),
	}
	// ALL_OPTIONS in rules allows all valid types (from hardcoded list)
	allOptions := []string{types.ALL_OPTIONS}
	rules := &types.FareMediaRules{
		FareType: types.RuleConfig{
			Options: &allOptions,
		},
	}
	validations.FareTypeValidation(fareMedia, row, rules)
	// "2" is in hardcoded valid options, and ALL_OPTIONS allows it, so should not error
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
		FareMediaId:   lib.Ptr("FM1"),
		FareMediaType: lib.Ptr(2),
	}
	// Rules are now optional - validation uses hardcoded valid options
	validations.FareTypeValidation(fareMedia, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid fare_media_type with nil rules should not error (uses hardcoded valid options)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareTypeValidation_WithNilOptions(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   lib.Ptr("FM1"),
		FareMediaType: lib.Ptr(2),
	}
	rules := &types.FareMediaRules{
		FareType: types.RuleConfig{
			Options: nil,
		},
	}
	// Rules with nil options are treated as no restrictions - uses hardcoded valid options
	validations.FareTypeValidation(fareMedia, row, rules)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid fare_media_type with nil options should not error (uses hardcoded valid options)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareTypeValidation_MultipleValidTypes(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	// Rules are optional - validation uses hardcoded valid options

	testCases := []struct {
		name          string
		fareMediaType int
		shouldError   bool
	}{
		{"Type 0 - None", 0, false},
		{"Type 1 - Physical paper ticket", 1, false},
		{"Type 2 - Physical transit card", 2, false},
		{"Type 3 - cEMV", 3, false},
		{"Type 4 - Mobile app", 4, false},
		{"Invalid type", 5, true},
		{"Invalid type", 99, true},
	}

	for _, tc := range testCases {
		services.AppMessageService.Clear()
		fareMedia := &types.FareMedia{
			FareMediaId:   lib.Ptr("FM1"),
			FareMediaType: lib.Ptr(tc.fareMediaType),
		}
		validations.FareTypeValidation(fareMedia, row, nil)
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

package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestDirectionPatternIdMatchValidation_ValidMatching(t *testing.T) {
	services.AppMessageService.Clear()
	patternId := "1001_0_1"
	directionId := 0
	trip := &types.Trip{
		PatternId:   lib.Ptr(patternId),
		DirectionId: &directionId,
	}
	gtfs := &types.Gtfs{}
	rules := &types.TripsRules{DirectionId: types.RuleConfig{Severity: types.SEVERITY_ERROR}}
	validations.DirectionPatternIdMatchValidation(trip, 1, gtfs, rules)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid matching pattern_id and direction_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDirectionPatternIdMatchValidation_ValidMatchingDirection1(t *testing.T) {
	services.AppMessageService.Clear()
	patternId := "1001_1_2"
	directionId := 1
	trip := &types.Trip{
		PatternId:   lib.Ptr(patternId),
		DirectionId: &directionId,
	}
	gtfs := &types.Gtfs{}
	rules := &types.TripsRules{DirectionId: types.RuleConfig{Severity: types.SEVERITY_ERROR}}
	validations.DirectionPatternIdMatchValidation(trip, 2, gtfs, rules)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid matching pattern_id with direction_id 1 should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDirectionPatternIdMatchValidation_InvalidPatternIdTooFewParts(t *testing.T) {
	services.AppMessageService.Clear()
	patternId := "1001_0"
	directionId := 0
	trip := &types.Trip{
		PatternId:   lib.Ptr(patternId),
		DirectionId: &directionId,
	}
	gtfs := &types.Gtfs{}
	rules := &types.TripsRules{DirectionId: types.RuleConfig{Severity: types.SEVERITY_ERROR}}
	validations.DirectionPatternIdMatchValidation(trip, 3, gtfs, rules)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Pattern ID with too few parts should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDirectionPatternIdMatchValidation_InvalidPatternIdTooManyParts(t *testing.T) {
	services.AppMessageService.Clear()
	patternId := "1001_0_1_2"
	directionId := 0
	trip := &types.Trip{
		PatternId:   lib.Ptr(patternId),
		DirectionId: &directionId,
	}
	gtfs := &types.Gtfs{}
	rules := &types.TripsRules{DirectionId: types.RuleConfig{Severity: types.SEVERITY_ERROR}}
	validations.DirectionPatternIdMatchValidation(trip, 4, gtfs, rules)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Pattern ID with too many parts should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDirectionPatternIdMatchValidation_InvalidDirectionIdOutOfRange(t *testing.T) {
	services.AppMessageService.Clear()
	patternId := "1001_2_1"
	directionId := 2
	trip := &types.Trip{
		PatternId:   lib.Ptr(patternId),
		DirectionId: &directionId,
	}
	gtfs := &types.Gtfs{}
	rules := &types.TripsRules{DirectionId: types.RuleConfig{Severity: types.SEVERITY_ERROR}}
	validations.DirectionPatternIdMatchValidation(trip, 5, gtfs, rules)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Pattern ID with direction_id out of range (2) should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDirectionPatternIdMatchValidation_InvalidDirectionIdNegative(t *testing.T) {
	services.AppMessageService.Clear()
	patternId := "1001_-1_1"
	directionId := -1
	trip := &types.Trip{
		PatternId:   lib.Ptr(patternId),
		DirectionId: &directionId,
	}
	gtfs := &types.Gtfs{}
	rules := &types.TripsRules{DirectionId: types.RuleConfig{Severity: types.SEVERITY_ERROR}}
	validations.DirectionPatternIdMatchValidation(trip, 6, gtfs, rules)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Pattern ID with negative direction_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDirectionPatternIdMatchValidation_InvalidDirectionIdNotInteger(t *testing.T) {
	services.AppMessageService.Clear()
	patternId := "1001_abc_1"
	directionId := 0
	trip := &types.Trip{
		PatternId:   lib.Ptr(patternId),
		DirectionId: &directionId,
	}
	gtfs := &types.Gtfs{}
	rules := &types.TripsRules{DirectionId: types.RuleConfig{Severity: types.SEVERITY_ERROR}}
	validations.DirectionPatternIdMatchValidation(trip, 7, gtfs, rules)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Pattern ID with non-integer direction_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDirectionPatternIdMatchValidation_NotMatchingDirectionId0(t *testing.T) {
	services.AppMessageService.Clear()
	patternId := "1001_0_1"
	directionId := 1
	trip := &types.Trip{
		PatternId:   lib.Ptr(patternId),
		DirectionId: &directionId,
	}
	gtfs := &types.Gtfs{}
	rules := &types.TripsRules{DirectionId: types.RuleConfig{Severity: types.SEVERITY_ERROR}}
	validations.DirectionPatternIdMatchValidation(trip, 8, gtfs, rules)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Pattern ID with direction_id 0 but trip direction_id 1 should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDirectionPatternIdMatchValidation_NotMatchingDirectionId1(t *testing.T) {
	services.AppMessageService.Clear()
	patternId := "1001_1_2"
	directionId := 0
	trip := &types.Trip{
		PatternId:   lib.Ptr(patternId),
		DirectionId: &directionId,
	}
	gtfs := &types.Gtfs{}
	rules := &types.TripsRules{DirectionId: types.RuleConfig{Severity: types.SEVERITY_ERROR}}
	validations.DirectionPatternIdMatchValidation(trip, 9, gtfs, rules)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Pattern ID with direction_id 1 but trip direction_id 0 should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDirectionPatternIdMatchValidation_Warning(t *testing.T) {
	services.AppMessageService.Clear()
	patternId := "1001_0_1"
	directionId := 1
	trip := &types.Trip{
		PatternId:   lib.Ptr(patternId),
		DirectionId: &directionId,
	}
	gtfs := &types.Gtfs{}
	rules := &types.TripsRules{DirectionId: types.RuleConfig{Severity: types.SEVERITY_WARNING}}
	validations.DirectionPatternIdMatchValidation(trip, 10, gtfs, rules)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Non-matching direction_id should produce warning when severity is WARNING",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDirectionPatternIdMatchValidation_Ignore(t *testing.T) {
	services.AppMessageService.Clear()
	patternId := "1001_0_1"
	directionId := 1
	trip := &types.Trip{
		PatternId:   lib.Ptr(patternId),
		DirectionId: &directionId,
	}
	gtfs := &types.Gtfs{}
	rules := &types.TripsRules{DirectionId: types.RuleConfig{Severity: types.SEVERITY_IGNORE}}
	validations.DirectionPatternIdMatchValidation(trip, 11, gtfs, rules)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Non-matching direction_id should be ignored when severity is IGNORE",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDirectionPatternIdMatchValidation_NoRules(t *testing.T) {
	services.AppMessageService.Clear()
	patternId := "1001_0_1"
	directionId := 0
	trip := &types.Trip{
		PatternId:   lib.Ptr(patternId),
		DirectionId: &directionId,
	}
	gtfs := &types.Gtfs{}
	validations.DirectionPatternIdMatchValidation(trip, 12, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid matching pattern_id and direction_id with nil rules should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

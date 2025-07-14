package fare_rules

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_rules/validations"
	"testing"
)

func TestContainsIdValidation_MissingContainsId(t *testing.T) {
	services.AppMessageService.Clear()
	fareRule := &types.FareRule{ContainsId: nil}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {},
		},
	}
	validations.ContainsIdValidation(fareRule, 1, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing contains_id (optional) should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContainsIdValidation_InvalidContainsId(t *testing.T) {
	services.AppMessageService.Clear()

	invalidContainsId := "INVALID"
	fareRule := &types.FareRule{ContainsId: &invalidContainsId}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {},
		},
	}

	validations.ContainsIdValidation(fareRule, 2, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid contains_id should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContainsIdValidation_ValidContainsId(t *testing.T) {
	services.AppMessageService.Clear()

	validContainsId := "CONTAIN1"
	fareRule := &types.FareRule{ContainsId: &validContainsId}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {"CONTAIN1": {1}},
		},
	}

	validations.ContainsIdValidation(fareRule, 3, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid contains_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContainsIdValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()

	fareRule := &types.FareRule{}
	gtfs := &types.Gtfs{}

	severity := types.SEVERITY_ERROR
	validations.ContainsIdValidation(fareRule, 3, gtfs, &types.FareRulesRules{ContainsId: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid contains_id should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestContainsIdValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()

	fareRule := &types.FareRule{}
	gtfs := &types.Gtfs{}

	severity := types.SEVERITY_WARNING
	validations.ContainsIdValidation(fareRule, 3, gtfs, &types.FareRulesRules{ContainsId: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Valid contains_id should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestHasBenchValidation_MissingHasBench_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasBench: nil}
	validations.HasBenchValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_bench with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasBenchValidation_MissingHasBench_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasBench: nil}
	severity := types.SEVERITY_ERROR
	validations.HasBenchValidation(stop, 2, &types.StopsRules{HasBench: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_bench with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasBenchValidation_MissingHasBench_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasBench: nil}
	severity := types.SEVERITY_WARNING
	validations.HasBenchValidation(stop, 3, &types.StopsRules{HasBench: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing has_bench with severity WARNING should warn")
	}
}

func TestHasBenchValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	val := 1
	stop := &types.Stop{HasBench: &val}
	validations.HasBenchValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_bench should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasBenchValidation_InvalidInput_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	val := 9999
	stop := &types.Stop{HasBench: &val}
	severity := types.SEVERITY_ERROR
	validations.HasBenchValidation(stop, 5, &types.StopsRules{HasBench: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid has_bench should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasBenchValidation_ValidInput_WithOptions(t *testing.T) {
	services.AppMessageService.Clear()
	val := 1
	stop := &types.Stop{HasBench: &val}
	severity := types.SEVERITY_ERROR
	validations.HasBenchValidation(stop, 7, &types.StopsRules{HasBench: types.RuleConfig{Severity: severity, Options: &[]string{"1", "2"}}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_bench with severity ERROR should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasBenchValidation_InValidInput_WithOptions(t *testing.T) {
	services.AppMessageService.Clear()
	val := 5
	stop := &types.Stop{HasBench: &val}
	severity := types.SEVERITY_ERROR
	validations.HasBenchValidation(stop, 7, &types.StopsRules{HasBench: types.RuleConfig{Severity: severity, Options: &[]string{"1", "2"}}})
	if services.AppMessageService.GetSummary().TotalErrors != 1 {
		t.Error("Valid has_bench with severity ERROR should error")
	}
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_bench with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

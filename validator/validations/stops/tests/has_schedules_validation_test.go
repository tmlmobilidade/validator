package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestHasSchedulesValidation_MissingHasSchedules_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasSchedules: nil}
	validations.HasSchedulesValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_schedules with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasSchedulesValidation_MissingHasSchedules_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasSchedules: nil}
	severity := types.SEVERITY_ERROR
	validations.HasSchedulesValidation(stop, 2, &types.StopsRules{HasSchedules: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_schedules with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasSchedulesValidation_MissingHasSchedules_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasSchedules: nil}
	severity := types.SEVERITY_WARNING
	validations.HasSchedulesValidation(stop, 3, &types.StopsRules{HasSchedules: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing has_schedules with severity WARNING should warn")
	}
}

func TestHasSchedulesValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	val := true
	stop := &types.Stop{HasSchedules: &val}
	validations.HasSchedulesValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_schedules should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

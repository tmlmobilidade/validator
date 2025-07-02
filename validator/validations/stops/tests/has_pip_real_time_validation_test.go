package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestHasPipRealTimeValidation_MissingHasPipRealTime_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasPipRealTime: nil}
	validations.HasPipRealTimeValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_pip_real_time with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasPipRealTimeValidation_MissingHasPipRealTime_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasPipRealTime: nil}
	severity := types.SEVERITY_ERROR
	validations.HasPipRealTimeValidation(stop, 2, &types.StopsRules{HasPipRealTime: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_pip_real_time with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasPipRealTimeValidation_MissingHasPipRealTime_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasPipRealTime: nil}
	severity := types.SEVERITY_WARNING
	validations.HasPipRealTimeValidation(stop, 3, &types.StopsRules{HasPipRealTime: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing has_pip_real_time with severity WARNING should warn")
	}
}

func TestHasPipRealTimeValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	val := true
	stop := &types.Stop{HasPipRealTime: &val}
	validations.HasPipRealTimeValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_pip_real_time should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

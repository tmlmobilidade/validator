package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestStopDescValidation_MissingStopDesc_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{StopDesc: nil}
	validations.StopDescValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing stop_desc with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopDescValidation_MissingStopDesc_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{StopDesc: nil}
	severity := types.SEVERITY_ERROR
	validations.StopDescValidation(stop, 2, &types.StopsRules{StopDesc: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing stop_desc with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopDescValidation_MissingStopDesc_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{StopDesc: nil}
	severity := types.SEVERITY_WARNING
	validations.StopDescValidation(stop, 3, &types.StopsRules{StopDesc: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Missing stop_desc with severity WARNING should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopDescValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	desc := "A description"
	stop := &types.Stop{StopDesc: &desc}
	validations.StopDescValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid stop_desc should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopDescValidation_DuplicateStopDesc(t *testing.T) {
	services.AppMessageService.Clear()
	val := "Duplicate"
	stop := &types.Stop{StopDesc: &val, StopName: &val}
	validations.StopDescValidation(stop, 5, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Duplicate stop_desc should not error, but should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Expected 1 warning for duplicate stop_desc")
	}
}

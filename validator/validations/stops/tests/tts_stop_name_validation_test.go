package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestTtsStopNameValidation_MissingTtsStopName_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{TtsStopName: nil}
	validations.TtsStopNameValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing tts_stop_name with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTtsStopNameValidation_MissingTtsStopName_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{TtsStopName: nil}
	severity := types.SEVERITY_ERROR
	validations.TtsStopNameValidation(stop, 2, &types.StopsRules{TtsStopName: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing tts_stop_name with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTtsStopNameValidation_MissingTtsStopName_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{TtsStopName: nil}
	severity := types.SEVERITY_WARNING
	validations.TtsStopNameValidation(stop, 3, &types.StopsRules{TtsStopName: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Missing tts_stop_name with severity WARNING should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTtsStopNameValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	val := "Main St"
	stop := &types.Stop{TtsStopName: &val}
	validations.TtsStopNameValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid tts_stop_name should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

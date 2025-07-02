package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestStopTimezoneValidation_MissingStopTimezone_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{StopTimezone: nil}
	validations.StopTimezoneValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing stop_timezone with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopTimezoneValidation_MissingStopTimezone_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{StopTimezone: nil}
	severity := types.SEVERITY_ERROR
	validations.StopTimezoneValidation(stop, 2, &types.StopsRules{StopTimezone: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing stop_timezone with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopTimezoneValidation_MissingStopTimezone_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{StopTimezone: nil}
	severity := types.SEVERITY_WARNING
	validations.StopTimezoneValidation(stop, 3, &types.StopsRules{StopTimezone: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Missing stop_timezone with severity WARNING should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopTimezoneValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	tz := "Europe/Lisbon"
	stop := &types.Stop{StopTimezone: &tz}
	validations.StopTimezoneValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid stop_timezone should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopTimezoneValidation_InvalidTimezone(t *testing.T) {
	services.AppMessageService.Clear()
	tz := "Invalid/Timezone"
	stop := &types.Stop{StopTimezone: &tz}
	validations.StopTimezoneValidation(stop, 5, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid stop_timezone should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

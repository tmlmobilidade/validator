package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestStopNameValidation_MissingRequiredStopName(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{
		LocationType: func() *int { i := 0; return &i }(), // stop
		StopName:     nil,
	}
	validations.StopNameValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing stop_name with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}

	// Now test with severity ERROR
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	validations.StopNameValidation(stop, 2, &types.StopsRules{StopName: types.RuleConfig{Severity: severity}})
	assertion = lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing stop_name with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopNameValidation_OptionalStopName(t *testing.T) {
	services.AppMessageService.Clear()
	lt := 3 // generic node
	stop := &types.Stop{
		LocationType: &lt,
		StopName:     nil,
	}
	validations.StopNameValidation(stop, 3, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing stop_name for optional location_type should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopNameValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	lt := 1 // station
	name := "Central Station"
	stop := &types.Stop{
		LocationType: &lt,
		StopName:     &name,
	}
	validations.StopNameValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid stop_name should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopNameValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()

	severity := types.SEVERITY_ERROR
	stop := &types.Stop{}
	validations.StopNameValidation(stop, 5, &types.StopsRules{StopName: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing stop_name with severity ERROR should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopNameValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()

	severity := types.SEVERITY_WARNING
	stop := &types.Stop{}
	validations.StopNameValidation(stop, 5, &types.StopsRules{StopName: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Missing stop_name with severity WARNING should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

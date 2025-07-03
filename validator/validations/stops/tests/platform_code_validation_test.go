package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestPlatformCodeValidation_MissingPlatformCode_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{PlatformCode: nil}
	validations.PlatformCodeValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing platform_code with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPlatformCodeValidation_MissingPlatformCode_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{PlatformCode: nil}
	severity := types.SEVERITY_ERROR
	validations.PlatformCodeValidation(stop, 2, &types.StopsRules{PlatformCode: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing platform_code with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPlatformCodeValidation_MissingPlatformCode_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{PlatformCode: nil}
	severity := types.SEVERITY_WARNING
	validations.PlatformCodeValidation(stop, 3, &types.StopsRules{PlatformCode: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing platform_code with severity WARNING should warn")
	}
}

func TestPlatformCodeValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	val := "3"
	stop := &types.Stop{PlatformCode: &val}
	validations.PlatformCodeValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid platform_code should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

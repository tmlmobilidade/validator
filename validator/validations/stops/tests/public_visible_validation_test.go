package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestPublicVisibleValidation_MissingPublicVisible_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{PublicVisible: nil}
	validations.PublicVisibleValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing public_visible with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPublicVisibleValidation_MissingPublicVisible_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{PublicVisible: nil}
	severity := types.SEVERITY_ERROR
	validations.PublicVisibleValidation(stop, 2, &types.StopsRules{PublicVisible: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing public_visible with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestPublicVisibleValidation_MissingPublicVisible_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{PublicVisible: nil}
	severity := types.SEVERITY_WARNING
	validations.PublicVisibleValidation(stop, 3, &types.StopsRules{PublicVisible: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing public_visible with severity WARNING should warn")
	}
}

func TestPublicVisibleValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	val := true
	stop := &types.Stop{PublicVisible: &val}
	validations.PublicVisibleValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid public_visible should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestParishIdValidation_MissingParishId_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{ParishId: nil}
	validations.ParishIdValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing parish_id with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestParishIdValidation_MissingParishId_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{ParishId: nil}
	severity := types.SEVERITY_ERROR
	validations.ParishIdValidation(stop, 2, &types.StopsRules{ParishId: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing parish_id with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestParishIdValidation_MissingParishId_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{ParishId: nil}
	severity := types.SEVERITY_WARNING
	validations.ParishIdValidation(stop, 3, &types.StopsRules{ParishId: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing parish_id with severity WARNING should warn")
	}
}

func TestParishIdValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	id := "PAR123"
	stop := &types.Stop{ParishId: &id}
	validations.ParishIdValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid parish_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

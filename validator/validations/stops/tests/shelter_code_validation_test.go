package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestShelterCodeValidation_MissingShelterCode_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{ShelterCode: nil}
	validations.ShelterCodeValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing shelter_code with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShelterCodeValidation_MissingShelterCode_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{ShelterCode: nil}
	severity := types.SEVERITY_ERROR
	validations.ShelterCodeValidation(stop, 2, &types.StopsRules{ShelterCode: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing shelter_code with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestShelterCodeValidation_MissingShelterCode_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{ShelterCode: nil}
	severity := types.SEVERITY_WARNING
	validations.ShelterCodeValidation(stop, 3, &types.StopsRules{ShelterCode: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing shelter_code with severity WARNING should warn")
	}
}

func TestShelterCodeValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	code := "SH123"
	stop := &types.Stop{ShelterCode: &code}
	validations.ShelterCodeValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid shelter_code should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

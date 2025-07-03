package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestHasShelterValidation_MissingHasShelter_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasShelter: nil}
	validations.HasShelterValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_Shelter with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasShelterValidation_MissingHasShelter_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasShelter: nil}
	severity := types.SEVERITY_ERROR
	validations.HasShelterValidation(stop, 2, &types.StopsRules{HasShelter: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_Shelter with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasShelterValidation_MissingHasShelter_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasShelter: nil}
	severity := types.SEVERITY_WARNING
	validations.HasShelterValidation(stop, 3, &types.StopsRules{HasShelter: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing has_Shelter with severity WARNING should warn")
	}
}

func TestHasShelterValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	val := 1
	stop := &types.Stop{HasShelter: &val}
	validations.HasShelterValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_Shelter should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasShelterValidation_InvalidInput_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	val := 9999
	stop := &types.Stop{HasShelter: &val}
	severity := types.SEVERITY_ERROR
	validations.HasShelterValidation(stop, 5, &types.StopsRules{HasShelter: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid has_Shelter should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasShelterValidation_ValidInput_WithOptions(t *testing.T) {
	services.AppMessageService.Clear()
	val := 1
	stop := &types.Stop{HasShelter: &val}
	severity := types.SEVERITY_ERROR
	validations.HasShelterValidation(stop, 7, &types.StopsRules{HasShelter: types.RuleConfig{Severity: severity, Options: &[]string{"1", "2"}}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_Shelter with severity ERROR should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasShelterValidation_InValidInput_WithOptions(t *testing.T) {
	services.AppMessageService.Clear()
	val := 5
	stop := &types.Stop{HasShelter: &val}
	severity := types.SEVERITY_ERROR
	validations.HasShelterValidation(stop, 7, &types.StopsRules{HasShelter: types.RuleConfig{Severity: severity, Options: &[]string{"1", "2"}}})
	if services.AppMessageService.GetSummary().TotalErrors != 1 {
		t.Error("Valid has_Shelter with severity ERROR should error")
	}
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_Shelter with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestHasTariffsInformationValidation_MissingHasTariffsInformation_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasTariffsInformation: nil}
	validations.HasTariffsInformationValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_tariffs_information with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasTariffsInformationValidation_MissingHasTariffsInformation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasTariffsInformation: nil}
	severity := types.SEVERITY_ERROR
	validations.HasTariffsInformationValidation(stop, 2, &types.StopsRules{HasTariffsInformation: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_tariffs_information with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasTariffsInformationValidation_MissingHasTariffsInformation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasTariffsInformation: nil}
	severity := types.SEVERITY_WARNING
	validations.HasTariffsInformationValidation(stop, 3, &types.StopsRules{HasTariffsInformation: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing has_tariffs_information with severity WARNING should warn")
	}
}

func TestHasTariffsInformationValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	val := true
	stop := &types.Stop{HasTariffsInformation: &val}
	validations.HasTariffsInformationValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_tariffs_information should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

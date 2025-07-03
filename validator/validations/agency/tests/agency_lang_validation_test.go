package agency

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyLangValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	agency := &types.Agency{AgencyLang: nil}
	validations.AgencyLangValidation(agency, 1, &types.AgencyRules{AgencyLang: types.RuleConfig{Severity: severity}})

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency language is required",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyLangValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	agency := &types.Agency{AgencyLang: nil}
	validations.AgencyLangValidation(agency, 2, &types.AgencyRules{AgencyLang: types.RuleConfig{Severity: severity}})

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Agency language is recommended",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyLangValidation_ValidLang(t *testing.T) {
	severity := types.SEVERITY_ERROR
	lang := "en"
	agency := &types.Agency{AgencyLang: &lang}
	validations.AgencyLangValidation(agency, 3, &types.AgencyRules{AgencyLang: types.RuleConfig{Severity: severity}})

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency language is valid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyLangValidation_InvalidLang(t *testing.T) {
	severity := types.SEVERITY_ERROR
	lang := "invalid-lang"
	agency := &types.Agency{AgencyLang: &lang}
	validations.AgencyLangValidation(agency, 4, &types.AgencyRules{AgencyLang: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency language is invalid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
} 
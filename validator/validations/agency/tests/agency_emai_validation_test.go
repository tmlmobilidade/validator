package agency

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyEmailValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	agency := &types.Agency{AgencyEmail: nil}
	validations.AgencyEmailValidation(agency, 1, &types.AgencyRules{AgencyEmail: types.RuleConfig{Severity: severity}})
	
	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency email is required",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}

	services.AppMessageService.Clear()
}

func TestAgencyEmailValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	agency := &types.Agency{AgencyEmail: nil}
	
	validations.AgencyEmailValidation(agency, 2, &types.AgencyRules{AgencyEmail: types.RuleConfig{Severity: severity}})
	
	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Agency email is recommended",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}

	services.AppMessageService.Clear()
}

func TestAgencyEmailValidation_ValidEmail(t *testing.T) {
	severity := types.SEVERITY_ERROR
	email := "test@example.com"
	agency := &types.Agency{AgencyEmail: &email}
	validations.AgencyEmailValidation(agency, 3, &types.AgencyRules{AgencyEmail: types.RuleConfig{Severity: severity}})

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency email is valid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}

	services.AppMessageService.Clear()
}

func TestAgencyEmailValidation_InvalidEmail(t *testing.T) {
	severity := types.SEVERITY_ERROR
	email := "invalid-email"
	agency := &types.Agency{AgencyEmail: &email}
	validations.AgencyEmailValidation(agency, 4, &types.AgencyRules{AgencyEmail: types.RuleConfig{Severity: severity}})
	
	// Assert	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency email is invalid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}

	services.AppMessageService.Clear()
}
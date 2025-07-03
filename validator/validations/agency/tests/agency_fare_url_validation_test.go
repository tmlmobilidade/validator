package agency

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyFareUrlValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	agency := &types.Agency{AgencyFareUrl: nil}
	validations.AgencyFareUrlValidation(agency, 1, &types.AgencyRules{AgencyFare: types.RuleConfig{Severity: severity}})

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency fare URL is required",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyFareUrlValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	agency := &types.Agency{AgencyFareUrl: nil}
	validations.AgencyFareUrlValidation(agency, 2, &types.AgencyRules{AgencyFare: types.RuleConfig{Severity: severity}})

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Agency fare URL is recommended",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyFareUrlValidation_ValidUrl(t *testing.T) {
	severity := types.SEVERITY_ERROR
	url := "https://example.com/fare"
	agency := &types.Agency{AgencyFareUrl: &url}
	validations.AgencyFareUrlValidation(agency, 3, &types.AgencyRules{AgencyFare: types.RuleConfig{Severity: severity}})

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency fare URL is valid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyFareUrlValidation_InvalidUrl(t *testing.T) {
	severity := types.SEVERITY_ERROR
	url := "invalid-url"
	agency := &types.Agency{AgencyFareUrl: &url}
	validations.AgencyFareUrlValidation(agency, 4, &types.AgencyRules{AgencyFare: types.RuleConfig{Severity: severity}})

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency fare URL is invalid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
} 
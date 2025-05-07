package agency

import (
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyFareUrlValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	agency := &types.Agency{AgencyFareUrl: nil}
	validations.AgencyFareUrlValidation(&severity, agency, 1, nil)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency fare URL is required",
	})
	services.AppMessageService.Clear()
}

func TestAgencyFareUrlValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	agency := &types.Agency{AgencyFareUrl: nil}
	validations.AgencyFareUrlValidation(&severity, agency, 2, nil)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Agency fare URL is recommended",
	})
	services.AppMessageService.Clear()
}

func TestAgencyFareUrlValidation_ValidUrl(t *testing.T) {
	severity := types.SEVERITY_ERROR
	url := "https://example.com/fare"
	agency := &types.Agency{AgencyFareUrl: &url}
	validations.AgencyFareUrlValidation(&severity, agency, 3, nil)
	assertMessage(t, Assertion{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency fare URL is valid",
	})
	services.AppMessageService.Clear()
}

func TestAgencyFareUrlValidation_InvalidUrl(t *testing.T) {
	severity := types.SEVERITY_ERROR
	url := "invalid-url"
	agency := &types.Agency{AgencyFareUrl: &url}
	validations.AgencyFareUrlValidation(&severity, agency, 4, nil)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency fare URL is invalid",
	})
	services.AppMessageService.Clear()
} 
package agency

import (
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyUrlValidation_Required(t *testing.T) {
	agency := &types.Agency{AgencyUrl: ""}
	validations.AgencyUrlValidation(agency, 1, nil)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency URL is required",
	})
	services.AppMessageService.Clear()
}

func TestAgencyUrlValidation_ValidUrl(t *testing.T) {
	agency := &types.Agency{AgencyUrl: "https://example.com"}
	validations.AgencyUrlValidation(agency, 2, nil)
	assertMessage(t, Assertion{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency URL is valid",
	})
	services.AppMessageService.Clear()
}

func TestAgencyUrlValidation_InvalidUrl(t *testing.T) {
	agency := &types.Agency{AgencyUrl: "invalid-url"}
	validations.AgencyUrlValidation(agency, 3, nil)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency URL is invalid",
	})
	services.AppMessageService.Clear()
} 
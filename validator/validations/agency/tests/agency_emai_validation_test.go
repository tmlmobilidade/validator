package agency

import (
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

type Assertion struct {
	Expected int
	Actual int
	Message string
}

func assertMessage(t *testing.T, assertion Assertion) {
	if assertion.Expected != assertion.Actual {
		t.Errorf("%s | Expected %v, got: %v", assertion.Message, assertion.Expected, assertion.Actual)
	}
}

func TestAgencyEmailValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	agency := &types.Agency{AgencyEmail: nil}
	validations.AgencyEmailValidation(&severity, agency, 1, nil)
	
	// Assert
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency email is required",
	})

	services.AppMessageService.Clear()
}

func TestAgencyEmailValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	agency := &types.Agency{AgencyEmail: nil}
	
	validations.AgencyEmailValidation(&severity, agency, 2, nil)
	
	// Assert
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Agency email is recommended",
	})

	services.AppMessageService.Clear()
}

func TestAgencyEmailValidation_ValidEmail(t *testing.T) {
	severity := types.SEVERITY_ERROR
	email := "test@example.com"
	agency := &types.Agency{AgencyEmail: &email}
	validations.AgencyEmailValidation(&severity, agency, 3, nil)

	// Assert
	assertMessage(t, Assertion{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency email is valid",
	})

	services.AppMessageService.Clear()
}

func TestAgencyEmailValidation_InvalidEmail(t *testing.T) {
	severity := types.SEVERITY_ERROR
	email := "invalid-email"
	agency := &types.Agency{AgencyEmail: &email}
	validations.AgencyEmailValidation(&severity, agency, 4, nil)
	
	// Assert
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency email is invalid",
	})

	services.AppMessageService.Clear()
} 
package agency

import (
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyPhoneValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	agency := &types.Agency{AgencyPhone: nil}
	validations.AgencyPhoneValidation(&severity, agency, 1, nil)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency phone is required",
	})
	services.AppMessageService.Clear()
}

func TestAgencyPhoneValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	agency := &types.Agency{AgencyPhone: nil}
	validations.AgencyPhoneValidation(&severity, agency, 2, nil)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Agency phone is recommended",
	})
	services.AppMessageService.Clear()
}

func TestAgencyPhoneValidation_ValidPhone(t *testing.T) {
	severity := types.SEVERITY_ERROR
	phone := "503-238-RIDE"
	agency := &types.Agency{AgencyPhone: &phone}
	validations.AgencyPhoneValidation(&severity, agency, 3, nil)
	assertMessage(t, Assertion{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency phone is valid",
	})
	services.AppMessageService.Clear()
}

func TestAgencyPhoneValidation_InvalidPhone(t *testing.T) {
	severity := types.SEVERITY_ERROR
	phone := "invalid-phone"
	agency := &types.Agency{AgencyPhone: &phone}
	validations.AgencyPhoneValidation(&severity, agency, 4, nil)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency phone is invalid",
	})
	services.AppMessageService.Clear()
} 
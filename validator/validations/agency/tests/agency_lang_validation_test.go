package agency

import (
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyLangValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	agency := &types.Agency{AgencyLang: nil}
	validations.AgencyLangValidation(&severity, agency, 1, nil)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency language is required",
	})
	services.AppMessageService.Clear()
}

func TestAgencyLangValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	agency := &types.Agency{AgencyLang: nil}
	validations.AgencyLangValidation(&severity, agency, 2, nil)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Agency language is recommended",
	})
	services.AppMessageService.Clear()
}

func TestAgencyLangValidation_ValidLang(t *testing.T) {
	severity := types.SEVERITY_ERROR
	lang := "en"
	agency := &types.Agency{AgencyLang: &lang}
	validations.AgencyLangValidation(&severity, agency, 3, nil)
	assertMessage(t, Assertion{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency language is valid",
	})
	services.AppMessageService.Clear()
}

func TestAgencyLangValidation_InvalidLang(t *testing.T) {
	severity := types.SEVERITY_ERROR
	lang := "invalid-lang"
	agency := &types.Agency{AgencyLang: &lang}
	validations.AgencyLangValidation(&severity, agency, 4, nil)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency language is invalid",
	})
	services.AppMessageService.Clear()
} 
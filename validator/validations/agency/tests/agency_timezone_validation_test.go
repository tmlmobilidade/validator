package agency

import (
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyTimezoneValidation_Required(t *testing.T) {
	agency := &types.Agency{AgencyTimezone: ""}
	validations.AgencyTimezoneValidation(agency, 1, nil)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency timezone is required",
	})
	services.AppMessageService.Clear()
}

func TestAgencyTimezoneValidation_ValidTimezone(t *testing.T) {
	agency := &types.Agency{AgencyTimezone: "America/New_York"}
	validations.AgencyTimezoneValidation(agency, 2, nil)
	assertMessage(t, Assertion{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency timezone is valid",
	})
	services.AppMessageService.Clear()
}

func TestAgencyTimezoneValidation_InvalidTimezone(t *testing.T) {
	agency := &types.Agency{AgencyTimezone: "Invalid/Timezone"}
	validations.AgencyTimezoneValidation(agency, 3, nil)
	assertMessage(t, Assertion{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency timezone is invalid",
	})
	services.AppMessageService.Clear()
} 
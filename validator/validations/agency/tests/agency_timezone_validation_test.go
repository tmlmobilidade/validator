package agency

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyTimezoneValidation_Required(t *testing.T) {
	agency := &types.Agency{AgencyTimezone: ""}
	validations.AgencyTimezoneValidation(agency, 1)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency timezone is required",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyTimezoneValidation_ValidTimezone(t *testing.T) {
	agency := &types.Agency{AgencyTimezone: "America/New_York"}
	validations.AgencyTimezoneValidation(agency, 2)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency timezone is valid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyTimezoneValidation_InvalidTimezone(t *testing.T) {
	agency := &types.Agency{AgencyTimezone: "Invalid/Timezone"}
	validations.AgencyTimezoneValidation(agency, 3)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency timezone is invalid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
} 
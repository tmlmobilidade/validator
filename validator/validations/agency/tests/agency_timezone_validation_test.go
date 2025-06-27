package agency

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyTimezoneValidation_Required(t *testing.T) {
	agency := &types.Agency{AgencyTimezone: nil}
	validations.AgencyTimezoneValidation(agency, 1, &types.AgencyRules{AgencyTimezone: types.RuleConfig{Severity: types.SEVERITY_ERROR}})

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
	agency := &types.Agency{AgencyTimezone: lib.Ptr("America/New_York")}
	validations.AgencyTimezoneValidation(agency, 2, &types.AgencyRules{AgencyTimezone: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
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
	agency := &types.Agency{AgencyTimezone: lib.Ptr("Invalid/Timezone")}
	validations.AgencyTimezoneValidation(agency, 3, &types.AgencyRules{AgencyTimezone: types.RuleConfig{Severity: types.SEVERITY_ERROR}})

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
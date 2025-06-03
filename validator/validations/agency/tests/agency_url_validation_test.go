package agency

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyUrlValidation_Required(t *testing.T) {
	agency := &types.Agency{AgencyUrl: nil}
	validations.AgencyUrlValidation(agency, 1)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency URL is required",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}

	services.AppMessageService.Clear()
}

func TestAgencyUrlValidation_ValidUrl(t *testing.T) {
	agency := &types.Agency{AgencyUrl: lib.Ptr("https://example.com")}
	validations.AgencyUrlValidation(agency, 2)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency URL is valid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyUrlValidation_InvalidUrl(t *testing.T) {
	agency := &types.Agency{AgencyUrl: lib.Ptr("invalid-url")}
	validations.AgencyUrlValidation(agency, 3)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency URL is invalid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	
	services.AppMessageService.Clear()
} 
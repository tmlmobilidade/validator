package agency

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyUrlValidation_Required(t *testing.T) {
	agency := &types.Agency{AgencyUrl: ""}
	validations.AgencyUrlValidation(agency, 1, nil)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency URL is required",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error("", assert)
	}

	services.AppMessageService.Clear()
}

func TestAgencyUrlValidation_ValidUrl(t *testing.T) {
	agency := &types.Agency{AgencyUrl: "https://example.com"}
	validations.AgencyUrlValidation(agency, 2, nil)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency URL is valid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error("", assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyUrlValidation_InvalidUrl(t *testing.T) {
	agency := &types.Agency{AgencyUrl: "invalid-url"}
	validations.AgencyUrlValidation(agency, 3, nil)

	// Assert
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Agency URL is invalid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error("", assert)
	}
	
	services.AppMessageService.Clear()
} 
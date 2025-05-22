package fare_attributes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_attributes/validations"
	"testing"
)

func TestAgencyIdValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	fareAttribute := &types.FareAttribute{AgencyId: nil}
	gtfs := &types.Gtfs{}
	validations.AgencyIdValidation(&severity, fareAttribute, 1, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "AgencyId is required",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyIdValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	fareAttribute := &types.FareAttribute{AgencyId: nil}
	gtfs := &types.Gtfs{}
	validations.AgencyIdValidation(&severity, fareAttribute, 2, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "AgencyId is recommended",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyIdValidation_Ignore(t *testing.T) {
	severity := types.SEVERITY_IGNORE
	fareAttribute := &types.FareAttribute{AgencyId: nil}
	gtfs := &types.Gtfs{}
	validations.AgencyIdValidation(&severity, fareAttribute, 3, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "AgencyId is ignored",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyIdValidation_ValidAgencyId(t *testing.T) {
	severity := types.SEVERITY_ERROR
	agencyId := "MY_AGENCY_ID"
	fareAttribute := &types.FareAttribute{AgencyId: &agencyId}
	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"agency": {
				"MY_AGENCY_ID": []int{1},
			},
		},
	}
	validations.AgencyIdValidation(&severity, fareAttribute, 4, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "AgencyId is valid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAgencyIdValidation_InvalidAgencyId(t *testing.T) {
	severity := types.SEVERITY_ERROR
	agencyId := "INVALID_AGENCY_ID"
	fareAttribute := &types.FareAttribute{AgencyId: &agencyId}
	gtfs := &types.Gtfs{}
	validations.AgencyIdValidation(&severity, fareAttribute, 5, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "AgencyId is invalid",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}





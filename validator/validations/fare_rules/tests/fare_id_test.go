package fare_rules

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_rules/validations"
	"testing"
)

func TestFareIdValidation_MissingFareId(t *testing.T) {
	services.AppMessageService.Clear()
	fareRule := &types.FareRule{FareId: nil}
	gtfs := &types.Gtfs{}
	validations.FareIdValidation(fareRule, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing fare_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareIdValidation_InvalidFareId(t *testing.T) {
	services.AppMessageService.Clear()
	fareId := "INVALID_ID"
	fareRule := &types.FareRule{FareId: &fareId}
	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_attributes": {},
		},
	}
	validations.FareIdValidation(fareRule, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid fare_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareIdValidation_ValidFareId(t *testing.T) {
	services.AppMessageService.Clear()
	fareId := "VALID_ID"
	fareRule := &types.FareRule{FareId: &fareId}
	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_attributes": {
				"VALID_ID": {1},
			},
		},
	}
	validations.FareIdValidation(fareRule, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid fare_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

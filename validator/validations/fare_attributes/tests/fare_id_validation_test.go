package fare_attributes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_attributes/validations"
	"testing"
)

func TestFareIdValidation_MissingFareId(t *testing.T) {
	services.AppMessageService.Clear()
	fareAttribute := &types.FareAttribute{FareId: nil}
	gtfs := &types.Gtfs{}
	validations.FareIdValidation(fareAttribute, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing fare_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareIdValidation_DuplicateFareId(t *testing.T) {
	services.AppMessageService.Clear()
	fareId := "DUPLICATE_ID"
	fareAttribute := &types.FareAttribute{FareId: &fareId}
	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_rules": {
				"DUPLICATE_ID": []int{1, 2},
			},
		},
	}
	validations.FareIdValidation(fareAttribute, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Duplicate fare_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareIdValidation_ValidFareId(t *testing.T) {
	services.AppMessageService.Clear()
	fareId := "VALID_ID"
	fareAttribute := &types.FareAttribute{FareId: &fareId}
	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_rules": {
				"VALID_ID": []int{1},
			},
		},
	}
	validations.FareIdValidation(fareAttribute, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid fare_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
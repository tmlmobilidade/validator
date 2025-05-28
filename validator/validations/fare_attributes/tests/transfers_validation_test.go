package fare_attributes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_attributes/validations"
	"testing"
)

func TestTransfersValidation_MissingTransfers(t *testing.T) {
	services.AppMessageService.Clear()
	fareAttribute := &types.FareAttribute{Transfers: nil}
	gtfs := &types.Gtfs{}
	validations.TransfersValidation(fareAttribute, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing transfers should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTransfersValidation_InvalidTransfers(t *testing.T) {
	services.AppMessageService.Clear()
	invalid := 5
	fareAttribute := &types.FareAttribute{Transfers: &invalid}
	gtfs := &types.Gtfs{}
	validations.TransfersValidation(fareAttribute, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid transfers should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTransfersValidation_ValidTransfers0(t *testing.T) {
	services.AppMessageService.Clear()
	val := 0
	fareAttribute := &types.FareAttribute{Transfers: &val}
	gtfs := &types.Gtfs{}
	validations.TransfersValidation(fareAttribute, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid transfers 0 should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTransfersValidation_ValidTransfers1(t *testing.T) {
	services.AppMessageService.Clear()
	val := 1
	fareAttribute := &types.FareAttribute{Transfers: &val}
	gtfs := &types.Gtfs{}
	validations.TransfersValidation(fareAttribute, 4, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid transfers 1 should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTransfersValidation_ValidTransfers2(t *testing.T) {
	services.AppMessageService.Clear()
	val := 2
	fareAttribute := &types.FareAttribute{Transfers: &val}
	gtfs := &types.Gtfs{}
	validations.TransfersValidation(fareAttribute, 5, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid transfers 2 should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
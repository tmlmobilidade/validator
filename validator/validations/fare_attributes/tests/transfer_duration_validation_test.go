package fare_attributes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_attributes/validations"
	"testing"
)

func TestTransferDurationValidation_MissingRequired(t *testing.T) {
	services.AppMessageService.Clear()
	fareAttribute := &types.FareAttribute{TransferDuration: nil}
	severity := types.SEVERITY_ERROR
	gtfs := &types.Gtfs{}
	validations.TransferDurationValidation(fareAttribute, 1, gtfs, &types.FareAttributesRules{TransferDuration: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing transfer_duration with error severity should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTransferDurationValidation_MissingWarning(t *testing.T) {
	services.AppMessageService.Clear()
	fareAttribute := &types.FareAttribute{TransferDuration: nil}
	severity := types.SEVERITY_WARNING
	gtfs := &types.Gtfs{}
	validations.TransferDurationValidation(fareAttribute, 2, gtfs, &types.FareAttributesRules{TransferDuration: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Missing transfer_duration with warning severity should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTransferDurationValidation_MissingIgnore(t *testing.T) {
	services.AppMessageService.Clear()
	fareAttribute := &types.FareAttribute{TransferDuration: nil}
	severity := types.SEVERITY_IGNORE
	gtfs := &types.Gtfs{}
	validations.TransferDurationValidation(fareAttribute, 3, gtfs, &types.FareAttributesRules{TransferDuration: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing transfer_duration with ignore severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTransferDurationValidation_NegativeDuration(t *testing.T) {
	services.AppMessageService.Clear()
	dur := -10
	fareAttribute := &types.FareAttribute{TransferDuration: &dur}
	severity := types.SEVERITY_ERROR
	gtfs := &types.Gtfs{}
	validations.TransferDurationValidation(fareAttribute, 4, gtfs, &types.FareAttributesRules{TransferDuration: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Negative transfer_duration should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTransferDurationValidation_ValidDuration(t *testing.T) {
	services.AppMessageService.Clear()
	dur := 3600
	fareAttribute := &types.FareAttribute{TransferDuration: &dur}
	severity := types.SEVERITY_ERROR
	gtfs := &types.Gtfs{}
	validations.TransferDurationValidation(fareAttribute, 5, gtfs, &types.FareAttributesRules{TransferDuration: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid transfer_duration should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

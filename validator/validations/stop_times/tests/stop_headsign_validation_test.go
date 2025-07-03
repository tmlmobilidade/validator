package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestStopHeadsignValidation_MissingStopHeadsign_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{StopHeadsign: nil}
	severity := types.SEVERITY_ERROR
	validations.StopHeadsignValidation(stopTime, 1, &types.StopTimesRules{StopHeadsign: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing stop_headsign with severity error should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopHeadsignValidation_MissingStopHeadsign_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{StopHeadsign: nil}
	severity := types.SEVERITY_WARNING
	validations.StopHeadsignValidation(stopTime, 2, &types.StopTimesRules{StopHeadsign: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Missing stop_headsign with severity warning should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopHeadsignValidation_MissingStopHeadsign_SeverityIgnore(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{StopHeadsign: nil}
	severity := types.SEVERITY_IGNORE
	validations.StopHeadsignValidation(stopTime, 3, &types.StopTimesRules{StopHeadsign: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Missing stop_headsign with severity ignore should not error or warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopHeadsignValidation_PresentStopHeadsign(t *testing.T) {
	services.AppMessageService.Clear()
	headsign := "Downtown"
	stopTime := &types.StopTime{StopHeadsign: &headsign}

	validations.StopHeadsignValidation(stopTime, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Present stop_headsign should not error or warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

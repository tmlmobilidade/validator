package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestHasNetworkMapValidation_MissingHasNetworkMap_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasNetworkMap: nil}
	validations.HasNetworkMapValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_network_map with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasNetworkMapValidation_MissingHasNetworkMap_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasNetworkMap: nil}
	severity := types.SEVERITY_ERROR
	validations.HasNetworkMapValidation(stop, 2, &types.StopsRules{HasNetworkMap: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing has_network_map with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestHasNetworkMapValidation_MissingHasNetworkMap_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{HasNetworkMap: nil}
	severity := types.SEVERITY_WARNING
	validations.HasNetworkMapValidation(stop, 3, &types.StopsRules{HasNetworkMap: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing has_network_map with severity WARNING should warn")
	}
}

func TestHasNetworkMapValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	val := true
	stop := &types.Stop{HasNetworkMap: &val}
	validations.HasNetworkMapValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid has_network_map should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestRegionIdValidation_MissingRegionId_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{RegionId: nil}
	validations.RegionIdValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing region_id with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRegionIdValidation_MissingRegionId_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{RegionId: nil}
	severity := types.SEVERITY_ERROR
	validations.RegionIdValidation(stop, 2, &types.StopsRules{RegionId: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing region_id with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestRegionIdValidation_MissingRegionId_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{RegionId: nil}
	severity := types.SEVERITY_WARNING
	validations.RegionIdValidation(stop, 3, &types.StopsRules{RegionId: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing region_id with severity WARNING should warn")
	}
}

func TestRegionIdValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	id := "REG123"
	stop := &types.Stop{RegionId: &id}
	validations.RegionIdValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid region_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

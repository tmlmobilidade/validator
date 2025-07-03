package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestLevelIdValidation_MissingLevelId_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{LevelId: nil}
	gtfs := types.Gtfs{IdMap: map[string]map[string][]int{"levels": {}}}
	validations.LevelIdValidation(stop, 1, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing level_id with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestLevelIdValidation_MissingLevelId_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{LevelId: nil}
	gtfs := types.Gtfs{IdMap: map[string]map[string][]int{"levels": {}}}
	severity := types.SEVERITY_ERROR
	validations.LevelIdValidation(stop, 2, gtfs, &types.StopsRules{LevelId: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing level_id with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestLevelIdValidation_MissingLevelId_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{LevelId: nil}
	gtfs := types.Gtfs{IdMap: map[string]map[string][]int{"levels": {}}}
	severity := types.SEVERITY_WARNING
	validations.LevelIdValidation(stop, 3, gtfs, &types.StopsRules{LevelId: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing level_id with severity WARNING should warn")
	}
}

func TestLevelIdValidation_InvalidForeignKey(t *testing.T) {
	services.AppMessageService.Clear()
	levelId := "L_INVALID"
	stop := &types.Stop{LevelId: &levelId}
	gtfs := types.Gtfs{IdMap: map[string]map[string][]int{"levels": {}}}
	validations.LevelIdValidation(stop, 4, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid level_id foreign key should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestLevelIdValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	levelId := "L1"
	stop := &types.Stop{LevelId: &levelId}
	gtfs := types.Gtfs{IdMap: map[string]map[string][]int{"levels": {levelId: {1}}}}
	validations.LevelIdValidation(stop, 5, gtfs, &types.StopsRules{LevelId: types.RuleConfig{Severity: types.SEVERITY_IGNORE}})
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid level_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

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
	validations.LevelIdValidation(nil, stop, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing level_id with default severity should not error",
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
	validations.LevelIdValidation(&severity, stop, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing level_id with severity ERROR should error",
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
	validations.LevelIdValidation(&severity, stop, 3, gtfs)
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing level_id with severity WARNING should warn")
	}
}

func TestLevelIdValidation_InvalidForeignKey(t *testing.T) {
	services.AppMessageService.Clear()
	levelId := "L_INVALID"
	stop := &types.Stop{LevelId: &levelId}
	gtfs := types.Gtfs{IdMap: map[string]map[string][]int{"levels": {}}}
	validations.LevelIdValidation(nil, stop, 4, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid level_id foreign key should error",
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
	validations.LevelIdValidation(nil, stop, 5, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid level_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
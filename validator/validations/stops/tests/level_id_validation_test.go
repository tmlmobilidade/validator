package stops

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestAllLevelIdValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericForeignKeyTestCases("level_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var levelId *string
			if tc.Id != nil {
				levelId = tc.Id
			}

			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"levels": {*tc.Id: {1}}}}.ToGtfs()
			if tc.Name == "ForeignKey_Invalid" {
				gtfs = test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"levels": {*tc.Id: {}}}}.ToGtfs()
			}
			validations.LevelIdValidation(&types.Stop{LevelId: levelId}, tc.Row, gtfs, &types.StopsRules{LevelId: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
		})
	}
	for _, tc := range test_helpers.GetGenericSeverityTestCases("level_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"levels": {"LEVEL1": {1}}}}.ToGtfs()
			validations.LevelIdValidation(&types.Stop{}, tc.Row, gtfs, &types.StopsRules{LevelId: types.RuleConfig{Severity: tc.Severity}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name, types.SEVERITY_WARNING)
		})
	}
	t.Run("Default_Severity", func(t *testing.T) {
		services.AppMessageService.Clear()
		gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"levels": {"LEVEL1": {}}}}.ToGtfs()
		validations.LevelIdValidation(&types.Stop{}, 1, gtfs, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Default severity should not error", types.SEVERITY_ERROR)
	})
}

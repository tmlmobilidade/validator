package stops

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestAllParentStationValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericForeignKeyTestCases("parent_station") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var parentStation *string
			if tc.Id != nil {
				parentStation = tc.Id
			}

			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": {*parentStation: {1}}}}.ToGtfs()
			if tc.Name == "ForeignKey_Invalid" {
				gtfs = test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": {*parentStation: {}}}}.ToGtfs()
			}
			stop := &types.Stop{ParentStation: parentStation}
			validations.ParentStationValidation(stop, tc.Row, gtfs, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
		})
	}
	for _, tc := range test_helpers.GetGenericSeverityTestCases("parent_station") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			stop := &types.Stop{}
			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": {}}}.ToGtfs()
			validations.ParentStationValidation(stop, tc.Row, gtfs, &types.StopsRules{ParentStation: types.RuleConfig{Severity: tc.Severity}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name, types.SEVERITY_WARNING)
		})
	}
	t.Run("DefaultSeverity", func(t *testing.T) {
		services.AppMessageService.Clear()
		stop := &types.Stop{}
		gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": {}}}.ToGtfs()
		validations.ParentStationValidation(stop, 1, gtfs, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "DefaultSeverity", types.SEVERITY_ERROR)
	})

	t.Run("TestLocationType2", func(t *testing.T) {
		services.AppMessageService.Clear()
		stop := &types.Stop{LocationType: lib.Ptr(2), ParentStation: nil}
		gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": {}}}.ToGtfs()
		validations.ParentStationValidation(stop, 1, gtfs, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "TestLocationType2", types.SEVERITY_ERROR)
	})

	t.Run("TestLocationType3", func(t *testing.T) {
		services.AppMessageService.Clear()
		stop := &types.Stop{LocationType: lib.Ptr(3), ParentStation: nil}
		gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": {}}}.ToGtfs()
		validations.ParentStationValidation(stop, 1, gtfs, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "TestLocationType3", types.SEVERITY_ERROR)
	})

	t.Run("TestLocationType4", func(t *testing.T) {
		services.AppMessageService.Clear()
		stop := &types.Stop{LocationType: lib.Ptr(4), ParentStation: nil}
		gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": {}}}.ToGtfs()
		validations.ParentStationValidation(stop, 1, gtfs, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "TestLocationType4", types.SEVERITY_ERROR)
	})

	t.Run("TestLocationType1_WithParentStation_ShouldError", func(t *testing.T) {
		services.AppMessageService.Clear()
		parentStation := "STATION1"
		stop := &types.Stop{LocationType: lib.Ptr(1), ParentStation: &parentStation}
		gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": {parentStation: {1}}}}.ToGtfs()
		validations.ParentStationValidation(stop, 1, gtfs, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "TestLocationType1_WithParentStation_ShouldError", types.SEVERITY_ERROR)
	})

	t.Run("TestLocationType0_WithParentStation_ValidForeignKey", func(t *testing.T) {
		services.AppMessageService.Clear()
		parentStation := "STATION1"
		stop := &types.Stop{LocationType: lib.Ptr(0), ParentStation: &parentStation}
		gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": {parentStation: {1}}}}.ToGtfs()
		validations.ParentStationValidation(stop, 1, gtfs, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "TestLocationType0_WithParentStation_ValidForeignKey", types.SEVERITY_ERROR)
	})

	t.Run("TestLocationType2_WithValidParentStation", func(t *testing.T) {
		services.AppMessageService.Clear()
		parentStation := "STATION1"
		stop := &types.Stop{LocationType: lib.Ptr(2), ParentStation: &parentStation}
		gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": {parentStation: {1}}}}.ToGtfs()
		validations.ParentStationValidation(stop, 1, gtfs, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "TestLocationType2_WithValidParentStation", types.SEVERITY_ERROR)
	})
}

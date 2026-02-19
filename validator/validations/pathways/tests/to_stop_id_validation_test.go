package pathways_tests

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/pathways/validations"
	"testing"
)

func TestAllToStopIdValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericForeignKeyTestCases("to_stop_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var toStopId *string
			if tc.Id != nil {
				toStopId = tc.Id
			}

			pathways := &types.Pathways{ToStopId: toStopId}

			var stopsIdMap map[string][]int
			if tc.Name == "ForeignKey_Invalid" {
				stopsIdMap = map[string][]int{}
			} else if toStopId != nil {
				stopsIdMap = map[string][]int{*toStopId: {1}}
			} else {
				stopsIdMap = map[string][]int{}
			}

			gtfs, cleanup, err := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": stopsIdMap}}.ToGtfsWithDB()
			if err != nil {
				t.Fatalf("failed to create mock gtfs: %v", err)
			}
			defer cleanup()

			validations.ToStopIdValidation(pathways, tc.Row, gtfs, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
		})
	}
}

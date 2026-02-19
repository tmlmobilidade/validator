package pathways_tests

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/pathways/validations"
	"testing"
)

func TestAllFromStopIdValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericForeignKeyTestCases("from_stop_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var fromStopId *string
			if tc.Id != nil {
				fromStopId = tc.Id
			}

			pathways := &types.Pathways{FromStopId: fromStopId}

			var stopsIdMap map[string][]int
			if tc.Name == "ForeignKey_Invalid" {
				stopsIdMap = map[string][]int{}
			} else if fromStopId != nil {
				stopsIdMap = map[string][]int{*fromStopId: {1}}
			} else {
				stopsIdMap = map[string][]int{}
			}

			gtfs, cleanup, err := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": stopsIdMap}}.ToGtfsWithDB()
			if err != nil {
				t.Fatalf("failed to create mock gtfs: %v", err)
			}
			defer cleanup()

			validations.FromStopIdValidation(pathways, tc.Row, gtfs, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
		})
	}
}

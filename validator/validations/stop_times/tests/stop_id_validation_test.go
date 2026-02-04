package stop_times

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestAllStopIdValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericForeignKeyTestCases("stop_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var stopTime *types.StopTime
			if tc.Name == "ForeignKey_Invalid" {
				stopTime = &types.StopTime{}
			} else if tc.Id != nil {
				stopTime = &types.StopTime{StopId: tc.Id}
			}

			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"stops": map[string][]int{*tc.Id: {1}}}}.ToGtfs()
			validations.StopIdValidation(stopTime, tc.Row, &gtfs, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
	t.Run("InvalidLocationType", func(t *testing.T) {
		services.AppMessageService.Clear()

		stopId := "S2"
		stopTime := &types.StopTime{
			StopId: &stopId,
		}

		gtfs := &types.Gtfs{
			IdMap: map[string]map[string][]int{
				"stops": {"S2": {0}},
			},
			Stop: []types.StopRaw{
				{StopId: "S2", LocationType: "1"},
			},
		}

		stopLocationTypeCache := map[string]string{"S2": "1"}
		validations.StopIdValidation(stopTime, 5, gtfs, stopLocationTypeCache)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "stop_id with invalid location_type should error")
	})
}

package trips

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestShapeIdValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	trip := &types.Trip{RouteId: lib.Ptr("route1"), TripId: lib.Ptr("trip1"), ShapeId: nil}
	gtfs := &types.Gtfs{
		Route:    []types.RouteRaw{{ContinuousPickup: "1"}},
		StopTime: []types.StopTimeRaw{},
		IdMap:    map[string]map[string][]int{},
	}
	validations.ShapeIdValidation(trip, 1, gtfs, &types.TripsRules{ShapeId: types.RuleConfig{Severity: severity}}, make(map[string][]types.StopTimeRaw))

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "shape_id is required when continuous pickup/dropoff is defined",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestShapeIdValidation_Recommended(t *testing.T) {
	severity := types.SEVERITY_WARNING
	trip := &types.Trip{RouteId: lib.Ptr("route1"), TripId: lib.Ptr("trip1"), ShapeId: nil}
	gtfs := &types.Gtfs{}
	validations.ShapeIdValidation(trip, 2, gtfs, &types.TripsRules{ShapeId: types.RuleConfig{Severity: severity}}, make(map[string][]types.StopTimeRaw))

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "shape_id is recommended",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestShapeIdValidation_Ignore(t *testing.T) {
	severity := types.SEVERITY_IGNORE
	trip := &types.Trip{RouteId: lib.Ptr("route1"), TripId: lib.Ptr("trip1"), ShapeId: nil}
	gtfs := &types.Gtfs{}
	validations.ShapeIdValidation(trip, 3, gtfs, &types.TripsRules{ShapeId: types.RuleConfig{Severity: severity}}, make(map[string][]types.StopTimeRaw))

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "shape_id is ignored, no error or warning should be reported",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestShapeIdValidation_DoesNotExist(t *testing.T) {
	services.AppMessageService.Clear()

	severity := types.SEVERITY_ERROR
	trip := &types.Trip{RouteId: lib.Ptr("route1"), TripId: lib.Ptr("trip1"), ShapeId: lib.Ptr("shape1")}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{},
	}
	validations.ShapeIdValidation(trip, 4, gtfs, &types.TripsRules{ShapeId: types.RuleConfig{Severity: severity}}, make(map[string][]types.StopTimeRaw))

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "shape_id does not exist in the shapes.txt file",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAllShapeIdValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericForeignKeyTestCases("shape_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var shapeId *string
			if tc.Id != nil {
				shapeId = tc.Id
			}

			trip := &types.Trip{
				RouteId: lib.Ptr("route1"),
				TripId:  lib.Ptr("trip1"),
				ShapeId: shapeId,
			}

			var gtfs *types.Gtfs
			if tc.Name == "ForeignKey_Present" && tc.Id != nil {
				gtfsVal := test_helpers.MockGtfs{
					IdMapData: types.GtfsIdMap{
						"shapes": {*tc.Id: []int{1}},
						"routes": {"route1": []int{1}},
					},
				}.ToGtfs()
				gtfs = &gtfsVal
			} else {
				gtfsVal := test_helpers.MockGtfs{}.ToGtfs()
				gtfs = &gtfsVal
			}

			validations.ShapeIdValidation(trip, tc.Row, gtfs, &types.TripsRules{ShapeId: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, make(map[string][]types.StopTimeRaw))
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
		})
	}
	for _, tc := range test_helpers.GetGenericSeverityTestCases("shape_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			trip := &types.Trip{ShapeId: nil}
			gtfsVal := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"shapes": {}}}.ToGtfs()
			gtfs := &gtfsVal
			validations.ShapeIdValidation(trip, tc.Row, gtfs, &types.TripsRules{ShapeId: types.RuleConfig{Severity: tc.Severity}}, make(map[string][]types.StopTimeRaw))
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, tc.Severity)
		})
	}
}

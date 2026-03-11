package trips

import (
	"main/i18n"
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

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

			if tc.Name == "ForeignKey_Invalid" {
				trip = &types.Trip{
					RouteId: lib.Ptr("route1"),
					TripId:  lib.Ptr("trip1"),
					ShapeId: nil,
				}
			}
			gtfs, cleanup, err := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"shapes": {*tc.Id: {1}}, "routes": {"route1": {0}}}}.ToGtfsWithDB()
			if err != nil {
				t.Fatalf("failed to create mock gtfs: %v", err)
			}
			defer cleanup()
			validations.ShapeIdValidation(trip, tc.Row, gtfs, &types.TripsRules{ShapeId: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, make(map[string][]types.StopTimeRaw))
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
		})
	}
	for _, tc := range test_helpers.GetGenericSeverityTestCases("shape_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			trip := &types.Trip{RouteId: lib.Ptr("route1"), TripId: lib.Ptr("trip1"), ShapeId: nil}
			gtfs, cleanup, err := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"routes": {"route1": {0}}}}.ToGtfsWithDB()
			if err != nil {
				t.Fatalf("failed to create mock gtfs: %v", err)
			}
			defer cleanup()
			validations.ShapeIdValidation(trip, tc.Row, gtfs, &types.TripsRules{ShapeId: types.RuleConfig{Severity: tc.Severity}}, make(map[string][]types.StopTimeRaw))
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name, types.SEVERITY_WARNING)
		})
	}
	t.Run("Valid_Cache", func(t *testing.T) {
		services.AppMessageService.Clear()
		trip := &types.Trip{RouteId: lib.Ptr("route1"), TripId: lib.Ptr("trip1"), ShapeId: nil}
		gtfs, cleanup, err := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"routes": {"route1": {0}}, "stop_times": {"trip1": {0}}}}.ToGtfsWithDB()
		if err != nil {
			t.Fatalf("failed to create mock gtfs: %v", err)
		}
		defer cleanup()
		validations.ShapeIdValidation(trip, 1, gtfs, nil, make(map[string][]types.StopTimeRaw))
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Shape ID is required", types.SEVERITY_ERROR)
	})
	t.Run("Invalid_Cache", func(t *testing.T) {
		services.AppMessageService.Clear()
		trip := &types.Trip{RouteId: lib.Ptr("route1"), TripId: lib.Ptr("trip1"), ShapeId: nil}
		gtfs, cleanup, err := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"routes": {"route1": {0}}}}.ToGtfsWithDB()
		if err != nil {
			t.Fatalf("failed to create mock gtfs: %v", err)
		}
		defer cleanup()
		validations.ShapeIdValidation(trip, 1, gtfs, &types.TripsRules{ShapeId: types.RuleConfig{Severity: types.SEVERITY_ERROR}}, make(map[string][]types.StopTimeRaw))
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Shape ID is required", types.SEVERITY_ERROR)
	})

	t.Run("Route_Not_Found_For_Trip", func(t *testing.T) {
		services.AppMessageService.Clear()
		trip := &types.Trip{
			RouteId: lib.Ptr("missing-route"),
			TripId:  lib.Ptr("trip1"),
			ShapeId: lib.Ptr("shape1"),
		}
		gtfs, cleanup, err := test_helpers.MockGtfs{
			IdMapData: types.GtfsIdMap{
				"shapes": {"shape1": {1}},
				"routes": {},
			},
		}.ToGtfsWithDB()
		if err != nil {
			t.Fatalf("failed to create mock gtfs: %v", err)
		}
		defer cleanup()

		validations.ShapeIdValidation(trip, 1, gtfs, nil, make(map[string][]types.StopTimeRaw))
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Missing route should produce one explicit error", types.SEVERITY_ERROR)
		test_helpers.AssertMessageContains(
			t,
			services.AppMessageService,
			i18n.AppTranslator.Get("shape_id_validation.route_not_found_for_trip", "missing-route"),
			"Missing route should produce route-specific error message",
		)
	})

	t.Run("Route_Lookup_Failed", func(t *testing.T) {
		services.AppMessageService.Clear()
		trip := &types.Trip{
			RouteId: lib.Ptr("route1"),
			TripId:  lib.Ptr("trip1"),
			ShapeId: lib.Ptr("shape1"),
		}

		// Empty Gtfs has nil db, forcing GetRowsById to fail and emit a clear lookup message.
		gtfs := &types.Gtfs{}
		validations.ShapeIdValidation(trip, 1, gtfs, nil, make(map[string][]types.StopTimeRaw))
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Lookup failure should produce one explicit error", types.SEVERITY_ERROR)
	})

	t.Run("Panic_Is_Recovered_With_Internal_Error", func(t *testing.T) {
		services.AppMessageService.Clear()
		trip := &types.Trip{
			RouteId: lib.Ptr("route1"),
			TripId:  lib.Ptr("trip1"),
			ShapeId: lib.Ptr("shape1"),
		}

		// Nil GTFS intentionally triggers panic path to assert graceful recovery.
		validations.ShapeIdValidation(trip, 1, nil, nil, make(map[string][]types.StopTimeRaw))
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Panic should be recovered and reported as one error", types.SEVERITY_ERROR)
	})
}

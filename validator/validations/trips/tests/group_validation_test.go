package trips

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func patternGroup(trips ...types.Trip) types.TripGroupedByPattern {
	return types.TripGroupedByPattern{"P1": {Trips: trips, Hash: []string{"same"}}}
}

func TestPatternIdValidation(t *testing.T) {
	t.Run("MissingPatternId", func(t *testing.T) {
		services.AppMessageService.Clear()
		trip := &types.Trip{PatternId: nil}

		validations.PatternIdValidation(trip, 1, nil, &types.TripsRules{PatternId: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "MissingPatternId", types.SEVERITY_ERROR)
	})

	t.Run("ValidPatternId", func(t *testing.T) {
		services.AppMessageService.Clear()
		trip := &types.Trip{PatternId: lib.Ptr("P1")}

		validations.PatternIdValidation(trip, 1, nil, &types.TripsRules{PatternId: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "ValidPatternId", types.SEVERITY_ERROR)
	})
}

func TestPatternIdGroupValidation(t *testing.T) {
	t.Run("MissingRequiredGroupingField", func(t *testing.T) {
		services.AppMessageService.Clear()
		groups := patternGroup(types.Trip{PatternId: lib.Ptr("P1"), Row: 7})

		validations.PatternIdGroupValidation(groups, nil, &types.TripsRules{PatternIdTripHasRequiredFieldsForGrouping: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "MissingRequiredGroupingField", types.SEVERITY_ERROR)
	})

	t.Run("MultipleTripSignatures", func(t *testing.T) {
		services.AppMessageService.Clear()
		groups := types.TripGroupedByPattern{"P1": {Trips: []types.Trip{{PatternId: lib.Ptr("P1"), ShapeId: lib.Ptr("S1"), RouteId: lib.Ptr("R1"), DirectionId: lib.Ptr(0), TripHeadsign: lib.Ptr("A"), Row: 7}}, Hash: []string{"a", "b"}}}

		validations.PatternIdGroupValidation(groups, nil, &types.TripsRules{
			PatternIdTripHasRequiredFieldsForGrouping: types.RuleConfig{Severity: types.SEVERITY_ERROR},
			PatternIdSingleTripSignaturePerPattern:    types.RuleConfig{Severity: types.SEVERITY_ERROR},
		})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "MultipleTripSignatures", types.SEVERITY_ERROR)
	})
}

func TestTripGroupConsistencyValidations(t *testing.T) {
	t.Run("DifferentRouteIdsInPattern", func(t *testing.T) {
		services.AppMessageService.Clear()
		groups := patternGroup(
			types.Trip{PatternId: lib.Ptr("P1"), RouteId: lib.Ptr("R1"), Row: 1},
			types.Trip{PatternId: lib.Ptr("P1"), RouteId: lib.Ptr("R2"), Row: 2},
		)

		validations.RouteIdGroupValidation(groups, nil, &types.TripsRules{RouteIdGroup: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "DifferentRouteIdsInPattern", types.SEVERITY_ERROR)
	})

	t.Run("DifferentDirectionIdsInPattern", func(t *testing.T) {
		services.AppMessageService.Clear()
		groups := patternGroup(
			types.Trip{PatternId: lib.Ptr("P1"), DirectionId: lib.Ptr(0), Row: 1},
			types.Trip{PatternId: lib.Ptr("P1"), DirectionId: lib.Ptr(1), Row: 2},
		)

		validations.DirectionIdGroupValidation(groups, nil, &types.TripsRules{DirectionIdGroup: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "DifferentDirectionIdsInPattern", types.SEVERITY_ERROR)
	})

	t.Run("DifferentHeadsignsInPattern", func(t *testing.T) {
		services.AppMessageService.Clear()
		groups := patternGroup(
			types.Trip{PatternId: lib.Ptr("P1"), TripHeadsign: lib.Ptr("A"), Row: 1},
			types.Trip{PatternId: lib.Ptr("P1"), TripHeadsign: lib.Ptr("B"), Row: 2},
		)

		validations.TripHeadsignGroupValidation(groups, nil, &types.TripsRules{TripHeadsignGroup: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "DifferentHeadsignsInPattern", types.SEVERITY_ERROR)
	})
}

func TestShapeIdGroupValidation(t *testing.T) {
	t.Run("MultipleShapeIdsForPattern", func(t *testing.T) {
		services.AppMessageService.Clear()
		patternGroups := patternGroup(
			types.Trip{PatternId: lib.Ptr("P1"), ShapeId: lib.Ptr("S1"), Row: 1},
			types.Trip{PatternId: lib.Ptr("P1"), ShapeId: lib.Ptr("S2"), Row: 2},
		)

		validations.ShapeIdGroupValidation(patternGroups, types.TripGroupedByShapeId{}, nil, &types.TripsRules{OneShapeIdPerPatternIdGroup: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "MultipleShapeIdsForPattern", types.SEVERITY_ERROR)
	})

	t.Run("MultiplePatternIdsForShape", func(t *testing.T) {
		services.AppMessageService.Clear()
		shapeGroups := types.TripGroupedByShapeId{"S1": {Trips: []types.Trip{
			{PatternId: lib.Ptr("P1"), ShapeId: lib.Ptr("S1"), Row: 1},
			{PatternId: lib.Ptr("P2"), ShapeId: lib.Ptr("S1"), Row: 2},
		}}}

		validations.ShapeIdGroupValidation(types.TripGroupedByPattern{}, shapeGroups, nil, &types.TripsRules{OnePatternIdPerShapeIdGroup: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "MultiplePatternIdsForShape", types.SEVERITY_ERROR)
	})
}

func TestShapeIdSamePatternIdValidation(t *testing.T) {
	services.AppMessageService.Clear()
	trip := &types.Trip{PatternId: lib.Ptr("P1"), ShapeId: lib.Ptr("S1")}

	validations.ShapeIdSamePatternIdValidation(trip, 1, nil, &types.TripsRules{ShapeIdSamePatternId: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "ShapeIdSamePatternIdValidation", types.SEVERITY_ERROR)
}

func TestValidatePatternGroups(t *testing.T) {
	services.AppMessageService.Clear()
	patternGroups := patternGroup(
		types.Trip{PatternId: lib.Ptr("P1"), ShapeId: lib.Ptr("S1"), RouteId: lib.Ptr("R1"), DirectionId: lib.Ptr(0), TripHeadsign: lib.Ptr("A"), Row: 1},
		types.Trip{PatternId: lib.Ptr("P1"), ShapeId: lib.Ptr("S2"), RouteId: lib.Ptr("R2"), DirectionId: lib.Ptr(1), TripHeadsign: lib.Ptr("B"), Row: 2},
	)
	shapeGroups := types.TripGroupedByShapeId{}
	rules := &types.TripsRules{
		PatternIdTripHasRequiredFieldsForGrouping: types.RuleConfig{Severity: types.SEVERITY_ERROR},
		PatternIdSingleTripSignaturePerPattern:    types.RuleConfig{Severity: types.SEVERITY_ERROR},
		RouteIdGroup:                              types.RuleConfig{Severity: types.SEVERITY_ERROR},
		DirectionIdGroup:                          types.RuleConfig{Severity: types.SEVERITY_ERROR},
		OneShapeIdPerPatternIdGroup:               types.RuleConfig{Severity: types.SEVERITY_ERROR},
		TripHeadsignGroup:                         types.RuleConfig{Severity: types.SEVERITY_ERROR},
	}

	validations.ValidatePatternGroups(patternGroups, shapeGroups, nil, rules)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 4, "ValidatePatternGroups", types.SEVERITY_ERROR)
}

func TestStopCoordinatesByTripIdValidation(t *testing.T) {
	t.Run("CachedStopTooFarFromShape", func(t *testing.T) {
		services.AppMessageService.Clear()
		trip := &types.Trip{TripId: lib.Ptr("T1"), ShapeId: lib.Ptr("SHP1")}
		stopTimes := map[string][]types.StopTimeRaw{"T1": {{TripId: "T1", StopId: "STOP1"}}}
		stops := map[string]types.StopCoordinatesValidation{"STOP1": {StopId: "STOP1", StopLat: "38.7", StopLon: "-9.1"}}
		closest := map[string]types.StopClosestShapePointsInfo{
			"STOP1|SHP1": {ShapeID: "SHP1", DistanceMeters: 1000, ClosestShapePtLat: 38.8, ClosestShapePtLon: -9.1, ClosestShapePtSeq: 1},
		}

		result := validations.StopCoordinatesByTripIdValidation(trip, 1, nil, stopTimes, stops, closest, &types.TripsRules{StopCoordinatesByTripId: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		if len(result) != 1 {
			t.Fatalf("expected one returned stop coordinate, got %d", len(result))
		}
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "CachedStopTooFarFromShape", types.SEVERITY_ERROR)
	})
}

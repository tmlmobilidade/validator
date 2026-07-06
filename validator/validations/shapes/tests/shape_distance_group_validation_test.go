package shapes

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/shapes/validations"
	"testing"
)

func TestShapeDistancesValidation(t *testing.T) {
	t.Run("Valid_BlockDistance", func(t *testing.T) {
		services.AppMessageService.Clear()
		shapes := []types.Shape{
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(0), ShapePtLat: lib.Ptr(float32(38.7000)), ShapePtLon: lib.Ptr(float32(-9.1000)), ShapeDistTraveled: lib.Ptr(0.0)},
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(1), ShapePtLat: lib.Ptr(float32(38.7090)), ShapePtLon: lib.Ptr(float32(-9.1000)), ShapeDistTraveled: lib.Ptr(1.0)},
		}

		validations.ShapeDistancesValidation(shapes, &types.ShapesRules{ShapeDistTraveledDeltaMismatchesHaversineBlock: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Valid_BlockDistance", types.SEVERITY_ERROR)
	})

	t.Run("Invalid_BlockDistance", func(t *testing.T) {
		services.AppMessageService.Clear()
		shapes := []types.Shape{
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(0), ShapePtLat: lib.Ptr(float32(38.7000)), ShapePtLon: lib.Ptr(float32(-9.1000)), ShapeDistTraveled: lib.Ptr(0.0)},
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(1), ShapePtLat: lib.Ptr(float32(38.7090)), ShapePtLon: lib.Ptr(float32(-9.1000)), ShapeDistTraveled: lib.Ptr(1.0)},
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(2), ShapePtLat: lib.Ptr(float32(38.7180)), ShapePtLon: lib.Ptr(float32(-9.1000)), ShapeDistTraveled: lib.Ptr(2.0)},
		}

		validations.ShapeDistancesValidation(shapes, &types.ShapesRules{ShapeDistTraveledDeltaMismatchesHaversineBlock: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Invalid_BlockDistance", types.SEVERITY_ERROR)
	})
}

func TestShapePointsCoordinatesConsistentValidation(t *testing.T) {
	t.Run("Valid_CloseConsecutivePoints", func(t *testing.T) {
		services.AppMessageService.Clear()
		shapes := []types.Shape{
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(0), ShapePtLat: lib.Ptr(float32(38.7000)), ShapePtLon: lib.Ptr(float32(-9.1000))},
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(1), ShapePtLat: lib.Ptr(float32(38.7001)), ShapePtLon: lib.Ptr(float32(-9.1001))},
		}

		validations.ShapePointsCoordinatesConsistentValidation(shapes, &types.ShapesRules{ShapePointsCoordinatesConsistent: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Valid_CloseConsecutivePoints", types.SEVERITY_ERROR)
	})

	t.Run("Invalid_FarConsecutivePoints", func(t *testing.T) {
		services.AppMessageService.Clear()
		shapes := []types.Shape{
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(0), ShapePtLat: lib.Ptr(float32(38.7000)), ShapePtLon: lib.Ptr(float32(-9.1000))},
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(1), ShapePtLat: lib.Ptr(float32(40.0000)), ShapePtLon: lib.Ptr(float32(-9.1000))},
		}

		validations.ShapePointsCoordinatesConsistentValidation(shapes, &types.ShapesRules{ShapePointsCoordinatesConsistent: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Invalid_FarConsecutivePoints", types.SEVERITY_ERROR)
	})
}

func TestShapePointsCoordinatesDistancesValidation(t *testing.T) {
	t.Run("Valid_SegmentDistance", func(t *testing.T) {
		services.AppMessageService.Clear()
		shapes := []types.Shape{
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(0), ShapePtLat: lib.Ptr(float32(38.7000)), ShapePtLon: lib.Ptr(float32(-9.1000)), ShapeDistTraveled: lib.Ptr(0.0)},
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(1), ShapePtLat: lib.Ptr(float32(38.7090)), ShapePtLon: lib.Ptr(float32(-9.1000)), ShapeDistTraveled: lib.Ptr(1.0)},
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(2), ShapePtLat: lib.Ptr(float32(38.7180)), ShapePtLon: lib.Ptr(float32(-9.1000)), ShapeDistTraveled: lib.Ptr(2.0)},
		}

		validations.ShapePointsCoordinatesDistancesValidation(shapes, &types.ShapesRules{ShapePointsCoordinatesDistances: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Valid_SegmentDistance", types.SEVERITY_ERROR)
	})

	t.Run("Invalid_SegmentDistance", func(t *testing.T) {
		services.AppMessageService.Clear()
		shapes := []types.Shape{
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(0), ShapePtLat: lib.Ptr(float32(38.7000)), ShapePtLon: lib.Ptr(float32(-9.1000)), ShapeDistTraveled: lib.Ptr(0.0)},
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(1), ShapePtLat: lib.Ptr(float32(38.7090)), ShapePtLon: lib.Ptr(float32(-9.1000)), ShapeDistTraveled: lib.Ptr(1.0)},
			{ShapeId: lib.Ptr("A"), ShapePtSequence: lib.Ptr(2), ShapePtLat: lib.Ptr(float32(38.9000)), ShapePtLon: lib.Ptr(float32(-9.1000)), ShapeDistTraveled: lib.Ptr(2.0)},
		}

		validations.ShapePointsCoordinatesDistancesValidation(shapes, &types.ShapesRules{ShapePointsCoordinatesDistances: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Invalid_SegmentDistance", types.SEVERITY_ERROR)
	})
}

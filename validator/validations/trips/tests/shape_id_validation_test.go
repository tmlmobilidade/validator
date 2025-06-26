package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestShapeIdValidation_Required(t *testing.T) {
	severity := types.SEVERITY_ERROR
	trip := &types.Trip{RouteId: lib.Ptr("route1"), TripId: lib.Ptr("trip1"), ShapeId: nil}
	gtfs := &types.Gtfs{
		Route: []types.RouteRaw{{ContinuousPickup: "1"}},
		StopTime: []types.StopTimeRaw{},
		IdMap: map[string]map[string][]int{},
	}
	validations.ShapeIdValidation(&severity, trip, 1, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "shape_id is required when continuous pickup/dropoff is defined",
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
	validations.ShapeIdValidation(&severity, trip, 2, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "shape_id is recommended",
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
	validations.ShapeIdValidation(&severity, trip, 3, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message: "shape_id is ignored, no error or warning should be reported",
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
	validations.ShapeIdValidation(&severity, trip, 4, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "shape_id does not exist in the shapes.txt file",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
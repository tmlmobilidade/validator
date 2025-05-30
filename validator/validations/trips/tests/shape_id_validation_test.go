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
		Files: map[string][]map[string]string{
			"routes": {{"continuous_pickup": "1"}},
			"stop_times": {},
		},
		IdMap: map[string]map[string][]int{
			"routes": {"route1": {0}},
		},
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
	shape := "shape1"
	trip := &types.Trip{RouteId: lib.Ptr("route1"), TripId: lib.Ptr("trip1"), ShapeId: &shape}
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
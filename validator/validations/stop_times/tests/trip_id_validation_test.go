package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestTripIdValidation_MissingTripId(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{TripId: nil}
	gtfs := &types.Gtfs{}
	validations.TripIdValidation(stopTime, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing trip_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTripIdValidation_EmptyTripId(t *testing.T) {
	services.AppMessageService.Clear()
	empty := ""
	stopTime := &types.StopTime{TripId: &empty}
	gtfs := &types.Gtfs{}
	validations.TripIdValidation(stopTime, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Empty trip_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTripIdValidation_InvalidForeignKey(t *testing.T) {
	services.AppMessageService.Clear()
	tripId := "invalid_trip"
	stopTime := &types.StopTime{TripId: &tripId}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"trips": {},
		},
	}
	validations.TripIdValidation(stopTime, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid trip_id foreign key should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestTripIdValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	tripId := "valid_trip"
	stopTime := &types.StopTime{TripId: &tripId}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"trips": {
				"valid_trip": {0},
			},
		},
	}
	validations.TripIdValidation(stopTime, 4, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid trip_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
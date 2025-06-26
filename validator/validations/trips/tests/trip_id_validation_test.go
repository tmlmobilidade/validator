package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestTripIdValidation_Required(t *testing.T) {
	trip := &types.Trip{TripId: nil}
	gtfs := &types.Gtfs{Trip: []types.TripRaw{{}, {}}} // >1 trip
	validations.TripIdValidation(trip, 1, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Trip ID is required when there is more than one trip",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestTripIdValidation_Unique(t *testing.T) {
	
	trip := &types.Trip{TripId: lib.Ptr("unique")}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"trips": {"unique": {1}}}}

	validations.TripIdValidation(trip, 2, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Trip ID should be unique, no error for unique trip_id",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestTripIdValidation_Duplicate(t *testing.T) {
	trip := &types.Trip{TripId: lib.Ptr("duplicate")}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"trips": {"duplicate": {1, 2}}}}
	validations.TripIdValidation(trip, 3, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Duplicate trip_id found. Trip IDs must be unique.",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
} 
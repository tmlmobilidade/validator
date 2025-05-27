package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestStopIdValidation_MissingStopId(t *testing.T) {
	services.AppMessageService.Clear()
	empty := ""
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {},
		},
	}
	stop := &types.Stop{StopId: &empty}
	validations.StopIdValidation(stop, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing stop_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopIdValidation_DuplicateStopId_Error(t *testing.T) {
	services.AppMessageService.Clear()
	stopId := "S1"
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {
				stopId: {1, 1},
			},
		},
	}
	stop := &types.Stop{StopId: &stopId}
	validations.StopIdValidation(stop, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Duplicate stop_id should not error, because it is not a duplicate in different rows",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopIdValidation_DuplicateStopId(t *testing.T) {
	services.AppMessageService.Clear()
	stopId := "S1"
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {
				stopId: {1, 2},
			},
		},
	}
	stop := &types.Stop{StopId: &stopId}
	validations.StopIdValidation(stop, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Duplicate stop_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopIdValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	stopId := "S1"
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {
				stopId: {1},
			},
		},
	}
	stop := &types.Stop{StopId: &stopId}
	validations.StopIdValidation(stop, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid stop_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
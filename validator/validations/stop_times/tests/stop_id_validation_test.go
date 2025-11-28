package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestStopIdValidation_Required(t *testing.T) {
	services.AppMessageService.Clear()
	stopId := "S1"
	stopTime := &types.StopTime{
		StopId: &stopId,
	}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {"S1": {0}},
		},
		Stop: []types.StopRaw{
			{StopId: "S1", LocationType: "0"},
		},
	}
	stopLocationTypeCache := map[string]string{"S1": "0"}
	validations.StopIdValidation(stopTime, 1, gtfs, stopLocationTypeCache)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid stop_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopIdValidation_Forbidden(t *testing.T) {
	services.AppMessageService.Clear()
	stopId := "S1"
	locationGroupId := "LG1"
	stopTime := &types.StopTime{
		StopId:          &stopId,
		LocationGroupId: &locationGroupId,
	}
	gtfs := &types.Gtfs{}
	stopLocationTypeCache := map[string]string{}
	validations.StopIdValidation(stopTime, 2, gtfs, stopLocationTypeCache)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "stop_id should be forbidden if location_group_id is defined",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopIdValidation_RequiredMissing(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	gtfs := &types.Gtfs{}
	stopLocationTypeCache := map[string]string{}
	validations.StopIdValidation(stopTime, 3, gtfs, stopLocationTypeCache)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing stop_id should error if location_group_id and location_id are not defined",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopIdValidation_InvalidForeignKey(t *testing.T) {
	services.AppMessageService.Clear()
	stopId := "INVALID"
	stopTime := &types.StopTime{
		StopId: &stopId,
	}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {},
		},
	}
	stopLocationTypeCache := map[string]string{}
	validations.StopIdValidation(stopTime, 4, gtfs, stopLocationTypeCache)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid stop_id foreign key should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopIdValidation_InvalidLocationType(t *testing.T) {
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
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "stop_id with invalid location_type should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

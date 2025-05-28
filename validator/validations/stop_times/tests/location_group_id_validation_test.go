package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestLocationGroupIdValidation_ForbiddenWithStopId(t *testing.T) {
	services.AppMessageService.Clear()
	stopId := "S1"
	locationGroupId := "LG1"
	stopTime := &types.StopTime{
		StopId: &stopId,
		LocationGroupId: &locationGroupId,
	}
	gtfs := &types.Gtfs{}
	validations.LocationGroupIdValidation(stopTime, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "location_group_id should be forbidden if stop_id is defined",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestLocationGroupIdValidation_ForbiddenWithLocationId(t *testing.T) {
	services.AppMessageService.Clear()
	locationId := "L1"
	locationGroupId := "LG1"
	stopTime := &types.StopTime{
		LocationId: &locationId,
		LocationGroupId: &locationGroupId,
	}
	gtfs := &types.Gtfs{}
	validations.LocationGroupIdValidation(stopTime, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "location_group_id should be forbidden if location_id is defined",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestLocationGroupIdValidation_InvalidForeignKey(t *testing.T) {
	services.AppMessageService.Clear()
	locationGroupId := "INVALID"
	stopTime := &types.StopTime{
		LocationGroupId: &locationGroupId,
	}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"location_groups": {},
		},
	}
	validations.LocationGroupIdValidation(stopTime, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid location_group_id foreign key should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestLocationGroupIdValidation_MissingLocationGroupsIndex(t *testing.T) {
	services.AppMessageService.Clear()
	locationGroupId := "LG1"
	stopTime := &types.StopTime{
		LocationGroupId: &locationGroupId,
	}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{},
	}
	validations.LocationGroupIdValidation(stopTime, 4, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing location_groups index should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestLocationGroupIdValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	locationGroupId := "LG1"
	stopTime := &types.StopTime{
		LocationGroupId: &locationGroupId,
	}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"location_groups": {
				"LG1": {0},
			},
		},
	}
	validations.LocationGroupIdValidation(stopTime, 5, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid location_group_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
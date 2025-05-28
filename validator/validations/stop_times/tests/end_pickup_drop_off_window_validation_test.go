package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestEndPickupDropOffWindowValidation_RequiredWithLocationGroupId(t *testing.T) {
	services.AppMessageService.Clear()
	locationGroupId := "LG1"
	stopTime := &types.StopTime{
		LocationGroupId: &locationGroupId,
	}
	gtfs := &types.Gtfs{}
	validations.EndPickupDropOffWindowValidation(nil, stopTime, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing end_pickup_drop_off_window with location_group_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_RequiredWithLocationId(t *testing.T) {
	services.AppMessageService.Clear()
	locationId := "L1"
	stopTime := &types.StopTime{
		LocationId: &locationId,
	}
	gtfs := &types.Gtfs{}
	validations.EndPickupDropOffWindowValidation(nil, stopTime, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing end_pickup_drop_off_window with location_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_RequiredWithStartPickupDropOffWindow(t *testing.T) {
	services.AppMessageService.Clear()
	startWindow := "07:00:00"
	stopTime := &types.StopTime{
		StartPickupDropOffWindow: &startWindow,
	}
	gtfs := &types.Gtfs{}
	validations.EndPickupDropOffWindowValidation(nil, stopTime, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing end_pickup_drop_off_window with start_pickup_drop_off_window should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_ForbiddenWithArrivalTime(t *testing.T) {
	services.AppMessageService.Clear()
	arrival := "08:00:00"
	endWindow := "10:00:00"
	stopTime := &types.StopTime{
		ArrivalTime: &arrival,
		EndPickupDropOffWindow: &endWindow,
	}
	gtfs := &types.Gtfs{}
	validations.EndPickupDropOffWindowValidation(nil, stopTime, 4, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "end_pickup_drop_off_window should be forbidden if arrival_time is defined",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_ForbiddenWithDepartureTime(t *testing.T) {
	services.AppMessageService.Clear()
	departure := "09:00:00"
	endWindow := "10:00:00"
	stopTime := &types.StopTime{
		DepartureTime: &departure,
		EndPickupDropOffWindow: &endWindow,
	}
	gtfs := &types.Gtfs{}
	validations.EndPickupDropOffWindowValidation(nil, stopTime, 5, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "end_pickup_drop_off_window should be forbidden if departure_time is defined",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	locationGroupId := "LG1"
	endWindow := "10:00:00"
	stopTime := &types.StopTime{
		LocationGroupId: &locationGroupId,
		EndPickupDropOffWindow: &endWindow,
	}
	gtfs := &types.Gtfs{}
	validations.EndPickupDropOffWindowValidation(nil, stopTime, 6, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid end_pickup_drop_off_window should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_InvalidTime(t *testing.T) {
	services.AppMessageService.Clear()
	locationGroupId := "LG1"
	endWindow := "INVALID"
	stopTime := &types.StopTime{
		LocationGroupId: &locationGroupId,
		EndPickupDropOffWindow: &endWindow,
	}
	gtfs := &types.Gtfs{}
	validations.EndPickupDropOffWindowValidation(nil, stopTime, 7, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid time for end_pickup_drop_off_window should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestEndPickupDropOffWindowValidation_Optional(t *testing.T) {
	services.AppMessageService.Clear()
	stopTime := &types.StopTime{}
	gtfs := &types.Gtfs{}
	validations.EndPickupDropOffWindowValidation(nil, stopTime, 8, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Optional end_pickup_drop_off_window should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
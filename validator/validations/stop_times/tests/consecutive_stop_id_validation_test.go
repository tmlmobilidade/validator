package stop_times

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stop_times/validations"
	"testing"
)

func TestConsecutiveStopIdValidation_ValidDifferentStopIds(t *testing.T) {
	services.AppMessageService.Clear()
	previousStopIdByTrip := make(map[string]*string)

	tripId1 := "T1"
	stopId1 := "S1"
	stopTime1 := &types.StopTime{
		TripId: &tripId1,
		StopId: &stopId1,
	}
	validations.ConsecutiveStopIdValidation(stopTime1, 1, previousStopIdByTrip)

	tripId2 := "T1"
	stopId2 := "S2"
	stopTime2 := &types.StopTime{
		TripId: &tripId2,
		StopId: &stopId2,
	}
	validations.ConsecutiveStopIdValidation(stopTime2, 2, previousStopIdByTrip)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Different stop_ids consecutively should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestConsecutiveStopIdValidation_InvalidConsecutiveStopIds(t *testing.T) {
	services.AppMessageService.Clear()
	previousStopIdByTrip := make(map[string]*string)

	tripId := "T1"
	stopId1 := "S1"
	stopTime1 := &types.StopTime{
		TripId: &tripId,
		StopId: &stopId1,
	}
	validations.ConsecutiveStopIdValidation(stopTime1, 1, previousStopIdByTrip)

	stopId2 := "S1" // Same stop_id
	stopTime2 := &types.StopTime{
		TripId: &tripId,
		StopId: &stopId2,
	}
	validations.ConsecutiveStopIdValidation(stopTime2, 2, previousStopIdByTrip)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Same stop_id consecutively should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestConsecutiveStopIdValidation_SameStopIdDifferentTrips(t *testing.T) {
	services.AppMessageService.Clear()
	previousStopIdByTrip := make(map[string]*string)

	tripId1 := "T1"
	stopId := "S1"
	stopTime1 := &types.StopTime{
		TripId: &tripId1,
		StopId: &stopId,
	}
	validations.ConsecutiveStopIdValidation(stopTime1, 1, previousStopIdByTrip)

	tripId2 := "T2"
	stopTime2 := &types.StopTime{
		TripId: &tripId2,
		StopId: &stopId, // Same stop_id but different trip
	}
	validations.ConsecutiveStopIdValidation(stopTime2, 2, previousStopIdByTrip)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Same stop_id in different trips should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestConsecutiveStopIdValidation_FirstStopInTrip(t *testing.T) {
	services.AppMessageService.Clear()
	previousStopIdByTrip := make(map[string]*string)

	tripId := "T1"
	stopId := "S1"
	stopTime := &types.StopTime{
		TripId: &tripId,
		StopId: &stopId,
	}
	validations.ConsecutiveStopIdValidation(stopTime, 1, previousStopIdByTrip)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "First stop in a trip should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestConsecutiveStopIdValidation_MissingStopId(t *testing.T) {
	services.AppMessageService.Clear()
	previousStopIdByTrip := make(map[string]*string)

	tripId := "T1"
	stopTime := &types.StopTime{
		TripId: &tripId,
		StopId: nil,
	}
	validations.ConsecutiveStopIdValidation(stopTime, 1, previousStopIdByTrip)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing stop_id should skip validation",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestConsecutiveStopIdValidation_EmptyStopId(t *testing.T) {
	services.AppMessageService.Clear()
	previousStopIdByTrip := make(map[string]*string)

	tripId := "T1"
	emptyStopId := ""
	stopTime := &types.StopTime{
		TripId: &tripId,
		StopId: &emptyStopId,
	}
	validations.ConsecutiveStopIdValidation(stopTime, 1, previousStopIdByTrip)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Empty stop_id should skip validation",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestConsecutiveStopIdValidation_MissingTripId(t *testing.T) {
	services.AppMessageService.Clear()
	previousStopIdByTrip := make(map[string]*string)

	stopId := "S1"
	stopTime := &types.StopTime{
		TripId: nil,
		StopId: &stopId,
	}
	validations.ConsecutiveStopIdValidation(stopTime, 1, previousStopIdByTrip)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing trip_id should skip validation",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestConsecutiveStopIdValidation_ValidAfterNonConsecutive(t *testing.T) {
	services.AppMessageService.Clear()
	previousStopIdByTrip := make(map[string]*string)

	tripId := "T1"
	stopId1 := "S1"
	stopTime1 := &types.StopTime{
		TripId: &tripId,
		StopId: &stopId1,
	}
	validations.ConsecutiveStopIdValidation(stopTime1, 1, previousStopIdByTrip)

	stopId2 := "S2"
	stopTime2 := &types.StopTime{
		TripId: &tripId,
		StopId: &stopId2,
	}
	validations.ConsecutiveStopIdValidation(stopTime2, 2, previousStopIdByTrip)

	stopId3 := "S1" // Same as first, but not consecutive
	stopTime3 := &types.StopTime{
		TripId: &tripId,
		StopId: &stopId3,
	}
	validations.ConsecutiveStopIdValidation(stopTime3, 3, previousStopIdByTrip)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Same stop_id appearing non-consecutively should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestConsecutiveStopIdValidation_MultipleConsecutiveErrors(t *testing.T) {
	services.AppMessageService.Clear()
	previousStopIdByTrip := make(map[string]*string)

	tripId := "T1"
	stopId1 := "S1"
	stopTime1 := &types.StopTime{
		TripId: &tripId,
		StopId: &stopId1,
	}
	validations.ConsecutiveStopIdValidation(stopTime1, 1, previousStopIdByTrip)

	stopId2 := "S1" // First consecutive error
	stopTime2 := &types.StopTime{
		TripId: &tripId,
		StopId: &stopId2,
	}
	validations.ConsecutiveStopIdValidation(stopTime2, 2, previousStopIdByTrip)

	stopId3 := "S1" // Second consecutive error
	stopTime3 := &types.StopTime{
		TripId: &tripId,
		StopId: &stopId3,
	}
	validations.ConsecutiveStopIdValidation(stopTime3, 3, previousStopIdByTrip)

	// Message service deduplicates messages with the same text, so we get 1 error message
	// with multiple rows (rows 2 and 3) instead of 2 separate error messages
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Multiple consecutive stop_ids should produce one error message with multiple rows",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}


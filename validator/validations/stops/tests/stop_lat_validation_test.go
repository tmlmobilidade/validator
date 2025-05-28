package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestStopLatValidation_MissingRequiredStopLat(t *testing.T) {
	services.AppMessageService.Clear()
	lt := 0 // stop
	stop := &types.Stop{
		LocationType: &lt,
		StopLat:      nil,
	}
	
	validations.StopLatValidation(nil, stop, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing stop_lat should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}

	// Now test with severity ERROR
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	validations.StopLatValidation(&severity, stop, 2)
	assertion = lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing stop_lat with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopLatValidation_OptionalStopLat(t *testing.T) {
	services.AppMessageService.Clear()
	lt := 3 // generic node
	stop := &types.Stop{
		LocationType: &lt,
		StopLat:      nil,
	}
	validations.StopLatValidation(nil, stop, 3)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing stop_lat for optional location_type should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopLatValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	lt := 1 // station
	lat := float32(40.1234)
	stop := &types.Stop{
		LocationType: &lt,
		StopLat:      &lat,
	}
	validations.StopLatValidation(nil, stop, 4)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid stop_lat should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopLatValidation_OutOfRange(t *testing.T) {
	services.AppMessageService.Clear()
	
	lt := 2 // entrance/exit
	lat := float32(100.0) // out of range
	stop := &types.Stop{
		LocationType: &lt,
		StopLat:      &lat,
	}
	validations.StopLatValidation(nil, stop, 5)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Out-of-range stop_lat should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 

func TestStopLatValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()

	stop := &types.Stop{}
	severity := types.SEVERITY_ERROR
	
	validations.StopLatValidation(&severity, stop, 6)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Out-of-range stop_lat with severity ERROR should error",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopLatValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{}

	severity := types.SEVERITY_WARNING
	validations.StopLatValidation(&severity, stop, 7)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Out-of-range stop_lat with severity WARNING should warn",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestStopLonValidation_MissingRequiredStopLon(t *testing.T) {
	services.AppMessageService.Clear()
	lt := 0 // stop
	stop := &types.Stop{
		LocationType: &lt,
		StopLon:      nil,
	}
	
	validations.StopLonValidation(nil, stop, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing stop_lon should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}

	// Now test with severity ERROR
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	validations.StopLonValidation(&severity, stop, 2)
	assertion = lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing stop_lon with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopLonValidation_OptionalStopLon(t *testing.T) {
	services.AppMessageService.Clear()
	lt := 3 // generic node
	stop := &types.Stop{
		LocationType: &lt,
		StopLon:      nil,
	}
	validations.StopLonValidation(nil, stop, 3)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing stop_lon for optional location_type should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopLonValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	lt := 1 // station
	lat := float32(40.1234)
	stop := &types.Stop{
		LocationType: &lt,
		StopLon:      &lat,
	}
	validations.StopLonValidation(nil, stop, 4)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid stop_lon should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopLonValidation_OutOfRange(t *testing.T) {
	services.AppMessageService.Clear()
	
	lt := 2 // entrance/exit
	lat := float32(200.0) // out of range
	stop := &types.Stop{
		LocationType: &lt,
		StopLon:      &lat,
	}
	validations.StopLonValidation(nil, stop, 5)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Out-of-range stop_lon should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 

func TestStopLonValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	
	stop := &types.Stop{}
	severity := types.SEVERITY_ERROR
	
	validations.StopLonValidation(&severity, stop, 6)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Out-of-range stop_lon with severity ERROR should error",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopLonValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	
	stop := &types.Stop{}
	severity := types.SEVERITY_WARNING
	
	validations.StopLonValidation(&severity, stop, 7)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Out-of-range stop_lon with severity WARNING should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
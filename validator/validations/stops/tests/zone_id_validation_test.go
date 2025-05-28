package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestZoneIdValidation_MissingZoneId(t *testing.T) {
	services.AppMessageService.Clear()
	
	stop := &types.Stop{}
	validations.ZoneIdValidation(nil, stop, 3)
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing zone_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestZoneIdValidation_ValidZoneId(t *testing.T) {
	services.AppMessageService.Clear()

	zone := "Z3"
	stop := &types.Stop{
		ZoneId:       &zone,
	}
	validations.ZoneIdValidation(nil, stop, 4)
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid zone_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 

func TestZoneIdValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	
	stop := &types.Stop{}
	severity := types.SEVERITY_ERROR
	validations.ZoneIdValidation(&severity, stop, 5)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing zone_id with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestZoneIdValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	
	stop := &types.Stop{}
	severity := types.SEVERITY_WARNING
	validations.ZoneIdValidation(&severity, stop, 6)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Missing zone_id with severity WARNING should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
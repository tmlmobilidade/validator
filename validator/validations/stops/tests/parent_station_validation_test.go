package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestParentStationValidation_MissingParentStation_LocationType0(t *testing.T) {
	services.AppMessageService.Clear()
	
	locationType := 0
	stop := &types.Stop{LocationType: &locationType, ParentStation: nil}
	severity := types.SEVERITY_ERROR
	
	validations.ParentStationValidation(&severity, stop, 1)
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing parent_station for location_type=0 should not error",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestParentStationValidation_ValidParentStation_LocationType0(t *testing.T) {
	services.AppMessageService.Clear()
	
	locationType := 0
	parent := "STATION1"
	stop := &types.Stop{LocationType: &locationType, ParentStation: &parent}
	severity := types.SEVERITY_ERROR
	
	validations.ParentStationValidation(&severity, stop, 1)
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid parent_station for location_type=0 should not error",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestParentStationValidation_ParentStationForbidden_LocationType1(t *testing.T) {
	services.AppMessageService.Clear()
	
	locationType := 1
	parent := "STATION1"
	stop := &types.Stop{LocationType: &locationType, ParentStation: &parent}
	severity := types.SEVERITY_ERROR
	
	validations.ParentStationValidation(&severity, stop, 2)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "parent_station is forbidden for stations (location_type=1)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestParentStationValidation_MissingParentStation_LocationType1(t *testing.T) {
	services.AppMessageService.Clear()
	
	locationType := 1
	stop := &types.Stop{LocationType: &locationType, ParentStation: nil}
	severity := types.SEVERITY_ERROR
	
	validations.ParentStationValidation(&severity, stop, 2)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "parent_station is forbidden for stations (location_type=1)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestParentStationValidation_MissingParentStation_LocationType2(t *testing.T) {
	services.AppMessageService.Clear()
	
	locationType := 2
	stop := &types.Stop{LocationType: &locationType, ParentStation: nil}
	severity := types.SEVERITY_ERROR
	
	validations.ParentStationValidation(&severity, stop, 3)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing parent_station for location_type=2 should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestParentStationValidation_ValidParentStation_LocationType2(t *testing.T) {
	services.AppMessageService.Clear()
	
	locationType := 2
	parent := "STATION1"
	stop := &types.Stop{LocationType: &locationType, ParentStation: &parent}
	severity := types.SEVERITY_ERROR
	
	validations.ParentStationValidation(&severity, stop, 4)
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid parent_station for location_type=2 should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
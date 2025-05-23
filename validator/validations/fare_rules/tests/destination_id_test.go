package fare_rules

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_rules/validations"
	"testing"
)

func TestDestinationIdValidation_MissingDestinationId(t *testing.T) {
	services.AppMessageService.Clear()
	fareRule := &types.FareRule{DestinationId: nil}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {},
		},
	}
	validations.DestinationIdValidation(fareRule, 1, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing destination_id (optional) should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDestinationIdValidation_InvalidDestinationId(t *testing.T) {
	services.AppMessageService.Clear()
	
	invalidDestinationId := "INVALID"
	fareRule := &types.FareRule{DestinationId: &invalidDestinationId}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {},
		},
	}
	
	validations.DestinationIdValidation(fareRule, 2, gtfs, nil)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid destination_id should error",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDestinationIdValidation_ValidDestinationId(t *testing.T) {
	services.AppMessageService.Clear()
	
	validDestinationId := "DEST1"
	fareRule := &types.FareRule{DestinationId: &validDestinationId}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {"DEST1": {1}},
		},
	}
	
	validations.DestinationIdValidation(fareRule, 3, gtfs, nil)
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid destination_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 

func TestDestinationIdValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	
	fareRule := &types.FareRule{}
	gtfs := &types.Gtfs{}

	severity := types.SEVERITY_ERROR
	validations.DestinationIdValidation(fareRule, 3, gtfs, &severity)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid destination_id should error",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDestinationIdValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	
	fareRule := &types.FareRule{}
	gtfs := &types.Gtfs{}
	
	severity := types.SEVERITY_WARNING
	validations.DestinationIdValidation(fareRule, 3, gtfs, &severity)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Valid destination_id should error",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
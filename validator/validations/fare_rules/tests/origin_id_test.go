package fare_rules

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_rules/validations"
	"testing"
)

func TestOriginIdValidation_MissingOriginId(t *testing.T) {
	services.AppMessageService.Clear()
	fareRule := &types.FareRule{OriginId: nil}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {},
		},
	}
	validations.OriginIdValidation(fareRule, 1, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing origin_id (optional) should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestOriginIdValidation_InvalidOriginId(t *testing.T) {
	services.AppMessageService.Clear()
	
	invalidOriginId := "INVALID"
	fareRule := &types.FareRule{OriginId: &invalidOriginId}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {},
		},
	}
	
	validations.OriginIdValidation(fareRule, 2, gtfs, nil)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid origin_id should error",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestOriginIdValidation_ValidOriginId(t *testing.T) {
	services.AppMessageService.Clear()
	
	validOriginId := "ORIGIN1"
	fareRule := &types.FareRule{OriginId: &validOriginId}
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {"ORIGIN1": {1}},
		},
	}
	
	validations.OriginIdValidation(fareRule, 3, gtfs, nil)
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid origin_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 

func TestOriginIdValidation_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	
	fareRule := &types.FareRule{}
	gtfs := &types.Gtfs{}

	severity := types.SEVERITY_ERROR
	validations.OriginIdValidation(fareRule, 3, gtfs, &severity)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid origin_id should error",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestOriginIdValidation_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	
	fareRule := &types.FareRule{}
	gtfs := &types.Gtfs{}
	
	severity := types.SEVERITY_WARNING
	validations.OriginIdValidation(fareRule, 3, gtfs, &severity)
	
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Valid origin_id should error",
	}
	
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
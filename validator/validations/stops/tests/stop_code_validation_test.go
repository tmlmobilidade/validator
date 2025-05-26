package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestStopCodeValidation_MissingStopCode(t *testing.T) {
	services.AppMessageService.Clear()
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stop_codes": {},
		},
	}
	stop := &types.Stop{StopCode: nil}
	validations.StopCodeValidation(nil, stop, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing stop_code should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestStopCodeValidation_DuplicateStopCode(t *testing.T) {
	services.AppMessageService.Clear()
	code := "C1"
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stop_codes": {
				code: {1, 2},
			},
		},
	}
	stop := &types.Stop{StopCode: &code}
	validations.StopCodeValidation(nil, stop, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0, // Should not error, but should warn
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Duplicate stop_code should not error (should warn)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Expected 1 warning for duplicate stop_code")
	}
}

func TestStopCodeValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	code := "C1"
	gtfs := &types.Gtfs{
		IdMap: map[string]map[string][]int{
			"stops": {
				"stop_code": {1},
			},
		},
	}
	stop := &types.Stop{StopCode: &code}
	validations.StopCodeValidation(nil, stop, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid stop_code should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	if services.AppMessageService.GetSummary().TotalWarnings != 0 {
		t.Error("Expected 0 warnings for valid stop_code")
	}
} 
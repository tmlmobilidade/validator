package calendar

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/calendar/validations"
	"testing"
)

func TestServiceIdValidation_Required(t *testing.T) {
	calendar := &types.Calendar{ServiceId: ""}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"calendar": {}}}
	validations.ServiceIdValidation(calendar, 1, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Service ID is required",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestServiceIdValidation_Unique(t *testing.T) {
	calendar := &types.Calendar{ServiceId: "service1"}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"calendar": {"service1": {1, 2}}}}
	validations.ServiceIdValidation(calendar, 2, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Duplicate service_id found. Service IDs must be unique.",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestServiceIdValidation_ValidUnique(t *testing.T) {
	calendar := &types.Calendar{ServiceId: "service1"}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"calendar": {"service1": {1}}}}
	validations.ServiceIdValidation(calendar, 3, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Service ID is unique, should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
} 
package trips

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestServiceIdValidation_Required(t *testing.T) {
	trip := &types.Trip{ServiceId: ""}
	gtfs := &types.Gtfs{}
	validations.ServiceIdValidation(trip, 1, gtfs)

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

func TestServiceIdValidation_ValidForeignKey(t *testing.T) {
	trip := &types.Trip{ServiceId: "service1"}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"calendar": {"service1": {1}}, "calendar_dates": {}}}
	validations.ServiceIdValidation(trip, 2, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Service ID references a valid service_id, should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestServiceIdValidation_InvalidForeignKey(t *testing.T) {
	trip := &types.Trip{ServiceId: "invalid"}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"calendar": {}, "calendar_dates": {}}}
	validations.ServiceIdValidation(trip, 3, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Service ID must reference a valid service_id from calendar.txt or calendar_dates.txt.",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
} 
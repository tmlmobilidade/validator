package calendar_dates

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/calendar_dates/validations"
	"testing"
)

func TestServiceIdValidation_Present(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	calendarDate := &types.CalendarDates{
		Date: "20240101",
		ExceptionType: nil,
		ServiceId: "S1",
	}
	gtfs := &types.Gtfs{}
	validations.ServiceIdValidation(calendarDate, row, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Present service_id should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestServiceIdValidation_Missing(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	calendarDate := &types.CalendarDates{
		Date: "20240101",
		ExceptionType: nil,
		ServiceId: "",
	}
	gtfs := &types.Gtfs{}
	validations.ServiceIdValidation(calendarDate, row, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing service_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
package calendar_dates

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/calendar_dates/validations"
	"testing"
)

func TestDateValidation_Valid(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	calendarDate := &types.CalendarDates{
		Date: "20240101",
		ExceptionType: nil,
		ServiceId: "S1",
	}
	
	gtfs := &types.Gtfs{}
	validations.DateValidation(calendarDate, row, gtfs)
	
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid date should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDateValidation_Empty(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	calendarDate := &types.CalendarDates{
		Date: "",
		ExceptionType: nil,
		ServiceId: "S1",
	}
	gtfs := &types.Gtfs{}
	validations.DateValidation(calendarDate, row, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Empty date should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDateValidation_InvalidFormat(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	calendarDate := &types.CalendarDates{
		Date: "2024-01-01",
		ExceptionType: nil,
		ServiceId: "S1",
	}
	gtfs := &types.Gtfs{}
	validations.DateValidation(calendarDate, row, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid date format should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
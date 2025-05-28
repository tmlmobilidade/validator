package calendar_dates

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/calendar_dates/validations"
	"testing"
)

func TestExceptionTypeValidation_Valid(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	validTypes := []int{1, 2}
	for _, et := range validTypes {
		calendarDate := &types.CalendarDates{
			Date: "20240101",
			ExceptionType: &et,
			ServiceId: "S1",
		}
		gtfs := &types.Gtfs{}
		validations.ExceptionTypeValidation(calendarDate, row, gtfs)
		assertion := lib.AssertionMessage{
			Expected: 0,
			Actual: services.AppMessageService.GetSummary().TotalErrors,
			Message: "Valid exception_type should not error",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Error(assert)
		}
		services.AppMessageService.Clear()
	}
}

func TestExceptionTypeValidation_Invalid(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	invalidType := 3
	calendarDate := &types.CalendarDates{
		Date: "20240101",
		ExceptionType: &invalidType,
		ServiceId: "S1",
	}
	gtfs := &types.Gtfs{}
	validations.ExceptionTypeValidation(calendarDate, row, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid exception_type should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
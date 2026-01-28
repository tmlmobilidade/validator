package calendar_dates

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/calendar_dates/validations"
	"testing"
)

func TestExceptionTypeValidation_Valid_Type1(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	gtfs := &types.Gtfs{}
	et := 1
	calendarDate := &types.CalendarDates{
		Date:          "20240101",
		ExceptionType: &et,
		ServiceId:     "S1",
	}
	validations.ExceptionTypeValidation(calendarDate, row, *gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid exception_type=1 should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestExceptionTypeValidation_Invalid(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	gtfs := &types.Gtfs{}
	invalidType := 3
	calendarDate := &types.CalendarDates{
		Date:          "20240101",
		ExceptionType: &invalidType,
		ServiceId:     "S1",
	}
	validations.ExceptionTypeValidation(calendarDate, row, *gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid exception_type should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

// ISO-1033: When only calendar_dates.txt is used (no calendar.txt), exception_type must be 1
func TestExceptionTypeValidation_OnlyCalendarDates_Type1_ShouldPass(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	gtfs := &types.Gtfs{} // Empty gtfs means no calendar.txt
	exceptionType := 1
	calendarDate := &types.CalendarDates{
		Date:          "20240101",
		ExceptionType: &exceptionType,
		ServiceId:     "S1",
	}
	validations.ExceptionTypeValidation(calendarDate, row, *gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "When only calendar_dates.txt is used, exception_type=1 should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

// ISO-1033: When only calendar_dates.txt is used (no calendar.txt), exception_type=2 should error
func TestExceptionTypeValidation_OnlyCalendarDates_Type2_ShouldError(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	gtfs := &types.Gtfs{} // Empty gtfs means no calendar.txt
	exceptionType := 2
	calendarDate := &types.CalendarDates{
		Date:          "20240101",
		ExceptionType: &exceptionType,
		ServiceId:     "S1",
	}
	validations.ExceptionTypeValidation(calendarDate, row, *gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "When only calendar_dates.txt is used, exception_type=2 should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

package calendar_dates

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/calendar_dates"
	"testing"
)



func TestParseCalendarDates_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	row := 2
	raw := map[string]string{
		"date": "20240101",
		"exception_type": "1",
		"service_id": "S1",
	}
	gtfs := &types.Gtfs{}
	calendarDate := validations.ParseCalendarDates(raw, row, gtfs)

	if calendarDate.Date != "20240101" {
		t.Errorf("Expected date to be 20240101, got %s", calendarDate.Date)
	}
	if *calendarDate.ExceptionType != 1 {
		t.Errorf("Expected exception_type to be 1, got %d", *calendarDate.ExceptionType)
	}
	if calendarDate.ServiceId != "S1" {
		t.Errorf("Expected service_id to be S1, got %s", calendarDate.ServiceId)
	}
}

func TestParseCalendarDates_InvalidExceptionType(t *testing.T) {
	services.AppMessageService.Clear()
	row := 2
	raw := map[string]string{
		"date": "20240101",
		"exception_type": "WRONG VALUE",
		"service_id": "S1",
	}
	gtfs := &types.Gtfs{}
	_ = validations.ParseCalendarDates(raw, row, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid exception_type (should error)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}


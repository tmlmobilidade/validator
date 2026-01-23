package calendar_dates

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/calendar_dates/validations"
	"testing"
)

func TestExceptionTypeOnlyCalendarDatesValidation_HasCalendar_ShouldNotApply(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	hasCalendar := true
	calendarDate := &types.CalendarDates{
		Date:          "20240101",
		ExceptionType: nil,
		ServiceId:     "S1",
	}
	validations.ExceptionTypeOnlyCalendarDatesValidation(calendarDate, hasCalendar, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "When calendar.txt exists, validation should not apply even if exception_type is nil",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestExceptionTypeOnlyCalendarDatesValidation_HasCalendar_WithInvalidValue_ShouldNotApply(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	hasCalendar := true
	invalidType := 2
	calendarDate := &types.CalendarDates{
		Date:          "20240101",
		ExceptionType: &invalidType,
		ServiceId:     "S1",
	}
	validations.ExceptionTypeOnlyCalendarDatesValidation(calendarDate, hasCalendar, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "When calendar.txt exists, validation should not apply even if exception_type is invalid",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestExceptionTypeOnlyCalendarDatesValidation_NoCalendar_ValidExceptionType_ShouldNotError(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	hasCalendar := false
	validType := 1
	calendarDate := &types.CalendarDates{
		Date:          "20240101",
		ExceptionType: &validType,
		ServiceId:     "S1",
	}
	validations.ExceptionTypeOnlyCalendarDatesValidation(calendarDate, hasCalendar, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "When only calendar_dates.txt exists and exception_type is 1, should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestExceptionTypeOnlyCalendarDatesValidation_NoCalendar_InvalidExceptionType_WithRules_ShouldError(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	hasCalendar := false
	invalidType := 2
	options := []string{"1"}
	rules := &types.CalendarDatesRules{
		ExceptionTypeOnlyCalendarDates: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
			Options:  &options,
		},
	}
	calendarDate := &types.CalendarDates{
		Date:          "20240101",
		ExceptionType: &invalidType,
		ServiceId:     "S1",
	}
	validations.ExceptionTypeOnlyCalendarDatesValidation(calendarDate, hasCalendar, row, rules)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "When only calendar_dates.txt exists and exception_type is not in allowed options, should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestExceptionTypeOnlyCalendarDatesValidation_NoCalendar_ValidExceptionType_WithRules_ShouldNotError(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	hasCalendar := false
	validType := 1
	options := []string{"1"}
	rules := &types.CalendarDatesRules{
		ExceptionTypeOnlyCalendarDates: types.RuleConfig{
			Severity: types.SEVERITY_ERROR,
			Options:  &options,
		},
	}
	calendarDate := &types.CalendarDates{
		Date:          "20240101",
		ExceptionType: &validType,
		ServiceId:     "S1",
	}
	validations.ExceptionTypeOnlyCalendarDatesValidation(calendarDate, hasCalendar, row, rules)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "When only calendar_dates.txt exists and exception_type is in allowed options, should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestExceptionTypeOnlyCalendarDatesValidation_NoCalendar_AllOptions_ShouldNotError(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	hasCalendar := false
	anyType := 2
	options := []string{types.ALL_OPTIONS}
	rules := &types.CalendarDatesRules{
		ExceptionTypeOnlyCalendarDates: types.RuleConfig{
			Options: &options,
		},
	}
	calendarDate := &types.CalendarDates{
		Date:          "20240101",
		ExceptionType: &anyType,
		ServiceId:     "S1",
	}
	validations.ExceptionTypeOnlyCalendarDatesValidation(calendarDate, hasCalendar, row, rules)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "When rules allow all options, any exception_type should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestExceptionTypeOnlyCalendarDatesValidation_Forbidden_ShouldError(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	hasCalendar := false
	validType := 1
	rules := &types.CalendarDatesRules{
		ExceptionTypeOnlyCalendarDates: types.RuleConfig{
			Severity: types.SEVERITY_FORBIDDEN,
		},
	}
	calendarDate := &types.CalendarDates{
		Date:          "20240101",
		ExceptionType: &validType,
		ServiceId:     "S1",
	}
	validations.ExceptionTypeOnlyCalendarDatesValidation(calendarDate, hasCalendar, row, rules)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "When field is forbidden but has a value, should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestExceptionTypeOnlyCalendarDatesValidation_WithSeverity_ShouldUseCustomSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	hasCalendar := false
	rules := &types.CalendarDatesRules{
		ExceptionTypeOnlyCalendarDates: types.RuleConfig{
			Severity: types.SEVERITY_WARNING,
		},
	}
	calendarDate := &types.CalendarDates{
		Date:          "20240101",
		ExceptionType: nil,
		ServiceId:     "S1",
	}
	validations.ExceptionTypeOnlyCalendarDatesValidation(calendarDate, hasCalendar, row, rules)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "When rules specify warning severity, should use custom severity",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

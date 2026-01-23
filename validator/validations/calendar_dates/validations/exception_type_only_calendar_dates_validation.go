package calendar_dates

import (
	"fmt"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
	"strconv"
	"strings"
)

/*
# Attributes
  - File: [calendar_dates.txt]
  - Field: exception_type
  - Presence: Optional
  - Type: Enum

# Description

If only the calendar_dates.txt file is used, this field must be set to 1.

Valid options are:

  - 1 - Service has been added for the specified date.

[calendar_dates.txt]: https://gtfs.org/schedule/reference/#calendar_datestxt
*/
func ExceptionTypeOnlyCalendarDatesValidation(calendarDate *types.CalendarDates, hasCalendar bool, row int, rules *types.CalendarDatesRules) {
	ctx := lib.NewValidationContext("exception_type", "calendar_dates.txt", "exception_type_only_calendar_dates_validation", row, services.AppMessageService)
	if rules != nil && rules.ExceptionTypeOnlyCalendarDates.Severity != "" {
		ctx.WithSeverity(rules.ExceptionTypeOnlyCalendarDates.Severity)
	}

	// Is null
	if calendarDate.ExceptionType == nil {
		if ctx.ShouldIgnore() { // The field is optional, so we can ignore it
			return
		}

		// The field is required but came empty, so we need to add an error
		message := ctx.GetRequiredMessage("exception_type_only_calendar_dates_validation.required", "exception_type_only_calendar_dates_validation.required")
		ctx.AddMessageWithSeverity(message)
		return
	}

	lib.AppLogger.Accent(fmt.Sprintf("===== > ctx.Severity: %s", string(ctx.Severity)))

	// The field is forbidden but came with a value, so we need to add an error
	if ctx.IsForbidden() {
		ctx.AddError(ctx.GetTranslatedMessage("exception_type_only_calendar_dates_validation.forbidden"))
		return
	}

	// This validation only applies when calendar_dates.txt exists without calendar.txt
	if hasCalendar {
		return
	}

	// Validate rules
	if rules != nil && rules.ExceptionTypeOnlyCalendarDates.Options != nil {
		if slices.Contains(*rules.ExceptionTypeOnlyCalendarDates.Options, types.ALL_OPTIONS) {
			return
		}

		// If the exception type is not allowed, add an error
		if !slices.Contains(*rules.ExceptionTypeOnlyCalendarDates.Options, strconv.Itoa(*calendarDate.ExceptionType)) {
			ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("exception_type_only_calendar_dates_validation.invalid", strings.Join(*rules.ExceptionTypeOnlyCalendarDates.Options, ", ")))
			return
		}
	}
}

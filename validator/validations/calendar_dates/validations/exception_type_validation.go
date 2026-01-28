package calendar_dates

import (
	"fmt"
	"main/services"
	"main/types"
	"slices"
	"strconv"
)

/*
# Attributes
  - File: [calendar_dates.txt]
  - Field: exception_type
  - Presence: Required
  - Type: Enum

# Description

Indicates whether service is available on the date specified in the date field.

Valid options are:

  - 1 - Service has been added for the specified date.
  - 2 - Service has been removed for the specified date.

# Example

Suppose a route has one set of trips available on holidays and another set of trips available on all other days.

One service_id could correspond to the regular service schedule and another service_id could correspond to the holiday schedule.

For a particular holiday, the [calendar_dates.txt] file could be used to add the holiday to the holiday service_id and to remove the holiday from the regular service_id schedule.

[calendar_dates.txt]: https://gtfs.org/schedule/reference/#calendar_datestxt
*/
func ExceptionTypeValidation(calendarDate *types.CalendarDates, row int, gtfs types.Gtfs, rules *types.CalendarDatesRules) {

	message := types.Message{
		Field:        "exception_type",
		FileName:     "calendar_dates.txt",
		Rows:         []int{row},
		Severity:     types.SEVERITY_ERROR,
		ValidationID: "exception_type_validation",
	}

	validExceptionTypes := []int{1, 2}

	if !slices.Contains(validExceptionTypes, *calendarDate.ExceptionType) {
		message.Message = fmt.Sprintf("Wrong exception_type value, must be 1 or 2, got %d", *calendarDate.ExceptionType)
		services.AppMessageService.AddMessage(message)
	}

	// ISO-1033: If only calendar_dates.txt is used (no calendar.txt), exception_type must be 1
	hasCalendar := gtfs.HasTable("calendar")
	if !hasCalendar && *calendarDate.ExceptionType == 2 {
		message.Message = "When only calendar_dates.txt is used (without calendar.txt), exception_type must be 1 (service added). Value 2 (service removed) requires a base calendar to remove from."
		services.AppMessageService.AddMessage(message)
	}

	// Validate Rule Options
	if rules != nil && rules.ExceptionType.Options != nil {
		if slices.Contains(*rules.ExceptionType.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.ExceptionType.Options, strconv.Itoa(*calendarDate.ExceptionType)) {
			message.Message = fmt.Sprintf("Exception type \"%d\" is not allowed", *calendarDate.ExceptionType)
			services.AppMessageService.AddMessage(message)
		}

		message.Message = fmt.Sprintf("Exception type \"%d\" is not allowed", *calendarDate.ExceptionType)
		services.AppMessageService.AddMessage(message)
		return
	}
}

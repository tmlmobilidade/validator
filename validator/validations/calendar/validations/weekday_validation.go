package trips

import (
	"main/services"
	"main/types"
)

/*
# Attributes

	- File: [calendar.txt]
	- Field: [monday, tuesday, wednesday, thursday, friday, saturday, sunday]
	- Presence: Required
	- Type: Enum

# Description

Indicates whether the service operates on all Mondays in the date range specified by the start_date and end_date fields.

Note that exceptions for particular dates may be listed in calendar_dates.txt.

Valid options are:
	- 1 - Service is available for all Mondays in the date range.
	- 0 - Service is not available for Mondays in the date range.

[calendar.txt]: https://gtfs.org/schedule/reference/#calendartxt
*/
func WeekdayValidation(calendar *types.Calendar, row int, weekday types.Weekday) {

	message := types.Message{
		Field: string(weekday),
		FileName: "calendar.txt",
		Message: "service_id is required",
		Rows: []int{row},
		Severity: types.SEVERITY_ERROR,
		ValidationID: "service_id_validation",
	}

	valid := false
	switch weekday {
	case types.WeekdayMonday:
		valid = calendar.Monday
	case types.WeekdayTuesday:
		valid = calendar.Tuesday
	case types.WeekdayWednesday:
		valid = calendar.Wednesday
	case types.WeekdayThursday:
		valid = calendar.Thursday
	case types.WeekdayFriday:
		valid = calendar.Friday
	case types.WeekdaySaturday:
		valid = calendar.Saturday
	case types.WeekdaySunday:
		valid = calendar.Sunday
	}

	if !valid {
		services.AppMessageService.AddMessage(message)
	}
}
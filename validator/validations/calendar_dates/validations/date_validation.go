package calendar_dates

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes
	- File: [calendar_dates.txt]
	- Field: date
	- Presence: Required
	- Type: Date

# Description

Date when service exception occurs.

[calendar_dates.txt]: https://gtfs.org/schedule/reference/#calendar_datestxt
*/
func DateValidation(calendarDate *types.CalendarDates, row int, gtfs *types.Gtfs) {	
	message := types.Message{
		Field: "date",
		FileName: "calendar_dates.txt",
		Rows: []int{row},
		Severity: types.SEVERITY_ERROR,
		ValidationID: "date_validation",
	}
	
	date := calendarDate.Date

	if date == "" {
		message.Message = "Date is required"
		services.AppMessageService.AddMessage(message)
	}

	if !lib.IsValidServiceDate(date) {
		lib.AppLogger.Accent("Invalid date format, expected format: YYYYMMDD, got:", date)
		message.Message = "Date is not valid service date format, should be YYYYMMDD"
		services.AppMessageService.AddMessage(message)
	}
}

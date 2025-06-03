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
	
	addMessage := func(message string) {
		services.AppMessageService.AddMessage(types.Message{
			Field: "date",
			FileName: "calendar_dates.txt",
			Rows: []int{row},
			Severity: types.SEVERITY_ERROR,
			ValidationID: "date_validation",
			Message: message,
		})
	}
	
	
	date := calendarDate.Date

	if date == "" {
		addMessage("Date is required")
		return
	}

	if !lib.IsValidServiceDate(date) {
		addMessage("Date is not valid service date format, should be YYYYMMDD")
		return
	}
}

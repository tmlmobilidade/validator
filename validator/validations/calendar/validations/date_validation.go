package trips

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

	- File: [calendar.txt]
	- Field: [start_date, end_date]
	- Presence: Required
	- Type: Date

# Description

Start service day for the service interval.

End service day for the service interval. This service day is included in the interval.

[calendar.txt]: https://gtfs.org/schedule/reference/#calendartxt
*/
func DateValidation(date string, dateType string,row int) {

	message := types.Message{
		Field: dateType,
		FileName: "calendar.txt",
		Message: "Service date is required",
		Rows: []int{row},
		Severity: types.SEVERITY_ERROR,
		ValidationID: "date_validation",
	}

	if date == "" {
		services.AppMessageService.AddMessage(message)
	}

	if !lib.IsValidServiceDate(date) {
		message.Message = "Invalid service date, must be in format YYYYMMDD"
		services.AppMessageService.AddMessage(message)
	}
}
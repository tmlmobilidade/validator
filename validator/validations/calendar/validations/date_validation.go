package trips

import (
	"main/i18n"
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
func DateValidation(date string, dateType string, row int) {

	addMessage := func(message string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        dateType,
			FileName:     "calendar.txt",
			Message:      message,
			Rows:         []int{row},
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "date_validation",
		})
	}

	if date == "" {
		addMessage(i18n.AppTranslator.Get("date_validation.required"))
		return
	}

	if !lib.IsValidServiceDate(date) {
		addMessage(i18n.AppTranslator.Get("date_validation.invalid", date))
		return
	}
}

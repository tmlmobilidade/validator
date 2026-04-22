package trips

import (
	"main/lib"
	"main/services"
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
	ctx := lib.NewValidationContext(dateType, "calendar.txt", "date_validation", "validate_service_interval_date", row, services.AppMessageService)

	if date == "" {
		ctx.AddError(ctx.GetTranslatedMessage("date_validation.required"))
		return
	}

	if !lib.IsValidServiceDate(date) {
		ctx.AddError(ctx.GetTranslatedMessage("date_validation.invalid", date))
		return
	}
}

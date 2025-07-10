package calendar_dates

import (
	"main/lib"
	"main/types"
	validations "main/validations/calendar_dates/validations"
)

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {

	lib.AppLogger.Debug("Running Calendar Dates Validations...")

	for i, rawCalendarDate := range gtfs.CalendarDates {
		calendarDate := ParseCalendarDates(rawCalendarDate, i)

		if calendarDate == (types.CalendarDates{}) {
			continue
		}

		var calendarDatesRules types.CalendarDatesRules
		if rules != nil {
			calendarDatesRules = rules.CalendarDates
		}

		// Validate service_id
		validations.ServiceIdValidation(&calendarDate, i)

		// Validate date
		validations.DateValidation(&calendarDate, i)

		// Validate exception_type
		validations.ExceptionTypeValidation(&calendarDate, i, &calendarDatesRules)

	}
}

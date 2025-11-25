package calendar_dates

import (
	"fmt"
	"main/lib"
	"main/types"
	validations "main/validations/calendar_dates/validations"
)

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {

	lib.AppLogger.Debug("Running Calendar Dates Validations...")

	err := gtfs.IterateCalendarDates(func(i int, rawCalendarDate types.CalendarDatesRaw) error {
		calendarDate := ParseCalendarDates(rawCalendarDate, i)

		if calendarDate == (types.CalendarDates{}) {
			return nil
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

		return nil
	})

	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error iterating calendar dates: %v", err))
	}
}

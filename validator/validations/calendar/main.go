package calendar

import (
	"main/lib"
	"main/types"
	validations "main/validations/calendar/validations"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Calendar Validations...")

	for i, rawCalendar := range gtfs.Files["calendar"] {
		calendar := ParseCalendar(rawCalendar, i, &gtfs)

		if calendar == (types.Calendar{}) {
			continue
		}

		// Validate service_id
		validations.ServiceIdValidation(&calendar, i, &gtfs)

		// Validate weekdays
		validations.WeekdayValidation(&calendar, i, types.WeekdayMonday)
		validations.WeekdayValidation(&calendar, i, types.WeekdayTuesday)
		validations.WeekdayValidation(&calendar, i, types.WeekdayWednesday)
		validations.WeekdayValidation(&calendar, i, types.WeekdayThursday)
		validations.WeekdayValidation(&calendar, i, types.WeekdayFriday)
		validations.WeekdayValidation(&calendar, i, types.WeekdaySaturday)
		validations.WeekdayValidation(&calendar, i, types.WeekdaySunday)

		// Validate service dates
		validations.DateValidation(calendar.StartDate, "start_date", i)
		validations.DateValidation(calendar.EndDate, "end_date", i)
		
	}
}

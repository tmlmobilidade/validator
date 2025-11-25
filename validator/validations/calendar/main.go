package calendar

import (
	"fmt"
	"main/lib"
	"main/types"
	validations "main/validations/calendar/validations"
)

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Calendar Validations...")

	err := gtfs.IterateCalendars(func(i int, rawCalendar types.CalendarRaw) error {
		calendar := ParseCalendar(rawCalendar, i, &gtfs)

		if calendar == (types.Calendar{}) {
			return nil
		}

		// Validate service_id
		validations.ServiceIdValidation(&calendar, i, &gtfs)

		// Validate service dates
		validations.DateValidation(calendar.StartDate, "start_date", i)
		validations.DateValidation(calendar.EndDate, "end_date", i)

		return nil
	})

	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error iterating calendars: %v", err))
	}
}

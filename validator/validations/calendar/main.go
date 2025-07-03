package calendar

import (
	"main/lib"
	"main/types"
	validations "main/validations/calendar/validations"
)

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Calendar Validations...")

	for i, rawCalendar := range gtfs.Calendar {
		calendar := ParseCalendar(rawCalendar, i, &gtfs)

		if calendar == (types.Calendar{}) {
			continue
		}

		// Validate service_id
		validations.ServiceIdValidation(&calendar, i, &gtfs)

		// Validate service dates
		validations.DateValidation(calendar.StartDate, "start_date", i)
		validations.DateValidation(calendar.EndDate, "end_date", i)
		
	}
}

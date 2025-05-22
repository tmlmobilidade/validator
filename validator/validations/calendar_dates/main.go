package calendar_dates

import (
	"main/lib"
	"main/types"
)

func RunValidations(gtfs types.Gtfs) {

	lib.AppLogger.Debug("Running Calendar Dates Validations...")

	for i, rawCalendarDate := range gtfs.Files["calendar_dates"] {
		calendarDate := ParseCalendarDates(rawCalendarDate, i, &gtfs)

		if calendarDate == (types.CalendarDates{}) {
			continue
		}
		
		
	}
}

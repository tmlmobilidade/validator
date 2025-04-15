package calendar_dates

import (
	"main/validator/lib"
	"main/validator/types"
)

type parseCalendarDatesValidation struct {
	*types.Validation
}

func NewParseCalendarDatesValidation(severity *types.Severity) *parseCalendarDatesValidation {
	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	return &parseCalendarDatesValidation{
		Validation: &types.Validation{
			ID:          "parse_calendar_dates",
			Description: "Validate calendar dates data",
			Severity:    s,
		},
	}
}

func (v *parseCalendarDatesValidation) Validate(gtfs types.Gtfs) (calendarDates []types.CalendarDates, messages []types.Message) {
	// Track unique service_id + date combinations
	serviceIdDatePairs := make(map[string]bool)

	// Check if calendar.txt exists to determine if calendar_dates.txt is required
	_, hasCalendar := gtfs.Files["calendar"]

	for i, calendarDate := range gtfs.Files["calendar_dates"] {
		calendarDate, calendarDateMessages := parseCalendarDate(calendarDate, hasCalendar)
		calendarDates = append(calendarDates, calendarDate)

		// Check for duplicate service_id + date combinations
		if calendarDate.ServiceId != "" && calendarDate.Date != "" {
			pairKey := calendarDate.ServiceId + "_" + calendarDate.Date
			if serviceIdDatePairs[pairKey] {
				messages = append(messages, types.Message{
					Field:        "service_id,date",
					FileName:     "calendar_dates.txt",
					Message:      "Duplicate (service_id, date) pair found. Each pair may only appear once in calendar_dates.txt.",
					Row:          i + 1,
					Severity:     v.Severity,
					ValidationID: v.ID,
				})
			}
			serviceIdDatePairs[pairKey] = true
		}

		// Update row number and other fields for each message
		for _, msg := range calendarDateMessages {
			msg.Row = i + 1
			msg.FileName = "calendar_dates.txt"
			msg.Severity = v.Severity
			msg.ValidationID = v.ID
			messages = append(messages, msg)
		}
	}

	// If calendar.txt is omitted, ensure calendar_dates.txt contains at least one entry
	if !hasCalendar && len(calendarDates) == 0 {
		messages = append(messages, types.Message{
			Field:        "",
			FileName:     "calendar_dates.txt",
			Message:      "calendar_dates.txt must contain at least one entry when calendar.txt is omitted",
			Row:          0,
			Severity:     v.Severity,
			ValidationID: v.ID,
		})
	}

	return calendarDates, messages
}

func parseCalendarDate(m map[string]string, hasCalendar bool) (calendarDate types.CalendarDates, messages []types.Message) {
	var parsingErrors []string

	// Parse Required Values
	lib.ParseStringToPrimitive(m["service_id"], &calendarDate.ServiceId, &parsingErrors)
	lib.ParseStringToPrimitive(m["date"], &calendarDate.Date, &parsingErrors)
	lib.ParseStringToPrimitive(m["exception_type"], &calendarDate.ExceptionType, &parsingErrors)

	if len(parsingErrors) > 0 {
		for _, err := range parsingErrors {
			messages = append(messages, types.Message{
				Field:   "N/A", //TODO: Add field name
				Message: err,
			})
		}
	}

	// Validate Required Fields
	if calendarDate.ServiceId == "" {
		messages = append(messages, types.Message{
			Field:   "service_id",
			Message: "Service ID is required.",
		})
	}

	if calendarDate.Date == "" {
		messages = append(messages, types.Message{
			Field:   "date",
			Message: "Date is required.",
		})
	} else {
		// Validate date format (YYYYMMDD)
		if !lib.IsValidServiceDate(calendarDate.Date) {
			messages = append(messages, types.Message{
				Field:   "date",
				Message: "Invalid date format. Date must be in YYYYMMDD format.",
			})
		}
	}

	// Validate exception_type enum values
	validExceptionTypes := map[int]bool{1: true, 2: true}
	if !validExceptionTypes[calendarDate.ExceptionType] {
		messages = append(messages, types.Message{
			Field:   "exception_type",
			Message: "Invalid exception_type value. Valid values are 1 (service added) or 2 (service removed).",
		})
	}

	return calendarDate, messages
}

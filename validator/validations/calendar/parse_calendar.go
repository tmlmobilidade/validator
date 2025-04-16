package calendar

import (
	"main/validator/lib"
	"main/validator/types"
	"time"
)

type parseCalendarValidation struct {
	*types.Validation
}

func NewParseCalendarValidation(severity *types.Severity) *parseCalendarValidation {
	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	return &parseCalendarValidation{
		Validation: &types.Validation{
			ID:          "parse_calendar",
			Description: "Validate calendar data",
			Severity:    s,
		},
	}
}

func (v *parseCalendarValidation) Validate(gtfs types.Gtfs) (calendars []types.Calendar, messages []types.Message) {
	serviceIds := make(map[string]bool)

	for i, calendar := range gtfs.Files["calendar"] {
		cal, calendarMessages := parseCalendar(calendar)
		calendars = append(calendars, cal)

		// Check for duplicate service IDs
		if cal.ServiceId != "" {
			if serviceIds[cal.ServiceId] {
				messages = append(messages, types.Message{
					Field:        "service_id",
					FileName:     "calendar.txt",
					Message:      "Duplicate service_id found. Service IDs must be unique.",
					Rows:         []int{i + 1},
					Severity:     v.Severity,
					ValidationID: v.ID,
				})
			}
			serviceIds[cal.ServiceId] = true
		}

		// Update row number and other fields for each message
		for _, msg := range calendarMessages {
			msg.Rows = []int{i + 1}
			msg.FileName = "calendar.txt"
			msg.Severity = v.Severity
			msg.ValidationID = v.ID
			messages = append(messages, msg)
		}
	}
	return calendars, messages
}

func parseCalendar(m map[string]string) (calendar types.Calendar, messages []types.Message) {
	var parsingErrors []string

	// Convert Required Values
	lib.ParseStringToPrimitive(m["service_id"], &calendar.ServiceId, &parsingErrors)

	// Parse day fields as integers
	var monday, tuesday, wednesday, thursday, friday, saturday, sunday int
	lib.ParseStringToPrimitive(m["monday"], &monday, &parsingErrors)
	lib.ParseStringToPrimitive(m["tuesday"], &tuesday, &parsingErrors)
	lib.ParseStringToPrimitive(m["wednesday"], &wednesday, &parsingErrors)
	lib.ParseStringToPrimitive(m["thursday"], &thursday, &parsingErrors)
	lib.ParseStringToPrimitive(m["friday"], &friday, &parsingErrors)
	lib.ParseStringToPrimitive(m["saturday"], &saturday, &parsingErrors)
	lib.ParseStringToPrimitive(m["sunday"], &sunday, &parsingErrors)

	calendar.Monday = monday == 1
	calendar.Tuesday = tuesday == 1
	calendar.Wednesday = wednesday == 1
	calendar.Thursday = thursday == 1
	calendar.Friday = friday == 1
	calendar.Saturday = saturday == 1
	calendar.Sunday = sunday == 1

	lib.ParseStringToPrimitive(m["start_date"], &calendar.StartDate, &parsingErrors)
	lib.ParseStringToPrimitive(m["end_date"], &calendar.EndDate, &parsingErrors)

	if len(parsingErrors) > 0 {
		for _, err := range parsingErrors {
			messages = append(messages, types.Message{
				Field:   "N/A",
				Message: err,
			})
		}
	}

	// Validate Required Fields
	if calendar.ServiceId == "" {
		messages = append(messages, types.Message{
			Field:   "service_id",
			Message: "Service ID is required.",
		})
	}

	// Validate day fields (must be 0 or 1)
	validDayValues := map[int]bool{0: true, 1: true}

	if !validDayValues[monday] {
		messages = append(messages, types.Message{
			Field:   "monday",
			Message: "Monday value must be either 0 or 1.",
		})
	}

	if !validDayValues[tuesday] {
		messages = append(messages, types.Message{
			Field:   "tuesday",
			Message: "Tuesday value must be either 0 or 1.",
		})
	}

	if !validDayValues[wednesday] {
		messages = append(messages, types.Message{
			Field:   "wednesday",
			Message: "Wednesday value must be either 0 or 1.",
		})
	}

	if !validDayValues[thursday] {
		messages = append(messages, types.Message{
			Field:   "thursday",
			Message: "Thursday value must be either 0 or 1.",
		})
	}

	if !validDayValues[friday] {
		messages = append(messages, types.Message{
			Field:   "friday",
			Message: "Friday value must be either 0 or 1.",
		})
	}

	if !validDayValues[saturday] {
		messages = append(messages, types.Message{
			Field:   "saturday",
			Message: "Saturday value must be either 0 or 1.",
		})
	}

	if !validDayValues[sunday] {
		messages = append(messages, types.Message{
			Field:   "sunday",
			Message: "Sunday value must be either 0 or 1.",
		})
	}

	// Validate dates
	if calendar.StartDate == "" {
		messages = append(messages, types.Message{
			Field:   "start_date",
			Message: "Start date is required.",
		})
	}

	if calendar.EndDate == "" {
		messages = append(messages, types.Message{
			Field:   "end_date",
			Message: "End date is required.",
		})
	}

	// Validate date format and range
	if calendar.StartDate != "" && calendar.EndDate != "" {
		startDate, err := time.Parse("20060102", calendar.StartDate)
		if err != nil {
			messages = append(messages, types.Message{
				Field:   "start_date",
				Message: "Invalid start date format. Expected YYYYMMDD.",
			})
		}

		endDate, err := time.Parse("20060102", calendar.EndDate)
		if err != nil {
			messages = append(messages, types.Message{
				Field:   "end_date",
				Message: "Invalid end date format. Expected YYYYMMDD.",
			})
		}

		if err == nil && startDate.After(endDate) {
			messages = append(messages, types.Message{
				Field:   "end_date",
				Message: "End date must be after or equal to start date.",
			})
		}
	}

	return calendar, messages
}

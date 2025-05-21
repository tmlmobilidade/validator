package calendar

import (
	"main/lib"
	"main/services"
	"main/types"
)

// Parses a trip row from the trips.txt file to a Trip struct
func ParseCalendar(rawCalendar map[string]string, row int, gtfs *types.Gtfs) (calendar types.Calendar) {
	message := types.Message{
		Field:   "",
		FileName: "calendar.txt",
		Rows: []int{row},
		Message: "",
		Severity: types.SEVERITY_ERROR,
		ValidationID: "calendar_parse",
	}

	var serviceId, startDate, endDate, monday, tuesday, wednesday, thursday, friday, saturday, sunday string

	fieldMappings := map[string]*string{
		"service_id":            &serviceId,
		"monday":                &monday,
		"tuesday":               &tuesday,
		"wednesday":             &wednesday,
		"thursday":              &thursday,
		"friday":                &friday,
		"saturday":              &saturday,
		"sunday":                &sunday,
		"start_date":            &startDate,
		"end_date":              &endDate,
	}

	// Loop through fields and parse each one
	for field, target := range fieldMappings {
		msg := lib.ParseStringToPrimitive(rawCalendar[field], target)
		if msg != "" {
			message.Message = msg
			message.Field = field
			services.AppMessageService.AddMessage(message)
			return types.Calendar{}
		}
	}

	calendar.ServiceId = serviceId
	calendar.StartDate = startDate
	calendar.EndDate = endDate
	calendar.Monday = monday == "1"
	calendar.Tuesday = tuesday == "1"
	calendar.Wednesday = wednesday == "1"
	calendar.Thursday = thursday == "1"
	calendar.Friday = friday == "1"
	calendar.Saturday = saturday == "1"
	calendar.Sunday = sunday == "1"

	return calendar
}


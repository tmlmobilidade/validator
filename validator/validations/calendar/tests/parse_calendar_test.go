package calendar

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/calendar"
	"testing"
)

func TestParseCalendar_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	raw := map[string]string{
		"service_id": "S1",
		"start_date": "20240101",
		"end_date":   "20241231",
		"monday":     "1",
		"tuesday":    "1",
		"wednesday":  "1",
		"thursday":   "1",
		"friday":     "1",
		"saturday":   "0",
		"sunday":     "0",
	}
	gtfs := &types.Gtfs{}
	calendar := validations.ParseCalendar(raw, row, gtfs)

	if calendar.ServiceId != "S1" {
		t.Errorf("Expected ServiceId 'S1', got '%s'", calendar.ServiceId)
	}
	if calendar.StartDate != "20240101" {
		t.Errorf("Expected StartDate '20240101', got '%s'", calendar.StartDate)
	}
	if calendar.EndDate != "20241231" {
		t.Errorf("Expected EndDate '20241231', got '%s'", calendar.EndDate)
	}
	if !calendar.Monday {
		t.Error("Expected Monday to be true")
	}
	if !calendar.Tuesday {
		t.Error("Expected Tuesday to be true")
	}
	if !calendar.Wednesday {
		t.Error("Expected Wednesday to be true")
	}
	if !calendar.Thursday {
		t.Error("Expected Thursday to be true")
	}
	if !calendar.Friday {
		t.Error("Expected Friday to be true")
	}
	if calendar.Saturday {
		t.Error("Expected Saturday to be false")
	}
	if calendar.Sunday {
		t.Error("Expected Sunday to be false")
	}
}

func TestParseCalendar_InvalidBooleanValues(t *testing.T) {
	services.AppMessageService.Clear()
	row := 3
	raw := map[string]string{
		"service_id": "S1",
		"start_date": "20240101",
		"end_date":   "20241231",
		"monday":     "2", // Invalid boolean value
		"tuesday":    "1",
		"wednesday":  "1",
		"thursday":   "2", // Invalid boolean value
		"friday":     "1",
		"saturday":   "0",
		"sunday":     "0",
	}
	gtfs := &types.Gtfs{}
	_ = validations.ParseCalendar(raw, row, gtfs)

	assertion := lib.AssertionMessage{
		Expected: 2,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid boolean value (should error)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestParseCalendar_AllDaysDisabled(t *testing.T) {
	services.AppMessageService.Clear()
	row := 4
	raw := map[string]string{
		"service_id": "S1",
		"start_date": "20240101",
		"end_date":   "20241231",
		"monday":     "0",
		"tuesday":    "0",
		"wednesday":  "0",
		"thursday":   "0",
		"friday":     "0",
		"saturday":   "0",
		"sunday":     "0",
	}
	gtfs := &types.Gtfs{}
	calendar := validations.ParseCalendar(raw, row, gtfs)

	if calendar.Monday || calendar.Tuesday || calendar.Wednesday || calendar.Thursday || 
	   calendar.Friday || calendar.Saturday || calendar.Sunday {
		t.Error("Expected all days to be false")
	}
}

func TestParseCalendar_AllDaysEnabled(t *testing.T) {
	services.AppMessageService.Clear()
	row := 5
	raw := map[string]string{
		"service_id": "S1",
		"start_date": "20240101",
		"end_date":   "20241231",
		"monday":     "1",
		"tuesday":    "1",
		"wednesday":  "1",
		"thursday":   "1",
		"friday":     "1",
		"saturday":   "1",
		"sunday":     "1",
	}
	gtfs := &types.Gtfs{}
	calendar := validations.ParseCalendar(raw, row, gtfs)

	if !calendar.Monday || !calendar.Tuesday || !calendar.Wednesday || !calendar.Thursday || 
	   !calendar.Friday || !calendar.Saturday || !calendar.Sunday {
		t.Error("Expected all days to be true")
	}
}

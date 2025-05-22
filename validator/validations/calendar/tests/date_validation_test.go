package calendar

import (
	"main/lib"
	"main/services"
	validations "main/validations/calendar/validations"
	"testing"
)

func TestDateValidation_Required(t *testing.T) {
	validations.DateValidation("", "start_date", 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Service date is required (should error if empty)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDateValidation_Valid(t *testing.T) {
	validations.DateValidation("20240101", "start_date", 2)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid service date (should not error)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestDateValidation_InvalidFormat(t *testing.T) {
	validations.DateValidation("2024-01-01", "end_date", 3)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid service date format (should error)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
} 
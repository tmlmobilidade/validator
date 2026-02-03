package calendar_dates

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/calendar_dates/validations"
	"testing"
)

func TestDateValidation_InvalidFormat(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	calendarDate := &types.CalendarDates{
		Date:          "2024-01-01",
		ExceptionType: nil,
		ServiceId:     "S1",
	}
	validations.DateValidation(calendarDate, row)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid date format should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAllDateValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetDateValidOptions()
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("date") {
		if tc.Name == "Recommended_Missing" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var date string
			if tc.Value != nil {
				date = validOptions[0]
			}
			validations.DateValidation(&types.CalendarDates{Date: date}, tc.Row)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
}

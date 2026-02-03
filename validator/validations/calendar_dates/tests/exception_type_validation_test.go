package calendar_dates

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/calendar_dates/validations"
	"testing"
)

func TestAllExceptionTypeValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetExceptionTypeValidOptions()
	dateValid := test_helpers.GetDateValidOptions()
	for _, tc := range test_helpers.GetGenericEnumIntTestCases("exception_type", validOptions) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var exceptionType *int
			if tc.Name == "Invalid_Value" {
				exceptionType = nil
			} else if tc.Value != nil {
				if ptr, ok := tc.Value.(*int); ok {
					exceptionType = ptr
				} else {
					exceptionType = tc.Value.(*int)
				}
			} else {
				exceptionType = nil
			}

			calendarDate := &types.CalendarDates{
				Date:          dateValid[0],
				ExceptionType: exceptionType,
				ServiceId:     "S1",
			}

			validations.ExceptionTypeValidation(calendarDate, tc.Row, nil)
			if tc.Name == "Missing_Value_Required" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, 1, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}
}

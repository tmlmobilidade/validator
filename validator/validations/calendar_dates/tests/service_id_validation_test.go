package calendar_dates

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/calendar_dates/validations"
	"testing"
)

func TestAllServiceIdValidationTestCases(t *testing.T) {
	dateValid := test_helpers.GetDateValidOptions()
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("service_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var serviceId string
			if tc.Value != nil {
				serviceId = *tc.Value
			}
			calendarDate := &types.CalendarDates{
				Date:          dateValid[0],
				ExceptionType: nil,
				ServiceId:     serviceId,
			}
			validations.ServiceIdValidation(calendarDate, tc.Row)
			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
		})
	}
}

func TestInvalidServiceIdValidation(t *testing.T) {
	dateValid := test_helpers.GetDateValidOptions()
	services.AppMessageService.Clear()
	calendarDate := &types.CalendarDates{
		Date:          dateValid[0],
		ExceptionType: nil,
		ServiceId:     "invalid",
	}
	validations.ServiceIdValidation(calendarDate, 1)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Invalid service_id should error")
}

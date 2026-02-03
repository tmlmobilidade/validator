package calendar_dates

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/calendar_dates/validations"
	"testing"
)

func TestAllServiceIdValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetDateValidOptions()
	invalidOptions := test_helpers.GetInvalidDateOptions()
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("service_id") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var serviceId string
			if tc.Name == "Invalid_Value" {
				serviceId = invalidOptions[0]
			} else if tc.Value != nil {
				serviceId = validOptions[0]
			} else {
				serviceId = ""
			}
			calendarDate := &types.CalendarDates{
				Date:          validOptions[0],
				ExceptionType: nil,
				ServiceId:     serviceId,
			}
			validations.ServiceIdValidation(calendarDate, tc.Row)
			if tc.Name == "Missing_Value_Required" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, 1, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}
}

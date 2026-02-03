package calendar

import (
	"main/lib/test_helpers"
	"main/services"
	validations "main/validations/calendar/validations"
	"testing"
)

func TestAllDateValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetDateValidOptions()
	invalidOptions := test_helpers.GetInvalidDateOptions()
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("start_date") {
		if tc.Name == "Recommended_Missing" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var date string
			if tc.Name == "Invalid_Value" {
				date = invalidOptions[0]
			} else if tc.Value != nil {
				date = validOptions[0]
			} else {
				date = ""
			}
			validations.DateValidation(date, "start_date", tc.Row)
			if tc.Name == "Recommended_Missing" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, 1, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}
}

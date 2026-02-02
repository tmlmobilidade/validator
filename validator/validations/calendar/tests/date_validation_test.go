package calendar

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	validations "main/validations/calendar/validations"
	"testing"
)

func TestDateValidation_InvalidFormat(t *testing.T) {
	validations.DateValidation("2024-01-01", "end_date", 3)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid service date format (should error)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestAllDateValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetDateValidOptions()
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("start_date") {
		if tc.Name == "Recommended_Missing" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var date string
			if tc.Value != nil {
				date = validOptions[0]
			}
			validations.DateValidation(date, "start_date", tc.Row)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
}

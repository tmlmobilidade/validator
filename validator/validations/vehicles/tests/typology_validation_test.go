package tests

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/vehicles/validations"
	"testing"
)

func TestAllTypologyValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetTypologyValidOptions()
	invalidOptions := test_helpers.GetFloat64InvalidOptions()
	for _, tc := range test_helpers.GetGenericEnumFloat64TestCases("typology", validOptions) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var typologyValue *float64
			if tc.Name == "Invalid_Option" {
				typologyValue = &invalidOptions[0]
			} else if tc.Value != nil {
				typologyValue = &validOptions[0]
			}

			validations.TypologyValidation(&types.Vehicle{Typology: typologyValue}, tc.Row, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
		})
	}
}

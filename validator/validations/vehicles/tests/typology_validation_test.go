package tests

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/vehicles/validations"
	"strconv"
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

			var vehicle *types.Vehicle
			if typologyValue != nil {
				typologyString := strconv.FormatFloat(float64(*typologyValue), 'f', -1, 64)
				vehicle = &types.Vehicle{Typology: lib.Ptr(typologyString)}
			} else {
				vehicle = &types.Vehicle{}
			}

			validations.TypologyValidation(vehicle, tc.Row, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
		})
	}
}

package stops

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestAllStopNameValidationTestCases(t *testing.T) {
	locationTypeValidOptions := test_helpers.GetLocationTypeValidOptions()

	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("stop_name") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			stop := &types.Stop{StopName: tc.Value, LocationType: &locationTypeValidOptions[0]}

			if tc.Name == "Required" {
				stop = &types.Stop{StopName: nil, LocationType: &locationTypeValidOptions[0]}
			}
			if tc.Name == "Invalid_Value" {
				stop = &types.Stop{StopName: lib.Ptr(""), LocationType: &locationTypeValidOptions[0]}
			}

			validations.StopNameValidation(stop, tc.Row, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
		})
	}

	for _, tc := range test_helpers.GetGenericSeverityTestCases("stop_name") {
		if tc.Name == "Severity_Forbidden_Missing" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			stop := &types.Stop{StopName: nil, LocationType: &locationTypeValidOptions[0]}
			validations.StopNameValidation(stop, tc.Row, &types.StopsRules{StopName: types.RuleConfig{Severity: tc.Severity}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name, types.SEVERITY_WARNING)
		})
	}

	t.Run("Optional_LocationType3_Missing", func(t *testing.T) {
		services.AppMessageService.Clear()
		stop := &types.Stop{StopName: nil, LocationType: &locationTypeValidOptions[3]}
		validations.StopNameValidation(stop, 2, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Optional_LocationType3_Missing", types.SEVERITY_ERROR)
	})

}

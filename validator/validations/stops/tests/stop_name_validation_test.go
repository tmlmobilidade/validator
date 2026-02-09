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
	// Only location_types 0 (stop/platform), 1 (station), 2 (entrance/exit) are required for stop_name
	locationTypeValidOptions := test_helpers.GetLocationTypeValidOptions()

	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("stop_name") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			// Default test stop for required field cases
			stop := &types.Stop{StopName: nil, LocationType: lib.Ptr(locationTypeValidOptions[0])}

			if tc.Name == "Required" {
				stop = &types.Stop{StopName: nil, LocationType: lib.Ptr(locationTypeValidOptions[0])}
			}
			if tc.Name == "Invalid_Value" {
				stop = &types.Stop{StopName: lib.Ptr(""), LocationType: lib.Ptr(locationTypeValidOptions[0])}
			}

			validations.StopNameValidation(stop, tc.Row, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
		})
	}

	// Test severity cases
	for _, tc := range test_helpers.GetGenericSeverityTestCases("stop_name") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			// Should be a "required" location_type for severity to apply to absence
			stop := &types.Stop{StopName: nil, LocationType: lib.Ptr(0)}
			validations.StopNameValidation(stop, tc.Row, &types.StopsRules{StopName: types.RuleConfig{Severity: tc.Severity}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name, types.SEVERITY_WARNING)
		})
	}

	// Test forbidden disables required check (should not error even if missing)
	t.Run("Severity_Forbidden_Missing", func(t *testing.T) {
		services.AppMessageService.Clear()
		stop := &types.Stop{StopName: nil, LocationType: lib.Ptr(0)}
		validations.StopNameValidation(stop, 1, &types.StopsRules{StopName: types.RuleConfig{Severity: types.SEVERITY_FORBIDDEN}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Severity_Forbidden_Missing", types.SEVERITY_ERROR)
	})

	// Test that missing stop_name for optional location_types returns no error
	t.Run("Optional_LocationType3_Missing", func(t *testing.T) {
		services.AppMessageService.Clear()
		stop := &types.Stop{StopName: nil, LocationType: lib.Ptr(3)}
		validations.StopNameValidation(stop, 2, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Optional_LocationType3_Missing", types.SEVERITY_ERROR)
	})
}

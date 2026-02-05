package stops

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"strconv"
	"testing"
)

func TestAllStopNameValidationTestCases(t *testing.T) {
	// Test required location types (0, 1, 2)
	for _, locationType := range []int{0, 1, 2} {
		t.Run("Required_LocationType_"+strconv.Itoa(locationType), func(t *testing.T) {
			services.AppMessageService.Clear()
			lt := locationType
			stop := &types.Stop{
				LocationType: &lt,
				StopName:     nil,
			}
			validations.StopNameValidation(stop, 1, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Required_LocationType", types.SEVERITY_ERROR)
		})
	}

	// Test optional location types (3, 4)
	for _, locationType := range []int{3, 4} {
		t.Run("Optional_LocationType_"+strconv.Itoa(locationType), func(t *testing.T) {
			services.AppMessageService.Clear()
			lt := locationType
			stop := &types.Stop{
				LocationType: &lt,
				StopName:     nil,
			}
			validations.StopNameValidation(stop, 1, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Optional_LocationType", types.SEVERITY_ERROR)
		})
	}

	// Test valid input
	t.Run("ValidInput", func(t *testing.T) {
		services.AppMessageService.Clear()
		lt := 1
		name := "Central Station"
		stop := &types.Stop{
			LocationType: &lt,
			StopName:     &name,
		}
		validations.StopNameValidation(stop, 1, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "ValidInput", types.SEVERITY_ERROR)
	})

	// Test severity cases
	for _, tc := range test_helpers.GetGenericSeverityTestCases("stop_name") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			stop := &types.Stop{StopName: nil}
			validations.StopNameValidation(stop, tc.Row, &types.StopsRules{StopName: types.RuleConfig{Severity: tc.Severity}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name, types.SEVERITY_WARNING)
		})
	}
}

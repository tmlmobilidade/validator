package trips

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestAllBikesAllowedValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetBikesAllowedValidOptions()
	for _, tc := range test_helpers.GetGenericEnumIntTestCases("bikes_allowed", validOptions) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var bikesAllowed *int
			if tc.Value != nil {
				if ptr, ok := tc.Value.(*int); ok {
					bikesAllowed = ptr
				}
			}

			trip := &types.Trip{BikesAllowed: bikesAllowed}
			gtfs := &types.Gtfs{}
			validations.BikesAllowedValidation(trip, tc.Row, gtfs, &types.TripsRules{BikesAllowed: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name, types.SEVERITY_ERROR)
		})
	}

	for _, tc := range test_helpers.GetGenericSeverityTestCases("bikes_allowed") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			trip := &types.Trip{BikesAllowed: nil}
			gtfs := &types.Gtfs{}
			validations.BikesAllowedValidation(trip, tc.Row, gtfs, &types.TripsRules{BikesAllowed: types.RuleConfig{Severity: tc.Severity}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedWarnings, tc.Name, types.SEVERITY_WARNING)
		})
	}
}

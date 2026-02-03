package feed_info

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/feed_info/validations"
	"testing"
)

func TestAllFeedVersionValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("feed_version") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var severity types.Severity
			if tc.ExpectedCode == "feed_version_validation.required" {
				severity = types.SEVERITY_ERROR
			} else if tc.ExpectedCode == "feed_version_validation.recommended" {
				severity = types.SEVERITY_WARNING
			} else {
				severity = types.SEVERITY_ERROR
			}
			validations.FeedVersionValidation(&severity, &types.FeedInfo{FeedVersion: tc.Value}, tc.Row)
			if tc.Name == "Recommended_Missing" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, 1, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}
	for _, tc := range test_helpers.GetGenericSeverityTestCases("feed_version") {
		if tc.Name != "Severity_Ignore_Missing" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			validations.FeedVersionValidation(&tc.Severity, &types.FeedInfo{FeedVersion: nil}, tc.Row)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
}

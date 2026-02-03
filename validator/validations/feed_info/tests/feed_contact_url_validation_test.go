package feed_info

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/feed_info/validations"
	"testing"
)

func TestAllFeedContactUrlValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericUrlTestCases("feed_contact_url") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var severity types.Severity
			if tc.ExpectedCode == "feed_contact_url_validation.required" {
				severity = types.SEVERITY_ERROR
			} else if tc.ExpectedCode == "feed_contact_url_validation.recommended" {
				severity = types.SEVERITY_WARNING
			} else {
				severity = types.SEVERITY_ERROR
			}
			feedInfo := &types.FeedInfo{FeedContactUrl: tc.Url}
			validations.FeedContactUrlValidation(&severity, feedInfo, tc.Row)
			if tc.Name == "Recommended_Missing" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, 1, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}
}

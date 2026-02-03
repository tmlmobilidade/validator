package feed_info

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/feed_info/validations"
	"testing"
)

func TestAllFeedContactEmailValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("feed_contact_email") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var severity types.Severity
			if tc.ExpectedCode == "feed_contact_email_validation.required" {
				severity = types.SEVERITY_ERROR
			} else if tc.ExpectedCode == "feed_contact_email_validation.recommended" {
				severity = types.SEVERITY_WARNING
			} else {
				severity = types.SEVERITY_ERROR
			}
			feedInfo := &types.FeedInfo{FeedContactEmail: tc.Value}
			validations.FeedContactEmailValidation(&severity, feedInfo, tc.Row)
			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
		})
	}
}

func TestInvalidFeedContactEmailValidation(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	invalid := "notanemail"
	feedInfo := &types.FeedInfo{FeedContactEmail: &invalid}
	validations.FeedContactEmailValidation(&severity, feedInfo, 1)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Invalid feed_contact_email should error")
}

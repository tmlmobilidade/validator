package feed_info

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/feed_info/validations"
	"testing"
)

func TestAllFeedPublisherNameValidationTestCases(t *testing.T) {
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("feed_publisher_name") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			feedInfo := &types.FeedInfo{FeedPublisherName: tc.Value}
			validations.FeedPublisherNameValidation(feedInfo, tc.Row)
			if tc.Name == "Recommended_Missing" {
				test_helpers.AssertMessageCount(t, services.AppMessageService, 1, tc.Name)
			} else {
				test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
			}
		})
	}
}

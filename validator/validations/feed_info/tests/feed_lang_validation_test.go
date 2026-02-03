package feed_info

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/feed_info/validations"
	"testing"
)

func TestAllFeedLangValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetValidLanguageCodes()
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("feed_lang") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var feedLang *string
			if tc.Name == "Invalid_Value" {
				feedLang = lib.Ptr("notalang")
			} else if tc.Value != nil {
				feedLang = lib.Ptr(validOptions[0])
			} else {
				feedLang = nil
			}
			validations.FeedLangValidation(&types.FeedInfo{FeedLang: feedLang}, tc.Row)
			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
		})
	}
}

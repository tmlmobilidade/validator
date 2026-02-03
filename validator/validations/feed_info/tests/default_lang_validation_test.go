package feed_info

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/feed_info/validations"
	"testing"
)

func TestAllDefaultLangValidationTestCases(t *testing.T) {
	validOptions := test_helpers.GetValidLanguageCodes()
	for _, tc := range test_helpers.GetGenericRequiredFieldTestCases("default_lang") {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var severity types.Severity
			if tc.ExpectedErrors > 0 {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}
			var defaultLang *string
			if tc.Name == "Invalid_Value" {
				defaultLang = lib.Ptr("notalang")
			} else if tc.Value != nil {
				defaultLang = lib.Ptr(validOptions[0])
			} else {
				defaultLang = nil
			}
			validations.DefaultLangValidation(&severity, &types.FeedInfo{DefaultLang: defaultLang}, tc.Row)
			expectedTotalMessages := tc.ExpectedErrors + tc.ExpectedWarnings
			test_helpers.AssertMessageCount(t, services.AppMessageService, expectedTotalMessages, tc.Name)
		})
	}
	for _, tc := range test_helpers.GetGenericSeverityTestCases("default_lang") {
		if tc.Name != "Severity_Ignore_Missing" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			validations.DefaultLangValidation(&tc.Severity, &types.FeedInfo{DefaultLang: nil}, tc.Row)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
}

package feed_info

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/feed_info/validations"
	"testing"
)

func TestFeedContactUrlValidation_MissingUrl_ErrorSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	feedInfo := &types.FeedInfo{FeedContactUrl: nil}
	validations.FeedContactUrlValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing feed_contact_url with error severity should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedContactUrlValidation_MissingUrl_WarningSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_WARNING
	feedInfo := &types.FeedInfo{FeedContactUrl: nil}
	validations.FeedContactUrlValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Missing feed_contact_url with warning severity should error (recommended)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedContactUrlValidation_MissingUrl_IgnoreSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_IGNORE
	feedInfo := &types.FeedInfo{FeedContactUrl: nil}
	validations.FeedContactUrlValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing feed_contact_url with ignore severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedContactUrlValidation_InvalidUrl(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	invalid := "not-a-url"
	feedInfo := &types.FeedInfo{FeedContactUrl: &invalid}
	validations.FeedContactUrlValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid feed_contact_url should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedContactUrlValidation_ValidUrl(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	valid := "https://example.com/contact"
	feedInfo := &types.FeedInfo{FeedContactUrl: &valid}
	validations.FeedContactUrlValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid feed_contact_url should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
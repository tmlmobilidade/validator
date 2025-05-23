package feed_info

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/feed_info/validations"
	"testing"
)

func TestFeedContactEmailValidation_MissingEmail_ErrorSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	feedInfo := &types.FeedInfo{FeedContactEmail: nil}
	validations.FeedContactEmailValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing feed_contact_email with error severity should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedContactEmailValidation_MissingEmail_WarningSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_WARNING
	feedInfo := &types.FeedInfo{FeedContactEmail: nil}
	validations.FeedContactEmailValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Missing feed_contact_email with warning severity should error (recommended)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedContactEmailValidation_MissingEmail_IgnoreSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_IGNORE
	feedInfo := &types.FeedInfo{FeedContactEmail: nil}
	validations.FeedContactEmailValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing feed_contact_email with ignore severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedContactEmailValidation_InvalidEmail(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	invalid := "not-an-email"
	feedInfo := &types.FeedInfo{FeedContactEmail: &invalid}
	validations.FeedContactEmailValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid feed_contact_email should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedContactEmailValidation_ValidEmail(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	valid := "contact@example.com"
	feedInfo := &types.FeedInfo{FeedContactEmail: &valid}
	validations.FeedContactEmailValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid feed_contact_email should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

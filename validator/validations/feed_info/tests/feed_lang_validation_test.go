package feed_info

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/feed_info/validations"
	"testing"
)


func TestFeedLangValidation_MissingFeedLang(t *testing.T) {
	services.AppMessageService.Clear()
	feedInfo := &types.FeedInfo{FeedLang: nil}
	validations.FeedLangValidation(feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing feed_lang should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedLangValidation_InvalidFeedLang(t *testing.T) {
	services.AppMessageService.Clear()
	invalid := "xx"
	feedInfo := &types.FeedInfo{FeedLang: &invalid}
	validations.FeedLangValidation(feedInfo, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid feed_lang should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedLangValidation_ValidFeedLang(t *testing.T) {
	services.AppMessageService.Clear()
	valid := "en"
	feedInfo := &types.FeedInfo{FeedLang: &valid}
	validations.FeedLangValidation(feedInfo, 3)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid feed_lang should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
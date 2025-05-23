package feed_info

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/feed_info/validations"
	"testing"
)

// --- FeedPublisherNameValidation ---
func TestFeedPublisherNameValidation_MissingPublisherName(t *testing.T) {
	services.AppMessageService.Clear()
	feedInfo := &types.FeedInfo{FeedPublisherName: nil}
	validations.FeedPublisherNameValidation(feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing feed_publisher_name should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedPublisherNameValidation_ValidPublisherName(t *testing.T) {
	services.AppMessageService.Clear()
	name := "Transit Co"
	feedInfo := &types.FeedInfo{FeedPublisherName: &name}
	validations.FeedPublisherNameValidation(feedInfo, 2)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid feed_publisher_name should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
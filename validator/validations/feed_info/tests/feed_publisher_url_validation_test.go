package feed_info

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/feed_info/validations"
	"testing"
)

func TestFeedPublisherUrlValidation_MissingPublisherUrl(t *testing.T) {
	services.AppMessageService.Clear()
	feedInfo := &types.FeedInfo{FeedPublisherUrl: nil}
	validations.FeedPublisherUrlValidation(feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing feed_publisher_url should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedPublisherUrlValidation_InvalidPublisherUrl(t *testing.T) {
	services.AppMessageService.Clear()
	invalid := "not_a_url"
	feedInfo := &types.FeedInfo{FeedPublisherUrl: &invalid}
	validations.FeedPublisherUrlValidation(feedInfo, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid feed_publisher_url should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedPublisherUrlValidation_ValidPublisherUrl(t *testing.T) {
	services.AppMessageService.Clear()
	url := "https://transit.example.com"
	feedInfo := &types.FeedInfo{FeedPublisherUrl: &url}
	validations.FeedPublisherUrlValidation(feedInfo, 3)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid feed_publisher_url should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}
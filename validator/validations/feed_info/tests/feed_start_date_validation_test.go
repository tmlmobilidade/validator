package feed_info

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/feed_info/validations"
	"testing"
)

func TestFeedStartDateValidation_MissingDate_ErrorSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	feedInfo := &types.FeedInfo{FeedStartDate: nil}
	validations.FeedStartDateValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing feed_start_date with error severity should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedStartDateValidation_MissingDate_WarningSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_WARNING
	feedInfo := &types.FeedInfo{FeedStartDate: nil}
	validations.FeedStartDateValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Missing feed_start_date with warning severity should error (recommended)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedStartDateValidation_MissingDate_IgnoreSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_IGNORE
	feedInfo := &types.FeedInfo{FeedStartDate: nil}
	validations.FeedStartDateValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing feed_start_date with ignore severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedStartDateValidation_InvalidDate(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	invalid := "2023-01-01" // not in YYYYMMDD format
	feedInfo := &types.FeedInfo{FeedStartDate: &invalid}
	validations.FeedStartDateValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid feed_start_date should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFeedStartDateValidation_ValidDate(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	valid := "20240101"
	feedInfo := &types.FeedInfo{FeedStartDate: &valid}
	validations.FeedStartDateValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid feed_start_date should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
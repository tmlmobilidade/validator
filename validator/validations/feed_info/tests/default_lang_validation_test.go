package feed_info

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/feed_info/validations"
	"testing"
)

func TestDefaultLangValidation_MissingLang_ErrorSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	feedInfo := &types.FeedInfo{DefaultLang: nil}
	validations.DefaultLangValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing default_lang with error severity should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDefaultLangValidation_MissingLang_WarningSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_WARNING
	feedInfo := &types.FeedInfo{DefaultLang: nil}
	validations.DefaultLangValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalWarnings,
		Message: "Missing default_lang with warning severity should error (recommended)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDefaultLangValidation_MissingLang_IgnoreSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_IGNORE
	feedInfo := &types.FeedInfo{DefaultLang: nil}
	validations.DefaultLangValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing default_lang with ignore severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDefaultLangValidation_InvalidLang(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	invalid := "notalang"
	feedInfo := &types.FeedInfo{DefaultLang: &invalid}
	validations.DefaultLangValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid default_lang should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestDefaultLangValidation_ValidLang(t *testing.T) {
	services.AppMessageService.Clear()
	severity := types.SEVERITY_ERROR
	valid := "en"
	feedInfo := &types.FeedInfo{DefaultLang: &valid}
	validations.DefaultLangValidation(&severity, feedInfo, 1)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid default_lang should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
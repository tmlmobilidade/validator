package lib

import (
	"main/types"
	"testing"
)

// TestMessageServiceInterface is the interface for test message services
// This avoids import cycles by not importing services directly
type TestMessageServiceInterface interface {
	AddMessage(message types.Message)
	AddMessages(messages []types.Message)
	GetSummary() types.Summary
	Clear()
}

// AssertValidationError checks if the expected number of errors occurred
func AssertValidationError(t *testing.T, ms TestMessageServiceInterface, expectedErrors int, message string) {
	t.Helper()
	summary := ms.GetSummary()
	if summary.TotalErrors != expectedErrors {
		t.Errorf("%s: Expected %d errors, got %d", message, expectedErrors, summary.TotalErrors)
	}
}

// AssertValidationWarning checks if the expected number of warnings occurred
func AssertValidationWarning(t *testing.T, ms TestMessageServiceInterface, expectedWarnings int, message string) {
	t.Helper()
	summary := ms.GetSummary()
	if summary.TotalWarnings != expectedWarnings {
		t.Errorf("%s: Expected %d warnings, got %d", message, expectedWarnings, summary.TotalWarnings)
	}
}

// AssertNoValidationIssues checks that there are no errors or warnings
func AssertNoValidationIssues(t *testing.T, ms TestMessageServiceInterface, message string) {
	t.Helper()
	summary := ms.GetSummary()
	if summary.TotalErrors > 0 || summary.TotalWarnings > 0 {
		t.Errorf("%s: Expected no issues, got %d errors and %d warnings", message, summary.TotalErrors, summary.TotalWarnings)
	}
}

// CreateTestGtfsRules creates a test GtfsRules with default values
func CreateTestGtfsRules() *types.GtfsRules {
	return &types.GtfsRules{
		Agency:      types.AgencyRules{},
		Stops:       types.StopsRules{},
		Routes:      types.RoutesRules{},
		Trips:       types.TripsRules{},
		StopTimes:   types.StopTimesRules{},
		Calendar:    types.CalendarRules{},
		CalendarDates: types.CalendarDatesRules{},
	}
}

// AssertMessageCount checks the total number of messages (errors + warnings)
func AssertMessageCount(t *testing.T, ms TestMessageServiceInterface, expectedCount int, message string) {
	t.Helper()
	summary := ms.GetSummary()
	total := summary.TotalErrors + summary.TotalWarnings
	if total != expectedCount {
		t.Errorf("%s: Expected %d total messages, got %d (errors: %d, warnings: %d)", 
			message, expectedCount, total, summary.TotalErrors, summary.TotalWarnings)
	}
}

// AssertMessageContains checks if any message contains the given text
func AssertMessageContains(t *testing.T, ms TestMessageServiceInterface, containsText string, message string) {
	t.Helper()
	summary := ms.GetSummary()
	for _, msg := range summary.Messages {
		if msg.Message == containsText {
			return // Found it
		}
	}
	t.Errorf("%s: Expected to find message containing '%s', but no such message found", message, containsText)
}


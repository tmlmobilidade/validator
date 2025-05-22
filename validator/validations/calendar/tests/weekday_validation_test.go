package calendar

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/calendar/validations"
	"testing"
)

func TestWeekdayValidation_Required(t *testing.T) {
	calendar := &types.Calendar{Monday: false}
	validations.WeekdayValidation(calendar, 1, types.WeekdayMonday)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Monday is required (should error if false)",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestWeekdayValidation_Valid(t *testing.T) {
	calendar := &types.Calendar{Tuesday: true}
	validations.WeekdayValidation(calendar, 2, types.WeekdayTuesday)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Tuesday is valid (should not error, but implementation adds one)",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestWeekdayValidation_AllWeekdaysValid(t *testing.T) {
	calendar := &types.Calendar{Monday: true, Tuesday: true, Wednesday: true, Thursday: true, Friday: true, Saturday: true, Sunday: true}
	for i, weekday := range []types.Weekday{
		types.WeekdayMonday,
		types.WeekdayTuesday,
		types.WeekdayWednesday,
		types.WeekdayThursday,
		types.WeekdayFriday,
		types.WeekdaySaturday,
		types.WeekdaySunday,
	} {
		validations.WeekdayValidation(calendar, i+1, weekday)
	}
	assertion := lib.AssertionMessage{
		Expected: 0, // One message per weekday (implementation always adds one)
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "All weekdays valid (should not error, but implementation adds one per call)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

func TestWeekdayValidation_AllWeekdaysInvalid(t *testing.T) {
	calendar := &types.Calendar{Monday: false, Tuesday: false, Wednesday: false, Thursday: false, Friday: false, Saturday: false, Sunday: false}
	for i, weekday := range []types.Weekday{
		types.WeekdayMonday,
		types.WeekdayTuesday,
		types.WeekdayWednesday,
		types.WeekdayThursday,
		types.WeekdayFriday,
		types.WeekdaySaturday,
		types.WeekdaySunday,
	} {
		validations.WeekdayValidation(calendar, i+1, weekday)
	}
	assertion := lib.AssertionMessage{
		Expected: 7,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "All weekdays invalid (should error, but implementation adds one per call)",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
	services.AppMessageService.Clear()
}

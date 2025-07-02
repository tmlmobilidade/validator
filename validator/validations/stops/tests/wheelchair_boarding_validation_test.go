package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestWheelchairBoardingValidation_MissingWheelchairBoarding_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{WheelchairBoarding: nil}
	validations.WheelchairBoardingValidation(stop, 1, nil)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing wheelchair_boarding with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestWheelchairBoardingValidation_MissingWheelchairBoarding_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{WheelchairBoarding: nil}
	severity := types.SEVERITY_ERROR
	validations.WheelchairBoardingValidation(stop, 2, &types.StopsRules{WheelchairBoarding: types.RuleConfig{Severity: severity}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing wheelchair_boarding with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestWheelchairBoardingValidation_MissingWheelchairBoarding_SeverityWarning(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{WheelchairBoarding: nil}
	severity := types.SEVERITY_WARNING
	validations.WheelchairBoardingValidation(stop, 3, &types.StopsRules{WheelchairBoarding: types.RuleConfig{Severity: severity}})
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing wheelchair_boarding with severity WARNING should warn")
	}
}

func TestWheelchairBoardingValidation_InvalidValue(t *testing.T) {
	services.AppMessageService.Clear()
	val := 5
	stop := &types.Stop{WheelchairBoarding: &val}
	validations.WheelchairBoardingValidation(stop, 4, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid wheelchair_boarding value should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestWheelchairBoardingValidation_ValidValues(t *testing.T) {
	for _, v := range []int{0, 1, 2} {
		services.AppMessageService.Clear()
		val := v
		stop := &types.Stop{WheelchairBoarding: &val}
		validations.WheelchairBoardingValidation(stop, 5+v, nil)
		assertion := lib.AssertionMessage{
			Expected: 0,
			Actual:   services.AppMessageService.GetSummary().TotalErrors,
			Message:  "Valid wheelchair_boarding value should not error",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Errorf("wheelchair_boarding %d: %s", v, assert)
		}
	}
}

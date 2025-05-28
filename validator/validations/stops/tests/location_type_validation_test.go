package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/stops/validations"
	"testing"
)

func TestLocationTypeValidation_MissingLocationType_DefaultSeverity(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{LocationType: nil}
	validations.LocationTypeValidation(nil, stop, 1)
	assertion := lib.AssertionMessage{
		Expected: 0, // Default severity is IGNORE, so should not error
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing location_type with default severity should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestLocationTypeValidation_MissingLocationType_SeverityError(t *testing.T) {
	services.AppMessageService.Clear()
	stop := &types.Stop{LocationType: nil}
	severity := types.SEVERITY_ERROR
	validations.LocationTypeValidation(&severity, stop, 2)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing location_type with severity ERROR should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestLocationTypeValidation_InvalidLocationType(t *testing.T) {
	services.AppMessageService.Clear()
	invalid := 99
	stop := &types.Stop{LocationType: &invalid}
	validations.LocationTypeValidation(nil, stop, 3)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid location_type should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestLocationTypeValidation_ValidLocationTypes(t *testing.T) {
	for _, val := range []int{0, 1, 2, 3, 4} {
		services.AppMessageService.Clear()
		lt := val
		stop := &types.Stop{LocationType: &lt}
		validations.LocationTypeValidation(nil, stop, 4+val)
		assertion := lib.AssertionMessage{
			Expected: 0,
			Actual: services.AppMessageService.GetSummary().TotalErrors,
			Message: "Valid location_type should not error",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Errorf("location_type %d: %s", val, assert)
		}
	}
} 
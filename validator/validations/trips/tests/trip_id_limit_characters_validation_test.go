package trips

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/trips/validations"
	"testing"
)

func TestAllTripIdLimitCharactersValidationTestCases(t *testing.T) {
	t.Run("Too_Long", func(t *testing.T) {
		services.AppMessageService.Clear()
		trip := &types.Trip{TripId: lib.Ptr("1234567890123456789012345678901234567890")}
		validations.TripIdLimitCharactersValidation(trip, 1, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Too_Long", types.SEVERITY_ERROR)
	})

	t.Run("Valid", func(t *testing.T) {
		services.AppMessageService.Clear()
		trip := &types.Trip{TripId: lib.Ptr("short_trip_id")}
		validations.TripIdLimitCharactersValidation(trip, 1, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Valid", types.SEVERITY_ERROR)
	})

	t.Run("Boundary_Exactly_32", func(t *testing.T) {
		services.AppMessageService.Clear()
		trip := &types.Trip{TripId: lib.Ptr("12345678901234567890123456789012")}
		validations.TripIdLimitCharactersValidation(trip, 1, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Boundary_Exactly_32", types.SEVERITY_ERROR)
	})

	t.Run("Nil_TripId", func(t *testing.T) {
		services.AppMessageService.Clear()
		trip := &types.Trip{TripId: nil}
		validations.TripIdLimitCharactersValidation(trip, 1, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Nil_TripId", types.SEVERITY_ERROR)
	})
}

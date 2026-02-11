package fare_media

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/fare_media/validations"
	"testing"
)

func TestAllGenericOptionsForFareMediaName(t *testing.T) {
	t.Run("Warning_Present", func(t *testing.T) {
		services.AppMessageService.Clear()
		row := 1
		fareMedia := &types.FareMedia{
			FareMediaId:   lib.Ptr("FM5"),
			FareMediaType: lib.Ptr(2),
			FareMediaName: nil,
		}
		validations.FareMediaNameValidation(fareMedia, row, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Warning_Recommended_Missing", types.SEVERITY_WARNING)
	})
	t.Run("Empty_Name", func(t *testing.T) {
		services.AppMessageService.Clear()
		row := 1
		fareMedia := &types.FareMedia{
			FareMediaId:   lib.Ptr("FM5"),
			FareMediaType: lib.Ptr(2),
			FareMediaName: lib.Ptr(""),
		}
		validations.FareMediaNameValidation(fareMedia, row, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Empty_Name", types.SEVERITY_WARNING)
	})
	t.Run("Invalid_Type", func(t *testing.T) {
		services.AppMessageService.Clear()
		row := 1
		fareMedia := &types.FareMedia{
			FareMediaId:   lib.Ptr("FM5"),
			FareMediaType: lib.Ptr(5),
			FareMediaName: lib.Ptr("Invalid Type"),
		}
		validations.FareMediaNameValidation(fareMedia, row, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Invalid_Type", types.SEVERITY_WARNING)
	})
	t.Run("Valid_Type", func(t *testing.T) {
		services.AppMessageService.Clear()
		row := 1
		fareMedia := &types.FareMedia{
			FareMediaId:   lib.Ptr("FM5"),
			FareMediaType: lib.Ptr(2),
			FareMediaName: lib.Ptr("Valid Type"),
		}
		validations.FareMediaNameValidation(fareMedia, row, nil)
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Valid_Type", types.SEVERITY_WARNING)
	})
}

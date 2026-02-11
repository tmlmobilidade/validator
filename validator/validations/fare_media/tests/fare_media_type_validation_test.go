package fare_media

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/fare_media/validations"
	"testing"
)

func TestAllGenericOptionsForFareMediaType(t *testing.T) {
	validOptions := []int{0, 1, 2, 3, 4}
	for _, tc := range test_helpers.GetGenericEnumIntTestCases("fare_media_type", validOptions) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()
			var fareMediaType *int
			if tc.Value != nil {
				if ptr, ok := tc.Value.(*int); ok && ptr != nil {
					fareMediaType = ptr
				}
			}
			fareMedia := &types.FareMedia{FareMediaType: fareMediaType}
			validations.FareMediaTypeValidation(fareMedia, tc.Row, nil)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name, types.SEVERITY_ERROR)
		})
	}
}

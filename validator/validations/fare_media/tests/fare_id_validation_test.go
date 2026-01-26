package fare_media

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_media/validations"
	"testing"
)

func TestFareIdValidation_MissingFareMediaId(t *testing.T) {
	services.AppMessageService.Clear()

	fareMedia := &types.FareMedia{
		FareMediaId: nil,
	}

	gtfs := &types.Gtfs{}

	validations.FareIdValidation(fareMedia, 1, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing fare_media_id should produce required error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareIdValidation_InvalidFareMediaId(t *testing.T) {
	services.AppMessageService.Clear()

	fareMedia := &types.FareMedia{
		FareMediaId: lib.Ptr("INVALID_ID"),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_media": {},
		},
	}

	validations.FareIdValidation(fareMedia, 2, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid fare_media_id should produce invalid error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareIdValidation_ValidFareMediaId(t *testing.T) {
	services.AppMessageService.Clear()

	fareMedia := &types.FareMedia{
		FareMediaId: lib.Ptr("VALID_ID"),
	}

	gtfs := &types.Gtfs{
		IdMap: types.GtfsIdMap{
			"fare_media": {
				"VALID_ID": {1},
			},
		},
	}

	validations.FareIdValidation(fareMedia, 3, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid fare_media_id should not produce errors",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

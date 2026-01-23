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
		FareMediaId: "",
	}

	gtfs := &types.Gtfs{}

	validations.FareIdValidation(fareMedia, 1, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 2, // required + invalid
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing fare_media_id should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareIdValidation_InvalidFareMediaId(t *testing.T) {
	services.AppMessageService.Clear()

	fareMedia := &types.FareMedia{
		FareMediaId: "INVALID_ID",
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
		Message:  "Invalid fare_media_id should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareIdValidation_ValidFareMediaId(t *testing.T) {
	services.AppMessageService.Clear()

	fareMedia := &types.FareMedia{
		FareMediaId: "VALID_ID",
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
		Message:  "Valid fare_media_id should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

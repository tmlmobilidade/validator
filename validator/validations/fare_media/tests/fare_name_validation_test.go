package fare_media

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/fare_media/validations"
	"testing"
)

func TestFareNameValidation_Type2WithEmptyName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   "FM1",
		FareMediaType: "2",
		FareMediaName: "",
	}
	validations.FareNameValidation(fareMedia, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Type 2 (transit cards) with empty fare_media_name should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareNameValidation_Type2WithName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   "FM1",
		FareMediaType: "2",
		FareMediaName: "Transit Card",
	}
	validations.FareNameValidation(fareMedia, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Type 2 (transit cards) with fare_media_name should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareNameValidation_Type4WithEmptyName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   "FM2",
		FareMediaType: "4",
		FareMediaName: "",
	}
	validations.FareNameValidation(fareMedia, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Type 4 (mobile apps) with empty fare_media_name should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareNameValidation_Type4WithName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   "FM2",
		FareMediaType: "4",
		FareMediaName: "Mobile App",
	}
	validations.FareNameValidation(fareMedia, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Type 4 (mobile apps) with fare_media_name should not warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestFareNameValidation_OtherTypesWithEmptyName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	otherTypes := []string{"1", "3", "5", "0", "99"}
	for _, fareMediaType := range otherTypes {
		fareMedia := &types.FareMedia{
			FareMediaId:   "FM3",
			FareMediaType: fareMediaType,
			FareMediaName: "",
		}
		validations.FareNameValidation(fareMedia, row, nil)
		assertion := lib.AssertionMessage{
			Expected: 0,
			Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
			Message:  "Other types with empty fare_media_name should not warn",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Error(assert)
		}
		services.AppMessageService.Clear()
	}
}

func TestFareNameValidation_OtherTypesWithName(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	otherTypes := []string{"1", "3", "5", "0", "99"}
	for _, fareMediaType := range otherTypes {
		fareMedia := &types.FareMedia{
			FareMediaId:   "FM3",
			FareMediaType: fareMediaType,
			FareMediaName: "Some Name",
		}
		validations.FareNameValidation(fareMedia, row, nil)
		assertion := lib.AssertionMessage{
			Expected: 0,
			Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
			Message:  "Other types with fare_media_name should not warn",
		}
		if assert := lib.Assert(assertion); assert != "" {
			t.Error(assert)
		}
		services.AppMessageService.Clear()
	}
}

func TestFareNameValidation_EmptyFareMediaType(t *testing.T) {
	services.AppMessageService.Clear()
	row := 1
	fareMedia := &types.FareMedia{
		FareMediaId:   "FM4",
		FareMediaType: "",
		FareMediaName: "",
	}
	validations.FareNameValidation(fareMedia, row, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors + services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Empty fare_media_type should not trigger validation",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

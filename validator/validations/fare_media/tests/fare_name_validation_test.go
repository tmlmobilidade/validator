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
		FareMediaId:   lib.Ptr("FM1"),
		FareMediaType: lib.Ptr(2),
		FareMediaName: lib.Ptr(""),
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
		FareMediaId:   lib.Ptr("FM1"),
		FareMediaType: lib.Ptr(2),
		FareMediaName: lib.Ptr("Transit Card"),
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
		FareMediaId:   lib.Ptr("FM2"),
		FareMediaType: lib.Ptr(4),
		FareMediaName: lib.Ptr(""),
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
		FareMediaId:   lib.Ptr("FM2"),
		FareMediaType: lib.Ptr(4),
		FareMediaName: lib.Ptr("Mobile App"),
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
	otherTypes := []int{1, 3, 5, 0, 99}
	for _, fareMediaType := range otherTypes {
		fareMedia := &types.FareMedia{
			FareMediaId:   lib.Ptr("FM3"),
			FareMediaType: lib.Ptr(fareMediaType),
			FareMediaName: lib.Ptr(""),
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
	otherTypes := []int{1, 3, 5, 0, 99}
	for _, fareMediaType := range otherTypes {
		fareMedia := &types.FareMedia{
			FareMediaId:   lib.Ptr("FM3"),
			FareMediaType: lib.Ptr(fareMediaType),
			FareMediaName: lib.Ptr("Some Name"),
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
		FareMediaId:   lib.Ptr("FM4"),
		FareMediaType: nil,
		FareMediaName: lib.Ptr(""),
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

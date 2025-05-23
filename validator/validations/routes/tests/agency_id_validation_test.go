package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestAgencyIdValidation_MissingAgencyId_MultipleAgencies(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{AgencyId: nil}
	agencyMap := map[string][]int{"A1": {1}, "A2": {2}}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"agency": agencyMap}}
	validations.AgencyIdValidation(route, 1, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing agency_id with multiple agencies should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAgencyIdValidation_InvalidAgencyId_MultipleAgencies(t *testing.T) {
	services.AppMessageService.Clear()
	agencyId := "INVALID"
	route := &types.Route{AgencyId: &agencyId}
	agencyMap := map[string][]int{"A1": {1}, "A2": {2}}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"agency": agencyMap}}
	validations.AgencyIdValidation(route, 2, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid agency_id with multiple agencies should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAgencyIdValidation_ValidAgencyId_MultipleAgencies(t *testing.T) {
	services.AppMessageService.Clear()
	agencyId := "A1"
	route := &types.Route{AgencyId: &agencyId}
	agencyMap := map[string][]int{"A1": {1}, "A2": {2}}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"agency": agencyMap}}
	validations.AgencyIdValidation(route, 3, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid agency_id with multiple agencies should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAgencyIdValidation_MissingAgencyId_OneAgency(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{AgencyId: nil}
	agencyMap := map[string][]int{"A1": {1}}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"agency": agencyMap}}
	validations.AgencyIdValidation(route, 4, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Missing agency_id with one agency should warn, not error",
	}
	if services.AppMessageService.GetSummary().TotalWarnings != 1 {
		t.Error("Missing agency_id with one agency should produce a warning")
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAgencyIdValidation_InvalidAgencyId_OneAgency(t *testing.T) {
	services.AppMessageService.Clear()
	agencyId := "INVALID"
	route := &types.Route{AgencyId: &agencyId}
	agencyMap := map[string][]int{"A1": {1}}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"agency": agencyMap}}
	validations.AgencyIdValidation(route, 5, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Invalid agency_id with one agency should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAgencyIdValidation_ValidAgencyId_OneAgency(t *testing.T) {
	services.AppMessageService.Clear()
	agencyId := "A1"
	route := &types.Route{AgencyId: &agencyId}
	agencyMap := map[string][]int{"A1": {1}}
	gtfs := &types.Gtfs{IdMap: map[string]map[string][]int{"agency": agencyMap}}
	validations.AgencyIdValidation(route, 6, gtfs)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual: services.AppMessageService.GetSummary().TotalErrors,
		Message: "Valid agency_id with one agency should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
} 
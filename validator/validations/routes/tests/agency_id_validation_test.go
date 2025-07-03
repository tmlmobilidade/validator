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
	gtfs := types.Gtfs{
		Agency: []types.AgencyRaw{
			{AgencyId: "A1"},
			{AgencyId: "A2"},
		}}

	validations.AgencyIdValidation(route, 1, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing agency_id with multiple agencies should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAgencyIdValidation_InvalidAgencyId_MultipleAgencies(t *testing.T) {
	services.AppMessageService.Clear()
	agencyId := "INVALID"
	route := &types.Route{AgencyId: &agencyId}
	gtfs := types.Gtfs{
		Agency: []types.AgencyRaw{
			{AgencyId: "A1"},
			{AgencyId: "A2"},
		},
	}

	validations.AgencyIdValidation(route, 2, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid agency_id with multiple agencies should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAgencyIdValidation_ValidAgencyId_MultipleAgencies(t *testing.T) {
	services.AppMessageService.Clear()
	agencyId := "A1"
	route := &types.Route{AgencyId: &agencyId}
	gtfs := types.Gtfs{
		Agency: []types.AgencyRaw{
			{AgencyId: "A1"},
			{AgencyId: "A2"},
		},
		IdMap: types.GtfsIdMap{
			"agency": {
				"A1": {1},
				"A2": {2},
			},
		},
	}

	validations.AgencyIdValidation(route, 3, gtfs, nil)

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid agency_id with multiple agencies should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAgencyIdValidation_MissingAgencyId_OneAgency(t *testing.T) {
	services.AppMessageService.Clear()
	route := &types.Route{AgencyId: nil}
	gtfs := types.Gtfs{
		Agency: []types.AgencyRaw{
			{AgencyId: "A1"},
		},
	}
	validations.AgencyIdValidation(route, 4, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing agency_id with one agency should warn, not error",
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
	gtfs := types.Gtfs{
		Agency: []types.AgencyRaw{
			{AgencyId: "A1"},
		},
	}
	validations.AgencyIdValidation(route, 5, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Invalid agency_id with one agency should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAgencyIdValidation_ValidAgencyId_OneAgency(t *testing.T) {
	services.AppMessageService.Clear()

	agencyId := "A1"
	route := &types.Route{AgencyId: &agencyId}
	gtfs := types.Gtfs{
		Agency: []types.AgencyRaw{
			{AgencyId: "A1"},
		},
		IdMap: types.GtfsIdMap{
			"agency": {
				"A1": {1},
			},
		},
	}

	validations.AgencyIdValidation(route, 6, gtfs, nil)
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid agency_id with one agency should not error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAgencyIdValidation_Severity_Warning(t *testing.T) {
	services.AppMessageService.Clear()

	route := &types.Route{}

	validations.AgencyIdValidation(route, 7, types.Gtfs{}, &types.RoutesRules{AgencyId: types.RuleConfig{Severity: types.SEVERITY_WARNING}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalWarnings,
		Message:  "Missing agency_id should warn",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAgencyIdValidation_Severity_Error(t *testing.T) {
	services.AppMessageService.Clear()

	route := &types.Route{}

	validations.AgencyIdValidation(route, 8, types.Gtfs{}, &types.RoutesRules{AgencyId: types.RuleConfig{Severity: types.SEVERITY_ERROR}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing agency_id should error",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

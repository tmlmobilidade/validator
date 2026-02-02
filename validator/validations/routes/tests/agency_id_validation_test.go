package routes

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestAllAgencyIdValidationTestCases(t *testing.T) {
	fieldName := "agency_id"

	for _, tc := range test_helpers.GetGenericIdTestCases(fieldName) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			var severity types.Severity
			if tc.ExpectedCode == "agency_id_required" {
				severity = types.SEVERITY_ERROR
			} else {
				severity = types.SEVERITY_WARNING
			}

			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"agency": tc.ExistingIds}}.ToGtfs()
			validations.AgencyIdValidation(&types.Route{AgencyId: tc.Id}, tc.Row, gtfs, &types.RoutesRules{AgencyId: types.RuleConfig{Severity: severity}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
}

func TestAgencyIdValidationTableCountUpperThan2(t *testing.T) {
	services.AppMessageService.Clear()
	gtfs := test_helpers.MockGtfs{TableCounts: map[string]int{"agency": 2}}.ToGtfs()
	validations.AgencyIdValidation(&types.Route{AgencyId: nil}, 1, gtfs, &types.RoutesRules{AgencyId: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Agency ID is required",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestAgencyIdValidationTableCountEqual1(t *testing.T) {
	services.AppMessageService.Clear()
	gtfs := test_helpers.MockGtfs{TableCounts: map[string]int{"agency": 1}}.ToGtfs()
	validations.AgencyIdValidation(&types.Route{AgencyId: nil}, 1, gtfs, &types.RoutesRules{AgencyId: types.RuleConfig{Severity: types.SEVERITY_WARNING}})
	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Agency ID is recommended",
	}
	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

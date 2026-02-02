package agency

import (
	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/agency/validations"
	"testing"
)

func TestAgencyIdValidation(t *testing.T) {
	fieldName := "agency_id"

	for _, tc := range test_helpers.GetGenericIdTestCases(fieldName) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			agency := &types.Agency{AgencyId: tc.Id}

			// Create a mock GTFS with the existing ID data
			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"agency": tc.ExistingIds}}.ToGtfs()
			validations.AgencyIdValidation(agency, tc.Row, gtfs, &types.AgencyRules{AgencyId: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, "Agency ID validation")
		})
	}
}

func TestAgencyIdValidationTableCountUpperThan2(t *testing.T) {
	services.AppMessageService.Clear()
	gtfs := test_helpers.MockGtfs{TableCounts: map[string]int{"agency": 2}}.ToGtfs()
	validations.AgencyIdValidation(&types.Agency{AgencyId: nil}, 1, gtfs, &types.AgencyRules{AgencyId: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
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

	validations.AgencyIdValidation(
		&types.Agency{AgencyId: nil},
		1,
		gtfs,
		&types.AgencyRules{
			AgencyId: types.RuleConfig{Severity: types.SEVERITY_WARNING},
		},
	)

	summary := services.AppMessageService.GetSummary()

	if assert := lib.Assert(lib.AssertionMessage{
		Expected: 0,
		Actual:   summary.TotalErrors,
		Message:  "Should not error when agency count == 1",
	}); assert != "" {
		t.Error(assert)
	}

	if assert := lib.Assert(lib.AssertionMessage{
		Expected: 1,
		Actual:   summary.TotalWarnings,
		Message:  "Should warn when agency count == 1",
	}); assert != "" {
		t.Error(assert)
	}
}

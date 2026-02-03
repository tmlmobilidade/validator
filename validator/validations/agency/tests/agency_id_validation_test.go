package agency

import (
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
	t.Run("TableCountUpperThan2", func(t *testing.T) {
		services.AppMessageService.Clear()
		gtfs := test_helpers.MockGtfs{TableCounts: map[string]int{"agency": 2}}.ToGtfs()
		validations.AgencyIdValidation(&types.Agency{AgencyId: nil}, 1, gtfs, &types.AgencyRules{AgencyId: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Agency ID validation")
	})
	t.Run("TableCountEqual1", func(t *testing.T) {
		services.AppMessageService.Clear()
		gtfs := test_helpers.MockGtfs{TableCounts: map[string]int{"agency": 1}}.ToGtfs()
		validations.AgencyIdValidation(&types.Agency{AgencyId: nil}, 1, gtfs, &types.AgencyRules{AgencyId: types.RuleConfig{Severity: types.SEVERITY_WARNING}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Agency ID validation")
	})
}

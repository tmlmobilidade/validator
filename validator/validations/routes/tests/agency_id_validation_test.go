package routes

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestAllAgencyIdValidationTestCases(t *testing.T) {
	fieldName := "agency_id"

	for _, tc := range test_helpers.GetGenericForeignKeyTestCases(fieldName) {
		if tc.Name == "ForeignKey_Invalid" {
			continue
		}
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"agency": map[string][]int{*tc.Id: {1}}}}.ToGtfs()
			validations.AgencyIdValidation(&types.Route{AgencyId: tc.Id}, tc.Row, gtfs, &types.RoutesRules{AgencyId: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
	t.Run("TestTableCountUpperThan2", func(t *testing.T) {
		services.AppMessageService.Clear()
		gtfs := test_helpers.MockGtfs{TableCounts: map[string]int{"agency": 2}}.ToGtfs()
		validations.AgencyIdValidation(&types.Route{AgencyId: nil}, 1, gtfs, &types.RoutesRules{AgencyId: types.RuleConfig{Severity: types.SEVERITY_ERROR}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Agency ID is required")
	})
	t.Run("TestTableCountEqual1", func(t *testing.T) {
		services.AppMessageService.Clear()
		gtfs := test_helpers.MockGtfs{TableCounts: map[string]int{"agency": 1}}.ToGtfs()
		validations.AgencyIdValidation(&types.Route{AgencyId: nil}, 1, gtfs, &types.RoutesRules{AgencyId: types.RuleConfig{Severity: types.SEVERITY_WARNING}})
		test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "Agency ID is recommended")
	})
}

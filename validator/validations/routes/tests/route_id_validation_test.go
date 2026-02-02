package routes

import (
	"main/lib/test_helpers"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestAllRouteIdValidationTestCases(t *testing.T) {
	fieldName := "route_id"
	for _, tc := range test_helpers.GetGenericIdTestCases(fieldName) {
		t.Run(tc.Name, func(t *testing.T) {
			services.AppMessageService.Clear()

			gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"routes": tc.ExistingIds}}.ToGtfs()
			validations.RouteIdValidation(&types.Route{RouteId: tc.Id}, tc.Row, &gtfs)
			test_helpers.AssertMessageCount(t, services.AppMessageService, tc.ExpectedErrors, tc.Name)
		})
	}
}

func TestRouteIdValidation_Forbidden(t *testing.T) {
	services.AppMessageService.Clear()
	routeId := "R1"
	route := &types.Route{RouteId: &routeId}
	gtfs := test_helpers.MockGtfs{IdMapData: types.GtfsIdMap{"routes": {"R1": {1, 2}}}}.ToGtfs()
	validations.RouteIdValidation(route, 2, &gtfs)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "Forbidden route_id should error")
}

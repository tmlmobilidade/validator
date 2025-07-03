package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	validations "main/validations/routes/validations"
	"testing"
)

func TestNetworkIdValidation_MissingNetworkId(t *testing.T) {
	services.AppMessageService.Clear()

	severity := types.SEVERITY_ERROR
	route := &types.Route{NetworkId: nil}
	gtfs := &types.Gtfs{}

	validations.NetworkIdValidation(route, 1, gtfs, &types.RoutesRules{NetworkId: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Missing network_id should error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestNetworkIdValidation_ForbiddenNetworkIdIfRouteNetworksExists(t *testing.T) {
	services.AppMessageService.Clear()

	severity := types.SEVERITY_ERROR
	networkId := "N1"
	route := &types.Route{NetworkId: &networkId}
	gtfs := &types.Gtfs{
		RouteNetwork: []types.RouteNetworkRaw{
			{NetworkId: "N1"},
		},
	}

	validations.NetworkIdValidation(route, 2, gtfs, &types.RoutesRules{NetworkId: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 1,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "network_id should be forbidden if route_networks.txt exists",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

func TestNetworkIdValidation_ValidInput(t *testing.T) {
	services.AppMessageService.Clear()

	severity := types.SEVERITY_ERROR
	networkId := "N2"
	route := &types.Route{NetworkId: &networkId}
	gtfs := &types.Gtfs{}

	validations.NetworkIdValidation(route, 3, gtfs, &types.RoutesRules{NetworkId: types.RuleConfig{Severity: severity}})

	assertion := lib.AssertionMessage{
		Expected: 0,
		Actual:   services.AppMessageService.GetSummary().TotalErrors,
		Message:  "Valid network_id should not error",
	}

	if assert := lib.Assert(assertion); assert != "" {
		t.Error(assert)
	}
}

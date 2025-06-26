package routes

import (
	"main/lib"
	"main/types"
	validations "main/validations/routes/validations"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Routes Validations...")

	for i, rawRoute := range gtfs.Route {
		route := validations.ParseRoutes(rawRoute, i)

		if route == (types.Route{}) {
			continue
		}

		// Validate route_id
		validations.RouteIdValidation(&route, i, &gtfs)

		// Validate agency_id
		validations.AgencyIdValidation(nil, &route, i, gtfs)

		// Validate route_short_name
		validations.RouteShortNameValidation(nil, &route, i)

		// Validate route_long_name
		validations.RouteLongNameValidation(nil, &route, i)

		// Validate route_desc
		validations.RouteDescValidation(nil, &route, i)
		
		// Validate route_type
		validations.RouteTypeValidation(&route, i)

		// Validate route_url
		validations.RouteUrlValidation(nil, &route, i, &gtfs)

		// Validate route_color
		validations.RouteColorValidation(nil, &route, i)

		// Validate route_text_color
		validations.RouteTextColorValidation(nil, &route, i)

		// Validate route_sort_order
		validations.RouteSortOrderValidation(nil, &route, i)

		// Validate continuous_drop_off
		validations.ContinuousDropOffValidation(nil, &route, i, &gtfs)

		// Validate continuous_pickup
		validations.ContinuousPickupValidation(nil, &route, i, &gtfs)

		// Validate network_id
		validations.NetworkIdValidation(nil, &route, i, &gtfs)
	}
}

package routes

import (
	"fmt"
	"main/config"
	"main/lib"
	"main/types"
	validations "main/validations/routes/validations"
	registry "main/validations"
)

func init() {
	registry.Register("routes", RunValidations)
}

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Routes Validations...")

	// Create progress tracker
	tracker := lib.CreateProgressTracker(gtfs, "routes.txt", config.ProgressThresholdLarge)

	err := gtfs.IterateRoutes(func(i int, rawRoute types.RouteRaw) error {
		tracker.Track()
		route := validations.ParseRoutes(rawRoute, i)

		if route == (types.Route{}) {
			return nil
		}

		var routeRules *types.RoutesRules
		if rules != nil {
			routeRules = &rules.Routes
		}

		// Validate route_id
		validations.RouteIdValidation(&route, i, &gtfs)

		// Validate agency_id
		validations.AgencyIdValidation(&route, i, gtfs, routeRules)

		// Validate route_short_name
		validations.RouteShortNameValidation(&route, i, routeRules)

		// Validate route_long_name
		validations.RouteLongNameValidation(&route, i, routeRules)

		// Validate route_desc
		validations.RouteDescValidation(&route, i, routeRules)

		// Validate route_type
		validations.RouteTypeValidation(&route, i, routeRules)

		// Validate route_url
		validations.RouteUrlValidation(&route, i, &gtfs, routeRules)

		// Validate route_color
		validations.RouteColorValidation(&route, i, routeRules)

		// Validate route_text_color
		validations.RouteTextColorValidation(&route, i, routeRules)

		// Validate route_sort_order
		validations.RouteSortOrderValidation(&route, i, routeRules)

		// Validate continuous_drop_off
		validations.ContinuousDropOffValidation(&route, i, &gtfs, routeRules)

		// Validate continuous_pickup
		validations.ContinuousPickupValidation(&route, i, &gtfs, routeRules)

		// Validate network_id
		validations.NetworkIdValidation(&route, i, &gtfs, routeRules)

		return nil
	})

	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error iterating routes: %v", err))
	} else {
		lib.AppLogger.Debug(fmt.Sprintf("Completed routes.txt validation: %d rows processed", tracker.GetProcessedCount()))
	}
}

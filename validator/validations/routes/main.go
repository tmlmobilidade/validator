package routes

import (
	"main/lib"
	"main/types"
	validations "main/validations/routes/validations"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Routes Validations...")

	for i, rawRoute := range gtfs.Files["routes"] {
		route := validations.ParseRoutes(rawRoute, i, &gtfs)

		if route == (types.Route{}) {
			continue
		}

		// Validate route_id
		validations.RouteIdValidation(&route, i, &gtfs)
	}
}

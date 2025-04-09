package routes

import (
	"main/src/lib"
	"main/src/services"
	"main/src/types"
	"strconv"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Routes Validations...")

	// Parsing Validation
	parseRouteValidation := NewParseRouteValidation(nil)
	routes, messages := parseRouteValidation.Validate(gtfs)
	for _, message := range messages {
		services.AppMessageService.AddMessage(message)
		lib.AppLogger.Error(message.Message)
	}

	// Print routes
	lib.AppLogger.Info("Total routes: parsed " + strconv.Itoa(len(routes)) + " routes")
}

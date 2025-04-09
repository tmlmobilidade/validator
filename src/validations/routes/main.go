package routes

import (
	"main/src/lib"
	"main/src/services"
	"main/src/types"
	"strconv"
)

func RunValidations(gtfsData types.Gtfs) {
	lib.AppLogger.Debug("Running Routes Validations...")

	// Parsing Validation
	parseRouteValidation := NewParseRouteValidation(nil)
	routes, messages := parseRouteValidation.Validate(gtfsData)
	for _, message := range messages {
		services.AppMessageService.AddMessage(message)
		lib.AppLogger.Error(message.Message)
	}

	// Print routes
	lib.AppLogger.Info("Total routes: parsed " + strconv.Itoa(len(routes)) + " routes")
}

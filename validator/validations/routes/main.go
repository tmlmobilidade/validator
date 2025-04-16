package routes

import (
	"main/lib"
	"main/services"
	"main/types"
	"strconv"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Routes Validations...")

	// Parsing Validation
	parseRouteValidation := NewParseRouteValidation(nil)
	routes, messages := parseRouteValidation.Validate(gtfs)
	for _, message := range messages {
		services.AppMessageService.AddMessage(message)
		lib.AppLogger.Error("[" + message.FileName + "] " + message.Message)
	}

	// Print routes
	lib.AppLogger.Info("Total routes: parsed " + strconv.Itoa(len(routes)) + " routes")
}

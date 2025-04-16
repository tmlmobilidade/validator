package stops

import (
	"main/lib"
	"main/services"
	"main/types"
	"strconv"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Validations for stops.txt")

	// Parsing Validation
	parseStopValidation := NewParseStopValidation(nil)
	stops, messages := parseStopValidation.Validate(gtfs)
	for _, message := range messages {
		services.AppMessageService.AddMessage(message)
		lib.AppLogger.Error("[" + message.FileName + "] " + message.Message)
	}

	// Print stops
	lib.AppLogger.Info("Total stops: parsed " + strconv.Itoa(len(stops)) + " stops")
}

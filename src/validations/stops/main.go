package stops

import (
	"main/src/lib"
	"main/src/services"
	"main/src/types"
)

func RunValidations(gtfsData types.Gtfs) {
	lib.AppLogger.Debug("Running Validations for stop.txt")

	// Parsing Validation
	parseStopValidation := NewParseStopValidation(nil)
	messages := parseStopValidation.Validate(gtfsData)
	for _, message := range messages {
		services.AppMessageService.AddMessage(message)
		lib.AppLogger.Error(message.Message)
	}
}

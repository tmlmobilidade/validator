package agency

import (
	"main/src/lib"
	"main/src/services"
	"main/src/types"
)

func RunValidations(gtfsData types.Gtfs) {
	lib.AppLogger.Debug("Running Validations for agency.txt")

	// Parsing Validation
	parseAgencyValidation := NewParseAgencyValidation(nil)
	agencies, messages := parseAgencyValidation.Validate(gtfsData)
	for _, message := range messages {
		services.AppMessageService.AddMessage(message)
		lib.AppLogger.Error(message.Message)
	}

	// Print agencies
	lib.PrintMap(agencies)
}

package agency

import (
	"main/src/lib"
	"main/src/services"
	"main/src/types"
	"strconv"
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
	lib.AppLogger.Info("Total agencies: parsed " + strconv.Itoa(len(agencies)) + " agencies")
}

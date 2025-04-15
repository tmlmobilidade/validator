package agency

import (
	"main/validator/lib"
	"main/validator/services"
	"main/validator/types"
	"strconv"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Validations for agency.txt")

	// Parsing Validation
	parseAgencyValidation := NewParseAgencyValidation(nil)
	agencies, messages := parseAgencyValidation.Validate(gtfs)
	for _, message := range messages {
		services.AppMessageService.AddMessage(message)
		lib.AppLogger.Error("[" + message.FileName + "] " + message.Message)
	}

	// Print agencies
	lib.AppLogger.Info("Total agencies: parsed " + strconv.Itoa(len(agencies)) + " agencies")
}

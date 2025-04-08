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
	messages := parseAgencyValidation.Validate(gtfsData)
	for _, message := range messages {
		services.AppMessageService.AddMessage(message)
	}
}

package stops

import (
	"main/src/services"
	"main/src/types"
)

func RunValidations(gtfsData types.Gtfs) {
	// Parsing Validation
	parseStopValidation := NewParseStopValidation(nil)
	messages := parseStopValidation.Validate(gtfsData)
	for _, message := range messages {
		services.AppMessageService.AddMessage(message)
	}
}

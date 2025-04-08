package stops

import (
	"main/src/services"
	"main/src/types"
)

func RunValidations(gtfsData types.Gtfs, messageService services.MessageService) {
	// Parsing Validation
	parseStopValidation := NewParseStopValidation(nil)
	messages := parseStopValidation.Validate(gtfsData)
	for _, message := range messages {
		messageService.AddMessage(message)
	}
}

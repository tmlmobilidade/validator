package stops

import (
	"main/src/types"
)

func RunValidations(gtfsData types.Gtfs) {
	// Parsing Validation
	parseStopValidation := NewParseStopValidation(nil)
	messages := parseStopValidation.Validate(gtfsData)
	for _, message := range messages {
		messageService.AddMessage(message)
	}
}

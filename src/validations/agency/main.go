package agency

import (
	"fmt"
	"main/src/types"
)

func RunValidations(gtfsData types.Gtfs) {
	fmt.Println("Running Validations for agency.txt")

	// Parsing Validation
	parseAgencyValidation := NewParseAgencyValidation(nil)
	messages := parseAgencyValidation.Validate(gtfsData)
	for _, message := range messages {
		messageService.AddMessage(message)
	}
}

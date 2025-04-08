package agency

import (
	"main/src/models"
	"main/src/services"
)

func RunValidations(gtfsData models.Gtfs, messageService services.MessageService) {
	// Parse Agency
	for idx, a := range gtfsData["agency"] {
		_, errors := ParseAgency(a)
		for _, error := range errors {
			messageService.AddMessage(services.Message{
				Field:    "agency",
				FileName: "agency.txt",
				Message:  error,
				Row:      idx,
				Severity: services.SEVERITY_ERROR,
			})
		}
	}
}

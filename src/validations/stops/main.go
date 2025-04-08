package stops

import (
	"fmt"
	"main/src/models"
	"main/src/services"
)

func RunValidations(gtfsData models.Gtfs, messageService services.MessageService) {
	fmt.Println("Running Stops Validations...")
	// Parse Agency
	for _, a := range gtfsData["stops"] {
		st, errors := ParseStop(a)
		if len(errors) > 0 {
			fmt.Println("Errors:", errors)
		} else {
			fmt.Printf("Stop: %+v\n", st)
		}
	}
}

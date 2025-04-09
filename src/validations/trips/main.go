package trips

import (
	"main/src/lib"
	"main/src/services"
	"main/src/types"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Trips Validations...")

	// Create validation with default severity
	validation := NewParseTripValidation(nil)

	// Run validation
	_, validationMessages := validation.Validate(gtfs)
	for _, message := range validationMessages {
		services.AppMessageService.AddMessage(message)
		lib.AppLogger.Error(message.Message)
	}
}

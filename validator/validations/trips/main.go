package trips

import (
	"main/validator/lib"
	"main/validator/services"
	"main/validator/types"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Trips Validations...")

	// Create validation with default severity
	validation := NewParseTripValidation(nil)

	// Run validation
	_, validationMessages := validation.Validate(gtfs)
	for _, message := range validationMessages {
		services.AppMessageService.AddMessage(message)
		lib.AppLogger.Error("[" + message.FileName + "] " + message.Message)
	}
}

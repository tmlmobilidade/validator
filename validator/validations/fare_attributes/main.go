package fare_attributes

import (
	"main/validator/lib"
	"main/validator/services"
	"main/validator/types"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Fare Attributes Validations...")

	// Create validation with default severity
	validation := NewParseFareAttributeValidation(nil)

	// Run validation
	_, validationMessages := validation.Validate(gtfs)
	for _, message := range validationMessages {
		services.AppMessageService.AddMessage(message)
		lib.AppLogger.Error("[" + message.FileName + "] " + message.Message)
	}
}

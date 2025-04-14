package calendar

import (
	"main/src/lib"
	"main/src/services"
	"main/src/types"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Calendar Validations...")

	// Create validation with default severity
	validation := NewParseCalendarValidation(nil)

	// Run validation
	_, validationMessages := validation.Validate(gtfs)
	for _, message := range validationMessages {
		services.AppMessageService.AddMessage(message)
		lib.AppLogger.Error("[" + message.FileName + "] " + message.Message)
	}
}

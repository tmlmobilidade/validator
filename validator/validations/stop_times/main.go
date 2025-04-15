package stop_times

import (
	"main/validator/lib"
	"main/validator/services"
	"main/validator/types"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running StopTimes Validations...")

	// Create validation with default severity
	validation := NewParseStopTimeValidation(nil)

	// Run validation
	_, validationMessages := validation.Validate(gtfs)
	for _, message := range validationMessages {
		services.AppMessageService.AddMessage(message)
		lib.AppLogger.Error("[" + message.FileName + "] " + message.Message)
	}
}

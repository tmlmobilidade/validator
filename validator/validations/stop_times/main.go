package stop_times

import (
	"main/lib"
	"main/types"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running StopTimes Validations...")

	// // Create validation with default severity
	// validation := NewParseStopTimeValidation(nil)

	// // Run validation
	// validationMessages := validation.Validate(gtfs)
	// for _, message := range validationMessages {
	// 	services.AppMessageService.AddMessage(message)
	// 	lib.AppLogger.Error("[" + message.FileName + "] " + message.Message)
	// }
}

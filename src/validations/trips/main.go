package trips

import (
	"main/src/lib"
	"main/src/types"
)

func RunValidations(gtfs types.Gtfs) (messages []types.Message) {
	lib.AppLogger.Debug("Running Trips Validations...")

	// Create validation with default severity
	validation := NewParseTripValidation(nil)

	// Run validation
	_, validationMessages := validation.Validate(gtfs)
	messages = append(messages, validationMessages...)

	return messages
}

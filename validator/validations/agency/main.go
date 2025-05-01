package agency

import (
	"main/lib"
	"main/services"
	"main/types"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Validations for agency.txt")

	for i, rawAgency := range gtfs.Files["agency"] {
		// Parse Agency Validation
		agency, msgs := ParseAgencyValidation(nil, rawAgency, i, &gtfs)
		if len(msgs) > 0 {
			for _, msg := range msgs {
				services.AppMessageService.AddMessage(msg)
			}
			break;
		}

		// Duplicate Agencies Validation
		msg := DuplicateAgenciesValidation(nil, &agency, i, &gtfs)
		if msg != nil {
			services.AppMessageService.AddMessage(*msg)
		}
	}

	lib.PrintMap(services.AppMessageService.GetSummary())
}

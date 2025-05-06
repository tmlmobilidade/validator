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
		agency := ParseAgencyValidation(rawAgency, i, &gtfs)

		// Duplicate Agencies Validation
		AgencyIdValidation(nil, &agency, i, &gtfs)

		// Validate Agency URL
		AgencyUrlValidation(&agency, i, &gtfs)

		// Validate Agency Timezone
		AgencyTimezoneValidation(&agency, i, &gtfs)

		// Validate Agency Lang
		AgencyLangValidation(nil, &agency, i, &gtfs)

		// Validate Agency Phone
		AgencyPhoneValidation(nil, &agency, i, &gtfs)

		// Validate Agency Fare URL
		AgencyFareUrlValidation(nil, &agency, i, &gtfs)

		// Validate Agency Email
		AgencyEmailValidation(nil, &agency, i, &gtfs)
		
	}

	lib.PrintMap(services.AppMessageService.GetSummary())
}

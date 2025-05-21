package agency

import (
	"main/lib"
	"main/types"
	validations "main/validations/agency/validations"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Validations for agency.txt")

	for i, rawAgency := range gtfs.Files["agency"] {
		// Parse Agency Validation
		agency := validations.ParseAgencyValidation(rawAgency, i, &gtfs)

		if agency == (types.Agency{}) {
			continue
		}

		// Duplicate Agencies Validation
		validations.AgencyIdValidation(nil, &agency, i, &gtfs)

		// Validate Agency URL
		validations.AgencyUrlValidation(&agency, i, &gtfs)

		// Validate Agency Timezone
		validations.AgencyTimezoneValidation(&agency, i, &gtfs)

		// Validate Agency Lang
		validations.AgencyLangValidation(nil, &agency, i, &gtfs)

		// Validate Agency Phone
		validations.AgencyPhoneValidation(nil, &agency, i, &gtfs)

		// Validate Agency Fare URL
		validations.AgencyFareUrlValidation(nil, &agency, i, &gtfs)

		// Validate Agency Email
		validations.AgencyEmailValidation(nil, &agency, i, &gtfs)
		
	}
}

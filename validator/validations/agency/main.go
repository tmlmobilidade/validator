package agency

import (
	"main/lib"
	"main/types"
	validations "main/validations/agency/validations"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Validations for agency.txt")

	for i, rawAgency := range gtfs.Agency {
		// Parse Agency Validation
		agency := validations.ParseAgency(rawAgency, i)

		if agency == (types.Agency{}) {
			continue
		}

		// Duplicate Agencies Validation
		validations.AgencyIdValidation(nil, &agency, i, gtfs)

		// Validate Agency URL
		validations.AgencyUrlValidation(&agency, i)

		// Validate Agency Timezone
		validations.AgencyTimezoneValidation(&agency, i)

		// Validate Agency Lang
		validations.AgencyLangValidation(nil, &agency, i)

		// Validate Agency Phone
		validations.AgencyPhoneValidation(nil, &agency, i)

		// Validate Agency Fare URL
		validations.AgencyFareUrlValidation(nil, &agency, i)

		// Validate Agency Email
		validations.AgencyEmailValidation(nil, &agency, i)
	}
}

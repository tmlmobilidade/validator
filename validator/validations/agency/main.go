package agency

import (
	"main/lib"
	"main/types"
	validations "main/validations/agency/validations"
)

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Validations for agency.txt")

	for i, rawAgency := range gtfs.Agency {
		// Parse Agency Validation
		agency := validations.ParseAgency(rawAgency, i)

		if agency == (types.Agency{}) {
			continue
		}

		// Duplicate Agencies Validation
		validations.AgencyIdValidation(&agency, i, gtfs, &rules.Agency)

		// Validate Agency Name
		validations.AgencyNameValidation(&agency, i, &rules.Agency)

		// Validate Agency URL
		validations.AgencyUrlValidation(&agency, i, &rules.Agency)

		// Validate Agency Timezone
		validations.AgencyTimezoneValidation(&agency, i, &rules.Agency)

		// Validate Agency Lang
		validations.AgencyLangValidation(&agency, i, &rules.Agency)

		// Validate Agency Phone
		validations.AgencyPhoneValidation(&agency, i, &rules.Agency)

		// Validate Agency Fare URL
		validations.AgencyFareUrlValidation(&agency, i, &rules.Agency)

		// Validate Agency Email
		validations.AgencyEmailValidation(&agency, i, &rules.Agency)
	}
}

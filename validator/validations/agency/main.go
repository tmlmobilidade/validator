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

		var agencyRules types.AgencyRules
		if rules != nil {
			agencyRules = rules.Agency
		}

		// Duplicate Agencies Validation
		validations.AgencyIdValidation(&agency, i, gtfs, &agencyRules)

		// Validate Agency Name
		validations.AgencyNameValidation(&agency, i, &agencyRules)

		// [CUSTOM VALIDATION] Check if agency_id matches agency_name
		validations.AgencyNameIdMatchValidation(&agency, i, &agencyRules)

		// Validate Agency URL
		validations.AgencyUrlValidation(&agency, i, &agencyRules)

		// Validate Agency Timezone
		validations.AgencyTimezoneValidation(&agency, i, &agencyRules)

		// Validate Agency Lang
		validations.AgencyLangValidation(&agency, i, &agencyRules)

		// Validate Agency Phone
		validations.AgencyPhoneValidation(&agency, i, &agencyRules)

		// Validate Agency Fare URL
		validations.AgencyFareUrlValidation(&agency, i, &agencyRules)

		// Validate Agency Email
		validations.AgencyEmailValidation(&agency, i, &agencyRules)
	}
}

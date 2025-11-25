package agency

import (
	"fmt"
	"main/lib"
	"main/types"
	validations "main/validations/agency/validations"
)

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Validations for agency.txt")

	// Get total count for progress tracking
	totalCount, err := gtfs.GetTableCount("agency")
	if err != nil {
		lib.AppLogger.Debug(fmt.Sprintf("Could not get table count for agency: %v", err))
		totalCount = 0
	}

	var processedCount int
	lastLoggedPercent := -1

	err = gtfs.IterateAgencies(func(i int, rawAgency types.AgencyRaw) error {
		processedCount++

		// Log progress every 10% or every 100 rows (whichever comes first)
		if totalCount > 0 {
			currentPercent := (processedCount * 100) / totalCount
			if currentPercent != lastLoggedPercent && (currentPercent%10 == 0 || processedCount%100 == 0) {
				lib.AppLogger.Debug(fmt.Sprintf("Validating agency.txt: %d/%d (%.1f%%)", processedCount, totalCount, float64(processedCount)*100.0/float64(totalCount)))
				lastLoggedPercent = currentPercent
			}
		} else if processedCount%100 == 0 {
			lib.AppLogger.Debug(fmt.Sprintf("Validating agency.txt: %d rows processed", processedCount))
		}
		// Parse Agency Validation
		agency := validations.ParseAgency(rawAgency, i)

		if agency == (types.Agency{}) {
			return nil
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

		return nil
	})

	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error iterating agencies: %v", err))
	} else {
		lib.AppLogger.Debug(fmt.Sprintf("Completed agency.txt validation: %d rows processed", processedCount))
	}
}

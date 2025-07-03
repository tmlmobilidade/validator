package fare_rules

import (
	"main/lib"
	"main/types"
	validations "main/validations/fare_rules/validations"
)

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running FareRules Validations...")
	
	for i, rawFareRule := range gtfs.FareRule {
		// Parse Fare Rule Validation
		fareRule := validations.ParseFareRule(rawFareRule, i)

		if fareRule == (types.FareRule{}) {
			continue
		}

		// validate contains_id
		validations.ContainsIdValidation(&fareRule, i, &gtfs, nil)

		// validate destination_id
		validations.DestinationIdValidation(&fareRule, i, &gtfs, nil)

		// validate origin_id
		validations.OriginIdValidation(&fareRule, i, &gtfs, nil)

		// validate fare_id
		validations.FareIdValidation(&fareRule, i, &gtfs)

		// validate route_id
		validations.RouteIdValidation(&fareRule, i, &gtfs)

	}
}

package fare_rules

import (
	"fmt"
	"main/lib"
	"main/types"
	validations "main/validations/fare_rules/validations"
)

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running FareRules Validations...")

	err := gtfs.IterateFareRules(func(i int, rawFareRule types.FareRuleRaw) error {
		// Parse Fare Rule Validation
		fareRule := validations.ParseFareRule(rawFareRule, i)

		if fareRule == (types.FareRule{}) {
			return nil
		}

		var fareRulesRules *types.FareRulesRules
		if rules != nil {
			fareRulesRules = &rules.FareRules
		}

		// validate contains_id
		validations.ContainsIdValidation(&fareRule, i, &gtfs, fareRulesRules)

		// validate destination_id
		validations.DestinationIdValidation(&fareRule, i, &gtfs, fareRulesRules)

		// validate origin_id
		validations.OriginIdValidation(&fareRule, i, &gtfs, fareRulesRules)

		// validate fare_id
		validations.FareIdValidation(&fareRule, i, &gtfs, fareRulesRules)

		// validate route_id
		validations.RouteIdValidation(&fareRule, i, &gtfs, fareRulesRules)

		return nil
	})

	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error iterating fare rules: %v", err))
	}
}

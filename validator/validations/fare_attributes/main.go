package fare_attributes

import (
	"fmt"
	"main/lib"
	"main/types"
	validations "main/validations/fare_attributes/validations"
)

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Fare Attributes Validations...")

	err := gtfs.IterateFareAttributes(func(i int, rawFareAttributes types.FareAttributeRaw) error {
		fareAttribute := ParseFareAttributes(rawFareAttributes, i)

		if fareAttribute == (types.FareAttribute{}) {
			return nil
		}

		var fareAttributesRules *types.FareAttributesRules
		if rules != nil {
			fareAttributesRules = &rules.FareAttributes
		}

		// Validate fare_id
		validations.FareIdValidation(&fareAttribute, i, &gtfs)

		// Validate price
		validations.PriceValidation(&fareAttribute, i)

		// Validate currency_type
		validations.CurrencyTypeValidation(&fareAttribute, i)

		// Validate payment_method
		validations.PaymentMethodValidation(&fareAttribute, i)

		// Validate transfers
		validations.TransfersValidation(&fareAttribute, i, &gtfs)

		// Validate agency_id
		validations.AgencyIdValidation(&fareAttribute, i, &gtfs, fareAttributesRules)

		// Validate transfer_duration
		validations.TransferDurationValidation(&fareAttribute, i, &gtfs, fareAttributesRules)

		return nil
	})

	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error iterating fare attributes: %v", err))
	}
}

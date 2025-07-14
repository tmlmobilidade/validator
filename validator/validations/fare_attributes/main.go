package fare_attributes

import (
	"main/lib"
	"main/types"
	validations "main/validations/fare_attributes/validations"
)

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Fare Attributes Validations...")

	for i, rawFareAttributes := range gtfs.FareAttribute {
		fareAttribute := ParseFareAttributes(rawFareAttributes, i)

		if fareAttribute == (types.FareAttribute{}) {
			continue
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
	}
}

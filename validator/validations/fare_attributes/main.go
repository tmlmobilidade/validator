package fare_attributes

import (
	"main/lib"
	"main/types"
	validations "main/validations/fare_attributes/validations"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Fare Attributes Validations...")

	for i, rawFareAttributes := range gtfs.Files["fare_attributes"] {
		fareAttribute := ParseFareAttributes(rawFareAttributes, i, &gtfs)

		if fareAttribute == (types.FareAttribute{}) {
			continue
		}

		// Validate fare_id
		validations.FareIdValidation(&fareAttribute, i, &gtfs)

		// Validate price
		validations.PriceValidation(&fareAttribute, i, &gtfs)
		
		
	}
}

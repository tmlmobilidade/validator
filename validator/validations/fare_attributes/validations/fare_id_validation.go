package fare_attributes

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [fare_attributes.txt]
  - Field: fare_id
  - Presence: Required
  - Type: Unique ID

# Description

Identifies a fare class.

[fare_attributes.txt]: https://gtfs.org/schedule/reference/#fare_attributestxt
*/
func FareIdValidation(fareAttribute *types.FareAttribute, row int, gtfs *types.Gtfs) {
	ctx := lib.NewValidationContext("fare_id", "fare_attributes.txt", "fare_id_validation", row, services.AppMessageService)

	if fareAttribute.FareId == nil {
		ctx.AddError(ctx.GetTranslatedMessage("fare_id_validation.required"))
		return
	}

	// Check if fare_id is Unique ID within fare_attributes table
	rows, err := gtfs.GetRowsById("fare_attributes", *fareAttribute.FareId)
	if err != nil {
		// Fallback to in-memory IdMap if database is not available
		if gtfs.IdMap != nil {
			if fareAttributeRows, exists := gtfs.IdMap["fare_attributes"]; exists {
				if indices, found := fareAttributeRows[*fareAttribute.FareId]; found {
					if len(indices) > 1 {
						ctx.AddError(ctx.GetTranslatedMessage("fare_id_validation.duplicate", *fareAttribute.FareId))
						return
					}
				}
			}
		}
		return
	}
	if len(rows) > 1 {
		ctx.AddError(ctx.GetTranslatedMessage("fare_id_validation.duplicate", *fareAttribute.FareId))
		return
	}
}

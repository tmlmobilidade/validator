package fare_attributes

import (
	"main/i18n"
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

	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "fare_id",
			FileName:     "fare_attributes.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "fare_id_validation",
		})
	}

	if fareAttribute.FareId == nil {
		addMessage(i18n.AppTranslator.Get("fare_id_validation.required"))
		return
	}

	if gtfs.IdMap["fare_rules"] != nil && len(gtfs.IdMap["fare_rules"][*fareAttribute.FareId]) > 1 {
		addMessage(i18n.AppTranslator.Get("fare_id_validation.duplicate", *fareAttribute.FareId))
		return
	}
}

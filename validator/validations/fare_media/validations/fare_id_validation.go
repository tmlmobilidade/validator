package fare_media

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [fare_media.txt]
  - Field: fare_id
  - Presence: Required
  - Type: unique ID referencing [fare_id]

# Description

Identifies a fare media.

[fare_media.txt]: https://gtfs.org/schedule/reference/#fare_mediatxt
*/

func FareIdValidation(fareMedia *types.FareMedia, row int, gtfs *types.Gtfs, rules *types.FareMediaRules) {

	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "fare_id",
			FileName:     "fare_media.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "fare_id_validation",
		})
	}

	// Validate presence
	if fareMedia.FareMediaId == "" {
		addMessage(i18n.AppTranslator.Get("fare_id_validation.required"))
		return
	}

	// Validate that fareMedia.FareMediaId exists in the fare_media.txt file
	if !lib.GtfsIdMapKeyExists(gtfs, "fare_media", *&fareMedia.FareMediaId) {
		addMessage(i18n.AppTranslator.Get("fare_id_validation.invalid"))
		return
	}
}

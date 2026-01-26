package fare_media

import (
	"main/i18n"
	"main/services"
	"main/types"
	"strconv"
)

/*
# Attributes

  - File: [fare_media.txt]
  - Field: fare_media_name
  - Presence: optional
  - Type: text

# Description

Name of the fare media.

For fare media which are transit cards (fare_media_type =2) or mobile apps (fare_media_type =4), the fare_media_name should be included and should match the rider-facing name used by the organizations delivering them.

[fare_media.txt]: https://gtfs.org/schedule/reference/#fare_mediatxt
*/

func FareNameValidation(fareMedia *types.FareMedia, row int, rules *types.FareMediaRules) {
	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "fare_name",
			FileName:     "fare_media.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_WARNING,
			ValidationID: "fare_name_validation",
		})
	}

	// Validate that fareMedia.FareMediaType is 2 (transit cards) or 4 (mobile apps)
	if fareMedia.FareMediaType != nil && (strconv.Itoa(*fareMedia.FareMediaType) == "2" || strconv.Itoa(*fareMedia.FareMediaType) == "4") {
		if fareMedia.FareMediaName == nil || *fareMedia.FareMediaName == "" {
			addMessage(i18n.AppTranslator.Get("fare_name_validation.warning"))
			return
		}
	}
}

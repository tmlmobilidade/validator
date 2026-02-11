package fare_media

import (
	"main/lib"
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

func FareMediaNameValidation(fareMedia *types.FareMedia, row int, gtfs *types.Gtfs, rules *types.FareMediaRules) {
	ctx := lib.NewValidationContext("fare_media_name", "fare_media.txt", "fare_media_name_validation", row, services.AppMessageService)
	if rules != nil && rules.FareMediaName.Severity != "" {
		ctx.WithSeverity(rules.FareMediaName.Severity)
	}

	// Validate that fareMedia.FareMediaType is 2 (transit cards) or 4 (mobile apps)
	if fareMedia.FareMediaType != nil && (strconv.Itoa(*fareMedia.FareMediaType) == "2" || strconv.Itoa(*fareMedia.FareMediaType) == "4") {
		if fareMedia.FareMediaName == nil || *fareMedia.FareMediaName == "" {
			ctx.AddWarning(ctx.GetTranslatedMessage("fare_media_name_validation.warning"))
			return
		}
		return
	}
}

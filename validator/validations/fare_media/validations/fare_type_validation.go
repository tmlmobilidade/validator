package fare_media

import (
	"main/i18n"
	"main/services"
	"main/types"
	"slices"
	"strconv"
)

/*
# Attributes

  - File: [fare_media.txt]
  - Field: fare_media_type
  - Presence: required
  - Type: enum

# Description

The type of fare media. Valid options are:

	0 - None. Used when there is no fare media involved in purchasing or validating a fare product, such as paying cash to a driver or conductor with no physical ticket provided.
	1 - Physical paper ticket that allows a passenger to take either a certain number of pre-purchased trips or unlimited trips within a fixed period of time.
	2 - Physical transit card that has stored tickets, passes or monetary value.
	3 - cEMV (contactless Europay, Mastercard and Visa) as an open-loop token container for account-based ticketing.
	4 - Mobile app that have stored virtual transit cards, tickets, passes, or monetary value.

[fare_media.txt]: https://gtfs.org/schedule/reference/#fare_mediatxt
*/

func FareTypeValidation(fareMedia *types.FareMedia, row int, rules *types.FareMediaRules) {
	addMessage := func(msg string) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "fare_type",
			FileName:     "fare_media.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "fare_type_validation",
		})
	}

	// Validate presence
	if fareMedia.FareMediaType == nil || strconv.Itoa(*fareMedia.FareMediaType) == "" {
		addMessage(i18n.AppTranslator.Get("fare_type_validation.required"))
		return
	}

	validTypeOptions := []string{"0", "1", "2", "3", "4"}
	// Validate that fareMedia.FareMediaType is in the valid options
	if !slices.Contains(validTypeOptions, strconv.Itoa(*fareMedia.FareMediaType)) {
		addMessage(i18n.AppTranslator.Get("fare_type_validation.invalid", fareMedia.FareMediaType))
		return
	}

	// Validate Rule Options
	if rules != nil && rules.FareType.Options != nil {
		if slices.Contains(*rules.FareType.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.FareType.Options, strconv.Itoa(*fareMedia.FareMediaType)) {
			addMessage(i18n.AppTranslator.Get("fare_type_validation.not_allowed", fareMedia.FareMediaType))
			return
		}
	}
}

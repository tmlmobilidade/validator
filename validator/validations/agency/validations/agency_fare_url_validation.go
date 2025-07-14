package agency

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
	"slices"
)

/*
# Attributes

  - File: [agency.txt]
  - Field: agency_phone
  - Presence: Optional
  - Type: URL

# Description

URL of a web page where a rider can purchase tickets or other fare instruments for that agency, or a web page containing information about that agency's fares.

[agency.txt]: https://gtfs.org/schedule/reference/#agencytxt
*/
func AgencyFareUrlValidation(agency *types.Agency, row int, rules *types.AgencyRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.AgencyFare.Severity != "" {
		s = rules.AgencyFare.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "agency_fare_url",
			FileName:     "agency.txt",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
			ValidationID: "agency_fare_url_validation",
		})
	}

	// Check if agency_fare_url is required
	if agency.AgencyFareUrl == nil && s != types.SEVERITY_IGNORE {
		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"agency_fare_url_validation.required",
				"agency_fare_url_validation.recommended",
			),
		)
		addMessage(message, s)
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("agency_fare_url_validation.forbidden"), s)
		return
	}

	// Check if agency_fare_url is valid
	if agency.AgencyFareUrl != nil && !lib.ValidateUrl(*agency.AgencyFareUrl) {
		addMessage(i18n.AppTranslator.Get("agency_fare_url_validation.invalid"), types.SEVERITY_ERROR)
		return
	}

	// Validate rules
	if rules != nil && rules.AgencyFare.Options != nil {
		if slices.Contains(*rules.AgencyFare.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.AgencyFare.Options, *agency.AgencyFareUrl) {
			addMessage(i18n.AppTranslator.Get("agency_fare_url_validation.not_allowed", *agency.AgencyFareUrl), s)
			return
		}
	}
}

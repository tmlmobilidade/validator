package agency

import (
	"fmt"
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
		message := lib.IfThenElse(s == types.SEVERITY_ERROR, "Agency fare URL is required", "Agency fare URL is recommended")
		addMessage(message, s)
	}

	// Check if agency_fare_url is valid
	if agency.AgencyFareUrl != nil {
		if urlErrors := lib.ValidateUrl(*agency.AgencyFareUrl); urlErrors != "" {
			addMessage(urlErrors, types.SEVERITY_ERROR)
		}
	}

	// Validate rules
	if rules != nil && rules.AgencyFare.Options != nil {
		if slices.Contains(*rules.AgencyFare.Options, types.ALL_OPTIONS) {
			return
		}

		if !slices.Contains(*rules.AgencyFare.Options, *agency.AgencyFareUrl) {
			addMessage(fmt.Sprintf("Agency fare URL is not allowed: %s", *agency.AgencyFareUrl), s)
			return
		}
	}
}

package trips

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [trips.txt]
  - Field: trip_short_name
  - Presence: Optional
  - Type: Text

# Description

Public facing text used to identify the trip to riders, for instance, to identify train numbers for commuter rail trips.
If riders do not commonly rely on trip names, `trip_short_name` should be empty.
A `trip_short_name` value, if provided, should uniquely identify a trip within a service day; it should not be used for destination names or limited/express designations.

[trips.txt]: https://gtfs.org/schedule/reference/#tripstxt
*/
func TripShortNameValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.TripShortName.Severity != "" {
		s = rules.TripShortName.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "trip_short_name",
			FileName:     "trips.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "trip_short_name_validation",
		})
	}

	// 1. Validate trip_short_name is required
	if trip.TripShortName == nil {
		if s == types.SEVERITY_IGNORE {
			return
		}

		message := i18n.AppTranslator.Get(
			lib.IfThenElse(s == types.SEVERITY_ERROR,
				"trip_short_name_validation.required",
				"trip_short_name_validation.recommended",
			),
		)
		addMessage(message, s)
		return
	}

	// 2. Validate trip_short_name is forbidden
	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("trip_short_name_validation.forbidden"), s)
		return
	}

}

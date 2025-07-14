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
  - Field: trip_headsign
  - Presence: Optional
  - Type: Text

# Description

Text that appears on signage identifying the trip's destination to riders.
This field is recommended for all services with headsign text displayed on the vehicle which may be used to distinguish amongst trips in a route.

If the headsign changes during a trip, values for `trip_headsign` may be overridden by defining values in `stop_times.stop_headsign` for specific `stop_time`s along the trip.

[trips.txt]: https://gtfs.org/schedule/reference/#tripstxt
*/
func TripHeadsignValidation(trip *types.Trip, row int, gtfs *types.Gtfs, rules *types.TripsRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.TripHeadsign.Severity != "" {
		s = rules.TripHeadsign.Severity
	}

	if s == types.SEVERITY_IGNORE {
		return
	}

	if trip.TripHeadsign != nil {
		return
	}

	message := types.Message{
		Field:        "trip_headsign",
		FileName:     "trips.txt",
		Message:      lib.IfThenElse(s == types.SEVERITY_ERROR, i18n.AppTranslator.Get("trip_headsign_validation.required"), i18n.AppTranslator.Get("trip_headsign_validation.recommended")),
		Rows:         []int{row},
		Severity:     s,
		ValidationID: "trip_headsign_validation",
	}

	services.AppMessageService.AddMessage(message)
}

package trips

import (
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
func TripShortNameValidation(severity *types.Severity, trip *types.Trip, row int, gtfs *types.Gtfs) {
	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}
	
	if s == types.SEVERITY_IGNORE {
		return;
	}

	if trip.TripShortName != nil {
		return;
	}

	message := types.Message{
		Field: "trip_short_name",
		FileName: "trips.txt",
		Message: lib.IfThenElse(s == types.SEVERITY_ERROR, "Trip short name is required", "Trip short name is recommended"),
		Rows: []int{row},
		Severity: s,
		ValidationID: "trip_short_name_validation",
	}

	services.AppMessageService.AddMessage(message)
}
package stop_times

import (
	"main/lib"
	"main/types"
	validations "main/validations/stop_times/validations"
)

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running StopTimes Validations...")

	for i, rawStopTimes := range gtfs.StopTime {
		stopTime := validations.ParseStopTimes(rawStopTimes, i)

		if stopTime == (types.StopTime{}) {
			continue
		}

		// Validate trip_id
		validations.TripIdValidation(&stopTime, i, &gtfs)
		
		// Validate arrival_time
		validations.ArrivalTimeValidation(nil, &stopTime, i, &gtfs)
		
		// Validate departure_time
		validations.DepartureTimeValidation(nil, &stopTime, i, &gtfs)
		
		// Validate stop_id
		validations.StopIdValidation(&stopTime, i, &gtfs)
		
		// Validate location_group_id
		validations.LocationGroupIdValidation(&stopTime, i, &gtfs)

		// Validate start_pickup_drop_off_window
		validations.StartPickupDropOffWindowValidation(nil, &stopTime, i, &gtfs)

		// Validate end_pickup_drop_off_window
		validations.EndPickupDropOffWindowValidation(nil, &stopTime, i, &gtfs)

		// Validate pickup_type
		validations.PickupTypeValidation(nil, &stopTime, i)

		// TODO: Validate location_id
		// validations.LocationIdValidation(&stopTime, i, &gtfs)

		// Validate stop_headsign
		validations.StopHeadsignValidation(nil, &stopTime, i)

		// Validate continuous_drop_off
		validations.ContinuousDropOffValidation(nil, &stopTime, i)

		// Validate continuous_pickup
		validations.ContinuousPickupValidation(nil, &stopTime, i)

		// Validate drop_off_type
		validations.DropOffTypeValidation(nil, &stopTime, i)

		// Validate shape_dist_traveled
		validations.ShapeDistTraveledValidation(nil, &stopTime, i, &gtfs)

		// Validate timepoint
		validations.TimepointValidation(nil, &stopTime, i)

		// Validate pickup_booking_rule_id
		validations.PickupBookingRuleIdValidation(nil, &stopTime, i, &gtfs)

		// Validate drop_off_booking_rule_id
		validations.DropOffBookingRuleIdValidation(nil, &stopTime, i, &gtfs)
	}
}

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
		validations.ArrivalTimeValidation(&stopTime, i, &gtfs, &rules.StopTimes)

		// Validate departure_time
		validations.DepartureTimeValidation(&stopTime, i, &gtfs, &rules.StopTimes)

		// Validate stop_id
		validations.StopIdValidation(&stopTime, i, &gtfs)

		// Validate location_group_id
		validations.LocationGroupIdValidation(&stopTime, i, &gtfs)

		// Validate start_pickup_drop_off_window
		validations.StartPickupDropOffWindowValidation(&stopTime, i, &rules.StopTimes)

		// Validate end_pickup_drop_off_window
		validations.EndPickupDropOffWindowValidation(&stopTime, i, &rules.StopTimes)

		// Validate pickup_type
		validations.PickupTypeValidation(&stopTime, i, &rules.StopTimes)

		// Validate stop_headsign
		validations.StopHeadsignValidation(&stopTime, i, &rules.StopTimes)

		// Validate continuous_drop_off
		validations.ContinuousDropOffValidation(&stopTime, i, &rules.StopTimes)

		// Validate continuous_pickup
		validations.ContinuousPickupValidation(&stopTime, i, &rules.StopTimes)

		// Validate drop_off_type
		validations.DropOffTypeValidation(&stopTime, i, &rules.StopTimes)

		// Validate shape_dist_traveled
		validations.ShapeDistTraveledValidation(&stopTime, i, &rules.StopTimes)

		// Validate timepoint
		validations.TimepointValidation(&stopTime, i, &rules.StopTimes)

		// Validate pickup_booking_rule_id
		validations.PickupBookingRuleIdValidation(&stopTime, i, &gtfs, &rules.StopTimes)

		// Validate drop_off_booking_rule_id
		validations.DropOffBookingRuleIdValidation(&stopTime, i, &gtfs, &rules.StopTimes)
	}
}

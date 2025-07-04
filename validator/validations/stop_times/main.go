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

		var stopTimesRules *types.StopTimesRules
		if rules != nil {
			stopTimesRules = &rules.StopTimes
		}

		// Validate trip_id
		validations.TripIdValidation(&stopTime, i, &gtfs)

		// Validate arrival_time
		validations.ArrivalTimeValidation(&stopTime, i, &gtfs, stopTimesRules)

		// Validate departure_time
		validations.DepartureTimeValidation(&stopTime, i, &gtfs, stopTimesRules)

		// Validate stop_id
		validations.StopIdValidation(&stopTime, i, &gtfs)

		// Validate location_group_id
		validations.LocationGroupIdValidation(&stopTime, i, &gtfs)

		// Validate start_pickup_drop_off_window
		validations.StartPickupDropOffWindowValidation(&stopTime, i, stopTimesRules)

		// Validate end_pickup_drop_off_window
		validations.EndPickupDropOffWindowValidation(&stopTime, i, stopTimesRules)

		// Validate pickup_type
		validations.PickupTypeValidation(&stopTime, i, stopTimesRules)

		// Validate stop_headsign
		validations.StopHeadsignValidation(&stopTime, i, stopTimesRules)

		// Validate continuous_drop_off
		validations.ContinuousDropOffValidation(&stopTime, i, stopTimesRules)

		// Validate continuous_pickup
		validations.ContinuousPickupValidation(&stopTime, i, stopTimesRules)

		// Validate drop_off_type
		validations.DropOffTypeValidation(&stopTime, i, stopTimesRules)

		// Validate shape_dist_traveled
		validations.ShapeDistTraveledValidation(&stopTime, i, stopTimesRules)

		// Validate timepoint
		validations.TimepointValidation(&stopTime, i, stopTimesRules)

		// Validate pickup_booking_rule_id
		validations.PickupBookingRuleIdValidation(&stopTime, i, &gtfs, stopTimesRules)

		// Validate drop_off_booking_rule_id
		validations.DropOffBookingRuleIdValidation(&stopTime, i, &gtfs, stopTimesRules)
	}
}

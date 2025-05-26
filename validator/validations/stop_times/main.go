package stop_times

import (
	"main/lib"
	"main/types"
	validations "main/validations/stop_times/validations"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running StopTimes Validations...")

	for i, rawStopTimes := range gtfs.Files["stop_times"] {
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

		// TODO: Validate location_id
		// validations.LocationIdValidation(&stopTime, i, &gtfs)

		// Validate stop_headsign
		validations.StopHeadsignValidation(nil, &stopTime, i)
	}
}

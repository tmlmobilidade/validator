package stops

import (
	"main/lib"
	"main/types"
	validations "main/validations/stops/validations"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Validations for stops.txt")

	for row, rawStop := range gtfs.Files["stops"] {
		stop := validations.ParseStop(rawStop, row)

		if stop == (types.Stop{}) {
			continue
		}
		
		// Validate stop_id
		validations.StopIdValidation(&stop, row, &gtfs)

		// Validate stop_code
		validations.StopCodeValidation(nil, &stop, row, &gtfs)
		
		// Validate stop_name
		validations.StopNameValidation(nil, &stop, row)
		
		// Validate tts_stop_name
		validations.TtsStopNameValidation(nil, &stop, row)

		// Validate stop_desc
		validations.StopDescValidation(nil, &stop, row)
		
		// Validate stop_lat
		validations.StopLatValidation(nil, &stop, row)

		// Validate stop_lon
		validations.StopLonValidation(nil, &stop, row)
		
		// Validate zone_id
		validations.ZoneIdValidation(nil, &stop, row)

		// Validate location_type
		validations.LocationTypeValidation(nil, &stop, row)

		// Validate parent_station
		validations.ParentStationValidation(nil, &stop, row, gtfs)

		// Validate stop_timezone
		validations.StopTimezoneValidation(nil, &stop, row)

		// Validate wheelchair_boarding
		validations.WheelchairBoardingValidation(nil, &stop, row)

		// Validate level_id
		validations.LevelIdValidation(nil, &stop, row, gtfs)

		// Validate platform_code
		validations.PlatformCodeValidation(nil, &stop, row)
	}
}

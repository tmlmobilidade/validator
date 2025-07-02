package stops

import (
	"main/lib"
	"main/types"
	validations "main/validations/stops/validations"
)

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Validations for stops.txt")

	for row, rawStop := range gtfs.Stop {
		stop := validations.ParseStop(rawStop, row)

		if stop == (types.Stop{}) {
			continue
		}

		// Validate stop_id
		validations.StopIdValidation(&stop, row, &gtfs, &rules.Stops)

		// Validate stop_code
		validations.StopCodeValidation(&stop, row, &gtfs, &rules.Stops)

		// Validate stop_name
		validations.StopNameValidation(&stop, row, &rules.Stops)

		// Validate tts_stop_name
		validations.TtsStopNameValidation(&stop, row, &rules.Stops)

		// Validate stop_desc
		validations.StopDescValidation(&stop, row, &rules.Stops)

		// Validate stop_lat
		validations.StopLatValidation(&stop, row, &rules.Stops)

		// Validate stop_lon
		validations.StopLonValidation(&stop, row, &rules.Stops)

		// Validate zone_id
		validations.ZoneIdValidation(&stop, row, &rules.Stops)

		// Validate location_type
		validations.LocationTypeValidation(&stop, row, &rules.Stops)

		// Validate parent_station
		validations.ParentStationValidation(&stop, row, gtfs, &rules.Stops)

		// Validate stop_timezone
		validations.StopTimezoneValidation(&stop, row, &rules.Stops)

		// Validate wheelchair_boarding
		validations.WheelchairBoardingValidation(&stop, row, &rules.Stops)

		// Validate level_id
		validations.LevelIdValidation(&stop, row, gtfs, &rules.Stops)

		// Validate platform_code
		validations.PlatformCodeValidation(&stop, row, &rules.Stops)

		// Validate region_id
		validations.RegionIdValidation(&stop, row, &rules.Stops)

		// Validate public_visible
		validations.PublicVisibleValidation(&stop, row, &rules.Stops)

		// Validate shelter_code
		validations.ShelterCodeValidation(&stop, row, &rules.Stops)

		// Validate shelter_maintainer
		validations.ShelterMaintainerValidation(&stop, row, &rules.Stops)

		// Validate stop_short_name
		validations.StopShortNameValidation(&stop, row, &rules.Stops)

		// Validate stop_url
		validations.StopUrlValidation(&stop, row, &rules.Stops)

		// Validate municipality_id
		validations.MunicipalityIdValidation(&stop, row, &rules.Stops)

		// Validate parish_id
		validations.ParishIdValidation(&stop, row, &rules.Stops)

		// Validate has_bench
		validations.HasBenchValidation(&stop, row, &rules.Stops)

		// Validate has_network_map
		validations.HasNetworkMapValidation(&stop, row, &rules.Stops)

		// Validate has_pip_real_time
		validations.HasPipRealTimeValidation(&stop, row, &rules.Stops)

		// Validate has_schedules
		validations.HasSchedulesValidation(&stop, row, &rules.Stops)

		// Validate has_shelter
		validations.HasShelterValidation(&stop, row, &rules.Stops)

		// Validate has_stop_sign
		validations.HasStopSignValidation(&stop, row, &rules.Stops)

		// Validate has_tariffs_information
		validations.HasTariffsInformationValidation(&stop, row, &rules.Stops)
	}
}

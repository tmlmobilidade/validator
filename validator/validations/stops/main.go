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
	}
}

package fare_media

import (
	"main/lib"
	"main/types"
	registry "main/validations"
)

func init() {
	registry.Register("fare_media", RunValidations)
}

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running FareMedia Validations...")
}

package levels

import (
	"main/lib"
	"main/types"
	registry "main/validations"
)

func init() {
	registry.Register("levels", RunValidations)
}

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Levels Validations...")
}

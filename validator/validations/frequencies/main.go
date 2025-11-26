package frequencies

import (
	"main/lib"
	"main/types"
	registry "main/validations"
)

func init() {
	registry.Register("frequencies", RunValidations)
}

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Frequencies Validations...")
}

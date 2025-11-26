package transfers

import (
	"main/lib"
	"main/types"
	registry "main/validations"
)

func init() {
	registry.Register("transfers", RunValidations)
}

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Transfers Validations...")
}

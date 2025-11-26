package pathways

import (
	"main/lib"
	"main/types"
	registry "main/validations"
)

func init() {
	registry.Register("pathways", RunValidations)
}

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Pathways Validations...")
}

package rider_categories

import (
	"main/lib"
	"main/types"
	registry "main/validations"
)

func init() {
	registry.Register("rider_categories", RunValidations)
}

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running RiderCategories Validations...")
}

package shapes

import (
	"main/lib"
	"main/types"
	validations "main/validations/shapes/validations"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Shapes Validations...")

	for row, rawShape := range gtfs.Files["shapes"] {
		shape := validations.ParseShape(rawShape, row)

		if shape == (types.Shape{}) {
			continue
		}
	}
}

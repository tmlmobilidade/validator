package shapes

import (
	"main/lib"
	"main/types"
	validations "main/validations/shapes/validations"
)

func RunValidations(gtfs types.Gtfs) {
	lib.AppLogger.Debug("Running Shapes Validations...")

	var allShapes []types.Shape

	for row, rawShape := range gtfs.Shape {
		shape := validations.ParseShape(rawShape, row)

		if shape == (types.Shape{}) {
			continue
		}

		// Validate shape_id
		validations.ShapeIdValidation(&shape, row)

		// Validate shape_pt_lat
		validations.ShapePtLatValidation(&shape, row)

		// Validate shape_pt_lon
		validations.ShapePtLonValidation(&shape, row)

		// Validate shape_pt_sequence
		validations.ShapePtSequenceValidation(&shape, row)

		// Validate shape_dist_traveled
		validations.ShapeDistTraveledValidation(nil, &shape, row)

		allShapes = append(allShapes, shape)
	}

	// Group-level validation: shape_pt_sequence must increase for each shape_id
	validations.ShapeSequenceValidation(allShapes)
}

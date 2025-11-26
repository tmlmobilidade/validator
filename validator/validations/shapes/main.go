package shapes

import (
	"fmt"
	"main/config"
	"main/lib"
	"main/types"
	validations "main/validations/shapes/validations"
	registry "main/validations"
)

func init() {
	registry.Register("shapes", RunValidations)
}

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Shapes Validations...")

	// Create progress tracker
	tracker := lib.CreateProgressTracker(gtfs, "shapes.txt", config.ProgressThresholdLarge)
	var allShapes []types.Shape

	err := gtfs.IterateShapes(func(row int, rawShape types.ShapeRaw) error {
		tracker.Track()
		shape := validations.ParseShape(rawShape, row)

		if shape == (types.Shape{}) {
			return nil
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
		return nil
	})

	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error iterating shapes: %v", err))
	} else {
		lib.AppLogger.Debug(fmt.Sprintf("Completed shapes.txt validation: %d rows processed", tracker.GetProcessedCount()))
	}

	// Group-level validation: shape_pt_sequence must increase for each shape_id
	validations.ShapeSequenceValidation(allShapes)
}

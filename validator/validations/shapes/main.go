package shapes

import (
	"fmt"
	"main/lib"
	"main/types"
	validations "main/validations/shapes/validations"
)

func RunValidations(gtfs types.Gtfs, rules *types.GtfsRules) {
	lib.AppLogger.Debug("Running Shapes Validations...")

	// Get total count for progress tracking
	totalCount, err := gtfs.GetTableCount("shapes")
	if err != nil {
		lib.AppLogger.Debug(fmt.Sprintf("Could not get table count for shapes: %v", err))
		totalCount = 0
	}

	var processedCount int
	lastLoggedPercent := -1
	var allShapes []types.Shape

	err = gtfs.IterateShapes(func(row int, rawShape types.ShapeRaw) error {
		processedCount++

		// Log progress every 10% or every 1000 rows (whichever comes first)
		if totalCount > 0 {
			currentPercent := (processedCount * 100) / totalCount
			if currentPercent != lastLoggedPercent && (currentPercent%10 == 0 || processedCount%1000 == 0) {
				lib.AppLogger.Debug(fmt.Sprintf("Validating shapes.txt: %d/%d (%.1f%%)", processedCount, totalCount, float64(processedCount)*100.0/float64(totalCount)))
				lastLoggedPercent = currentPercent
			}
		} else if processedCount%1000 == 0 {
			lib.AppLogger.Debug(fmt.Sprintf("Validating shapes.txt: %d rows processed", processedCount))
		}
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
		lib.AppLogger.Debug(fmt.Sprintf("Completed shapes.txt validation: %d rows processed", processedCount))
	}

	// Group-level validation: shape_pt_sequence must increase for each shape_id
	validations.ShapeSequenceValidation(allShapes)
}

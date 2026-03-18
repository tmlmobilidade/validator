package shapes

import (
	"sort"

	"main/lib"
	"main/services"
	shapes_coordinates "main/services/geo/shapes"
	"main/types"
)

type shapeConsistentPoint struct {
	id       string
	row      int
	sequence int
	lat      float64
	lon      float64
}

func buildShapeFromConsistentPoint(point shapeConsistentPoint) *types.Shape {
	return &types.Shape{
		ShapePtLat: lib.Ptr(float32(point.lat)),
		ShapePtLon: lib.Ptr(float32(point.lon)),
	}
}

type consistentViolation struct {
	id          string
	row         int
	currentLat  float64
	currentLon  float64
	currentSeq  int
	previousLat float64
	previousLon float64
	previousSeq int
}

func uniqueConsistentRows(rows []int) []int {
	seen := make(map[int]struct{}, len(rows))
	unique := make([]int, 0, len(rows))
	for _, row := range rows {
		if _, ok := seen[row]; ok {
			continue
		}
		seen[row] = struct{}{}
		unique = append(unique, row)
	}
	return unique
}

// ShapeCoordinatesDistanceValidation validates if consecutive points from the same shape are not too far apart.
func ShapeCoordinatesConsistentValidation(shapes []types.Shape) {
	shapeGroups := map[string][]shapeConsistentPoint{}
	violations := []consistentViolation{}

	for i, shape := range shapes {
		if shape.ShapeId == nil || *shape.ShapeId == "" {
			continue
		}
		if shape.ShapePtSequence == nil || shape.ShapePtLat == nil || shape.ShapePtLon == nil {
			continue
		}

		shapeGroups[*shape.ShapeId] = append(shapeGroups[*shape.ShapeId], shapeConsistentPoint{
			row:      i,
			sequence: *shape.ShapePtSequence,
			lat:      float64(*shape.ShapePtLat),
			lon:      float64(*shape.ShapePtLon),
		})
	}

	for _, shapeGroup := range shapeGroups {
		sort.Slice(shapeGroup, func(i, j int) bool {
			return shapeGroup[i].sequence < shapeGroup[j].sequence
		})

		for i := 1; i < len(shapeGroup); i++ {
			if shapeGroup[i].sequence == 0 {
				i++
			}
			prev := shapeGroup[i-1]
			current := shapeGroup[i]
			prevShapePoint := buildShapeFromConsistentPoint(prev)
			currentShapePoint := buildShapeFromConsistentPoint(current)

			closeEnough, _ := shapes_coordinates.ShapePointIsCloseToBeforeShapePoint(prevShapePoint, currentShapePoint)
			if closeEnough {
				continue
			}

			violations = append(violations, consistentViolation{
				row:         current.row,
				currentLat:  current.lat,
				currentLon:  current.lon,
				currentSeq:  current.sequence,
				previousLat: prev.lat,
				previousLon: prev.lon,
				previousSeq: prev.sequence,
			})
		}
	}

	if len(violations) > 100 {
		rows := make([]int, 0, len(violations))
		for _, violation := range violations {
			rows = append(rows, violation.row)
		}

		for _, row := range uniqueConsistentRows(rows) {
			ctx := lib.NewValidationContext("coordinates", "shapes.txt", "coordinates_consistent_validation", row, services.AppMessageService)
			ctx.AddError(ctx.GetTranslatedMessage("coordinates_consistent_validation.ManyErrors"))
		}
		return
	}

	for _, violation := range violations {
		ctx := lib.NewValidationContext("coordinates", "shapes.txt", "coordinates_consistent_validation", violation.row, services.AppMessageService)
		ctx.AddError(ctx.GetTranslatedMessage(
			"coordinates_consistent_validation.invalid_consistent_distance",
			violation.id,
			violation.currentLat,
			violation.currentLon,
			violation.currentSeq,
			violation.previousLat,
			violation.previousLon,
			violation.previousSeq,
		))
	}
}

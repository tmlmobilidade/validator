package shapes

import (
	"sort"

	"main/lib"
	"main/services"
	shapes_coordinates "main/services/geo/shapes"
	"main/types"
)

type shapeCoordinatePoint struct {
	row      int
	sequence int
	lat      float64
	lon      float64
}

func buildShapeFromPoint(point shapeCoordinatePoint) *types.Shape {
	return &types.Shape{
		ShapePtLat: lib.Ptr(float32(point.lat)),
		ShapePtLon: lib.Ptr(float32(point.lon)),
	}
}

// ShapeCoordinatesDistanceValidation validates if consecutive points from the same shape are not too far apart.
func ShapeCoordinatesDistanceValidation(shapes []types.Shape) {
	shapeGroups := map[string][]shapeCoordinatePoint{}

	for i, shape := range shapes {
		if shape.ShapeId == nil || *shape.ShapeId == "" {
			continue
		}
		if shape.ShapePtSequence == nil || shape.ShapePtLat == nil || shape.ShapePtLon == nil {
			continue
		}

		shapeGroups[*shape.ShapeId] = append(shapeGroups[*shape.ShapeId], shapeCoordinatePoint{
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
			prev := shapeGroup[i-1]
			current := shapeGroup[i]
			prevShape := buildShapeFromPoint(prev)
			currentShape := buildShapeFromPoint(current)
			closeEnough, _ := shapes_coordinates.ShapeIsCloseToOtherShape(prevShape, currentShape)
			if closeEnough {
				continue
			}

			// If the next point is close to the previous one, the current point is likely an isolated outlier.
			// Mark only this point and skip the immediate next check to avoid cascaded false positives.
			if i+1 < len(shapeGroup) {
				next := shapeGroup[i+1]
				nextShape := buildShapeFromPoint(next)
				prevAndNextAreClose, _ := shapes_coordinates.ShapeIsCloseToOtherShape(prevShape, nextShape)
				if prevAndNextAreClose {
					ctx := lib.NewValidationContext("coordenates", "shapes.txt", "coordenates_distance_validation", current.row, services.AppMessageService)
					ctx.AddError(ctx.GetTranslatedMessage("coordenates_distance_validation.invalid_distance", current.lat, current.lon))
					i++
					continue
				}
			}

			ctx := lib.NewValidationContext("coordenates", "shapes.txt", "coordenates_distance_validation", current.row, services.AppMessageService)
			ctx.AddError(ctx.GetTranslatedMessage("coordenates_distance_validation.invalid_distance", current.lat, current.lon))
		}
	}
}

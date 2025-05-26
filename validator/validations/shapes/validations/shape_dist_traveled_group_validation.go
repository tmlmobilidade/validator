package shapes

import (
	"fmt"
	"main/services"
	"main/types"
)

// ShapeDistTraveledGroupValidation checks that shape_dist_traveled increases for each shape_id
func ShapeDistTraveledGroupValidation(shapes []types.Shape) {
	shapeGroups := make(map[string][]struct {
		sequence int
		dist     *float64
		row      int
	})

	// Group shapes by shape_id, collecting sequence, distance and row information
	for i, shape := range shapes {
		// Skip shapes with missing required fields
		if shape.ShapeId == nil || shape.ShapePtSequence == nil {
			continue
		}

		// Add shape point to its group
		shapeId := *shape.ShapeId
		shapeGroups[shapeId] = append(shapeGroups[shapeId], struct {
			sequence int
			dist     *float64
			row      int
		}{
			sequence: *shape.ShapePtSequence,
			dist:     shape.ShapeDistTraveled,
			row:      i,
		})
	}

	for shapeId, points := range shapeGroups {
		prevDist := -1.0
		for i, pt := range points {
			if pt.dist == nil {
				continue
			}
			if i == 0 {
				prevDist = *pt.dist
				continue
			}
			if *pt.dist < prevDist {
				msg := types.Message{
					Field:        "shape_dist_traveled",
					FileName:     "shapes.txt",
					Rows:         []int{pt.row},
					Message:      fmt.Sprintf("shape_dist_traveled for shape_id '%s' must increase along the trip (found %f after %f)", shapeId, *pt.dist, prevDist),
					Severity:     types.SEVERITY_ERROR,
					ValidationID: "shape_dist_traveled_group_validation",
				}
				services.AppMessageService.AddMessage(msg)
			}
			prevDist = *pt.dist
		}
	}
} 
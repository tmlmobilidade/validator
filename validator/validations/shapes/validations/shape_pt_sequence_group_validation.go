package shapes

import (
	"fmt"
	"main/services"
	"main/types"
)

// ShapePtSequenceGroupValidation checks that shape_pt_sequence values increase for each shape_id
func ShapePtSequenceGroupValidation(shapes []types.Shape) {
	// Group shapes by shape_id
	shapeGroups := make(map[string][]struct {
		sequence int
		row      int
	})

	for i, shape := range shapes {
		if shape.ShapeId == nil || shape.ShapePtSequence == nil {
			continue // skip invalid rows
		}
		shapeId := *shape.ShapeId
		shapeGroups[shapeId] = append(shapeGroups[shapeId], struct {
			sequence int
			row      int
		}{sequence: *shape.ShapePtSequence, row: i})
	}

	for shapeId, points := range shapeGroups {
		// Sort by row order (as in file)
		prev := -1
		for _, pt := range points {
			if prev != -1 && pt.sequence <= prev {
				msg := types.Message{
					Field:        "shape_pt_sequence",
					FileName:     "shapes.txt",
					Rows:         []int{pt.row},
					Message:      fmt.Sprintf("shape_pt_sequence for shape_id '%s' must increase along the trip (found %d after %d)", shapeId, pt.sequence, prev),
					Severity:     types.SEVERITY_ERROR,
					ValidationID: "shape_pt_sequence_group_validation",
				}
				services.AppMessageService.AddMessage(msg)
			}
			prev = pt.sequence
		}
	}
} 
package shapes

import (
	"main/services"
	"main/types"
	"sort"
)

type ShapePtSequenceGroup struct {
	shapeId  string
	sequence int
	dist     float64
	row      int
}

/*
Validate the shape sequence, based on shape_pt_sequence and shape_dist_traveled.

https://gtfs.org/schedule/reference/#shapestxt
*/
func ShapeSequenceValidation(shapes []types.Shape) {

	addMessage := func(msg string, row int) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "shape_pt_sequence",
			FileName:     "shapes.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "shape_pt_sequence_validation",
		})
	}
	

	// Group shapes by shape_id
	shapeGroups := make(map[string][]ShapePtSequenceGroup)

	for i, shape := range shapes {
		if shape.ShapeId == nil || shape.ShapePtSequence == nil {
			addMessage("shape_id and shape_pt_sequence are required and must not be empty.", i)
			return
		}

		// Only add to group if shape_id and shape_pt_sequence are present
		group := ShapePtSequenceGroup{
			shapeId:  *shape.ShapeId,
			sequence: *shape.ShapePtSequence,
			row:      i,
		}
		if shape.ShapeDistTraveled != nil {
			group.dist = *shape.ShapeDistTraveled
		} else {
			group.dist = -1 // Use -1 to indicate missing distance
		}
		shapeGroups[*shape.ShapeId] = append(shapeGroups[*shape.ShapeId], group)
	}

	// Sort shapeGroups by sequence
	for _, shapeGroup := range shapeGroups {
		sort.Slice(shapeGroup, func(i, j int) bool {
			return shapeGroup[i].sequence < shapeGroup[j].sequence
		})

		// Check if the shape_pt_sequence values are increasing
		for i, shape := range shapeGroup {
			if i > 0 {
				if shape.sequence <= shapeGroup[i-1].sequence {
					addMessage("shape_pt_sequence for shape_id '"+shape.shapeId+"' must increase along the trip", shape.row)
				}
				// Only check dist if both current and previous are present
				if shape.dist >= 0 && shapeGroup[i-1].dist >= 0 {
					if shape.dist < shapeGroup[i-1].dist {
						addMessage("shape_dist_traveled for shape_id '"+shape.shapeId+"' must increase along the trip", shape.row)
					}
				}
			}
		}
	}
} 
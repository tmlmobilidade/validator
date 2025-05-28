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

// ShapePtSequenceGroupValidation checks that shape_pt_sequence values increase alongside shape_dist_traveled for each shape_id
func ShapePtSequenceGroupValidation(shapes []types.Shape) {

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
		if shape.ShapeId == nil || shape.ShapePtSequence == nil || shape.ShapeDistTraveled == nil {
			addMessage("shape_id, shape_pt_sequence, and shape_dist_traveled are required and must not be empty.", i)
			return
		}

		shapeGroups[*shape.ShapeId] = append(shapeGroups[*shape.ShapeId], ShapePtSequenceGroup{
			shapeId:  *shape.ShapeId,
			sequence: *shape.ShapePtSequence,
			dist:     *shape.ShapeDistTraveled,
			row:      i,
		})
	}

	// Sort shapeGroups by sequence
	for _, shapeGroup := range shapeGroups {
		sort.Slice(shapeGroup, func(i, j int) bool {
			return shapeGroup[i].sequence < shapeGroup[j].sequence
		})

		// Check if the shape_pt_sequence values are increasing
		for i, shape := range shapeGroup {
			if i > 0 && (shape.sequence < shapeGroup[i-1].sequence || shape.dist < shapeGroup[i-1].dist) {
				addMessage("shape_pt_sequence for shape_id '" + shape.shapeId + "' must increase along the trip", shape.row)
			}
		}
	}
} 
package shapes

import (
	"main/services"
	"main/types"
)

/*
# Attributes

 - File: [shapes.txt]
 - Field: shape_id
 - Presence: Required
 - Type: ID

# Description

Identifies a shape.

[shapes.txt]: https://gtfs.org/schedule/reference/#shapestxt
*/
func ShapeIdValidation(shape *types.Shape, row int) {
	if shape.ShapeId == nil || *shape.ShapeId == "" {
		message := types.Message{
			Field:        "shape_id",
			FileName:     "shapes.txt",
			Rows:         []int{row},
			Message:      "shape_id is required and must not be empty.",
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "shape_id_validation",
		}
		services.AppMessageService.AddMessage(message)
	}
} 
package shapes

import (
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

 - File: [shapes.txt]
 - Field: shape_pt_lat
 - Presence: Required
 - Type: Latitude

# Description

Latitude of a shape point. Each record in shapes.txt represents a shape point used to define the shape.

[shapes.txt]: https://gtfs.org/schedule/reference/#shapestxt
*/
func ShapePtLatValidation(shape *types.Shape, row int) {

	addMessage := func(msg string) {
		message := types.Message{
			Field:        "shape_pt_lat",
			FileName:     "shapes.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "shape_pt_lat_validation",
		}
		services.AppMessageService.AddMessage(message)
	}
	
	if shape.ShapePtLat == nil {
		addMessage("shape_pt_lat is required and must not be empty.")
		return
	}
	
	if errMsg := lib.ValidateLatitude(*shape.ShapePtLat); errMsg != "" {
		addMessage(errMsg)
	}
} 
package shapes

import (
	"main/i18n"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [shapes.txt]
  - Field: shape_pt_sequence
  - Presence: Required
  - Type: Non-negative integer

# Description

Sequence in which the shape points connect to form the shape.

Values must increase along the trip but do not need to be consecutive.

# Example

If the shape "A_shp" has three points in its definition, the [shapes.txt] file might contain these records to define the shape:

	shape_id,shape_pt_lat,shape_pt_lon,shape_pt_sequence
	A_shp,37.61956,-122.48161,0
	A_shp,37.64430,-122.41070,6
	A_shp,37.65863,-122.30839,11

[shapes.txt]: https://gtfs.org/schedule/reference/#shapestxt
*/
func ShapePtSequenceValidation(shape *types.Shape, row int) {

	addMessage := func(msg string) {
		message := types.Message{
			Field:        "shape_pt_sequence",
			FileName:     "shapes.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "shape_pt_sequence_validation",
		}
		services.AppMessageService.AddMessage(message)
	}

	if shape.ShapePtSequence == nil {
		addMessage(i18n.AppTranslator.Get("shape_pt_sequence_validation.required"))
		return
	}

	if *shape.ShapePtSequence < 0 {
		addMessage(i18n.AppTranslator.Get("shape_pt_sequence_validation.invalid"))
	}
}

package shapes

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [shapes.txt]
  - Field: shape_pt_lon
  - Presence: Required
  - Type: Longitude

# Description

Longitude of a shape point.

[shapes.txt]: https://gtfs.org/schedule/reference/#shapestxt
*/
func ShapePtLonValidation(shape *types.Shape, row int) {

	addMessage := func(msg string) {
		message := types.Message{
			Field:        "shape_pt_lon",
			FileName:     "shapes.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     types.SEVERITY_ERROR,
			ValidationID: "shape_pt_lon_validation",
		}
		services.AppMessageService.AddMessage(message)
	}

	if shape.ShapePtLon == nil {
		addMessage(i18n.AppTranslator.Get("shape_pt_lon_validation.required"))
		return
	}

	if !lib.ValidateLongitude(*shape.ShapePtLon) {
		addMessage(i18n.AppTranslator.Get("shape_pt_lon_validation.invalid"))
	}
}

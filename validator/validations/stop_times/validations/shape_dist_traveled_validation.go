package stop_times

import (
	"main/i18n"
	"main/lib"
	"main/services"
	"main/types"
)

/*
# Attributes

  - File: [stop_times.txt]
  - Field: shape_dist_traveled
  - Presence: Optional
  - Type: Non-negative Float

# Description

Actual distance traveled along the associated shape, from the first stop to the stop specified in this record.

This field specifies how much of the shape to draw between any two stops during a trip.

Must be in the same units used in [shapes.txt].

Values used for shape_dist_traveled must increase along with stop_sequence; they must not be used to show reverse travel along a route.

Recommended for routes that have looping or inlining (the vehicle crosses or travels over the same portion of alignment in one trip).

See shapes.shape_dist_traveled.

# Example

If a bus travels a distance of 5.25 kilometers from the start of the shape to the stop,shape_dist_traveled=5.25.

[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
[shapes.txt]: https://gtfs.org/schedule/reference/#shapestxt
*/
func ShapeDistTraveledValidation(stopTime *types.StopTime, row int, rules *types.StopTimesRules) {
	s := types.SEVERITY_IGNORE
	if rules != nil && rules.ShapeDistTraveled.Severity != "" {
		s = rules.ShapeDistTraveled.Severity
	}

	addMessage := func(msg string, severity types.Severity) {
		services.AppMessageService.AddMessage(types.Message{
			Field:        "shape_dist_traveled",
			FileName:     "stop_times.txt",
			ValidationID: "shape_dist_traveled_validation",
			Message:      msg,
			Rows:         []int{row},
			Severity:     severity,
		})
	}

	if stopTime.ShapeDistTraveled == nil {
		if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
			return
		}
		warn := lib.IfThenElse(s == types.SEVERITY_WARNING, i18n.AppTranslator.Get("shape_dist_traveled_validation.recommended"), i18n.AppTranslator.Get("shape_dist_traveled_validation.required"))
		addMessage(warn, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("shape_dist_traveled_validation.forbidden"), s)
		return
	}

	if *stopTime.ShapeDistTraveled < 0 {
		addMessage(i18n.AppTranslator.Get("shape_dist_traveled_validation.negative"), types.SEVERITY_ERROR)
		return
	}
}

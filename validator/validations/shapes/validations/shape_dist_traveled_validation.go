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
  - Field: shape_dist_traveled
  - Presence: Optional
  - Type: Non-negative float

# Description

Actual distance traveled along the shape from the first shape point to the point specified in this record.

Used by trip planners to show the correct portion of the shape on a map.

Values must increase along with `shape_pt_sequence`; they must not be used to show reverse travel along a route.

Distance units must be consistent with those used in [stop_times.txt].

Recommended for routes that have looping or inlining (the vehicle crosses or travels over the same portion of alignment in one trip).

If a vehicle retraces or crosses the route alignment at points in the course of a trip, `shape_dist_traveled` is important to clarify how portions of the points in [shapes.txt] line up correspond with records in [stop_times.txt].

# Example

If a bus travels along the three points defined above for A_shp, the additional `shape_dist_traveled` values (shown here in kilometers) would look like this:

	shape_id,shape_pt_lat,shape_pt_lon,shape_pt_sequence,shape_dist_traveled
	A_shp,37.61956,-122.48161,0,0
	A_shp,37.64430,-122.41070,6,6.8310
	A_shp,37.65863,-122.30839,11,15.8765

[shapes.txt]: https://gtfs.org/schedule/reference/#shapestxt
[stop_times.txt]: https://gtfs.org/schedule/reference/#stoptimetxt
*/
func ShapeDistTraveledValidation(severity *types.Severity, shape *types.Shape, row int) {

	s := types.SEVERITY_IGNORE
	if severity != nil {
		s = *severity
	}

	// Add message to the message service
	addMessage := func(msg string, severity types.Severity) {
		message := types.Message{
			Field:        "shape_dist_traveled",
			FileName:     "shapes.txt",
			Rows:         []int{row},
			Message:      msg,
			Severity:     severity,
			ValidationID: "shape_dist_traveled_validation",
		}
		services.AppMessageService.AddMessage(message)
	}

	if shape.ShapeDistTraveled == nil {
		if s == types.SEVERITY_IGNORE || s == types.SEVERITY_FORBIDDEN {
			return
		}

		warn := lib.IfThenElse(s == types.SEVERITY_ERROR, i18n.AppTranslator.Get("shape_dist_traveled_validation.required"), i18n.AppTranslator.Get("shape_dist_traveled_validation.recommended"))
		addMessage(warn, s)
		return
	}

	if s == types.SEVERITY_FORBIDDEN {
		addMessage(i18n.AppTranslator.Get("shape_dist_traveled_validation.forbidden"), s)
		return
	}

	// Validate shape_dist_traveled
	if *shape.ShapeDistTraveled < 0 {
		addMessage(i18n.AppTranslator.Get("shape_dist_traveled_validation.invalid"), types.SEVERITY_ERROR)
		return
	}
}

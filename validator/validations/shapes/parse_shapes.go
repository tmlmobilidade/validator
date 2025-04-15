package shapes

import (
	"main/validator/lib"
	"main/validator/types"
)

type parseShapeValidation struct {
	*types.Validation
}

func NewParseShapeValidation(severity *types.Severity) *parseShapeValidation {
	s := types.SEVERITY_ERROR
	if severity != nil {
		s = *severity
	}

	return &parseShapeValidation{
		Validation: &types.Validation{
			ID:          "parse_shape",
			Description: "Validate shape data",
			Severity:    s,
		},
	}
}

func (v *parseShapeValidation) Validate(gtfs types.Gtfs) (shapes []types.Shape, messages []types.Message) {
	shapeIds := make(map[string]bool)

	for i, shape := range gtfs.Files["shapes"] {
		shape, shapeMessages := parseShape(shape)
		shapes = append(shapes, shape)

		// Check for duplicate trip IDs
		if shape.ShapeId != "" {
			if shapeIds[shape.ShapeId] {
				messages = append(messages, types.Message{
					Field:        "shape_id",
					FileName:     "shapes.txt",
					Message:      "Duplicate shape_id found. Shape IDs must be unique.",
					Row:          i + 1,
					Severity:     v.Severity,
					ValidationID: v.ID,
				})
			}
			shapeIds[shape.ShapeId] = true
		}

		// Update row number and other fields for each message
		for _, msg := range shapeMessages {
			msg.Row = i + 1
			msg.FileName = "shapes.txt"
			msg.Severity = v.Severity
			msg.ValidationID = v.ID
			messages = append(messages, msg)
		}
	}
	return shapes, messages
}

// ParseShapes validates and parses a shapes.txt row into a Shape struct
func parseShape(m map[string]string) (types.Shape, []types.Message) {
	var shape types.Shape
	var messages []types.Message
	var parsingErrors []string

	// Parse required fields
	lib.ParseStringToPrimitive(m["shape_id"], &shape.ShapeId, &parsingErrors)

	// Convert Optional Primitive Values
	var shapePtLat, shapePtLon float32
	var shapePtSequence int
	var shapeDistTraveled float64

	lib.ParseStringToPrimitive(m["shape_pt_lat"], &shapePtLat, &parsingErrors)
	lib.ParseStringToPrimitive(m["shape_pt_lon"], &shapePtLon, &parsingErrors)
	lib.ParseStringToPrimitive(m["shape_pt_sequence"], &shapePtSequence, &parsingErrors)
	lib.ParseStringToPrimitive(m["shape_dist_traveled"], &shapeDistTraveled, &parsingErrors)

	shape.ShapePtLat = lib.IfThenElse(m["shape_pt_lat"] != "", &shapePtLat, nil)
	shape.ShapePtLon = lib.IfThenElse(m["shape_pt_lon"] != "", &shapePtLon, nil)
	shape.ShapePtSequence = lib.IfThenElse(m["shape_pt_sequence"] != "", &shapePtSequence, nil)
	shape.ShapeDistTraveled = lib.IfThenElse(m["shape_dist_traveled"] != "", &shapeDistTraveled, nil)

	// Handle parsing errors
	if len(parsingErrors) > 0 {
		for _, err := range parsingErrors {
			messages = append(messages, types.Message{
				Field:   "N/A",
				Message: err,
			})
		}
	}

	// Validate required fields
	if shape.ShapeId == "" {
		messages = append(messages, types.Message{
			Field:   "shape_id",
			Message: "Shape ID is required.",
		})
	}

	// Validate latitude shape lat
	if shape.ShapePtLat == nil {
		messages = append(messages, types.Message{
			Field:   "shape_pt_lat",
			Message: "Shape point latitude is required.",
		})
	} else {
		v := lib.ValidateLatitude(*shape.ShapePtLat)
		if v != "" {
			messages = append(messages, types.Message{
				Field:   "shape_pt_lat",
				Message: v,
			})
		}
	}

	// Validate longitude shape lon
	if shape.ShapePtLon == nil {
		messages = append(messages, types.Message{
			Field:   "shape_pt_lon",
			Message: "Shape point longitude is required.",
		})
	} else {
		v := lib.ValidateLongitude(*shape.ShapePtLon)
		if v != "" {
			messages = append(messages, types.Message{
				Field:   "shape_pt_lon",
				Message: v,
			})
		}
	}

	// Validate shape pt sequence
	if shape.ShapePtSequence == nil || *shape.ShapePtSequence < 0 {
		messages = append(messages, types.Message{
			Field:   "shape_pt_sequence",
			Message: "Shape point sequence must be non-negative integer.",
		})
	}

	// Validate shape dist traveled
	if shape.ShapeDistTraveled == nil || *shape.ShapeDistTraveled < 0 {
		messages = append(messages, types.Message{
			Field:   "shape_dist_traveled",
			Message: "Shape distance traveled must be non-negative float.",
		})
	}

	return shape, messages
}

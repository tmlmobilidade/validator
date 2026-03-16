package shapes

import (
	"math"
	"sort"
	"strconv"

	"main/lib"
	"main/services"
	shapes_coordinates "main/services/geo/shapes"
	"main/types"
)

type shapeDistancePoint struct {
	row          int
	sequence     int
	lat          float64
	lon          float64
	distTraveled *float64
}

const maxDistanceDeltaToleranceMeters = 20.0

func getShapeDistanceToleranceMeters(rules *types.ShapesRules) float64 {
	if rules == nil || rules.ShapeDistTolerance.Options == nil || len(*rules.ShapeDistTolerance.Options) == 0 {
		return maxDistanceDeltaToleranceMeters
	}

	value, err := strconv.ParseFloat((*rules.ShapeDistTolerance.Options)[0], 64)
	if err != nil || value < 0 {
		return maxDistanceDeltaToleranceMeters
	}

	return value
}

// ShapeCoordinatesDistancesValidation validates if the geometric distance between
// consecutive points matches the shape_dist_traveled increment.
func ShapeCoordinatesDistancesValidation(shapes []types.Shape, rules *types.ShapesRules) {
	severity := types.SEVERITY_ERROR
	if rules != nil && rules.ShapeDistTolerance.Severity != "" {
		severity = types.Severity(rules.ShapeDistTolerance.Severity)
	}

	shapeGroups := map[string][]shapeDistancePoint{}
	toleranceMeters := getShapeDistanceToleranceMeters(rules)

	for i, shape := range shapes {
		ctx := lib.NewValidationContext("coordinates", "shapes.txt", "coordinates_distances_validation", i, services.AppMessageService)
		ctx.WithSeverity(severity)

		if shape.ShapeId == nil || *shape.ShapeId == "" {
			continue
		}
		if shape.ShapePtSequence == nil || shape.ShapePtLat == nil || shape.ShapePtLon == nil {
			continue
		}

		shapeGroups[*shape.ShapeId] = append(shapeGroups[*shape.ShapeId], shapeDistancePoint{
			row:          i,
			sequence:     *shape.ShapePtSequence,
			lat:          float64(*shape.ShapePtLat),
			lon:          float64(*shape.ShapePtLon),
			distTraveled: shape.ShapeDistTraveled,
		})
	}

	for _, shapeGroup := range shapeGroups {
		sort.Slice(shapeGroup, func(i, j int) bool {
			return shapeGroup[i].sequence < shapeGroup[j].sequence
		})

		for i := 1; i < len(shapeGroup); i++ {
			prev := shapeGroup[i-1]
			current := shapeGroup[i]

			if prev.distTraveled == nil || current.distTraveled == nil {
				continue
			}

			ctx := lib.NewValidationContext("coordinates", "shapes.txt", "coordinates_distances_validation", current.row, services.AppMessageService)
			ctx.WithSeverity(severity)

			realDistanceMeters := shapes_coordinates.GetDistanceBetweenPositionsMeters(
				shapes_coordinates.ShapesDistance{ShapePtLat: prev.lat, ShapePtLon: prev.lon},
				shapes_coordinates.ShapesDistance{ShapePtLat: current.lat, ShapePtLon: current.lon},
			)
			distTraveledDeltaMeters := *current.distTraveled - *prev.distTraveled

			if math.Abs(realDistanceMeters-distTraveledDeltaMeters) <= toleranceMeters && math.Abs(realDistanceMeters-distTraveledDeltaMeters) >= -toleranceMeters {
				continue
			}

			if distTraveledDeltaMeters != 0.0 {
				ctx.AddMessageWithSeverity(ctx.GetTranslatedMessage("coordinates_distances_validation.invalid_distances", strconv.FormatFloat(prev.lat, 'f', -1, 64), strconv.FormatFloat(prev.lon, 'f', -1, 64), strconv.FormatFloat(current.lat, 'f', -1, 64), strconv.FormatFloat(current.lon, 'f', -1, 64), distTraveledDeltaMeters, realDistanceMeters))
			}
		}
	}
}

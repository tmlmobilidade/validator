package services

import (
	"math"
	"sort"
	"strconv"

	"main/types"
)

const EARTH_RADIUS_METERS = 6371000.0
const SEGMENT_LENGTH = 50.0 // 50m segments reduce Haversine calls 5x vs 10m while still validating 100m stop distance
const MAX_STOP_DISTANCE_TO_CLOSEST_SHAPE_POINT_METERS = 100.0
const MAX_SHAPE_POINT_DISTANCE_METERS = 1000.0

const shapeDistTraveledKilometersThreshold = 800.0

type shapePointWithSequence struct {
	sequence   int
	coordinate types.ShapesDistance
}

type ShapeChunkedData struct {
	ChunkedCoordinates []types.ShapesDistance
	OriginalPoints     []shapePointWithSequence
}

// FindClosestOriginalPoint returns the sequence, lat, lon of the original shape point
// closest to stopPoint. Used when resolving chunked coords back to original for validation messages.
func (d *ShapeChunkedData) FindClosestOriginalPoint(stopPoint types.ShapesDistance) (seq int, lat, lon float64) {
	minDist := math.MaxFloat64
	for _, pt := range d.OriginalPoints {
		dist := getDistanceBetweenPositions(stopPoint, pt.coordinate)
		if dist < minDist {
			minDist = dist
			seq = pt.sequence
			lat = pt.coordinate.ShapePtLat
			lon = pt.coordinate.ShapePtLon
		}
	}
	return seq, lat, lon
}

// ShapeDistTraveledKilometersThreshold returns the threshold: if max shape_dist_traveled < this, units are km.
func ShapeDistTraveledKilometersThreshold() float64 {
	return shapeDistTraveledKilometersThreshold
}

// ShapeDistTraveledToMeters converts shape_dist_traveled to meters.
// Uses heuristic based on last (max) value in shape: if < 800 assume kilometers, else meters.
func ShapeDistTraveledToMeters(value float64, maxInShape float64) float64 {
	if maxInShape < shapeDistTraveledKilometersThreshold {
		return value * 1000 // km to m
	}
	return value // already in meters
}

// hasConsecutiveShapeDistanceInconsistency checks if the distance between two consecutive shapes is greater than the maximum allowed distance
func hasConsecutiveShapeDistanceInconsistency(orderedCoordinates []types.ShapesDistance) bool {
	if len(orderedCoordinates) != 2 {
		return false
	}

	distanceMeters := getDistanceBetweenPositions(orderedCoordinates[0], orderedCoordinates[1])
	if distanceMeters > MAX_SHAPE_POINT_DISTANCE_METERS {
		return true
	}

	return false
}

// getDistanceBetweenPositions calculates the distance between two shapes in meters
func getDistanceBetweenPositions(a, b types.ShapesDistance) float64 {
	lat1 := a.ShapePtLat * math.Pi / 180
	lon1 := a.ShapePtLon * math.Pi / 180
	lat2 := b.ShapePtLat * math.Pi / 180
	lon2 := b.ShapePtLon * math.Pi / 180

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	sinLat := math.Sin(dLat / 2)
	sinLon := math.Sin(dLon / 2)
	h := sinLat*sinLat + math.Cos(lat1)*math.Cos(lat2)*sinLon*sinLon

	return 2 * EARTH_RADIUS_METERS * math.Atan2(math.Sqrt(h), math.Sqrt(1-h))
}

// GetDistanceBetweenPositionsMeters calculates the distance between two shapes in meters
func GetDistanceBetweenPositionsMeters(a, b types.ShapesDistance) float64 {
	return getDistanceBetweenPositions(a, b)
}

// interpolatePositions interpolates the position between two shapes
func interpolatePositions(a, b types.ShapesDistance, ratio float64) types.ShapesDistance {
	if ratio < 0 {
		ratio = 0
	}
	if ratio > 1 {
		ratio = 1
	}

	return types.ShapesDistance{
		ShapePtLat: a.ShapePtLat + (b.ShapePtLat-a.ShapePtLat)*ratio,
		ShapePtLon: a.ShapePtLon + (b.ShapePtLon-a.ShapePtLon)*ratio,
	}
}

// ChunkShapesDistances chunks the shapes distances into segments
func ChunkShapesDistances(distances []types.ShapesDistance) []types.ShapesDistance {

	if len(distances) == 0 {
		return distances
	}

	coordinates := make([]types.ShapesDistance, 0, len(distances))
	for _, distance := range distances {
		coordinates = append(coordinates, types.ShapesDistance{
			ShapePtLat: distance.ShapePtLat,
			ShapePtLon: distance.ShapePtLon,
		})
	}

	cumDist := make([]float64, 0, len(coordinates))
	cumDist = append(cumDist, 0)
	for i := 0; i < len(coordinates)-1; i++ {
		cumDist = append(cumDist, cumDist[i]+getDistanceBetweenPositions(coordinates[i], coordinates[i+1]))
	}

	totalLength := cumDist[len(cumDist)-1]
	if totalLength == 0 {
		return distances
	}

	nodeCount := int(math.Floor(totalLength/SEGMENT_LENGTH)) + 1
	result := make([]types.ShapesDistance, 0, nodeCount+1)
	segIdx := 0

	for i := range nodeCount {
		targetDist := SEGMENT_LENGTH * float64(i)

		for segIdx < len(coordinates)-2 && cumDist[segIdx+1] < targetDist {
			segIdx++
		}

		segStart := cumDist[segIdx]
		segEnd := cumDist[segIdx+1]
		ratio := 0.0
		if segEnd > segStart {
			ratio = (targetDist - segStart) / (segEnd - segStart)
		}

		result = append(result, interpolatePositions(coordinates[segIdx], coordinates[segIdx+1], ratio))
	}

	lastCoord := coordinates[len(coordinates)-1]
	lastResult := result[len(result)-1]
	if lastResult.ShapePtLat != lastCoord.ShapePtLat || lastResult.ShapePtLon != lastCoord.ShapePtLon {
		result = append(result, lastCoord)
	}

	return result
}

// ShapePointIsCloseToBeforeShapePoint checks if the shape point is close to the before shape point
func ShapePointIsCloseToBeforeShapePoint(beforeShapePoint *types.Shape, shapePoint *types.Shape) bool {
	if shapePoint == nil || beforeShapePoint == nil || shapePoint.ShapePtLat == nil || shapePoint.ShapePtLon == nil || beforeShapePoint.ShapePtLat == nil || beforeShapePoint.ShapePtLon == nil {
		return false
	}

	beforeShapePointCoordinate := types.ShapesDistance{
		ShapePtLat: float64(*beforeShapePoint.ShapePtLat),
		ShapePtLon: float64(*beforeShapePoint.ShapePtLon),
	}

	shapePointCoordinate := types.ShapesDistance{
		ShapePtLat: float64(*shapePoint.ShapePtLat),
		ShapePtLon: float64(*shapePoint.ShapePtLon),
	}

	coordinatesToValidate := []types.ShapesDistance{shapePointCoordinate, beforeShapePointCoordinate}

	if hasConsecutiveShapeDistanceInconsistency(coordinatesToValidate) {
		return false
	}

	distanceMeters := getDistanceBetweenPositions(shapePointCoordinate, beforeShapePointCoordinate)
	return distanceMeters <= MAX_SHAPE_POINT_DISTANCE_METERS
}

// BuildShapeChunkedData builds chunked coordinates and original points from shape coordinates.
// Used to populate the shape cache for performance.
func BuildShapeChunkedData(points []types.ShapeCoordinatesValidation) *ShapeChunkedData {
	if len(points) == 0 {
		return nil
	}
	sorted := make([]types.ShapeCoordinatesValidation, len(points))
	copy(sorted, points)
	sort.Slice(sorted, func(i, j int) bool {
		seqI, _ := strconv.Atoi(sorted[i].ShapePtSeq)
		seqJ, _ := strconv.Atoi(sorted[j].ShapePtSeq)
		return seqI < seqJ
	})

	orderedCoordinates := make([]types.ShapesDistance, 0, len(sorted))
	origPoints := make([]shapePointWithSequence, 0, len(sorted))
	for _, p := range sorted {
		lat, err := strconv.ParseFloat(p.ShapePtLat, 64)
		if err != nil {
			continue
		}
		lon, err := strconv.ParseFloat(p.ShapePtLon, 64)
		if err != nil {
			continue
		}
		seq, _ := strconv.Atoi(p.ShapePtSeq)
		coord := types.ShapesDistance{ShapePtLat: lat, ShapePtLon: lon}
		orderedCoordinates = append(orderedCoordinates, coord)
		origPoints = append(origPoints, shapePointWithSequence{sequence: seq, coordinate: coord})
	}
	if len(orderedCoordinates) == 0 {
		return nil
	}
	chunked := ChunkShapesDistances(orderedCoordinates)
	return &ShapeChunkedData{ChunkedCoordinates: chunked, OriginalPoints: origPoints}
}

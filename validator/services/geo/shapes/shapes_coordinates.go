package services

import (
	"math"
	"sort"
	"strconv"

	"main/lib"
	"main/types"
)

const SEGMENT_LENGTH = 50.0 // 50m segments reduce Haversine calls 5x vs 10m while still validating 100m stop distance
const MAX_STOP_DISTANCE_TO_CLOSEST_SHAPE_POINT_METERS = 100.0
const MAX_SHAPE_POINT_DISTANCE_METERS = 1000.0

const shapeDistTraveledKilometersThreshold = 800.0

type shapePointWithSequence struct {
	sequence   int
	coordinate types.Coordinates
}

type ShapeChunkedData struct {
	ChunkedCoordinates []types.Coordinates
	OriginalPoints     []shapePointWithSequence
}

/*
Finds the original shape point in the slice whose coordinates are closest to the provided stopPoint.

It returns the sequence number, latitude, and longitude of this closest original point.

This is primarily used to map a chunked/interpolated coordinate back to its source GTFS shape row
for generating clear validation messages.
*/
func (d *ShapeChunkedData) FindClosestOriginalPoint(stopPoint types.Coordinates) (seq int, lat, lon float64) {
	// Initialize minimum distance to the highest possible value.
	minDist := math.MaxFloat64

	// Iterate through each original shape point and find the one closest to stopPoint.
	for _, pt := range d.OriginalPoints {
		dist := lib.HaversineDistance(stopPoint, pt.coordinate)
		if dist < minDist {
			minDist = dist
			seq = pt.sequence
			lat = pt.coordinate.Lat
			lon = pt.coordinate.Lng
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
func hasConsecutiveShapeDistanceInconsistency(orderedCoordinates []types.Coordinates) bool {
	if len(orderedCoordinates) != 2 {
		return false
	}

	distanceMeters := lib.HaversineDistance(orderedCoordinates[0], orderedCoordinates[1])
	if distanceMeters > MAX_SHAPE_POINT_DISTANCE_METERS {
		return true
	}

	return false
}

// interpolatePositions interpolates the position between two shapes
func interpolatePositions(a, b types.Coordinates, ratio float64) types.Coordinates {
	if ratio < 0 {
		ratio = 0
	}
	if ratio > 1 {
		ratio = 1
	}

	return types.Coordinates{
		Lat: a.Lat + (b.Lat-a.Lat)*ratio,
		Lng: a.Lng + (b.Lng-a.Lng)*ratio,
	}
}

// ChunkCoordinatess chunks the shapes distances into segments
func ChunkCoordinatess(distances []types.Coordinates) []types.Coordinates {

	if len(distances) == 0 {
		return distances
	}

	coordinates := make([]types.Coordinates, 0, len(distances))
	for _, distance := range distances {
		coordinates = append(coordinates, types.Coordinates{
			Lat: distance.Lat,
			Lng: distance.Lng,
		})
	}

	cumDist := make([]float64, 0, len(coordinates))
	cumDist = append(cumDist, 0)
	for i := 0; i < len(coordinates)-1; i++ {
		cumDist = append(cumDist, cumDist[i]+lib.HaversineDistance(coordinates[i], coordinates[i+1]))
	}

	totalLength := cumDist[len(cumDist)-1]
	if totalLength == 0 {
		return distances
	}

	nodeCount := int(math.Floor(totalLength/SEGMENT_LENGTH)) + 1
	result := make([]types.Coordinates, 0, nodeCount+1)
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
	if lastResult.Lat != lastCoord.Lat || lastResult.Lng != lastCoord.Lng {
		result = append(result, lastCoord)
	}

	return result
}

// ShapePointIsCloseToBeforeShapePoint checks if the shape point is close to the before shape point
func ShapePointIsCloseToBeforeShapePoint(beforeShapePoint *types.Shape, shapePoint *types.Shape) bool {
	if shapePoint == nil || beforeShapePoint == nil || shapePoint.ShapePtLat == nil || shapePoint.ShapePtLon == nil || beforeShapePoint.ShapePtLat == nil || beforeShapePoint.ShapePtLon == nil {
		return false
	}

	beforeShapePointCoordinate := types.Coordinates{
		Lat: float64(*beforeShapePoint.ShapePtLat),
		Lng: float64(*beforeShapePoint.ShapePtLon),
	}

	shapePointCoordinate := types.Coordinates{
		Lat: float64(*shapePoint.ShapePtLat),
		Lng: float64(*shapePoint.ShapePtLon),
	}

	coordinatesToValidate := []types.Coordinates{shapePointCoordinate, beforeShapePointCoordinate}

	if hasConsecutiveShapeDistanceInconsistency(coordinatesToValidate) {
		return false
	}

	distanceMeters := lib.HaversineDistance(shapePointCoordinate, beforeShapePointCoordinate)
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

	orderedCoordinates := make([]types.Coordinates, 0, len(sorted))
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
		coord := types.Coordinates{Lat: lat, Lng: lon}
		orderedCoordinates = append(orderedCoordinates, coord)
		origPoints = append(origPoints, shapePointWithSequence{sequence: seq, coordinate: coord})
	}
	if len(orderedCoordinates) == 0 {
		return nil
	}
	chunked := ChunkCoordinatess(orderedCoordinates)
	return &ShapeChunkedData{ChunkedCoordinates: chunked, OriginalPoints: origPoints}
}

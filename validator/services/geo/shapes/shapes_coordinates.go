package services

import "math"

type shapesDistances struct {
	ShapePtLat float64
	ShapePtLon float64
}

const earthRadiusMeters = 6371000.0
const segmentLength = 10.0

func getDistanceBetweenPositions(a, b shapesDistances) float64 {
	lat1 := a.ShapePtLat * math.Pi / 180
	lon1 := a.ShapePtLon * math.Pi / 180
	lat2 := b.ShapePtLat * math.Pi / 180
	lon2 := b.ShapePtLon * math.Pi / 180

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	sinLat := math.Sin(dLat / 2)
	sinLon := math.Sin(dLon / 2)
	h := sinLat*sinLat + math.Cos(lat1)*math.Cos(lat2)*sinLon*sinLon

	return 2 * earthRadiusMeters * math.Atan2(math.Sqrt(h), math.Sqrt(1-h))
}

func interpolatePositions(a, b shapesDistances, ratio float64) shapesDistances {
	if ratio < 0 {
		ratio = 0
	}
	if ratio > 1 {
		ratio = 1
	}

	return shapesDistances{
		ShapePtLat: a.ShapePtLat + (b.ShapePtLat-a.ShapePtLat)*ratio,
		ShapePtLon: a.ShapePtLon + (b.ShapePtLon-a.ShapePtLon)*ratio,
	}
}

func chunkShapesDistances(distances []shapesDistances) []shapesDistances {

	if len(distances) == 0 {
		return distances
	}

	coordinates := make([]shapesDistances, 0, len(distances))
	for _, distance := range distances {
		coordinates = append(coordinates, shapesDistances{
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

	nodeCount := int(math.Floor(totalLength/segmentLength)) + 1
	result := make([]shapesDistances, 0, nodeCount+1)
	segIdx := 0

	for i := 0; i < nodeCount; i++ {
		targetDist := segmentLength * float64(i)

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

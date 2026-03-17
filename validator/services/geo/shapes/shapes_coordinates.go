package services

import (
	"math"
	"sort"
	"strconv"

	"main/types"
)

const earthRadiusMeters = 6371000.0
const segmentLength = 10.0
const MaxStopDistanceToClosestShapeMeters = 100.0
const MaxShapePointDistanceMeters = 1000.0

type StopClosestShapeInfo struct {
	ShapeID             string
	DistanceMeters      float64
	ClosestShapePtLat   float64
	ClosestShapePtLon   float64
	ClosestShapePtSeq   int
}

type ShapesDistance struct {
	ShapePtLat float64
	ShapePtLon float64
}

type shapePointWithSequence struct {
	sequence   int
	coordinate ShapesDistance
}

func hasConsecutiveShapeDistanceInconsistency(orderedCoordinates []ShapesDistance) bool {
	if len(orderedCoordinates) < 2 {
		return false
	}

	for i := 1; i < len(orderedCoordinates); i++ {
		distanceMeters := getDistanceBetweenPositions(orderedCoordinates[i-1], orderedCoordinates[i])
		if distanceMeters > MaxShapePointDistanceMeters {
			return true
		}
	}

	return false
}

func getDistanceBetweenPositions(a, b ShapesDistance) float64 {
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

func GetDistanceBetweenPositionsMeters(a, b ShapesDistance) float64 {
	return getDistanceBetweenPositions(a, b)
}

func interpolatePositions(a, b ShapesDistance, ratio float64) ShapesDistance {
	if ratio < 0 {
		ratio = 0
	}
	if ratio > 1 {
		ratio = 1
	}

	return ShapesDistance{
		ShapePtLat: a.ShapePtLat + (b.ShapePtLat-a.ShapePtLat)*ratio,
		ShapePtLon: a.ShapePtLon + (b.ShapePtLon-a.ShapePtLon)*ratio,
	}
}

func ChunkShapesDistances(distances []ShapesDistance) []ShapesDistance {

	if len(distances) == 0 {
		return distances
	}

	coordinates := make([]ShapesDistance, 0, len(distances))
	for _, distance := range distances {
		coordinates = append(coordinates, ShapesDistance{
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
	result := make([]ShapesDistance, 0, nodeCount+1)
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

func ShapeIsCloseToOtherShape(shape *types.Shape, otherShape *types.Shape, nextShape ...*types.Shape) (bool, error) {
	if shape == nil || otherShape == nil || shape.ShapePtLat == nil || shape.ShapePtLon == nil || otherShape.ShapePtLat == nil || otherShape.ShapePtLon == nil {
		return false, nil
	}

	shapeCoordinate := ShapesDistance{
		ShapePtLat: float64(*shape.ShapePtLat),
		ShapePtLon: float64(*shape.ShapePtLon),
	}
	otherShapeCoordinate := ShapesDistance{
		ShapePtLat: float64(*otherShape.ShapePtLat),
		ShapePtLon: float64(*otherShape.ShapePtLon),
	}
	coordinatesToValidate := []ShapesDistance{shapeCoordinate, otherShapeCoordinate}

	if len(nextShape) > 0 && nextShape[0] != nil && nextShape[0].ShapePtLat != nil && nextShape[0].ShapePtLon != nil {
		coordinatesToValidate = append(coordinatesToValidate, ShapesDistance{
			ShapePtLat: float64(*nextShape[0].ShapePtLat),
			ShapePtLon: float64(*nextShape[0].ShapePtLon),
		})
	}

	if hasConsecutiveShapeDistanceInconsistency(coordinatesToValidate) {
		return false, nil
	}

	distanceMeters := getDistanceBetweenPositions(shapeCoordinate, otherShapeCoordinate)
	return distanceMeters <= MaxShapePointDistanceMeters, nil
}

func BuildStopClosestShapeDistanceMap(gtfs *types.Gtfs) (map[string]StopClosestShapeInfo, error) {
	stopClosestShapeDistance := map[string]StopClosestShapeInfo{}
	if gtfs == nil {
		return stopClosestShapeDistance, nil
	}

	tripToShapeID := make(map[string]string)
	if err := gtfs.IterateTrips(func(_ int, rawTrip types.TripRaw) error {
		if rawTrip.TripId == "" || rawTrip.ShapeId == "" {
			return nil
		}
		tripToShapeID[rawTrip.TripId] = rawTrip.ShapeId
		return nil
	}); err != nil {
		return nil, err
	}

	stopToShapeIDs := make(map[string]map[string]struct{})
	if err := gtfs.IterateStopTimes(func(_ int, rawStopTime types.StopTimeRaw) error {
		if rawStopTime.StopId == "" || rawStopTime.TripId == "" {
			return nil
		}

		shapeID, ok := tripToShapeID[rawStopTime.TripId]
		if !ok {
			return nil
		}

		if _, exists := stopToShapeIDs[rawStopTime.StopId]; !exists {
			stopToShapeIDs[rawStopTime.StopId] = map[string]struct{}{}
		}

		stopToShapeIDs[rawStopTime.StopId][shapeID] = struct{}{}
		return nil
	}); err != nil {
		return nil, err
	}

	shapeCoordinatesByID := make(map[string][]shapePointWithSequence)
	if err := gtfs.IterateShapes(func(_ int, rawShape types.ShapeRaw) error {
		if rawShape.ShapeId == "" || rawShape.ShapePtSequence == "" || rawShape.ShapePtLat == "" || rawShape.ShapePtLon == "" {
			return nil
		}

		sequence, err := strconv.Atoi(rawShape.ShapePtSequence)
		if err != nil {
			return nil
		}

		lat, err := strconv.ParseFloat(rawShape.ShapePtLat, 64)
		if err != nil {
			return nil
		}

		lon, err := strconv.ParseFloat(rawShape.ShapePtLon, 64)
		if err != nil {
			return nil
		}

		shapeCoordinatesByID[rawShape.ShapeId] = append(shapeCoordinatesByID[rawShape.ShapeId], shapePointWithSequence{
			sequence: sequence,
			coordinate: ShapesDistance{
				ShapePtLat: lat,
				ShapePtLon: lon,
			},
		})

		return nil
	}); err != nil {
		return nil, err
	}

	chunkedShapeCoordinates := make(map[string][]ShapesDistance, len(shapeCoordinatesByID))
	for shapeID, shapePoints := range shapeCoordinatesByID {
		sort.Slice(shapePoints, func(i, j int) bool {
			return shapePoints[i].sequence < shapePoints[j].sequence
		})

		orderedCoordinates := make([]ShapesDistance, 0, len(shapePoints))
		for _, point := range shapePoints {
			orderedCoordinates = append(orderedCoordinates, point.coordinate)
		}

		chunkedShapeCoordinates[shapeID] = ChunkShapesDistances(orderedCoordinates)
	}

	if err := gtfs.IterateStops(func(_ int, rawStop types.StopRaw) error {
		if rawStop.StopId == "" {
			return nil
		}

		stopShapeIDs, ok := stopToShapeIDs[rawStop.StopId]
		if !ok || len(stopShapeIDs) == 0 {
			return nil
		}

		lat, err := strconv.ParseFloat(rawStop.StopLat, 64)
		if err != nil {
			return nil
		}

		lon, err := strconv.ParseFloat(rawStop.StopLon, 64)
		if err != nil {
			return nil
		}

		stopPoint := ShapesDistance{ShapePtLat: lat, ShapePtLon: lon}
		minDistance := math.MaxFloat64
		closestShapeID := ""
		closestCoord := ShapesDistance{}
		foundDistance := false

		for shapeID := range stopShapeIDs {
			chunkedCoordinates, hasCoordinates := chunkedShapeCoordinates[shapeID]
			if !hasCoordinates {
				continue
			}

			for _, coordinate := range chunkedCoordinates {
				distance := getDistanceBetweenPositions(stopPoint, coordinate)
				if distance < minDistance {
					minDistance = distance
					closestShapeID = shapeID
					closestCoord = coordinate
					foundDistance = true
				}
			}
		}

		if foundDistance {
			closestSeq := 0
			closestLat := closestCoord.ShapePtLat
			closestLon := closestCoord.ShapePtLon
			if shapePoints, ok := shapeCoordinatesByID[closestShapeID]; ok {
				minDistToOrig := math.MaxFloat64
				for _, pt := range shapePoints {
					d := getDistanceBetweenPositions(stopPoint, pt.coordinate)
					if d < minDistToOrig {
						minDistToOrig = d
						closestSeq = pt.sequence
						closestLat = pt.coordinate.ShapePtLat
						closestLon = pt.coordinate.ShapePtLon
					}
				}
			}

			stopClosestShapeDistance[rawStop.StopId] = StopClosestShapeInfo{
				ShapeID:           closestShapeID,
				DistanceMeters:     minDistance,
				ClosestShapePtLat: closestLat,
				ClosestShapePtLon: closestLon,
				ClosestShapePtSeq: closestSeq,
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return stopClosestShapeDistance, nil
}

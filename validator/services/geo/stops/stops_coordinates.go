package stops

import (
	"main/lib"
	shapes_coordinates "main/services/geo/shapes"
	"main/types"
	"math"
	"strconv"
)

// BuildStopClosestShapePointsDistanceMap builds the stop closest shape distance map.
// For each stop, it considers only shapes from stopToShapeIDs (shapes used on trips that visit that stop).
func BuildStopClosestShapePointsDistanceMap(
	stops []types.StopCoordinatesValidation,
	stopToShapeIDs map[string][]string,
	shapeCoordinatesByID map[string][]types.ShapeCoordinatesValidation,
) (map[string]types.StopClosestShapePointsInfo, error) {
	chunkedByID := make(map[string]*shapes_coordinates.ShapeChunkedData)
	for shapeID, points := range shapeCoordinatesByID {
		if data := shapes_coordinates.BuildShapeChunkedData(points); data != nil {
			chunkedByID[shapeID] = data
		}
	}
	return BuildStopClosestShapePointsDistanceMapFromCache(stops, stopToShapeIDs, chunkedByID)
}

// BuildStopClosestShapePointsDistanceMapFromCache uses pre-built chunked shape data for performance.
func BuildStopClosestShapePointsDistanceMapFromCache(
	stops []types.StopCoordinatesValidation,
	stopToShapeIDs map[string][]string,
	shapeChunkedCache map[string]*shapes_coordinates.ShapeChunkedData,
) (map[string]types.StopClosestShapePointsInfo, error) {
	stopClosestShapePointsDistance := make(map[string]types.StopClosestShapePointsInfo)
	lib.AppLogger.Accent("Building stop closest shape points distance map from cache")

	for _, stop := range stops {
		if stop.StopId == "" {
			continue
		}
		stopShapeIDs, ok := stopToShapeIDs[stop.StopId]
		if !ok || len(stopShapeIDs) == 0 {
			continue
		}

		lat, err := strconv.ParseFloat(stop.StopLat, 64)
		if err != nil {
			continue
		}
		lon, err := strconv.ParseFloat(stop.StopLon, 64)
		if err != nil {
			continue
		}

		stopPoint := types.Coordinates{Lat: lat, Lng: lon}
		minDistance := math.MaxFloat64
		closestShapeID := ""
		closestCoord := types.Coordinates{}
		foundDistance := false

		for _, shapeID := range stopShapeIDs {
			data := shapeChunkedCache[shapeID]
			if data == nil {
				continue
			}

			for _, coordinate := range data.ChunkedCoordinates {
				distance := lib.HaversineDistance(stopPoint, coordinate)
				if distance < minDistance {
					minDistance = distance
					closestShapeID = shapeID
					closestCoord = coordinate
					foundDistance = true
					if distance < shapes_coordinates.MAX_STOP_DISTANCE_TO_CLOSEST_SHAPE_POINT_METERS {
						break
					}
				}
			}
			if foundDistance && minDistance < shapes_coordinates.MAX_STOP_DISTANCE_TO_CLOSEST_SHAPE_POINT_METERS {
				break
			}
		}

		if foundDistance {
			if minDistance < shapes_coordinates.MAX_STOP_DISTANCE_TO_CLOSEST_SHAPE_POINT_METERS {
				continue
			}
			closestSeq := 0
			closestLat := closestCoord.Lat
			closestLon := closestCoord.Lng
			if data := shapeChunkedCache[closestShapeID]; data != nil {
				closestSeq, closestLat, closestLon = data.FindClosestOriginalPoint(stopPoint)
			}

			stopClosestShapePointsDistance[stop.StopId] = types.StopClosestShapePointsInfo{
				ShapeID:           closestShapeID,
				DistanceMeters:    minDistance,
				ClosestShapePtLat: closestLat,
				ClosestShapePtLon: closestLon,
				ClosestShapePtSeq: closestSeq,
			}
		}
	}
	lib.AppLogger.Accent("Stop closest shape points distance map built")
	return stopClosestShapePointsDistance, nil
}

// StopShapeCacheKey returns the cache key for (stop_id, shape_id) pairs.
func StopShapeCacheKey(stopID, shapeID string) string {
	return stopID + "|" + shapeID
}

// BuildStopClosestShapePointsDistanceMapPerStopShape builds a cache of stop-to-shape distance
// violations keyed by "stop_id|shape_id". Only stores entries where distance > MAX_STOP_DISTANCE_TO_CLOSEST_SHAPE_POINT_METERS.
// Used for performance: compute once for all unique (stop, shape) pairs instead of per-trip.
func BuildStopClosestShapePointsDistanceMapPerStopShape(
	stops []types.StopCoordinatesValidation,
	stopToShapeIDs map[string][]string,
	shapeChunkedCache map[string]*shapes_coordinates.ShapeChunkedData,
) map[string]types.StopClosestShapePointsInfo {
	result := make(map[string]types.StopClosestShapePointsInfo)
	lib.AppLogger.Accent("Building stop-to-shape distance cache (per stop,shape)...")

	for _, stop := range stops {
		if stop.StopId == "" {
			continue
		}
		stopShapeIDs, ok := stopToShapeIDs[stop.StopId]
		if !ok || len(stopShapeIDs) == 0 {
			continue
		}

		lat, err := strconv.ParseFloat(stop.StopLat, 64)
		if err != nil {
			continue
		}
		lon, err := strconv.ParseFloat(stop.StopLon, 64)
		if err != nil {
			continue
		}

		stopPoint := types.Coordinates{Lat: lat, Lng: lon}

		for _, shapeID := range stopShapeIDs {
			data := shapeChunkedCache[shapeID]
			if data == nil {
				continue
			}

			minDistance := math.MaxFloat64
			closestCoord := types.Coordinates{}
			foundDistance := false

			for _, coordinate := range data.ChunkedCoordinates {
				distance := lib.HaversineDistance(stopPoint, coordinate)
				if distance < minDistance {
					minDistance = distance
					closestCoord = coordinate
					foundDistance = true
					if distance < shapes_coordinates.MAX_STOP_DISTANCE_TO_CLOSEST_SHAPE_POINT_METERS {
						break
					}
				}
			}

			if foundDistance && minDistance > shapes_coordinates.MAX_STOP_DISTANCE_TO_CLOSEST_SHAPE_POINT_METERS {
				closestSeq := 0
				closestLat := closestCoord.Lat
				closestLon := closestCoord.Lng
				if data != nil {
					closestSeq, closestLat, closestLon = data.FindClosestOriginalPoint(stopPoint)
				}

				key := StopShapeCacheKey(stop.StopId, shapeID)
				result[key] = types.StopClosestShapePointsInfo{
					ShapeID:           shapeID,
					DistanceMeters:    minDistance,
					ClosestShapePtLat: closestLat,
					ClosestShapePtLon: closestLon,
					ClosestShapePtSeq: closestSeq,
				}
			}
		}
	}
	lib.AppLogger.Accent("Stop-to-shape distance cache built")
	return result
}

package services

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"

	"main/lib"
	"main/types"
)

var (
	municipalityCoordinatesMutex sync.RWMutex
	municipalityGeometries       []municipalityGeometry
)

type point struct {
	Lon float64
	Lat float64
}

type municipalityPolygon []point

type municipalityGeometry struct {
	MunicipalityID string
	Polygons       []municipalityPolygon
	MinLat         float64
	MaxLat         float64
	MinLon         float64
	MaxLon         float64
}

func MunicipalityIDFromRaw(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}

	var asString string
	if err := json.Unmarshal(raw, &asString); err == nil {
		return strings.TrimSpace(asString)
	}

	var asFloat float64
	if err := json.Unmarshal(raw, &asFloat); err == nil {
		return strings.TrimSpace(strconv.FormatFloat(asFloat, 'f', -1, 64))
	}

	return ""
}

func polygonBBox(points municipalityPolygon) (minLat, maxLat, minLon, maxLon float64) {
	minLat, maxLat = points[0].Lat, points[0].Lat
	minLon, maxLon = points[0].Lon, points[0].Lon
	for _, p := range points[1:] {
		if p.Lat < minLat {
			minLat = p.Lat
		}
		if p.Lat > maxLat {
			maxLat = p.Lat
		}
		if p.Lon < minLon {
			minLon = p.Lon
		}
		if p.Lon > maxLon {
			maxLon = p.Lon
		}
	}
	return minLat, maxLat, minLon, maxLon
}

const epsilon = 1e-9

// Municipality polygons are simplified and represent land boundaries.
// A wider tolerance avoids false negatives for shape points near coastlines
// and estuaries that still belong to Portugal's operational transit area.
const portugalBoundaryToleranceMeters = 1000.0
const metersPerDegreeLatitude = 111320.0

func pointOnSegment(lon, lat float64, a, b point) bool {
	// Check collinearity first using cross product.
	cross := (b.Lon-a.Lon)*(lat-a.Lat) - (b.Lat-a.Lat)*(lon-a.Lon)
	if math.Abs(cross) > epsilon {
		return false
	}

	// Then ensure projected point lies on segment bounds.
	minLon := math.Min(a.Lon, b.Lon) - epsilon
	maxLon := math.Max(a.Lon, b.Lon) + epsilon
	minLat := math.Min(a.Lat, b.Lat) - epsilon
	maxLat := math.Max(a.Lat, b.Lat) + epsilon

	return lon >= minLon && lon <= maxLon && lat >= minLat && lat <= maxLat
}

func pointInPolygon(lon, lat float64, polygon municipalityPolygon) bool {
	if len(polygon) < 3 {
		return false
	}

	inside := false
	j := len(polygon) - 1
	for i := 0; i < len(polygon); i++ {
		pi := polygon[i]
		pj := polygon[j]
		if pointOnSegment(lon, lat, pj, pi) {
			return true
		}
		intersects := ((pi.Lat > lat) != (pj.Lat > lat)) &&
			(lon < (pj.Lon-pi.Lon)*(lat-pi.Lat)/(pj.Lat-pi.Lat)+pi.Lon)
		if intersects {
			inside = !inside
		}
		j = i
	}
	return inside
}

func metersPerDegreeLongitude(lat float64) float64 {
	return metersPerDegreeLatitude * math.Cos(lat*math.Pi/180.0)
}

func approxDistancePointToSegmentMeters(lon, lat float64, a, b point) float64 {
	scaleX := metersPerDegreeLongitude(lat)
	if math.Abs(scaleX) < epsilon {
		scaleX = epsilon
	}
	scaleY := metersPerDegreeLatitude

	px := lon * scaleX
	py := lat * scaleY
	ax := a.Lon * scaleX
	ay := a.Lat * scaleY
	bx := b.Lon * scaleX
	by := b.Lat * scaleY

	abX := bx - ax
	abY := by - ay
	abLenSq := abX*abX + abY*abY
	if abLenSq <= epsilon {
		dx := px - ax
		dy := py - ay
		return math.Sqrt(dx*dx + dy*dy)
	}

	t := ((px-ax)*abX + (py-ay)*abY) / abLenSq
	if t < 0 {
		t = 0
	} else if t > 1 {
		t = 1
	}

	closestX := ax + t*abX
	closestY := ay + t*abY
	dx := px - closestX
	dy := py - closestY
	return math.Sqrt(dx*dx + dy*dy)
}

func pointNearPolygonBoundary(lon, lat float64, polygon municipalityPolygon, toleranceMeters float64) bool {
	if len(polygon) < 2 {
		return false
	}

	j := len(polygon) - 1
	for i := 0; i < len(polygon); i++ {
		if approxDistancePointToSegmentMeters(lon, lat, polygon[j], polygon[i]) <= toleranceMeters {
			return true
		}
		j = i
	}

	return false
}

func findMunicipalityByCoordinates(lat, lon float64) (municipalityID string, found bool) {
	for _, geometry := range municipalityGeometries {
		if lat < geometry.MinLat || lat > geometry.MaxLat || lon < geometry.MinLon || lon > geometry.MaxLon {
			continue
		}

		for _, polygon := range geometry.Polygons {
			if pointInPolygon(lon, lat, polygon) {
				return geometry.MunicipalityID, true
			}
		}
	}

	return "", false
}

func findMunicipalityByCoordinatesWithTolerance(lat, lon, toleranceMeters float64) (municipalityID string, found bool) {
	if municipalityID, found = findMunicipalityByCoordinates(lat, lon); found {
		return municipalityID, true
	}

	latPadding := toleranceMeters / metersPerDegreeLatitude
	lonScale := metersPerDegreeLongitude(lat)
	if math.Abs(lonScale) < epsilon {
		lonScale = epsilon
	}
	lonPadding := toleranceMeters / lonScale
	if lonPadding < 0 {
		lonPadding = -lonPadding
	}

	for _, geometry := range municipalityGeometries {
		if lat < geometry.MinLat-latPadding || lat > geometry.MaxLat+latPadding || lon < geometry.MinLon-lonPadding || lon > geometry.MaxLon+lonPadding {
			continue
		}

		for _, polygon := range geometry.Polygons {
			if pointNearPolygonBoundary(lon, lat, polygon, toleranceMeters) {
				return geometry.MunicipalityID, true
			}
		}
	}

	return "", false
}

func parsePolygonCoordinates(raw json.RawMessage) ([]municipalityPolygon, error) {
	var rings [][][]float64
	if err := json.Unmarshal(raw, &rings); err != nil {
		return nil, err
	}

	polygons := make([]municipalityPolygon, 0, len(rings))
	for _, ring := range rings {
		if len(ring) < 3 {
			continue
		}
		polygon := make(municipalityPolygon, 0, len(ring))
		for _, coord := range ring {
			if len(coord) < 2 {
				continue
			}
			polygon = append(polygon, point{Lon: coord[0], Lat: coord[1]})
		}
		if len(polygon) >= 3 {
			polygons = append(polygons, polygon)
		}
	}
	return polygons, nil
}

func parseMultiPolygonCoordinates(raw json.RawMessage) ([]municipalityPolygon, error) {
	var multi [][][][]float64
	if err := json.Unmarshal(raw, &multi); err != nil {
		return nil, err
	}

	polygons := make([]municipalityPolygon, 0)
	for _, polygonRings := range multi {
		polygonRaw, err := json.Marshal(polygonRings)
		if err != nil {
			return nil, err
		}
		parsed, err := parsePolygonCoordinates(polygonRaw)
		if err != nil {
			return nil, err
		}
		polygons = append(polygons, parsed...)
	}
	return polygons, nil
}

func parseFeatureGeometry(raw json.RawMessage) ([]municipalityPolygon, error) {
	var geometry types.MunicipalityCoordinatesGeometry
	if err := json.Unmarshal(raw, &geometry); err == nil && geometry.Type != "" {
		switch geometry.Type {
		case "Polygon":
			return parsePolygonCoordinates(geometry.Coordinates)
		case "MultiPolygon":
			return parseMultiPolygonCoordinates(geometry.Coordinates)
		default:
			return nil, fmt.Errorf("unsupported geometry type %q", geometry.Type)
		}
	}

	// Support legacy payloads where geometry is directly an array of [lon, lat] points.
	var ring [][]float64
	if err := json.Unmarshal(raw, &ring); err == nil {
		polygon := make(municipalityPolygon, 0, len(ring))
		for _, coord := range ring {
			if len(coord) < 2 {
				continue
			}
			polygon = append(polygon, point{Lon: coord[0], Lat: coord[1]})
		}
		if len(polygon) >= 3 {
			return []municipalityPolygon{polygon}, nil
		}
	}

	// Support raw polygon coordinates without the geometry object wrapper.
	if polygons, err := parsePolygonCoordinates(raw); err == nil && len(polygons) > 0 {
		return polygons, nil
	}

	// Support raw multipolygon coordinates without the geometry object wrapper.
	if polygons, err := parseMultiPolygonCoordinates(raw); err == nil && len(polygons) > 0 {
		return polygons, nil
	}

	return nil, fmt.Errorf("geometry is not a supported object, polygon, or multipolygon format")
}

func loadMunicipalityCoordinates(collection types.MunicipalityCoordinatesFeatureCollection) error {
	geometries := make([]municipalityGeometry, 0, len(collection.Features))

	for idx, feature := range collection.Features {
		municipalityID := MunicipalityIDFromRaw(feature.ID)
		if municipalityID == "" {
			return fmt.Errorf("municipality coordinates: missing _id at feature index %d", idx)
		}

		polygons, err := parseFeatureGeometry(feature.Geometry)
		if err != nil {
			return fmt.Errorf("municipality coordinates: invalid geometry at feature index %d: %w", idx, err)
		}

		if len(polygons) == 0 {
			return fmt.Errorf("municipality coordinates: empty polygons for municipality_id %q", municipalityID)
		}

		minLat, maxLat, minLon, maxLon := polygonBBox(polygons[0])
		for _, poly := range polygons[1:] {
			pMinLat, pMaxLat, pMinLon, pMaxLon := polygonBBox(poly)
			if pMinLat < minLat {
				minLat = pMinLat
			}
			if pMaxLat > maxLat {
				maxLat = pMaxLat
			}
			if pMinLon < minLon {
				minLon = pMinLon
			}
			if pMaxLon > maxLon {
				maxLon = pMaxLon
			}
		}

		geometries = append(geometries, municipalityGeometry{
			MunicipalityID: municipalityID,
			Polygons:       polygons,
			MinLat:         minLat,
			MaxLat:         maxLat,
			MinLon:         minLon,
			MaxLon:         maxLon,
		})
	}

	municipalityCoordinatesMutex.Lock()
	municipalityGeometries = geometries
	municipalityCoordinatesMutex.Unlock()

	return nil
}

func LoadMunicipalityCoordinatesFromFile(path string) error {
	trimmedPath := strings.TrimSpace(path)
	if trimmedPath == "" {
		municipalityCoordinatesMutex.Lock()
		municipalityGeometries = nil
		municipalityCoordinatesMutex.Unlock()
		lib.AppLogger.Debug("Municipality coordinates path not set; skipping coordinates preload.")
		return nil
	}

	data, err := os.ReadFile(trimmedPath)
	if err != nil {
		return fmt.Errorf("failed to read municipality coordinates file: %w", err)
	}

	var collection types.MunicipalityCoordinatesFeatureCollection
	if err := json.Unmarshal(data, &collection); err != nil {
		return fmt.Errorf("failed to parse municipality coordinates JSON: %w", err)
	}

	if err := loadMunicipalityCoordinates(collection); err != nil {
		return err
	}

	lib.AppLogger.Info(fmt.Sprintf("Loaded %d municipality geometries.", len(collection.Features)))
	return nil
}

func ResolveMunicipalityByCoordinates(lat, lon float32) (municipalityID string, found bool, enabled bool) {
	municipalityCoordinatesMutex.RLock()
	defer municipalityCoordinatesMutex.RUnlock()

	if municipalityGeometries == nil {
		return "", false, false
	}

	municipalityID, found = findMunicipalityByCoordinates(float64(lat), float64(lon))
	return municipalityID, found, true
}

func MunicipalityCoordinatesEnabled() bool {
	municipalityCoordinatesMutex.RLock()
	defer municipalityCoordinatesMutex.RUnlock()

	return municipalityGeometries != nil
}

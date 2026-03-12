package services

import (
	"encoding/json"
	"fmt"
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

func pointInPolygon(lon, lat float64, polygon municipalityPolygon) bool {
	if len(polygon) < 3 {
		return false
	}

	inside := false
	j := len(polygon) - 1
	for i := 0; i < len(polygon); i++ {
		pi := polygon[i]
		pj := polygon[j]
		intersects := ((pi.Lat > lat) != (pj.Lat > lat)) &&
			(lon < (pj.Lon-pi.Lon)*(lat-pi.Lat)/(pj.Lat-pi.Lat)+pi.Lon)
		if intersects {
			inside = !inside
		}
		j = i
	}
	return inside
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

	lat64 := float64(lat)
	lon64 := float64(lon)

	for _, geometry := range municipalityGeometries {
		if lat64 < geometry.MinLat || lat64 > geometry.MaxLat || lon64 < geometry.MinLon || lon64 > geometry.MaxLon {
			continue
		}

		for _, polygon := range geometry.Polygons {
			if pointInPolygon(lon64, lat64, polygon) {
				return geometry.MunicipalityID, true, true
			}
		}
	}

	return "", false, true
}

func MunicipalityCoordinatesEnabled() bool {
	municipalityCoordinatesMutex.RLock()
	defer municipalityCoordinatesMutex.RUnlock()

	return municipalityGeometries != nil
}

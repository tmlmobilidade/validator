package stops

import (
	"os"
	"path/filepath"
	"testing"

	"main/lib"
	"main/lib/test_helpers"
	"main/services"
	municipality_coordinates "main/services/geo/municipalities"
	"main/types"
	validations "main/validations/stops/validations"
)

func writeMunicipalityCoordinatesFixture(t *testing.T, content string) string {
	t.Helper()
	filePath := filepath.Join(t.TempDir(), "municipality_coordinates.json")
	if err := os.WriteFile(filePath, []byte(content), 0o644); err != nil {
		t.Fatalf("failed to write fixture: %v", err)
	}
	return filePath
}

func TestCoordinatesValidation_ExactMatch(t *testing.T) {
	services.AppMessageService.Clear()
	defer municipality_coordinates.LoadMunicipalityCoordinatesFromFile("")

	filePath := writeMunicipalityCoordinatesFixture(t, `{
		"type": "FeatureCollection",
		"features": [
			{
				"_id": "1506",
				"type": "Feature",
				"properties": { "name": "Lisboa", "area_ha": 1000, "district_id": "11" },
				"geometry": {
					"type": "Polygon",
					"coordinates": [[
						[-9.20, 38.65], [-9.05, 38.65], [-9.05, 38.80], [-9.20, 38.80], [-9.20, 38.65]
					]]
				}
			}
		]
	}`)
	if err := municipality_coordinates.LoadMunicipalityCoordinatesFromFile(filePath); err != nil {
		t.Fatalf("failed to load municipality coordinates fixture: %v", err)
	}

	stop := &types.Stop{
		StopLat:        lib.Ptr(float32(38.71667)),
		StopLon:        lib.Ptr(float32(-9.13333)),
		MunicipalityId: lib.Ptr("1506"),
	}

	validations.CoordinatesValidation(stop, 1, nil, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "exact match should not raise error", types.SEVERITY_ERROR)
}

func TestCoordinatesValidation_MunicipalityMismatch(t *testing.T) {
	services.AppMessageService.Clear()
	defer municipality_coordinates.LoadMunicipalityCoordinatesFromFile("")

	filePath := writeMunicipalityCoordinatesFixture(t, `{
		"type": "FeatureCollection",
		"features": [
			{
				"_id": "1506",
				"type": "Feature",
				"properties": { "name": "Lisboa", "area_ha": 1000, "district_id": "11" },
				"geometry": {
					"type": "Polygon",
					"coordinates": [[
						[-9.20, 38.65], [-9.05, 38.65], [-9.05, 38.80], [-9.20, 38.80], [-9.20, 38.65]
					]]
				}
			}
		]
	}`)
	if err := municipality_coordinates.LoadMunicipalityCoordinatesFromFile(filePath); err != nil {
		t.Fatalf("failed to load municipality coordinates fixture: %v", err)
	}

	stop := &types.Stop{
		StopLat:        lib.Ptr(float32(38.71667)),
		StopLon:        lib.Ptr(float32(-9.13333)),
		MunicipalityId: lib.Ptr("1111"),
	}

	validations.CoordinatesValidation(stop, 1, nil, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "mismatched municipality should raise error", types.SEVERITY_ERROR)
}

func TestCoordinatesValidation_CoordinateNotMapped(t *testing.T) {
	services.AppMessageService.Clear()
	defer municipality_coordinates.LoadMunicipalityCoordinatesFromFile("")

	filePath := writeMunicipalityCoordinatesFixture(t, `{
		"type": "FeatureCollection",
		"features": [
			{
				"_id": "1506",
				"type": "Feature",
				"properties": { "name": "Lisboa", "area_ha": 1000, "district_id": "11" },
				"geometry": {
					"type": "Polygon",
					"coordinates": [[
						[-9.20, 38.65], [-9.05, 38.65], [-9.05, 38.80], [-9.20, 38.80], [-9.20, 38.65]
					]]
				}
			}
		]
	}`)
	if err := municipality_coordinates.LoadMunicipalityCoordinatesFromFile(filePath); err != nil {
		t.Fatalf("failed to load municipality coordinates fixture: %v", err)
	}

	stop := &types.Stop{
		StopLat:        lib.Ptr(float32(38.90000)),
		StopLon:        lib.Ptr(float32(-9.40000)),
		MunicipalityId: lib.Ptr("1506"),
	}

	validations.CoordinatesValidation(stop, 1, nil, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "missing coordinate mapping should raise error", types.SEVERITY_ERROR)
}

func TestCoordinatesValidation_LegacyGeometryArrayFormat(t *testing.T) {
	services.AppMessageService.Clear()
	defer municipality_coordinates.LoadMunicipalityCoordinatesFromFile("")

	filePath := writeMunicipalityCoordinatesFixture(t, `{
		"type": "FeatureCollection",
		"features": [
			{
				"_id": "1506",
				"properties": { "name": "Lisboa", "area_ha": 1000, "district_id": "11" },
				"geometry": [
					[-9.20, 38.65], [-9.05, 38.65], [-9.05, 38.80], [-9.20, 38.80], [-9.20, 38.65]
				]
			}
		]
	}`)
	if err := municipality_coordinates.LoadMunicipalityCoordinatesFromFile(filePath); err != nil {
		t.Fatalf("failed to load municipality coordinates fixture: %v", err)
	}

	stop := &types.Stop{
		StopLat:        lib.Ptr(float32(38.71667)),
		StopLon:        lib.Ptr(float32(-9.13333)),
		MunicipalityId: lib.Ptr("1506"),
	}

	validations.CoordinatesValidation(stop, 1, nil, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "legacy geometry array format should be accepted", types.SEVERITY_ERROR)
}

func TestCoordinatesValidation_MapNotLoaded_SkipsValidation(t *testing.T) {
	services.AppMessageService.Clear()
	_ = municipality_coordinates.LoadMunicipalityCoordinatesFromFile("")
	defer municipality_coordinates.LoadMunicipalityCoordinatesFromFile("")

	stop := &types.Stop{
		StopLat:        lib.Ptr(float32(38.71667)),
		StopLon:        lib.Ptr(float32(-9.13333)),
		MunicipalityId: lib.Ptr("9999"),
	}

	validations.CoordinatesValidation(stop, 1, nil, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "validation should be skipped when municipality map is not loaded", types.SEVERITY_ERROR)
}

func TestCoordinatesValidation_MapLoadError_SkipsValidation(t *testing.T) {
	services.AppMessageService.Clear()
	_ = municipality_coordinates.LoadMunicipalityCoordinatesFromFile("")
	defer municipality_coordinates.LoadMunicipalityCoordinatesFromFile("")

	if err := municipality_coordinates.LoadMunicipalityCoordinatesFromFile(filepath.Join(t.TempDir(), "does-not-exist.json")); err == nil {
		t.Fatalf("expected an error when loading a non-existing municipality coordinates file")
	}

	stop := &types.Stop{
		StopLat:        lib.Ptr(float32(38.71667)),
		StopLon:        lib.Ptr(float32(-9.13333)),
		MunicipalityId: lib.Ptr("9999"),
	}

	validations.CoordinatesValidation(stop, 1, nil, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "validation should be skipped when municipality map preload fails", types.SEVERITY_ERROR)
}

func TestCoordinatesValidation_PointOnPolygonBoundary_IsMapped(t *testing.T) {
	services.AppMessageService.Clear()
	defer municipality_coordinates.LoadMunicipalityCoordinatesFromFile("")

	filePath := writeMunicipalityCoordinatesFixture(t, `{
		"type": "FeatureCollection",
		"features": [
			{
				"_id": "1506",
				"type": "Feature",
				"properties": { "name": "Lisboa", "area_ha": 1000, "district_id": "11" },
				"geometry": {
					"type": "Polygon",
					"coordinates": [[
						[-9.20, 38.65], [-9.05, 38.65], [-9.05, 38.80], [-9.20, 38.80], [-9.20, 38.65]
					]]
				}
			}
		]
	}`)
	if err := municipality_coordinates.LoadMunicipalityCoordinatesFromFile(filePath); err != nil {
		t.Fatalf("failed to load municipality coordinates fixture: %v", err)
	}

	// This point lies exactly on the left polygon edge.
	stop := &types.Stop{
		StopLat:        lib.Ptr(float32(38.72)),
		StopLon:        lib.Ptr(float32(-9.20)),
		MunicipalityId: lib.Ptr("1506"),
	}

	validations.CoordinatesValidation(stop, 1, nil, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "point on polygon boundary should be mapped", types.SEVERITY_ERROR)
}

func TestCoordinatesValidation_StopDistanceToShapeWithinLimit(t *testing.T) {
	services.AppMessageService.Clear()
	_ = municipality_coordinates.LoadMunicipalityCoordinatesFromFile("")
	defer municipality_coordinates.LoadMunicipalityCoordinatesFromFile("")

	stop := &types.Stop{
		StopId:  lib.Ptr("stop_1"),
		StopLat: lib.Ptr(float32(38.71667)),
		StopLon: lib.Ptr(float32(-9.13333)),
	}

	validations.CoordinatesValidation(stop, 1, nil, map[string]float64{
		"stop_1": 80.0,
	})
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "stop within 100m from closest shape should not raise error", types.SEVERITY_ERROR)
}

func TestCoordinatesValidation_StopDistanceToShapeTooFar(t *testing.T) {
	services.AppMessageService.Clear()
	_ = municipality_coordinates.LoadMunicipalityCoordinatesFromFile("")
	defer municipality_coordinates.LoadMunicipalityCoordinatesFromFile("")

	stop := &types.Stop{
		StopId:  lib.Ptr("stop_1"),
		StopLat: lib.Ptr(float32(38.71667)),
		StopLon: lib.Ptr(float32(-9.13333)),
	}

	validations.CoordinatesValidation(stop, 1, nil, map[string]float64{
		"stop_1": 120.0,
	})
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "stop farther than 100m from closest shape should raise error", types.SEVERITY_ERROR)
}

package stops

import (
	"os"
	"path/filepath"
	"testing"

	"main/lib"
	"main/lib/test_helpers"
	"main/services"
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

func TestCoordenatesValidation_ExactMatch(t *testing.T) {
	services.AppMessageService.Clear()
	defer services.LoadMunicipalityCoordinatesFromFile("")

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
	if err := services.LoadMunicipalityCoordinatesFromFile(filePath); err != nil {
		t.Fatalf("failed to load municipality coordinates fixture: %v", err)
	}

	stop := &types.Stop{
		StopLat:        lib.Ptr(float32(38.71667)),
		StopLon:        lib.Ptr(float32(-9.13333)),
		MunicipalityId: lib.Ptr("1506"),
	}

	validations.CoordenatesValidation(stop, 1, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 0, "exact match should not raise error", types.SEVERITY_ERROR)
}

func TestCoordenatesValidation_MunicipalityMismatch(t *testing.T) {
	services.AppMessageService.Clear()
	defer services.LoadMunicipalityCoordinatesFromFile("")

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
	if err := services.LoadMunicipalityCoordinatesFromFile(filePath); err != nil {
		t.Fatalf("failed to load municipality coordinates fixture: %v", err)
	}

	stop := &types.Stop{
		StopLat:        lib.Ptr(float32(38.71667)),
		StopLon:        lib.Ptr(float32(-9.13333)),
		MunicipalityId: lib.Ptr("1111"),
	}

	validations.CoordenatesValidation(stop, 1, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "mismatched municipality should raise error", types.SEVERITY_ERROR)
}

func TestCoordenatesValidation_CoordinateNotMapped(t *testing.T) {
	services.AppMessageService.Clear()
	defer services.LoadMunicipalityCoordinatesFromFile("")

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
	if err := services.LoadMunicipalityCoordinatesFromFile(filePath); err != nil {
		t.Fatalf("failed to load municipality coordinates fixture: %v", err)
	}

	stop := &types.Stop{
		StopLat:        lib.Ptr(float32(38.70000)),
		StopLon:        lib.Ptr(float32(-9.10000)),
		MunicipalityId: lib.Ptr("1506"),
	}

	validations.CoordenatesValidation(stop, 1, nil)
	test_helpers.AssertMessageCount(t, services.AppMessageService, 1, "missing coordinate mapping should raise error", types.SEVERITY_ERROR)
}

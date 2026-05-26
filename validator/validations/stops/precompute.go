package stops

import (
	"encoding/json"
	"fmt"
	"main/lib"
	"main/types"
	"os"
	"path/filepath"
)

type StopsDataFlag struct {
	AgencyIDs    []string `json:"agency_ids"`
	IsHarmonized bool     `json:"is_harmonized"`
	ShortName    string   `json:"short_name"`
	StopID       string   `json:"stop_id"`
}

type StopsDataEntry struct {
	Name      string          `json:"name"`
	Latitude  float64         `json:"latitude"`
	Longitude float64         `json:"longitude"`
	Flags     []StopsDataFlag `json:"flags"`
}

// buildStopsIds loads the root-level stops_data.json and indexes stop_id values from flags.
func BuildStopsDataCache() *types.StopsDataCache {
	lib.AppLogger.Debug("Pre-computing stops_data cache...")

	fileBytes, sourcePath, err := readStopsDataFile()
	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error reading stops_data.json: %v", err))
		return &types.StopsDataCache{ByStopID: make(map[string]types.StopsDataRecord)}
	}
	lib.AppLogger.Debug(fmt.Sprintf("Loaded stops_data.json from %s", sourcePath))

	var entries []StopsDataEntry
	if err := json.Unmarshal(fileBytes, &entries); err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error parsing stops_data.json: %v", err))
		return &types.StopsDataCache{ByStopID: make(map[string]types.StopsDataRecord)}
	}

	cache := &types.StopsDataCache{
		ByStopID: make(map[string]types.StopsDataRecord),
	}

	for _, entry := range entries {
		for _, flag := range entry.Flags {
			if flag.StopID != "" {
				if _, exists := cache.ByStopID[flag.StopID]; !exists {
					cache.ByStopID[flag.StopID] = types.StopsDataRecord{
						Name:      entry.Name,
						Latitude:  entry.Latitude,
						Longitude: entry.Longitude,
						Flags: []types.StopsDataFlag{
							{
								AgencyIDs:    flag.AgencyIDs,
								IsHarmonized: flag.IsHarmonized,
								ShortName:    flag.ShortName,
								StopID:       flag.StopID,
							},
						},
					}
				}
			}
		}
	}

	lib.AppLogger.Debug(fmt.Sprintf("Pre-computed stops_data cache for %d stops", len(cache.ByStopID)))
	return cache
}

func readStopsDataFile() ([]byte, string, error) {
	possiblePaths := []string{
		"stops_data.json",
		filepath.Join("..", "stops_data.json"),
		filepath.Join("..", "..", "stops_data.json"),
	}

	for _, path := range possiblePaths {
		fileBytes, err := os.ReadFile(path)
		if err == nil {
			return fileBytes, path, nil
		}
	}

	return nil, "", os.ErrNotExist
}

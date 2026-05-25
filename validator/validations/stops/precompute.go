package stops

import (
	"encoding/json"
	"fmt"
	"main/lib"
	"os"
	"path/filepath"
	"strconv"
)

// buildValidStopIDsSet loads the root-level stops_ids.json into a set for fast stop_id lookups.
//
// Accepted JSON formats:
//   - [100001, 100002, ...]
//   - [{"_id":100001}, {"_id":100002}, ...] (legacy support)
func buildValidStopIDsSet() map[string]struct{} {
	lib.AppLogger.Debug("Pre-computing valid stop IDs from stops_ids.json...")

	fileBytes, sourcePath, err := readStopIDsFile()
	if err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error reading stops_ids.json: %v", err))
		return map[string]struct{}{}
	}
	lib.AppLogger.Debug(fmt.Sprintf("Loaded stops_ids.json from %s", sourcePath))

	validIDs := make(map[string]struct{})

	var ids []int
	if err := json.Unmarshal(fileBytes, &ids); err == nil {
		for _, id := range ids {
			validIDs[strconv.Itoa(id)] = struct{}{}
		}
		lib.AppLogger.Debug(fmt.Sprintf("Pre-computed %d valid stop IDs", len(validIDs)))
		return validIDs
	}

	var legacyIDs []struct {
		ID int `json:"_id"`
	}
	if err := json.Unmarshal(fileBytes, &legacyIDs); err != nil {
		lib.AppLogger.Error(fmt.Sprintf("Error parsing stops_ids.json: %v", err))
		return map[string]struct{}{}
	}

	for _, entry := range legacyIDs {
		validIDs[strconv.Itoa(entry.ID)] = struct{}{}
	}

	lib.AppLogger.Debug(fmt.Sprintf("Pre-computed %d valid stop IDs", len(validIDs)))
	return validIDs
}

func readStopIDsFile() ([]byte, string, error) {
	possiblePaths := []string{
		"stops_ids.json",
		filepath.Join("..", "stops_ids.json"),
		filepath.Join("..", "..", "stops_ids.json"),
	}

	for _, path := range possiblePaths {
		fileBytes, err := os.ReadFile(path)
		if err == nil {
			return fileBytes, path, nil
		}
	}

	return nil, "", os.ErrNotExist
}
